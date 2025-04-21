package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	linesFlag, wordsFlag, charsFlag, filePath := parseFlags()

	content, err := validateFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	text := string(content)
	printCounts(text, filePath, linesFlag, wordsFlag, charsFlag)
}

func validateFile(filePath string) (string, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		if os.IsPermission(err) {
			fmt.Fprintf(os.Stderr, "wc: %s: Permission denied\n", filePath)
		} else {
			fmt.Fprintf(os.Stderr, "wc: %s: %v\n", filePath, err)
		}
		os.Exit(1)
	}

	if info.IsDir() {
		fmt.Fprintf(os.Stderr, "wc: %s: read: Is a directory\n", filePath)
		os.Exit(1)
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsPermission(err) {
			fmt.Fprintf(os.Stderr, "wc: %s: Permission denied\n", filePath)
		} else {
			fmt.Fprintf(os.Stderr, "wc: %s: %v\n", filePath, err)
		}
		os.Exit(1)
	}

	return string(content), nil
}

func parseFlags() (linesFlag, wordsFlag, charsFlag bool, filePath string) {
	lFlag := flag.Bool("l", false, "Count lines")
	wFlag := flag.Bool("w", false, "Count words")
	cFlag := flag.Bool("c", false, "Count characters")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "Please provide a file path")
		os.Exit(1)
	}
	filePath = flag.Arg(0)

	if !*lFlag && !*wFlag && !*cFlag {
		*lFlag, *wFlag, *cFlag = true, true, true
	}

	return *lFlag, *wFlag, *cFlag, filePath
}

func printCounts(text, filePath string, linesFlag, wordsFlag, charsFlag bool) {
	if linesFlag {
		fmt.Printf("%8d", countLines(text))
	}
	if wordsFlag {
		fmt.Printf("%8d", countWords(text))
	}
	if charsFlag {
		fmt.Printf("%8d", countChars(text))
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
	return utf8.RuneCountInString(text)
}
