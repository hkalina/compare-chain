package main

import (
	"bufio"
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
	for scanner.Scan() {
		callback(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
