package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func printUsage() {
	fmt.Println("Usage: ./csvshark <FILENAME>.csv")
}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		printUsage()
		return

	}
	// Get the filename
	csvFile := os.Args[1]

	file, err := os.Open(csvFile) // -> (*os.File, error)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// get file info
	fileInfo, err := file.Stat() // -> (os.FileInfo, error)
	if err != nil {
		log.Println("Cannot read file stats:", err)
	}

	// create the byte array & include file size length
	data := make([]byte, fileInfo.Size())

	// Read in the bytes
	b, err := file.Read(data)
	if err != nil {
		log.Println("Cannot read file:", err)
	}
	fmt.Printf("Read %d bytes from file\n", b)

	// Create a csv reader for the Read() method. Read returns a sliced []record
	r := csv.NewReader(strings.NewReader(string(data)))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record) // []string
	}
}
