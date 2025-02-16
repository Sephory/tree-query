package tree

import (
	ts "github.com/tree-sitter/go-tree-sitter"
)

type Match struct {
	Captures []Capture `json:"captures"`
}

func (m Match) ToMap() map[string][]string {
	captures := make(map[string][]string, len(m.Captures))
	for _, c := range m.Captures {
		if captures[c.Name] == nil {
			captures[c.Name] = []string{c.Node.Text}
		} else {
			captures[c.Name] = append(captures[c.Name], c.Node.Text)
		}
	}
	return captures
}

type Capture struct {
	Name string `json:"name"`
	Node Node   `json:"node"`
}

func QueryTree(code []byte, language Language, query string) ([]Match, error) {
	var matches []Match
	t := loadTree(code, language)
	q, err := ts.NewQuery(getLanguage(language), query)
	if err != nil {
		return nil, err
	}
	cursor := ts.NewQueryCursor()
	defer cursor.Close()
	tsMatches := cursor.Matches(q, t.RootNode(), code)
	m := tsMatches.Next()
	for m != nil {
		matches = append(matches, makeMatch(m, code, q.CaptureNames()))
		m = tsMatches.Next()
	}
	return matches, nil
}

func makeMatch(m *ts.QueryMatch, code []byte, captureNames []string) Match {
	var captures []Capture
	for _, c := range m.Captures {
		captures = append(captures, Capture{
			Name: captureNames[c.Index],
			Node: newNode(c.Node.Walk(), code, true),
		})
	}
	return Match{
		Captures: captures,
	}
}
