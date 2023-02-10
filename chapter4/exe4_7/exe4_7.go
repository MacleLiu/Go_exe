// 修改reverse函，翻转一个UTF-8编码的字符串中的字符元素
// 传入参数是该字符串对应的字节slice类型
// 能做到不重新分配内存吗
package main

import (
	"fmt"
	"unicode/utf8"
)

func rev(b []byte) {
	size := len(b)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[size-1-i] = b[size-1-i], b[i]
	}
}
func revUTF8(b []byte) []byte {
	//先翻转每个字符的字节，再整体翻转所有字节
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		rev(b[i : i+size])
		i += size
	}
	rev(b)
	return b
}

func main() {
	s := "北京  欢 迎 您"
	fmt.Println([]byte(s))
	fmt.Println(string(revUTF8([]byte(s))))
	fmt.Println(revUTF8([]byte(s)))
}
