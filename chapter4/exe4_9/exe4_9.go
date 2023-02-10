// wordfreq统计输入文本文件中每个单词出现的次数
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		counts[scanner.Text()]++
	}
	for k, v := range counts {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
