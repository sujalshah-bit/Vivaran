package filehandling

import (
	"testing"
)

const (
	expectedSize  int64 = 342190
	expectedWords int64 = 58164
	expectedLines int64 = 7145
	expectedChars int64 = 339292
)

func TestFileMethods(t *testing.T) {
	f := NewFile("test.txt")

	t.Cleanup(func() { f.Close() })

	t.Run("File size should match expected value", func(t *testing.T) {
		got := f.GetSize()
		if got != expectedSize {
			t.Errorf("Mismatch: expected %d, got %d", expectedSize, got)
		}
	})

	t.Run("Word count should match expected value", func(t *testing.T) {
		got := f.GetWords()
		if got != expectedWords {
			t.Errorf("Mismatch: expected %d, got %d", expectedWords, got)
		}
	})

	t.Run("Line count should match expected value", func(t *testing.T) {
		got := f.GetLines()
		if got != expectedLines {
			t.Errorf("Mismatch: expected %d, got %d", expectedLines, got)
		}
	})

	t.Run("Character count should match expected Unicode value (wc -m behavior)", func(t *testing.T) {
		got := f.GetChars()
		if got != expectedChars {
			t.Errorf("Mismatch: expected %d, got %d", expectedChars, got)
		}
	})
}
