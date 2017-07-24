package template

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/AlecAivazis/survey"
	"github.com/Masterminds/sprig"
)

func templatePath(path string, data map[string]interface{}) (string, error) {
	raw := strings.NewReader(path)

	var tpl bytes.Buffer
	err := renderTemplate(raw, &tpl, data)

	return tpl.String(), err
}

// renderTemplate is our all purpose wrapper around template redering
func renderTemplate(src io.Reader, dest io.Writer, data map[string]interface{}) error {
	raw, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}

	t, err := template.New("template").Funcs(sprig.TxtFuncMap()).Parse(string(raw))
	if err != nil {
		return err
	}

	return t.Execute(dest, data)
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
