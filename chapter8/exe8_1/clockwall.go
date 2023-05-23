package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

type clock struct{ name, host string }

func (c clock) watch(w io.Writer, r io.Reader) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		fmt.Fprintf(w, "%s: %s\n", c.name, s.Text())
	}
	fmt.Println(c.name, "done")
	if s.Err() != nil {
		log.Printf("can't read from %s: %s", c.name, s.Err())
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "usage: clockwall TimeZoneNAME=HOST:PORT(Shanghai=localhost:8000) ...")
		os.Exit(1)
	}
	clocks := []clock{}
	for _, v := range os.Args[1:] {
		fields := strings.Split(v, "=")
		if len(fields) != 2 {
			fmt.Fprintf(os.Stderr, "bad arg: %s\n", v)
			os.Exit(1)
		}
		clocks = append(clocks, clock{fields[0], fields[1]})
	}
	for _, c := range clocks {
		conn, err := net.Dial("tcp", c.host)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go c.watch(os.Stdout, conn)
	}
	// Sleep while other goroutines do the work.
	/* for {
		time.Sleep(time.Minute)
	} */
}
