package main

import (
	"fmt"
	"os"

	"github.com/bkr-dev/setup-clean/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing cli-tool: %s", err)
		os.Exit(1)
	}
}
