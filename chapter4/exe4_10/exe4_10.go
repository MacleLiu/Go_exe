// 修改issuess，安装时间来输出结果，比如一个月以内，一年以内或者超过一年
package main

import (
	"Go_exe/chapter4/github"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		if daysAgo(item.CreatedAt) < 30 { //30天以内
			fmt.Printf("#%-5d  %s %9.9s %.55s\n",
				item.Number, item.CreatedAt, item.User.Login, item.Title)
		}
	}
}
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
