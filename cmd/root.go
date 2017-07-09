package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var cfgDir string

// RootCmd is the root Cobra Command for the CLI. It has persistent flags to
// set the config directory.
var RootCmd = &cobra.Command{
	Use:   "tmpltr",
	Short: "Tmpltr is a template manager that generates files or directories from templates.",
	Long: `Tmpltr is a template manager that generates files and directories from templates.

Tmpltr stores template files in the configuration directory.
`,
}

func init() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println("Unable to get home dir for the current user")
		os.Exit(1)
	}

	defaultCfgDir := filepath.Join(home, ".tmpltr")

	RootCmd.PersistentFlags().StringVarP(&cfgDir, "config", "c", defaultCfgDir, "Tmpltr configuration directory")
}
