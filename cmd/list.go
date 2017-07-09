package cmd

import (
	"fmt"
	"os"

	"github.com/andytom/tmpltr/template"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print a list of all installed templates",
	Long: `Print a list of all installed template

Prints out a list of all templates that are stored in the config directory.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := template.OpenStore(cfgDir)
		if err != nil {
			return fmt.Errorf("Unable to access the template directory %q: %q", cfgDir, err)
		}

		templates, err := s.List()
		if err != nil {
			return fmt.Errorf("Unable to list the templates: %q", err)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Description"})

		for key, template := range templates {
			table.Append([]string{key, template.Meta.Description})
		}

		table.Render()
		return nil
	},
}
