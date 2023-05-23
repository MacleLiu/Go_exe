// 已知一个HTML节点树和零个或多个名字
// ElementsByTagName函数返回所有符合给出名字的函数
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	if len(name) == 0 {
		return nil
	}
	var nodes []*html.Node
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode {
			if InSlice(name, n.Data) {
				nodes = append(nodes, n)
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return nodes
}

// InSlice函数判断字符串s是否包含在字符串切片strs中
func InSlice(strs []string, s string) bool {
	for _, str := range strs {
		if str == s {
			return true
		}
	}
	return false
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func main() {
	url := "http://localhost:8000"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Get: %s Error: %v", url, err)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("parsing %s as HTML: %v", url, err)
	}
	nodes := ElementsByTagName(doc, "h1", "p", "h3")
	for _, node := range nodes {
		fmt.Println(node.Data)
	}
}
