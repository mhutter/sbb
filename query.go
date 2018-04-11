package sbb

import (
	"net/url"
	"regexp"
	"strings"
)

// Information about the OpenTransport API
const (
	APIBase = "https://transport.opendata.ch/v1/connections"
)

var wordTranslations = map[string][]string{
	"from": []string{"from", "von"},
	"to":   []string{"to", "nach"},
	"via":  []string{"via"},
	"time": []string{"at", "um"},
	"date": []string{"on", "am"},
}

var keywords = make(map[string]string)

func init() {
	// generate `keywords` from `wordTranslations` for lookup
	for kw, terms := range wordTranslations {
		for _, term := range terms {
			keywords[term] = kw
		}
	}
}

// Query represents a timetable enquiry
type Query struct {
	From string
	To   string
	Via  []string
	Time string
	Date string
}

type token struct {
	name   string
	values []string
}

// ParseQuery tries to interpret user input into a query
func ParseQuery(in []string) *Query {
	tokens := tokenize(in)
	query := &Query{}

	for _, t := range tokens {
		value := strings.Join(t.values, " ")
		switch t.name {
		case "from":
			query.From = value
		case "to":
			query.To = value
		case "via":
			query.Via = append(query.Via, value)
		case "time":
			query.Time = normalizeDateTime(value, ":")
		case "date":
			query.Date = normalizeDateTime(value, "-")
		}
	}

	return query
}

var reDateTimeNormalize = regexp.MustCompile(`[^0-9]+`)

// normalizeDateTime replaces all non-numbers by `delim`
func normalizeDateTime(in, delim string) string {
	return reDateTimeNormalize.ReplaceAllString(in, delim)
}

func tokenize(in []string) []token {
	tokens := make([]token, 0)
	var curr token

	for _, word := range in {
		if kw, isKW := getKeyword(word); isKW {
			if curr.name != "" {
				tokens = append(tokens, curr)
			}
			curr = token{name: kw}
			continue
		}

		curr.values = append(curr.values, word)
	}

	return append(tokens, curr)
}

// getKeyword determines if `word` is a keyword and returns
// its canonical form and `true` if so, or `""` and `false`
// otherwise.
func getKeyword(word string) (string, bool) {
	word = strings.ToLower(word)
	if keywords[word] != "" {
		return keywords[word], true
	}
	return "", false
}

// URL returns the complete url for the open transport API
func (q *Query) URL() string {
	params := &url.Values{
		"from":     []string{q.From},
		"to":       []string{q.To},
		"via[]":    q.Via,
		"time":     []string{q.Time},
		"fields[]": RequiredFields,
	}
	return APIBase + "?" + params.Encode()
}
