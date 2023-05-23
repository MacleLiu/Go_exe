package main

import "io"

type LimitedReader struct {
	r io.Reader
	n int64
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {
	if l.n <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.n {
		p = p[:l.n]
	}
	n, err = l.r.Read(p)
	l.n -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader { return &LimitedReader{r, n} }
