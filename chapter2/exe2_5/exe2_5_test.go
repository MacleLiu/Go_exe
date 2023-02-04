package main

import "testing"

func BenchmarkPopcount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Popcount(23)
	}
}
