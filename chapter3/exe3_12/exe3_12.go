// isomery函数判断两个字符串是否同文异构
package main

import (
	"fmt"
	"reflect"
)

func isomery(s string, t string) bool {
	counts_s := make(map[string]int)
	counts_t := make(map[string]int)
	if len(s) == len(t) && s != t {
		for _, c := range s {
			counts_s[string(c)]++
		}
		for _, c := range t {
			counts_t[string(c)]++
		}
		return reflect.DeepEqual(counts_s, counts_t)
	}
	//Golang中要比较两个map实例需要使用reflect包的DeepEqual()方法。
	return false
}
func main() {
	fmt.Println(isomery("345623lm天天ccyuhbfhgh", "clmc56y天uhgh43天23bfh"))
	//fmt.Println(isomery("交付3456fg", "交付3456fg"))
}
