package sbb

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

var testCases = []struct {
	input string
	conn  *Query
}{
	{
		input: "from Brig to Horn",
		conn:  &Query{From: "Brig", To: "Horn"},
	},
	{
		input: "von Brig nach Horn",
		conn:  &Query{From: "Brig", To: "Horn"},
	},
	{
		input: "to Bern from Zurich",
		conn:  &Query{From: "Zurich", To: "Bern"},
	},
	{
		input: "from St. Gallen to St. Margrethen",
		conn:  &Query{From: "St. Gallen", To: "St. Margrethen"},
	},
	{
		input: "from Zürich HB to Bern via Brig",
		conn:  &Query{From: "Zürich HB", To: "Bern", Via: []string{"Brig"}},
	},
	{
		input: "from A to B via C via D via E",
		conn:  &Query{From: "A", To: "B", Via: []string{"C", "D", "E"}},
	},
	{
		input: "from Brig to Horn on 2017/01/02",
		conn:  &Query{From: "Brig", To: "Horn", Date: "2017-01-02"},
	},
}

func TestParseQuery(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			act := ParseQuery(strings.Split(tc.input, " "))
			if !reflect.DeepEqual(act, tc.conn) {
				t.Errorf("got %#v\n     expected %#v", act, tc.conn)
			}
		})
	}
}

type testQueryCase struct {
	q Query
	p []string
}

func TestQueryURL(t *testing.T) {
	cases := []testQueryCase{
		testQueryCase{
			q: Query{From: "Foo", To: "Bar"},
			p: []string{"from=Foo", "to=Bar"},
		},
		testQueryCase{
			q: Query{From: "A", To: "B", Via: []string{"C"}},
			p: []string{"from=A", "to=B", "via[]=C"},
		},
		testQueryCase{
			q: Query{From: "A", To: "B", Via: []string{"C", "D", "E"}},
			p: []string{"from=A", "to=B", "via[]=C", "via[]=D", "via[]=E"},
		},
		testQueryCase{
			q: Query{From: "Bern", To: "Brig", Time: "3:42"},
			p: []string{"from=Bern", "to=Brig", "time=3%3A42"},
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%s", c.q), func(t *testing.T) {
			testQueryURL(t, &c)
		})
	}

}

func testQueryURL(t *testing.T, c *testQueryCase) {
	actual := c.q.URL()
	c.p = append(c.p, APIBase, "?")
	for _, field := range RequiredFields {
		encodedField := strings.Replace(field, "/", "%2F", -1)
		c.p = append(c.p, "fields%5B%5D="+encodedField)
	}

	for i := range c.p {
		c.p[i] = strings.Replace(c.p[i], "[]", "%5B%5D", -1)
	}

	for _, part := range c.p {
		if !strings.Contains(actual, part) {
			t.Errorf("Expected '%s' to contain '%s'", actual, part)
		}
	}
}
