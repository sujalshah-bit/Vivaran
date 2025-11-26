package main

import (
	"fmt"

	"github.com/sujalshah-bit/Vivaran/config"
	filehandling "github.com/sujalshah-bit/Vivaran/pkg/file_handling"
	stdinhandling "github.com/sujalshah-bit/Vivaran/pkg/stdin_handling"
	"github.com/sujalshah-bit/Vivaran/util"
)

func main() {
	cfg := config.LoadConfig()

	var f filehandling.File
	var d *stdinhandling.Data

	// Open file only when not stdin AND positional arg exists
	if !cfg.IsStdin {
		if len(cfg.Args) == 0 {
			panic("filename required when input is not from stdin")
		}

		f = filehandling.NewFile(cfg.Args[0])

		defer f.Close()
	} else {
		d = stdinhandling.NewDataFromStdin()
	}

	if cfg.Lines || cfg.Size || cfg.Words || cfg.Char {
		if cfg.Size {
			value := util.If(cfg.IsStdin, d.Size(), f.GetSize())
			fmt.Printf("%d ", value)
		}
		if cfg.Lines {
			value := util.If(cfg.IsStdin, d.Lines(), f.GetLines())
			fmt.Printf("%d ", value)
		}
		if cfg.Words {
			value := util.If(cfg.IsStdin, d.Words(), f.GetWords())
			fmt.Printf("%d ", value)
		}
		if cfg.Char {
			value := util.If(cfg.IsStdin, d.UnicodeChars(), f.GetChars())
			fmt.Printf("%d ", value)
		}
	} else {
		if cfg.IsStdin {
			fmt.Printf("%d %d %d %d", d.Size(), d.Lines(), d.Words(), d.UnicodeChars())
		} else {
			fmt.Printf("%d %d %d %d", f.GetSize(), f.GetLines(), f.GetWords(), f.GetChars())
		}
	}

}
