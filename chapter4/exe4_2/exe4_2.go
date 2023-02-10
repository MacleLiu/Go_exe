// 输出输入值的SHA散列值
// 默认输出SHA256散列值，可使用命令行参数指定输出SHA384或SHA512散列
// eg: exe4_2.exe -SHA384 23
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	var s, t string
	if len(os.Args) == 1 || len(os.Args) > 3 {
		fmt.Println("参数错误")
		os.Exit(0)
	} else if len(os.Args) == 2 {
		s = os.Args[1]
	} else {
		if os.Args[1] != "-SHA384" && os.Args[1] != "-SHA512" && os.Args[1] != "-SHA256" {
			fmt.Println("参数错误")
			os.Exit(0)
		}
		t = os.Args[1]
		s = os.Args[2]
	}
	switch t {
	case "-SHA384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(s)))
	case "-SHA512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(s)))
	default:
		fmt.Printf("%x\n", sha256.Sum256([]byte(s)))
	}
}
