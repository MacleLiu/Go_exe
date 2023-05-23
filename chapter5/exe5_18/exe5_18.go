// Fetch下载URL并返回本地文件的名字和长度
// 使用defer语句关闭打开的可写文件
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	fmt.Println(resp.Request.URL.Path)
	if local == "/" || local == "." {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		//关闭文件，并保留错误消息
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
	n, err = io.Copy(f, resp.Body)
	return local, n, err
}

func main() {
	url := "http://www.baidu.com"
	local, n, err := fetch(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
	}
	fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", url, local, n)
}
