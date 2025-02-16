package tree

import (
	ts "github.com/tree-sitter/go-tree-sitter"
)

type Node struct {
	Kind     string `json:"kind"`
	Name     string `json:"name,omitempty"`
	Text     string `json:"text"`
	Children []Node `json:"children,omitempty"`
}

func newNode(c *ts.TreeCursor, code []byte, withChildren bool) Node {
	var children []Node
	n := c.Node()
	if withChildren {
		children = getChildren(c, code)
		for i := range children {
			children[i].Name = n.FieldNameForChild(uint32(i))
		}
	}
	return Node{
		Kind:     n.Kind(),
		Text:     n.Utf8Text(code),
		Children: children,
	}
}

func getChildren(c *ts.TreeCursor, code []byte) []Node {
	var children []Node
	if !c.GotoFirstChild() {
		return children
	}
	n := newNode(c, code, true)
	if n.Kind != n.Text {
		children = append(children, n)
	}
	for c.GotoNextSibling() {
		n = newNode(c, code, true)
		if n.Kind != n.Text {
			children = append(children, n)
		}
	}
	c.GotoParent()
	return children
}
