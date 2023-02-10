// 函数expand替换参数s中每一个字符串"$foo"为f("foo")的返回值
package main

import (
	"fmt"
	"strings"
)

func f(s string) string {
	return s
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("foo"), -1)
}
func main() {
	s := "2jnkc8$foolkand0090mnd$foo021ji$foo$foofn$foo"
	fmt.Println(expand(s, f))
}
