package template

import (
	"reflect"
	"testing"

	"github.com/AlecAivazis/survey"
)

func TestNewTemplate(t *testing.T) {
	testCases := [...]struct {
		name     string
		input    string
		expected Template
		hasError bool
	}{
		{
			name:  "Basic Template",
			input: "test_fixture/basic",
			expected: Template{
				baseDir: "test_fixture/basic",
				Meta: meta{
					Description: "A basic template",
				},
			},
		},
		{
			name:     "Invalid Template - Missing Directory",
			input:    "no/such/template",
			expected: Template{baseDir: "no/such/template"},
			hasError: true,
		},
		{
			name:     "Invalid Template - Empty Dir",
			input:    "test_fixture/empty",
			expected: Template{baseDir: "test_fixture/empty"},
			hasError: true,
		},
		{
			name:     "Invalid Config - Missing fields",
			input:    "test_fixture/invalid_config",
			expected: Template{baseDir: "test_fixture/invalid_config"},
			hasError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := New(tc.input)

			if tc.hasError {
				if err == nil {
					t.Fatalf("Expected an error but didn't get one!")
				}
			} else {
				if err != nil {
					t.Fatalf("Got an unexpected error %q", err)
				}
			}

			if !reflect.DeepEqual(result.Meta, tc.expected.Meta) {
				t.Fatalf("Meta Data is different. Expected %+v but got %+v",
					tc.expected.Meta,
					result.Meta,
				)
			}

			if !reflect.DeepEqual(result.Questions, tc.expected.Questions) &&
				len(tc.expected.Questions) > 0 {

				t.Fatalf("Questions are different. Expected %+v but got %+v",
					tc.expected.Questions,
					result.Questions,
				)
			}

		})
	}
}

func TestGetQuestions(t *testing.T) {

	testCases := [...]struct {
		name     string
		template Template
		expected []*survey.Question
	}{
		{
			name: "Single Question",
			template: Template{
				Questions: []question{
					{
						Name:    "Test",
						Type:    "input",
						Message: "Message",
					},
				},
			},
			expected: []*survey.Question{
				{
					Name: "Test",
					Prompt: &survey.Input{
						Message: "Message",
					},
				},
			},
		},
		{
			name: "Multiple Question",
			template: Template{
				Questions: []question{
					{
						Name:    "Test 1",
						Type:    "input",
						Message: "Message",
					},
					{
						Name:    "Test 2",
						Type:    "input",
						Message: "Message",
					},
				},
			},
			expected: []*survey.Question{
				{
					Name: "Test 1",
					Prompt: &survey.Input{
						Message: "Message",
					},
				},
				{
					Name: "Test 2",
					Prompt: &survey.Input{
						Message: "Message",
					},
				},
			},
		},
		{
			name: "Dummy Type",
			template: Template{
				Questions: []question{
					{
						Type: "Dummy",
					},
				},
			},
			expected: []*survey.Question{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.template.GetQuestions()

			if len(tc.expected) != len(result) {
				t.Fatalf("Expected %d questions, but got %d", len(tc.expected), len(result))
			}

			for i, q := range tc.expected {
				if result[i].Name != q.Name {
					t.Errorf("Names do not match got %q but wanted %q", result[i].Name, q.Name)
				}

				if !reflect.DeepEqual(result[i].Prompt, q.Prompt) {
					t.Errorf("Prompts do not match got %+v but wanted %+v", result[i].Prompt, q.Prompt)
				}
			}

		})
	}
}
