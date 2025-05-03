package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Packet struct {
	No       string
	Time     string
	Src      string
	Dst      string
	Protocol string
	Length   int
	Info     string
}

func (p *Packet) NetworkPacket() {
	fmt.Println(p.Dst) // NICE! 👍🏻🚀🐊👏🏻
}

func NewPacket(rows [][]string) []Packet {
	var networkData []Packet
	for _, row := range rows {
		length, _ := strconv.ParseInt(row[0], 0, 0)
		p := Packet{
			No:       row[0],
			Time:     row[1],
			Src:      row[2],
			Dst:      row[3],
			Protocol: row[4],
			Length:   int(length),
			Info:     row[6],
		}
		networkData = append(networkData, p)
	}
	return networkData
}

func printUsage() {
	fmt.Println("Usage: ./csvshark <FILENAME>.csv")
}

// ----------------------------------------------

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		printUsage()
		return
	}

	// Get the filename
	csvFile := os.Args[1]

	file, err := os.Open(csvFile) // -> (*os.File, error) file, err := os.Open(csvFile) // -> (*os.File, error)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Read() // use Read to remove the header

	rows, err := reader.ReadAll() // use ReadAll to read resto file
	if err != nil {
		log.Println("Cannot read CSV file:", err)
	}

	// Construct the struct
	networkData := NewPacket(rows)

	// To call the method on all Packets, you need a loop:
	for i := range networkData {
		networkData[i].NetworkPacket()
	}
}
