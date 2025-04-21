package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	linesFlag := flag.Bool("l", false, "Count lines")
	wordsFlag := flag.Bool("w", false, "Count words")
	charsFlag := flag.Bool("c", false, "Count characters")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "Please provide a file path")
		os.Exit(1)
	}

	filePath := flag.Arg(0)
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	text := string(content)
	lines := countLines(text)
	words := countWords(text)
	chars := countChars(text)

	// If no flag is set, show all counts (default wc behavior)
	if !*linesFlag && !*wordsFlag && !*charsFlag {
		*linesFlag, *wordsFlag, *charsFlag = true, true, true
	}

	// Print in wc-style right-aligned format
	if *linesFlag {
		fmt.Printf("%8d", lines)
	}
	if *wordsFlag {
		fmt.Printf("%8d", words)
	}
	if *charsFlag {
		fmt.Printf("%8d", chars)
	}
	fmt.Printf(" %s\n", filePath)
}

func countLines(text string) int {
	if len(text) == 0 {
		return 0
	}
	scanner := bufio.NewScanner(strings.NewReader(text))
	lines := 0
	for scanner.Scan() {
		lines++
	}
	return lines - 1
}

func countWords(text string) int {
	return len(strings.Fields(text))
}

func countChars(text string) int {
	return len([]rune(text))
}
