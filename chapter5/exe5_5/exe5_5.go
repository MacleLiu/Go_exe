// CountWordsAndImages发送一个HTTP GET请求，并获取文档的字数与图片数量
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML:%s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}
func countWordsAndImages(n *html.Node) (words, images int) {
	texts, images := visit(nil, 0, n)
	for _, v := range texts {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		} else {
			scanner := bufio.NewScanner(strings.NewReader(v))
			scanner.Split(bufio.ScanWords)
			for scanner.Scan() {
				words++
			}
		}

	}
	//bare return
	return
}

// 递归循环html
func visit(texts []string, imgs int, n *html.Node) ([]string, int) {
	//文本
	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}
	//图片
	if n.Type == html.ElementNode && (n.Data == "img") {
		imgs++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "script" || c.Data == "style" {
			continue
		}

		texts, imgs = visit(texts, imgs, c)
	}
	//多返回值
	return texts, imgs
}
