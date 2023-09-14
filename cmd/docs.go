package cmd

import (
	"fmt"

	"github.com/bkr-dev/setup-clean/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(docs)
}

var url = "https://www.wikipedia.com"

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
	util.OpenSingleLinkInBrowser(url)
}
