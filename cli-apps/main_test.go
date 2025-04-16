package main

import (
	"io/fs"
	"testing"
)

func TestFindLineCount(t *testing.T) {
	fileMap := []struct {
		name      string
		filepath  string
		expOutput int
		wantErr   bool
		expErr    error
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
			expErr:   fs.ErrNotExist,
		},
	}

	for _, tt := range fileMap {
		t.Run(tt.name, func(t *testing.T) {
			gotCount, gotErr := FindLineCount(tt.filepath)

			if tt.wantErr {
				if gotErr == nil {
					t.Fatalf("Epected error but got nil")
				}

				if tt.expErr != gotErr {
					t.Errorf("Expected %v but got %v", tt.expErr, gotErr)
				}

				return
			}

			if tt.expOutput != gotCount {
				t.Errorf("Expected %v but got %v", tt.expOutput, gotCount)
			}
		})
	}
}
