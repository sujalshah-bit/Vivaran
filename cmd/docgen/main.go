//go:generate go run .
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sujalshah-bit/Vivaran/pkg/flags"
)

func main() {
	var b strings.Builder

	b.WriteString("## **Supported Flags**\n\n")
	b.WriteString("| Flag | Name | Description | Example |\n")
	b.WriteString("| ---- | ---- | ----------- | ------- |\n")

	for _, f := range flags.Supported {
		fmt.Fprintf(
			&b,
			"| `-%s` | %s | %s | `%s` |\n",
			f.Short, f.Name, f.Desc, f.Example,
		)
	}

	// write standalone file at repo root
	if err := os.WriteFile("../../SUPPORTED_FLAGS.md", []byte(b.String()), 0644); err != nil {
		panic(err)
	}

	if err := updateReadme(b.String()); err != nil {
		panic(err)
	}
}

func updateReadme(flagsTable string) error {
	data, err := os.ReadFile("../../README.md")
	if err != nil {
		return err
	}

	readme := string(data)

	start := "<!-- BEGIN:SUPPORTED_FLAGS -->"
	end := "<!-- END:SUPPORTED_FLAGS -->"

	before, after, found := strings.Cut(readme, start)
	if !found {
		return fmt.Errorf("start marker not found")
	}

	_, after, found = strings.Cut(after, end)
	if !found {
		return fmt.Errorf("end marker not found")
	}

	newReadme := before +
		start + "\n\n" +
		flagsTable + "\n" +
		end +
		after

	return os.WriteFile("../../README.md", []byte(newReadme), 0644)
}
