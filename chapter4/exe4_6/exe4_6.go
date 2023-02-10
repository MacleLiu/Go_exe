// spaceConver函数将一个UTF-8编码的字节slice中所有相邻的Unicode空白符缩减为一个ASCII空白字符
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func spaceConver(s []byte) []byte {
	var before bool = false //标志前一个字符是否为Unicode空白符
	var i, l int = 0, 0
	for l < len(s) {
		v, size := utf8.DecodeRune(s[i:])
		l += size //记录已读取字节长度，控制循环条件
		if unicode.IsSpace(v) {
			if before {
				//前一个字符是Unicode空白符，则将s[i]之后的字节向前移一字节，覆盖掉当前空白符
				copy(s[i:], s[i+size:])
			} else {
				s[i] = byte(32)
				before = true
				i += size
			}
		} else {
			before = false
			i += size
		}

	}
	return s[0:i]
}

func main() {
	b := []byte("北京\t\n\f欢\f迎\n您")
	fmt.Printf("%s\n%d\n", b, b)
	b = spaceConver(b)
	fmt.Printf("%s\n%d\n", b, b)
}
