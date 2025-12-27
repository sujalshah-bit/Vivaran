package api

import (
	"os"
)

type Counter interface {
	GetFile() *os.File
	Close()
}

type Output struct {
	Str string
}
