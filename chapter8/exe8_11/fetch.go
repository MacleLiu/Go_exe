// 这是一个fetch变种，它并发请求多个url
// 当第一个响应返回时，取消其他请求
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var reqCancel = make(chan struct{})

func main() {
	responses := make(chan string, 3)
	go func() { responses <- request("www.baidu.com") }()
	go func() { responses <- request("www.jd.com") }()
	go func() { responses <- request("www.bilibili.com") }()
	if s, ok := <-responses; ok {
		close(reqCancel)
		fmt.Print(s)
	}

}

func request(hostname string) (responses string) {
	req, err := http.NewRequest("GET", "http://"+hostname, nil)
	if err != nil {
		log.Print(err)
		return ""
	}
	req.Cancel = reqCancel

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
		return ""
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Print(hostname, ": ", resp.Status)
		return ""
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}
