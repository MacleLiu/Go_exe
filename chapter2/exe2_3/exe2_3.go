// 使用“种群统计”，获取一个数字中被置位的个数
package main

// pc[i]是i的种群统计
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// Popcount返回x的种群统计值(置位的个数)
func Popcount1(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// 使用循环重新Popcount来代替单个表达式
func Popcount2(x uint64) int {
	var n = 0
	for i := 0; i < 8; i++ {
		n = n + int(pc[byte(x>>(i*8))])
	}
	return n
}
