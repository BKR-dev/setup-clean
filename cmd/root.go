package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "setup-cli",
		Short: "use to setup local environemnt",
		Long: `setup-cli is used to setup your local
		machine in a state that you are ready to
		connect to the clusters. follow steps carefully and attentively.`,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func Execute() error {
	return rootCmd.Execute()
}
