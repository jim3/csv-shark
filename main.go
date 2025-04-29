package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type fileSizeError struct {
	message string
	size    int
}

func main() {
	// Check if an argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./csvshark <FILENAME>.csv")
		return
	}

	// TODO: Error checking
	// Check if the file exists and is readable and has a .csv extension

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
