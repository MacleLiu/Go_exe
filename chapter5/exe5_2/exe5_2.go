// htmlEleCount函数统计HTMl文档树中所有的元素个数
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	count := map[string]int{}
	htmlEleCount(count, doc)
}

func htmlEleCount(count map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		count[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		htmlEleCount(count, c)
	}
}
