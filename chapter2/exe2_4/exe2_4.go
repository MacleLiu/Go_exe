// 通过移位操作，每次判断最右边的位是否为1
package main

func Popcount(x uint64) int {
	var n int
	for i := 0; i < 64; i++ {
		if (x & 1) == 1 {
			n++
		}
		x = x >> 1
	}
	return n
}
