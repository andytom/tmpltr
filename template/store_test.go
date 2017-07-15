package template

import (
	"testing"
)

func TestGetPath(t *testing.T) {
	testCases := []struct {
		name     string
		store    Store
		input    string
		expected string
	}{
		{
			name:     "Basic test case",
			store:    Store{"/tmp"},
			input:    "test",
			expected: "/tmp/test",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.store.getPath(tc.input)

			if result != tc.expected {
				t.Fatalf("Expected %q but got %q", tc.expected, result)
			}
		})
	}
}
