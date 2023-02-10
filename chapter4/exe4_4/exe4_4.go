// rotate使用一次遍历完成元素旋转
package main

import "fmt"

func rotate(s []int) (t []int) {
	t = make([]int, len(s))
	for i, v := range s {
		t[(len(s) - 1 - i)] = v
	}
	return
}
func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	s = rotate(s[:])
	fmt.Println(s)
}
