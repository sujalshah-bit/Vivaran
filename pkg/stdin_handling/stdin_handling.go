package stdinhandling

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"unicode/utf8"
)

type Data struct {
	data []byte
}

func NewDataFromStdin() *Data {
	d, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return &Data{data: d}
}

func (d *Data) Size() int64 {
	return int64(len(d.data))
}

func (d *Data) Lines() int64 {
	scanner := bufio.NewScanner(bufio.NewReader(d.ToReader()))
	buf := make([]byte, 64*1024)
	scanner.Buffer(buf, 5*1024*1024)

	var count int64
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return count
}

func (d *Data) Words() int64 {
	scanner := bufio.NewScanner(bufio.NewReader(d.ToReader()))
	scanner.Split(bufio.ScanWords)
	buf := make([]byte, 64*1024)
	scanner.Buffer(buf, 2*1024*1024)

	var count int64
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return count
}

func (d *Data) UnicodeChars() int64 {
	return int64(utf8.RuneCount(d.data))
}

func (d *Data) ToReader() io.Reader {
	r := bytes.NewReader(d.data)
	r.Seek(0, io.SeekStart)
	return r
}
