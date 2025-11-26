package filehandling

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type File struct {
	name string
	file *os.File
	size int64
	stat os.FileInfo
}

func NewFile(path string) File {
	if strings.TrimSpace(path) == "" {
		check(fmt.Errorf("empty path"))
	}

	f, err := os.Open(path)
	check(err)
	stat, err := f.Stat()
	check(err)

	return File{
		name: f.Name(),
		file: f,
		size: stat.Size(),
		stat: stat,
	}
}

func (f *File) GetChars() int64 {
	if _, err := f.file.Seek(0, io.SeekStart); err != nil {
		check(err)
	}

	data, err := io.ReadAll(f.file)
	if err != nil {
		check(err)
	}

	return int64(utf8.RuneCount(data))
}

func (f *File) GetWords() int64 {
	_, err := f.file.Seek(0, io.SeekStart)
	check(err)

	scanner := bufio.NewScanner(f.file)
	scanner.Split(bufio.ScanWords)
	const maxCapacity = 1024 * 1024 // 1MB
	buf := make([]byte, 64*1024)
	scanner.Buffer(buf, maxCapacity)

	var count int64
	for scanner.Scan() {
		count++
	}

	check(scanner.Err())
	return count
}

func (f *File) GetLines() int64 {
	_, err := f.file.Seek(0, io.SeekStart)
	check(err)

	scanner := bufio.NewScanner(f.file)

	const maxCapacity = 1024 * 1024 // 1MB
	buf := make([]byte, 64*1024)
	scanner.Buffer(buf, maxCapacity)

	var count int64
	for scanner.Scan() {
		count++
	}

	check(scanner.Err())

	return count
}

func (f *File) GetSize() int64 {
	return f.size
}

func (f *File) Close() {
	f.file.Close()
}
