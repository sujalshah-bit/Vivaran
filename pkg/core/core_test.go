package core

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCounterBasic(t *testing.T) {
	tests := []struct {
		name      string
		wantLines int64
		wantWords int64
		wantChars int64
		wantSize  int64
	}{
		{"empty.txt", 0, 0, 0, 0},
		{"single-word.txt", 0, 1, 5, 5},
		{"multi-line.txt", 2, 4, 20, 20},
		{"unicode.txt", 1, 2, 6, 14},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir, _ := os.Getwd()
			filePath := filepath.Join(dir, "test-data/", tt.name)
			file, err := os.Open(filePath)
			if err != nil {
				t.Error(err)
			}
			defer file.Close()
			reader := &File{file: file}
			buf := make([]byte, 32) // small buffer to test chunking
			got := Counter(reader, buf)

			if got.Lines != tt.wantLines {
				t.Errorf("Lines = %d, wantLines %d", got.Lines, tt.wantLines)
			}
			if got.Words != tt.wantWords {
				t.Errorf("Words = %d, wantWords %d", got.Words, tt.wantWords)
			}
			if got.Chars != tt.wantChars {
				t.Errorf("Chars = %d, wantChars %d", got.Chars, tt.wantChars)
			}
			if got.Size != tt.wantSize {
				t.Errorf("Size = %d, wantSize %d", got.Size, tt.wantSize)
			}
		})
	}
}
