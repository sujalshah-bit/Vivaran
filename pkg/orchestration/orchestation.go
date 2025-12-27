package orchestration

import (
	"errors"
	"fmt"

	"github.com/sujalshah-bit/Vivaran/config"
	"github.com/sujalshah-bit/Vivaran/pkg/api"
	"github.com/sujalshah-bit/Vivaran/pkg/core"

	"github.com/sujalshah-bit/Vivaran/util"
)

const (
	B  = 1 << (10 * iota) // iota = 0 → 1 << 0 = 1 (discarded)
	KB                    // iota = 1 → 1 << 10 = 1024
	MB                    // iota = 2 → 1 << 20 = 1,048,576
)

func resolveInput(cfg config.FileConfig) (api.Counter, error) {
	if cfg.IsStdin {
		return core.NewDataFromStdin(), nil
	}

	if len(cfg.Args) == 0 {
		return nil, errors.New("filename required when input is not stdin")
	}

	return core.NewFile(cfg.Args[0]), nil
}

func Orchestrate(cfg config.FileConfig) *api.Output {
	input, err := resolveInput(cfg)
	util.Check(err)

	defer input.Close()

	// TODO: Make measurement units configurable.
	buf := make([]byte, cfg.BufferSize*KB)
	inputStats := core.Counter(input, buf)

	var str string

	if cfg.Lines || cfg.Default {
		str += fmt.Sprintf("Lines: %d ", inputStats.Lines)
	}

	if cfg.Words || cfg.Default {
		str += fmt.Sprintf("Words: %d ", inputStats.Words)
	}

	if cfg.Size || cfg.Default {
		str += fmt.Sprintf("Size: %d ", inputStats.Size)
	}

	if cfg.Char {
		str += fmt.Sprintf("Char: %d ", inputStats.Chars)
	}

	fmt.Print(str)
	return &api.Output{Str: str}

}
