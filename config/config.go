package config

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

type FileConfig struct {
	Size    bool
	Lines   bool
	Char    bool
	Words   bool
	Args    []string // non-flag CLI arguments
	IsStdin bool
}

func LoadConfig() *FileConfig {
	cfg := &FileConfig{}

	pflag.BoolVar(&cfg.Size, "c", false, "File size")
	pflag.BoolVar(&cfg.Lines, "l", false, "Total lines in the file")
	pflag.BoolVar(&cfg.Words, "w", false, "Total Words in the file")
	pflag.BoolVar(&cfg.Char, "m", false, "Total Characters in the file")
	pflag.Parse()
	cfg.Args = pflag.Args()

	stat, _ := os.Stdin.Stat()
	fromStdin := (stat.Mode() & os.ModeCharDevice) == 0

	fmt.Println("Positional args:", cfg.Args)
	fmt.Println("Is input from stdin?", fromStdin)

	if len(cfg.Args) > 1 && fromStdin {
		panic(fmt.Errorf("only one input is allowed"))
	}
	if len(cfg.Args) < 1 && !fromStdin {
		panic(fmt.Errorf("atleast one input needed"))
	}
	if len(cfg.Args) < 1 && fromStdin {
		cfg.IsStdin = true
	}

	return cfg
}
