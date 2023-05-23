package main

/* 在64位机器上面，^uint(0)返回的其实就是18446744073709551615,其实也就是2^64次方，
如果实在32位机器上面会返回2^32次方，这个时候先向右移63位，64位机会的到1，而32位会
得到0，这个时候32向左移1或者0（根据^uint(0)得到的结果），如果是1，便会得到64,0会
得到32，这样一来就可以判断是32位或者是64位机器了。 */

const wordSize = 32 << (^uint(0) >> 63)

type IntSet struct {
	words []uint
}

// Has方法的返回值表示是否存在非负数x
func (s *IntSet) Has(x int) bool {
	word, bit := x/wordSize, uint(x%wordSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add添加非负数x到集合中
func (s *IntSet) Add(x int) {
	word, bit := x/wordSize, uint(x%wordSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}
