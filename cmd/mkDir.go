package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(mkDir)
}

var (
	mkDir = &cobra.Command{
		Use:     "mkd",
		Aliases: []string{"mkdir"},
		Short:   "creates config directories",
		Run:     mkDirCmd,
	}
)

func mkDirCmd(cmd *cobra.Command, args []string) {
	// util.MakeDirSkeleton()
	// util.CleanUpAllDirs()
}
