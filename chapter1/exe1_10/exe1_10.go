// fetchall并发获取URL并报告时间和大小
// fetchll将内容输出到文件
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	filename := strings.TrimPrefix(url, "http://") + ".html" //使用URL作为文件名
	file, err := OpenFile(filename)
	if err != nil {
		ch <- fmt.Sprintf("while open file: %v", err)
		return
	}
	nbytes, err := io.Copy(file, resp.Body)
	file.Close()
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

// OpenFile 判断文件是否存在  存在则OpenFile 不存在则Create
func OpenFile(filename string) (*os.File, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("创建文件：%s\n", filename)
		return os.Create(filename) //创建文件
	}
	fmt.Printf("清空并打开文件：%s\n", filename)
	return os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0666)
}
