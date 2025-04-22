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
	linesFlag, wordsFlag, charsFlag, filePaths := parseFlags()
	fileContents := validateFiles(filePaths)

	var totalLines, totalWords, totalChars int

	for path, content := range fileContents {
		printCounts(content, path, linesFlag, wordsFlag, charsFlag)

		if linesFlag {
			totalLines += countLines(content)
		}
		if wordsFlag {
			totalWords += countWords(content)
		}
		if charsFlag {
			totalChars += countChars(content)
		}
	}

	if len(fileContents) > 1 {
		if linesFlag {
			fmt.Printf("%8d", totalLines)
		}
		if wordsFlag {
			fmt.Printf("%8d", totalWords)
		}
		if charsFlag {
			fmt.Printf("%8d", totalChars)
		}
		fmt.Printf(" total\n")
	}
}

func parseFlags() (linesFlag, wordsFlag, charsFlag bool, filePaths []string) {
	lFlag := flag.Bool("l", false, "Count lines")
	wFlag := flag.Bool("w", false, "Count words")
	cFlag := flag.Bool("c", false, "Count characters")
	flag.Parse()
	filePaths = flag.Args()
	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "Please provide a file path")
		os.Exit(1)
	}

	if !*lFlag && !*wFlag && !*cFlag {
		*lFlag, *wFlag, *cFlag = true, true, true
	}

	return *lFlag, *wFlag, *cFlag, filePaths
}

func validateFiles(filePaths []string) map[string]string {
	results := make(map[string]string)
	for _, path := range filePaths {
		content, ok := checkFileValid(path)
		if ok {
			results[path] = content
		}
	}
	return results
}

func checkFileValid(filePath string) (string, bool) {
	info, err := os.Stat(filePath)
	if err != nil {
		if os.IsPermission(err) {
			fmt.Fprintf(os.Stderr, "wc: %s: Permission denied\n", filePath)
		} else {
			fmt.Fprintf(os.Stderr, "wc: %s: %v\n", filePath, err)
		}
		return "", false
	}

	if info.IsDir() {
		fmt.Fprintf(os.Stderr, "wc: %s: read: Is a directory\n", filePath)
		return "", false
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsPermission(err) {
			fmt.Fprintf(os.Stderr, "wc: %s: Permission denied\n", filePath)
		} else {
			fmt.Fprintf(os.Stderr, "wc: %s: %v\n", filePath, err)
		}
		return "", false
	}

	return string(content), true
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
