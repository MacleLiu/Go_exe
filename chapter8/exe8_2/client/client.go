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

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	//mustCopy(os.Stdout, conn)
	fmt.Fprint(os.Stdout, conn)

	for {
		cmd := parseCmd(os.Stdin)
		switch cmd[0] {
		case "ls":
			fmt.Println("I am ls!")
			fmt.Fprint(conn, "ls")
			fmt.Fprint(os.Stdout, conn)
		case "get":
			fmt.Fprintln(conn, "get")
			receiveFile(cmd[1], conn)
		}
	}
	/* switch cmd[0] {
	case "ls":
		fmt.Println("I am ls!")
		conn.Write([]byte("ls"))
	case "get":
		fmt.Println("I am get!")
		conn.Write([]byte(strings.Join(cmd, " ")))
		receiveFile("test.txt", conn)
	default:
		fmt.Println("default")
	} */

	//go mustCopy(os.Stdout, conn)
	//go receiveFile("test.txt", conn)

	//mustCopy(conn, os.Stdin)
}

func parseCmd(input io.Reader) []string {
	in := bufio.NewReader(input)
	s, _ := in.ReadString('\n')
	cmd := strings.Fields(s)
	return cmd
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func receiveFile(filename string, conn net.Conn) {
	//新建文件，file相当于文件句柄(可读可写)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("os.Create err = ", err)
		return
	}
	defer file.Close()

	buf := make([]byte, 4*1024)
	for {
		num, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完毕")
				return
			} else {
				fmt.Println("conn.Read err = ", err)
				return
			}
		}

		//写入内容
		file.Write(buf[:num])
	}
}
