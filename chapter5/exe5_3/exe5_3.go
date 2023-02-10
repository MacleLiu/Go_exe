// outHtmlTxt函数输出HTMl文档树中所有文本节点的内容，但不包括<script>和<style>
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
	outHtmlTxt(doc)
}

func outHtmlTxt(n *html.Node) {
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	} else if n.Type == html.TextNode {
		fmt.Print(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outHtmlTxt(c)
	}
}
