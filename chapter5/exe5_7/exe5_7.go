// forEachNode 调用pre(x)和post(x)遍历以n为根的树中的每个节点x
// 两个函数是可选的
// pre在子节点被访问2前(前序)调用
// post在访问后(后序)调用
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode { //元素节点
		attrStr := ""
		for _, a := range n.Attr { //遍历元素属性，将属性和属性值按格式保存到变量attrStr
			attrStr = attrStr + " " + a.Key + "=" + "'" + a.Val + "'"
		}
		if n.FirstChild == nil { //若元素节点没有子节点，在尾部添加"/"，实现简短形式输出
			attrStr = attrStr + "/"
		}
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data+attrStr)
		depth++
	}
	if n.Type == html.TextNode { //文本节点
		n.Data = strings.TrimSpace(n.Data)
		if n.Data != "" {
			text := strings.Split(n.Data, "\n")
			for _, t := range text {
				fmt.Printf("%*s%s\n", depth*2, "", t)
			}
		}
	}
	if n.Type == html.CommentNode { //注释节点
		comment := strings.Split(n.Data, "\n")
		for i, c := range comment {
			if c != "" {
				if i == 0 {
					fmt.Printf("%*s<!--%s\n", depth*2, "", c)
				} else {
					fmt.Printf("%*s%s\n", depth*2, "", c)
				}
			}
		}
		fmt.Printf("%*s-->\n", depth*2, "")
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil { //没有子节点的元素输出简短形式，没有结束标签
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
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
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "outline2: %v\n", err)
			os.Exit(1)
		}
		forEachNode(doc, startElement, endElement)
	}
}
