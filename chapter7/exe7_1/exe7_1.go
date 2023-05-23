// 使用类似ByteCounter的想法，实现单词和行计数器。
// 实现时考虑使用bufio.ScanWords
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	count := 0
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count++
	}
	*c += WordCounter(count)
	return count, nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	count := 0
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		count++
	}
	*c += LineCounter(count)
	return count, nil
}

func main() {
	var c WordCounter
	c.Write([]byte("I want to travel around the world in future. "))
	fmt.Println(c)
	c.Write([]byte("I want to travel around the world in future. "))
	fmt.Println(c)

	var l LineCounter
	l.Write([]byte("I want to travel around the world in future. \nI think this is a good idear! "))
	fmt.Println(l)
	l.Write([]byte("I want to travel around the world in future. \nI think this is a good idear! "))
	fmt.Println(l)
}
