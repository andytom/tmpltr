package template

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/AlecAivazis/survey"
)

func templatePath(path string, data map[string]interface{}) (string, error) {
	t, err := template.New("path").Parse(path)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, data)

	return tpl.String(), err
}

func parseQuestion(q question) (*survey.Question, error) {

	sq := survey.Question{Name: q.Name}

	switch q.Type {
	case "input":
		sq.Prompt = &survey.Input{
			Message: q.Message,
			Help:    q.Help,
			Default: q.Default,
		}
	case "select":
		sq.Prompt = &survey.Select{
			Message: q.Message,
			Help:    q.Help,
			Options: q.Options,
			Default: q.Default,
		}
	default:
		return &sq, fmt.Errorf("Unknown question type %s", q.Type)
	}

	if q.Required {
		sq.Validate = survey.Required
	}

	return &sq, nil
}
