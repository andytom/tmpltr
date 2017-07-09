package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var (
	version string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Long:  `All software has versions this is ours's.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
