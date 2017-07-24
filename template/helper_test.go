package template

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/AlecAivazis/survey"
)

func TestRenderTemplate(t *testing.T) {
	testCases := [...]struct {
		name          string
		inputTemplate io.Reader
		inputData     map[string]interface{}
		expected      string
		hasError      bool
	}{
		{
			name:          "Non-template input",
			inputTemplate: strings.NewReader("test"),
			expected:      "test",
			hasError:      false,
		},
		{
			name:          "Basic Input",
			inputTemplate: strings.NewReader("{{ .a }}"),
			inputData:     map[string]interface{}{"a": "test"},
			expected:      "test",
			hasError:      false,
		},
		{
			name: "Invalid Template - Broken input template",
			// This creates a nil pointer to an os.File which will
			// act like an io.Reader but it will error when you try
			// to read from it.
			inputTemplate: (*os.File)(nil),
			hasError:      true,
		},
		{
			name:          "Invalid template - Missing final '}'",
			inputTemplate: strings.NewReader("{{ .a }"),
			hasError:      true,
		},
		{
			name:          "Invalid Template - Invalid function",
			inputTemplate: strings.NewReader("{{ NoSuchFunction }}"),
			hasError:      true,
		},
		{
			name:          "Invalid Template - Broken Import",
			inputTemplate: strings.NewReader("{{ template \"No Such Template\" . }}"),
			hasError:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			var out bytes.Buffer

			err := renderTemplate(tc.inputTemplate, &out, tc.inputData)

			if tc.hasError {
				if err == nil {
					t.Errorf("Expected an error but didn't get one!")
				}
			} else {
				if err != nil {
					t.Errorf("Got an unexpected error %q", err)
				}
			}

			if out.String() != tc.expected {
				t.Errorf("Expected %q but got %q", tc.expected, out.String())
			}

		})
	}
}

func TestTemplatePath(t *testing.T) {
	testCases := [...]struct {
		name     string
		input    string
		expected string
		data     map[string]interface{}
		hasError bool
	}{
		{
			name:     "Non-template input",
			input:    "test",
			expected: "test",
			hasError: false,
		},
		{
			name:     "Super simple basic template",
			input:    "{{ .a }}",
			expected: "test",
			data:     map[string]interface{}{"a": "test"},
			hasError: false,
		},
		{
			name:     "Realistic example template",
			input:    "/path/to/some/{{ .a }}/dir",
			expected: "/path/to/some/test/dir",
			data:     map[string]interface{}{"a": "test"},
			hasError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := templatePath(tc.input, tc.data)

			if tc.hasError {
				if err == nil {
					t.Errorf("Expected an error but didn't get one!")
				}
			} else {
				if err != nil {
					t.Errorf("Got an unexpected error %q", err)
				}
			}

			if result != tc.expected {
				t.Errorf("Got %q but wanted %q", result, tc.expected)
			}
		})
	}
}

func TestParseQuestion(t *testing.T) {

	testCases := [...]struct {
		name     string
		input    question
		expected survey.Question
		hasError bool
	}{
		{
			name:     "Question",
			input:    question{},
			hasError: true,
		},
		{
			name: "Input",
			input: question{
				Name:    "Test",
				Type:    "input",
				Message: "Message",
				Help:    "Help",
				Default: "Default",
				//Required: false,
			},
			expected: survey.Question{
				Name: "Test",
				Prompt: &survey.Input{
					Message: "Message",
					Help:    "Help",
					Default: "Default",
				},
			},
			hasError: false,
		},
		{
			name: "Required Input",
			input: question{
				Name:     "Test",
				Type:     "input",
				Message:  "Message",
				Help:     "Help",
				Default:  "Default",
				Required: true,
			},
			expected: survey.Question{
				Name: "Test",
				Prompt: &survey.Input{
					Message: "Message",
					Help:    "Help",
					Default: "Default",
				},
			},
			hasError: false,
		},
		{
			name: "Select",
			input: question{
				Name:    "Test",
				Type:    "select",
				Message: "Message",
				Help:    "Help",
				Default: "Default",
				Options: []string{"a", "b", "c"},
			},
			expected: survey.Question{
				Name: "Test",
				Prompt: &survey.Select{
					Message: "Message",
					Help:    "Help",
					Default: "Default",
					Options: []string{"a", "b", "c"},
				},
			},
			hasError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := parseQuestion(tc.input)

			if tc.hasError {
				if err == nil {
					t.Fatalf("Expected an error but didn't get one!")
				}
			} else {
				if err != nil {
					t.Fatalf("Got an unexpected error %q", err)
				}
			}

			// These comparisons are probably lacking somethnig but
			// they seem to cover most cases so think we are good
			// here for now.
			if result.Name != tc.expected.Name {
				t.Errorf("Names do not match got %q but wanted %q", result.Name, tc.expected.Name)
			}

			if !reflect.DeepEqual(result.Prompt, tc.expected.Prompt) {
				t.Errorf("Prompts do not match got %#v but wanted %#v", result.Prompt, tc.expected.Prompt)
			}

			if tc.input.Required && result.Validate == nil {
				t.Errorf("Validates is nil but required is %t", tc.input.Required)
			}
		})
	}
}
