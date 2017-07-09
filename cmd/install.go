package cmd

import (
	"fmt"

	"github.com/andytom/tmpltr/template"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(installCmd)
	installCmd.Flags().BoolVarP(&force, "force", "f", false, "Force the installation of a template")
}

var force bool

var installCmd = &cobra.Command{
	Use:   "install NAME DIR",
	Short: "Install a template from a directory",
	Long: `Install a template from a directory

This will install the template found in DIR with the name NAME

If there is already a template with that name you will need to force the
installation which will over write the existing template.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("Need a directory and template")
		}

		templateKey := args[0]
		srcDir := args[1]

		s, err := template.OpenStore(cfgDir)
		if err != nil {
			return fmt.Errorf("Unable to access the template directory %q: %q", cfgDir, err)
		}

		t, err := template.New(srcDir)
		if err != nil {
			return fmt.Errorf("Unable to install template: %q", err)
		}

		_, err = s.Get(templateKey)
		if err == nil {
			// We have an existing template so we need to check if we have to replace or not.
			if force {
				fmt.Printf("Replacing %q with template from %q\n", templateKey, srcDir)
				return s.Replace(templateKey, &t)
			}
			return fmt.Errorf("There is already a template with the name %q", templateKey)
		}

		fmt.Printf("Installing template from %q as %q\n", srcDir, templateKey)
		return s.Create(templateKey, &t)
	},
}
