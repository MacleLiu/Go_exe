// 去除[]string slice中相邻的重复字符串
package main

import "fmt"

func disrepet(s []string) []string {
	dist := s[0:1]
	for i := 1; i < len(s); i++ {
		if s[i] != dist[len(dist)-1] {
			dist = append(dist, s[i])
		}
	}
	return dist
}
func main() {
	s := []string{"a", "a", "f", "e", "s", "f", "a", "s", "ss", "ss", "s", "w", "f", "s", "r"}
	s = disrepet(s)
	fmt.Println(s)
}
