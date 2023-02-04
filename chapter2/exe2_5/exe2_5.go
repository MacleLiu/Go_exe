// 利用x&(x-1)可以清除最右边的非零位
package main

func Popcount(x uint64) int {
	var n int
	for x != 0 {
		n++
		x = x & (x - 1)
	}
	return n
}
