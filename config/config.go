package config

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"github.com/sujalshah-bit/Vivaran/pkg/flags"
)

type FileConfig struct {
	Size       bool
	Lines      bool
	Char       bool
	Words      bool
	Default    bool
	BufferSize int
	Args       []string // non-flag CLI arguments
	IsStdin    bool
}

func LoadConfig() *FileConfig {
	cfg := &FileConfig{}

	for _, f := range flags.Supported {
		switch f.Short {
		case "c":
			pflag.BoolVar(&cfg.Size, f.Short, false, f.Desc)
		case "l":
			pflag.BoolVar(&cfg.Lines, f.Short, false, f.Desc)
		case "w":
			pflag.BoolVar(&cfg.Words, f.Short, false, f.Desc)
		case "m":
			pflag.BoolVar(&cfg.Char, f.Short, false, f.Desc)
		case "bs":
			pflag.IntVar(&cfg.BufferSize, f.Short, 32, f.Desc)
		}
	}
	pflag.Parse()
	cfg.Args = pflag.Args()

	stat, _ := os.Stdin.Stat()
	fromStdin := (stat.Mode() & os.ModeCharDevice) == 0

	// fmt.Println("Positional args:", cfg.Args)
	// fmt.Println("Is input from stdin?", fromStdin)

	if len(cfg.Args) > 1 && fromStdin {
		panic(fmt.Errorf("only one input is allowed"))
	}
	if len(cfg.Args) < 1 && !fromStdin {
		panic(fmt.Errorf("atleast one input needed"))
	}
	if len(cfg.Args) < 1 && fromStdin {
		cfg.IsStdin = true
	}
	if !cfg.Size && !cfg.Char && !cfg.Lines && !cfg.Words {
		// Default opt l, w, c
		cfg.Default = true
	}

	return cfg
}
