package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(mkDir)
	mkDir.PersistentFlags().String("rm-all", "", "remove all directories")
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
	fmt.Println(args)
	rmAll, err := cmd.Flags().GetString("rm-all")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rmAll)
	rm := cmd.Flag("rm-all")
	fmt.Println(rm)
	os.Exit(1)
	// util.MakeDirSkeleton()
	// util.CleanUpAllDirs()
}
