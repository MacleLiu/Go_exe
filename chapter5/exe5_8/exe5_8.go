// forEachNode 调用pre(x)和post(x)遍历以n为根的树中的每个节点x
// 两个函数是可选的
// pre在子节点被访问2前(前序)调用
// post在访问后(后序)调用
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode { //元素节点
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id { //id属性，且值为目标id字符串，返回true
				return true
			}
		}
	}
	return false
}

func forEachNode(n *html.Node, id string, pre func(n *html.Node, id string) bool) *html.Node {
	var ok bool
	var element *html.Node
	if pre != nil {
		if ok = pre(n, id); ok { //返回值为真表示匹配成功，返回当前节点
			return n
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if element != nil { //element不为空表示已经匹配到目标节点，结束遍历
			break
		}
		element = forEachNode(c, id, pre)
	}
	return element
}

func ElementById(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, startElement)
}

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Parse: %v\n", err)
			os.Exit(1)
		}
		r := ElementById(doc, "ceshi")
		fmt.Println(r.Data)
	}
}
