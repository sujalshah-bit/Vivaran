package core

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/sujalshah-bit/Vivaran/pkg/api"
	"github.com/sujalshah-bit/Vivaran/util"
)

type Stats struct {
	Lines int64
	Words int64
	Chars int64
	Size  int64
}

type File struct {
	name string
	file *os.File
	size int64
	stat os.FileInfo
}

func NewFile(path string) *File {
	if strings.TrimSpace(path) == "" {
		util.Check(fmt.Errorf("empty path"))
	}

	f, err := os.Open(path)
	util.Check(err)
	stat, err := f.Stat()
	util.Check(err)

	return &File{
		name: f.Name(),
		file: f,
		size: stat.Size(),
		stat: stat,
	}
}

func (f *File) GetFile() *os.File {
	return f.file
}

func (f *File) Close() {
	f.file.Close()
}

type Stdin struct {
	file *os.File
}

func NewDataFromStdin() *Stdin {
	return &Stdin{file: os.Stdin}
}

func (s *Stdin) GetFile() *os.File {
	return s.file
}

func (f *Stdin) Close() {
	// no-op
}

func Counter(input api.Counter, buf []byte) *Stats {
	var inWord bool
	var countChars, countLines, countSize, countWords int64
	for {
		n, err := input.GetFile().Read(buf)
		if n > 0 {
			chunk := buf[:n]
			countSize += int64(n)
			countChars += int64(utf8.RuneCount(chunk))

			for _, b := range chunk {
				if b == '\n' {
					countLines++
				}
				if b == ' ' || b == '\n' || b == '\t' || b == '\r' {
					inWord = false
				} else if !inWord {
					inWord = true
					countWords++
				}
			}
		}
		if err == io.EOF {
			break
		}
		util.Check(err)
	}

	return &Stats{Chars: countChars, Lines: countLines, Size: countSize, Words: countWords}
}
