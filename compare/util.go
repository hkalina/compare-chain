package main

import (
	"bufio"
	"log"
	"os"
)

func readFileRows(filename string, callback func(string)()) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // skip first row (table header)
	if skipRows != 0 {
		for i:= int64(0); i < skipRows; i++ {
			scanner.Scan()
		}
		log.Printf("Skipped %d rows", skipRows)
	}
	for scanner.Scan() {
		callback(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
