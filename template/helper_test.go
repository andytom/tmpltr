package template

import (
	"reflect"
	"testing"

	"github.com/AlecAivazis/survey"
)

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
		{
			name:     "Invalid template - Missing final '}'",
			input:    "{{ .a }",
			hasError: true,
		},
		{
			name:     "Invalid Template - Invalid function",
			input:    "{{ NoSuchFunction }}",
			hasError: true,
		},
		{
			name:     "Invalid Template - Broken Import",
			input:    "{{ template \"No Such Template\" . }}",
			hasError: true,
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
