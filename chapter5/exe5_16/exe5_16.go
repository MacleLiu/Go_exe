// 变长版本的strings.Join函数
package main

import "fmt"

func Join(sep string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	var str string
	for _, s := range strs[:len(strs)-1] {
		str = str + s + sep
	}
	str = str + strs[len(strs)-1]
	return str
}

func main() {
	s := []string{"this", "is", "a", "method"}
	fmt.Println(Join(","))
	fmt.Println(Join(",", "abc"))
	fmt.Println(Join(",", s...))
}
