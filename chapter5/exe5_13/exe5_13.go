// 根据需要创建目录保存来自相同域名下的页面
// 当前程序实际是根据Host进行判断，而非Domain
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"Go_exe/chapter5/links"
)

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

var origHost string

// save函数判断当前页面是否是origHost下的页面，若是则保存
func save(rawurl string) error {
	url, err := url.Parse(rawurl)
	if err != nil {
		return fmt.Errorf("bad url: %s", err)
	}
	if origHost == "" {
		origHost = url.Host
	}
	if origHost != url.Host {
		return nil
	}
	dir := url.Host
	var filename string
	if filepath.Ext(url.Path) == "" {
		dir = filepath.Join(dir, url.Path)
		filename = filepath.Join(dir, "index.html")
	} else {
		dir = filepath.Join(dir, filepath.Dir(url.Path))
		filename = filepath.Join(dir, filepath.Base(url.Path))
	}
	if !pathExist(filename) {
		if !pathExist(dir) {
			err = os.MkdirAll(dir, 0777)
			if err != nil {
				return err
			}
		}
		resp, err := http.Get(rawurl)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return err
		}
		// Check for delayed write errors, as mentioned at the end of section 5.8.
		err = file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func crawl(url string) []string {
	fmt.Println(url)
	err := save(url)
	if err != nil {
		log.Printf(`can't cache "%s": %s`, url, err)
	}
	list, err := links.Extract(url)
	if err != nil {
		log.Printf(`can't extract links from "%s": %s`, url, err)
	}
	return list
}

// pathExist函数判断文件或文件夹是否存在
func pathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}
