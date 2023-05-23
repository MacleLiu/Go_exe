// 使用panic和recover写一个函数，它没有return语句，但是能够返回一个非零的值
package main

import "fmt"

func hello() (r string) {
	defer func() {
		s := recover()
		r = fmt.Sprintf("%v", s)
	}()
	panic("hello")
}

func main() {
	fmt.Println(hello())
}
