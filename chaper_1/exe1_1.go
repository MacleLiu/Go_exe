// 修改echo程序输出os.Args[0]，即命令的名字。
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1])
}
