package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(docs)
}

var url = "https://www.doculink.com"

var (
	docs = &cobra.Command{
		Use:     "docs",
		Aliases: []string{"doc"},
		Short:   "links documentation",
		Run:     docCmd,
	}
)

func docCmd(cmd *cobra.Command, args []string) {
	fmt.Println("Opening browser")
	err := exec.Command("open", url).Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening browser to show documentation %s", err)
	}
}
