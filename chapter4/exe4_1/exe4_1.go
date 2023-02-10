// 使用种群统计方法，统计SHA256散列中不同的位数
package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		//fmt.Println(i, pc[i], pc[i/2], byte(i&1))
	}
}
func SHA256_PopCount(x [32]byte) int {
	var r int
	for i := 0; i < 32; i++ {
		r = r + int(pc[byte(int(x[i])>>(0*8))]+
			pc[byte(int(x[i])>>(1*8))]+
			pc[byte(int(x[i])>>(2*8))]+
			pc[byte(int(x[i])>>(3*8))]+
			pc[byte(int(x[i])>>(4*8))]+
			pc[byte(int(x[i])>>(5*8))]+
			pc[byte(int(x[i])>>(6*8))]+
			pc[byte(int(x[i])>>(7*8))])
	}
	return r
}
func main() {
	c := sha256.Sum256([]byte("x"))
	//fmt.Printf("%b\n%x\n", c, c)
	fmt.Printf("置1的位个数：%d\n", SHA256_PopCount(c))
}
