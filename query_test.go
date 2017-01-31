package sbb

import (
	"strings"
	"testing"
)

type testCase struct {
	input    []string
	from, to string
}

func TestParseQuery(t *testing.T) {
	testCases := []testCase{
		testCase{
			input: []string{"from", "Brig", "to", "Horn"},
			from:  "Brig",
			to:    "Horn",
		},
		testCase{
			input: []string{"to", "Bern", "from", "Zurich"},
			from:  "Zurich",
			to:    "Bern",
		},
		testCase{
			input: []string{"from", "St.", "Gallen", "to", "St.", "Margrethen"},
			from:  "St. Gallen",
			to:    "St. Margrethen",
		},
		testCase{
			input: []string{"to", "Zürich", "HB", "from", "Schaffhausen"},
			from:  "Schaffhausen",
			to:    "Zürich HB",
		},
	}

	for _, tc := range testCases {
		act := ParseQuery(tc.input)
		if act.From != tc.from {
			t.Errorf("From is `%s` (expected `%s`)", act.From, tc.from)
		}

		if act.To != tc.to {
			t.Errorf("To is `%s` (expected `%s`)", act.To, tc.to)
		}
	}
}

func TestQueryURL(t *testing.T) {
	actual := (&Query{From: "Foo", To: "Bar"}).URL()
	expectedParts := []string{
		APIBase,
		"?",
		"from=Foo",
		"to=Bar",
	}
	for _, field := range RequiredFields {
		encodedField := strings.Replace(field, "/", "%2F", -1)
		expectedParts = append(expectedParts, "fields%5B%5D="+encodedField)
	}

	for _, part := range expectedParts {
		if !strings.Contains(actual, part) {
			t.Errorf("Expected '%s' to contain '%s'", actual, part)
		}
	}
}
