package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		if n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}
		if n.Data == "img" {
			for _, img := range n.Attr {
				if img.Key == "src" {
					links = append(links, img.Val)
				}
			}
		}
		if n.Data == "script" {
			for _, script := range n.Attr {
				if script.Key == "src" {
					links = append(links, script.Val)
				}
			}

		}
		if n.Data == "link" {
			for _, link := range n.Attr {
				if link.Key == "href" {
					links = append(links, link.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v/n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
