package core

import (
	"testing"
)

func TestFileMethods(t *testing.T) {
	f := NewFile("test.txt")

	t.Cleanup(func() { f.Close() })

	// TODO
}
