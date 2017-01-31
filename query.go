package sbb

import (
	"net/url"
	"strings"
)

// Information about the OpenTransport API
const (
	APIBase = "https://transport.opendata.ch/v1/connections"
)

// Query represents a timetable enquiry
type Query struct {
	From string
	To   string
}

// ParseQuery tries to interpret user input into a query
func ParseQuery(in []string) *Query {
	var current string
	labels := map[string][]string{}

	for _, part := range in {
		p := strings.ToLower(part)
		switch p {
		case "from", "to":
			current = p
			continue
		}

		labels[current] = append(labels[current], part)
	}
	return &Query{
		From: strings.Join(labels["from"], " "),
		To:   strings.Join(labels["to"], " "),
	}
}

// URL returns the complete url for the open transport API
func (q *Query) URL() string {
	params := &url.Values{
		"from":     []string{q.From},
		"to":       []string{q.To},
		"fields[]": RequiredFields,
	}
	return APIBase + "?" + params.Encode()
}
