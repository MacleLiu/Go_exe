package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:8000")
	if err != nil {
		log.Fatal()
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	//conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal()
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	//go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
	conn.CloseWrite()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
