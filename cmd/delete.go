package cmd

import (
	"errors"
	"fmt"

	"github.com/andytom/tmpltr/template"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete NAME",
	Short: "Delete a template",
	Long: `Delete a template

This command removes the template with the name NAME from tmpltr's internal store.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Need a template name to delete")
		}

		templateKey := args[0]

		s, err := template.OpenStore(cfgDir)
		if err != nil {
			return fmt.Errorf("Unable to access the template directory %q: %q", cfgDir, err)
		}

		fmt.Printf("Deleting template %q\n", templateKey)

		return s.Delete(templateKey)
	},
}
