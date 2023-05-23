// IsPalindrome函数判断一个序列是否是回文
// 假定对于下标分别为i、j的元素，如果!s.Less(i,j)&&!s.Less(j,i)，那么两个元素相等
package main

import (
	"fmt"
	"sort"
)

type ByteSlice []byte

func (s ByteSlice) Len() int           { return len(s) }
func (s ByteSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s ByteSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if !s.Less(i, j) && !s.Less(j, i) {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	name := ByteSlice([]byte("aaccffccaa"))
	fmt.Println(IsPalindrome(name))
}
