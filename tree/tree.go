package tree

import (
	"path/filepath"
	"unsafe"

	ts "github.com/tree-sitter/go-tree-sitter"
	ts_go "github.com/tree-sitter/tree-sitter-go/bindings/go"
	ts_php "github.com/tree-sitter/tree-sitter-php/bindings/go"
)

type Language int

const (
	Go Language = iota
	PHP
)

type Tree struct {
	Root Node `json:"root"`
}

func makeTree(t *ts.Tree, code []byte) Tree {
	cursor := t.RootNode().Walk()
	defer cursor.Close()
	return Tree{
		Root: newNode(cursor, code, true),
	}
}

func LoadTree(code []byte, language Language) Tree {
	return makeTree(loadTree(code, language), code)
}

func loadTree(code []byte, language Language) *ts.Tree {
	parser := ts.NewParser()
	defer parser.Close()
	parser.SetLanguage(getLanguage(language))
	return parser.Parse(code, nil)
}

func GetLanguageForFile(filename string) Language {
	var language Language
	switch filepath.Ext(filename) {
	case ".go":
		language = Go
	case ".php":
		language = PHP
	}
	return language
}

func getLanguage(language Language) *ts.Language {
	var ptr unsafe.Pointer
	switch language {
	case PHP:
		ptr = ts_php.LanguagePHP()
	case Go:
		ptr = ts_go.Language()
	}
	if ptr == nil {
		panic("Invalid Language")
	}
	return ts.NewLanguage(ptr)
}
