package main

import (
	"github.com/sujalshah-bit/Vivaran/config"
	"github.com/sujalshah-bit/Vivaran/pkg/orchestration"
)

func main() {
	cfg := config.LoadConfig()
	orchestration.Orchestrate(*cfg)
}
