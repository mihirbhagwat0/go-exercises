package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	lineFlag := flag.Bool("l", false, "Print line count")

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("provide a file path")
		os.Exit(1)
	}

	filePath := args[0]

	if *lineFlag {
		lineCount, err := FindLineCount(filePath)
		if err != nil {
			fmt.Printf("error %v", err)
		}
		fmt.Printf("line count is %d", lineCount)
	}
}

func FindLineCount(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	return lineCount, nil
}
