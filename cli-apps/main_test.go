package main

import (
	"os"
	"testing"
)

func createTempFile(t *testing.T, content string) string {
	t.Helper()
	tmpFile, err := os.CreateTemp("", "testfile*.txt")
	if err != nil {
		t.Fatalf("Error creating temp file: %v", err)
	}
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Could not write to temp file: %v", err)
	}
	tmpFile.Close()
	return tmpFile.Name()
}

func readFileContent(t *testing.T, filePath string) string {
	t.Helper()
	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}
	return string(data)
}

func TestCountLines(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected int
	}{
		{"Empty file", "", 0},
		{"One line", "Hello, World!", 0},
		{"Multiple lines", "Line 1\nLine 2\nLine 3\n", 2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			path := createTempFile(t, tc.content)
			defer os.Remove(path)
			got := countLines(readFileContent(t, path))
			if got != tc.expected {
				t.Errorf("Expected %d lines, got %d", tc.expected, got)
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected int
	}{
		{"Empty file", "", 0},
		{"One line with multiple words", "Hello World from Go", 4},
		{"Multiple lines", "First line\nSecond line\nThird line", 6},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			path := createTempFile(t, tc.content)
			defer os.Remove(path)
			got := countWords(readFileContent(t, path))
			if got != tc.expected {
				t.Errorf("Expected %d words, got %d", tc.expected, got)
			}
		})
	}
}

func TestCountChars(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected int
	}{
		{"Empty file", "", 0},
		{"ASCII text", "Hello", 5},
		{"Unicode text", "अअअअअअ", 6},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			path := createTempFile(t, tc.content)
			defer os.Remove(path)
			got := countChars(readFileContent(t, path))
			if got != tc.expected {
				t.Errorf("Expected %d characters, got %d", tc.expected, got)
			}
		})
	}
}
