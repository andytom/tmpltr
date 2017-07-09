package cmd

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/andytom/tmpltr/template"

	"github.com/AlecAivazis/survey"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(useCmd)
}

var useCmd = &cobra.Command{
	Use:   "use NAME DIR",
	Short: "Apply a template to a directory",
	Long: `Apply a template to a directory

This applies the template with the name NAME in directory DIR.
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("Need a directory and template")
		}

		templateKey := args[0]
		dest := filepath.Clean(args[1])

		s, err := template.OpenStore(cfgDir)
		if err != nil {
			return fmt.Errorf("Unable to access the template directory %q: %q", cfgDir, err)
		}

		// Load the template
		t, err := s.Get(templateKey)
		if err != nil {
			return fmt.Errorf("Was unable to load template %q: %q", templateKey, err)
		}

		fmt.Printf("Loaded template %q\n", templateKey)

		// Ask the questions
		questions := t.GetQuestions()
		answers := map[string]interface{}{}

		err = survey.Ask(questions, &answers)
		if err != nil {
			return fmt.Errorf("Unable to get answers to questions: %q", err)
		}

		// Execute the template
		fmt.Printf("Executing in %q\n", dest)
		return t.Execute(dest, answers)
	},
}
