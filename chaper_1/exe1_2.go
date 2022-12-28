// 修改echo程序，输出参数的索引和值，每行一个。
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, val := range os.Args[1:] {
		fmt.Printf("Index: %d  Value: %s\n", i, val)
	}
}
