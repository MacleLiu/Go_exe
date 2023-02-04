package main

import "testing"

func BenchmarkPopcount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Popcount1(23)
	}
}

func BenchmarkPopcount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Popcount2(23)
	}
}
