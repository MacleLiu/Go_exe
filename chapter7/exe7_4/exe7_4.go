package main

import (
	"fmt"
	"io"
)

type StringReader struct {
	s string
	i int
}

func (r *StringReader) Read(b []byte) (n int, err error) {
	if r.i >= len(r.s) {
		return 0, io.EOF
	}
	n = copy(b, r.s[r.i:])
	r.i += n
	return
}

func NewReader(s string) StringReader { return StringReader{s, 0} }

func main() {
	str := "I love you."
	reader := NewReader(str)
	r := make([]byte, 5)
	reader.Read(r)
	/* n, err := reader.Read(r)
	for err == nil {
		fmt.Println(n, string(r[0:n]))
		n, err = reader.Read(r)
	} */
	fmt.Println(string(r))
	fmt.Println(reader.i)
	fmt.Println(reader.s)
}
