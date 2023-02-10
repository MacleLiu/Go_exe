package main

import "fmt"

func reverse(s *[5]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
func main() {
	var a [5]int = [5]int{1, 2, 45, 7, 5}
	reverse(&a)
	fmt.Println(a)
}
