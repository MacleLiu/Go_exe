// 修改dup2程序，输出出现重复行的文件的名称
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	fileName, seq := "", ""
	if len(files) != 0 {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			if isDuplicate(f) {
				fileName += seq + arg
				seq = " "
			}
			f.Close()
		}
		fmt.Println(fileName)
	} else {
		fmt.Println("未指定文件")
	}

}
func isDuplicate(f *os.File) bool {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			return true
		}
	}
	return false
}
