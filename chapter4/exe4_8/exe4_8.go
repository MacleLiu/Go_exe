// 统计字母，数字和其他在Unicode分类中字符数量
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	counts := make(map[string]int)
	file, err := os.Open("./test.txt") //打开文本文件test.txt
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	in := bufio.NewReader(file)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if unicode.IsDigit(r) {
			counts["Digit"]++
		} else if unicode.IsLetter(r) {
			counts["Letter"]++
		} else if unicode.IsControl(r) {
			counts["Control"]++
		} else if r == unicode.ReplacementChar && n == 1 {
			counts["Invalid"]++
		} else {
			counts["Others"]++
		}
	}
	for k, v := range counts {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
