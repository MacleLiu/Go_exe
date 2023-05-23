package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	elepara, attrpara := parsePara()
	dec := xml.NewDecoder(os.Stdin)
	var stack []string // stack of element names
	var attr []string  //slice of element attributes
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push
			if len(tok.Attr) > 0 {
				for _, b := range tok.Attr {
					s := fmt.Sprintf("%s=%s", b.Name.Local, b.Value)
					attr = append(attr, s)
				}
			}
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
			attr = nil                   //Empty attribute slice
		case xml.CharData:
			if containsAll(stack, elepara) && containsAllAttr(attr, attrpara) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

// containsAllAttr reports whether x contains the elements of y.
func containsAllAttr(x, y []string) bool {
	if len(y) == 0 {
		return true
	}
	eleAttrStr := strings.Join(x, " ")
	for _, a := range y {
		if strings.Contains(eleAttrStr, a) {
			continue
		} else {
			return false
		}
	}
	return true
}

// parsePara parse the input parameters, output element name clice and attribute slice
func parsePara() (ele, attr []string) {
	for _, s := range os.Args[1:] {
		if strings.Contains(s, "=") {
			attr = append(attr, s)
		} else {
			ele = append(ele, s)
		}
	}
	return
}
