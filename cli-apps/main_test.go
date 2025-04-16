package main

import (
	"testing"
)

func TestFindLineCount(t *testing.T) {
	fileMap := []struct {
		name      string
		filepath  string
		expOutput int
		wantErr   bool
	}{
		{
			name:      "valid file with 2 lines",
			filepath:  "test.txt",
			expOutput: 2,
			wantErr:   false,
		},
		{
			name:      "empty file",
			filepath:  "empty_file.txt",
			expOutput: 0,
			wantErr:   false,
		},
		{
			name:     "non-existent file",
			filepath: "non_existent.txt",
			wantErr:  true,
		},
	}

	for _, tt := range fileMap {
		t.Run(tt.name, func(t *testing.T) {
			count, err := FindLineCount(tt.filepath)

			if tt.wantErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if count != tt.expOutput {
				t.Errorf("expected %d lines, got %d", tt.expOutput, count)
			}
		})
	}
}
