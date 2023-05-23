package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Node any //CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func parse(r io.Reader) (Node, error) {
	dec := xml.NewDecoder(r)
	var stack []*Element
	var root Node
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			ele := &Element{tok.Name, tok.Attr, nil}
			if len(stack) == 0 {
				root = ele
			} else {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, ele)
			}
			stack = append(stack, ele) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if len(stack) != 0 {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, CharData(tok))
			}
		}
	}
	return root, nil
}

func (n *Element) String() string {
	b := &bytes.Buffer{}
	visit(n, b, 0)
	return b.String()
}

func visit(n Node, w io.Writer, depth int) {
	switch n := n.(type) {
	case *Element:
		var attrs string
		if len(n.Attr) != 0 {
			for _, b := range n.Attr {
				attrs += fmt.Sprintf("%s=%q ", b.Name.Local, b.Value)
			}
		}
		fmt.Fprintf(w, "%*s%s %s\n", depth*2, "", n.Type.Local, attrs)
		for _, c := range n.Children {
			visit(c, w, depth+1)
		}
	case CharData:
		if len(strings.TrimSpace(string(n))) != 0 {
			fmt.Fprintf(w, "%*s%q\n", depth*2, "", n)
		}
	default:
		panic(fmt.Sprintf("got %T", n))
	}
}

func main() {
	node, err := parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	fmt.Println(node)
}
