// 使用下面命令运行测试
// go test -bench=Echo2
// go test -bench=Echo3
package main

import (
	"fmt"
	"strings"
	"testing"
)

func Echo2(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Sprintln(s)
}

func Echo3(agrs []string) {
	fmt.Sprintln(strings.Join(agrs[:], " "))
}

var para []string = []string{"a", "ab", "abc", "abcd", "abcde", "saad", "sdasdaasf", "sad", "wqwx", "sadv", "sfxczv", "saf", "Sdasf", "sdfas", "swf"}

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo2(para)
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo3(para)
	}
}
