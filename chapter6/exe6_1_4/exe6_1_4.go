// Intset是一个包含非负整数的集合
// 零值代表空的集合
package main

import (
	"Go_exe/chapter6"
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

// Has方法的返回值表示是否存在非负数x
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add添加非负数x到集合中
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith将会对s和t做并集并将结果存在s中
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String方法以字符串"{1 2 3}"的形式返回集中
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len返回元素个数
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		count += chapter6.Popcount(word)
	}
	return count
}

// 从集合去除元素x
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] &^= 1 << bit
}

// 删除所有元素
func (s *IntSet) Clear() {
	for word := range s.words {
		s.words[word] = 0
	}
}

// 返回集合的副本
func (s *IntSet) Copy() *IntSet {
	newSet := &IntSet{}
	newSet.words = make([]uint64, len(s.words))
	copy(newSet.words, s.words)
	return newSet
}

// 变长方法AddAll允许接受一串整型值作为参数
func (s *IntSet) AddAll(nums ...int) {
	for _, n := range nums {
		s.Add(n)
	}
}

// IntersectWith将会对s和t做交集并将结果存在s中
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			break
		}
	}
	if len(t.words) < len(s.words) {
		copy(s.words, s.words[:len(t.words)])
	}
}

// DifferWith求s对于t的差集并将结果存在s中
func (s *IntSet) DifferWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			break
		}
	}
	if len(t.words) < len(s.words) {
		copy(s.words, s.words[:len(t.words)])
	}
}

// SymmetryDifferWith求s和t的对称差并将结果存在s中
func (s *IntSet) SymmetryDifferWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Elems返回包含集合元素的slice
func (s *IntSet) Elems() []int {
	e := make([]int, 0)
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				e = append(e, 64*i+j)
			}
		}
	}
	return e
}

func main() {
	var x, y IntSet
	x.AddAll(2, 30, 5, 61, 1, 78, 45, 6)
	y.AddAll(34, 6, 9, 2, 6, 9, 434, 78)
	//x.SymmetryDifferWith(&y)
	fmt.Println(x.Elems())
}
