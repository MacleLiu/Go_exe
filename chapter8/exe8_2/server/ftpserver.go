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
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		printPrompt(conn)
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	fmt.Println("handleConn")
	input := bufio.NewScanner(c)
	for input.Scan() {
		cmd := parseCmd(input.Text())
		fmt.Println(cmd[0])
		switch cmd[0] {
		case "pwd":
			fmt.Fprintln(c, getwd())
			printPrompt(c)
		case "cd":
			changeDir(cmd[1])
			printPrompt(c)
		case "ls":
			fmt.Fprintln(c, getDir())
			printPrompt(c)
		case "get":
			sendFile(cmd[1], c)
			printPrompt(c)
		case "close":
			c.Close()
		default:
			fmt.Fprintln(c, "Unknown Command!")
			printPrompt(c)
		}
	}
	c.Close()
}

/*
对用户输入的命令进行解析，使用一个slice来保存命令和参数。
便于进行处理。
*/
func parseCmd(cmd string) []string {
	return strings.Split(cmd, " ")
}

// 打印模拟终端的命令行提示符
func printPrompt(w io.Writer) {
	fmt.Fprint(w, getwd()+"> ")
}

func getwd() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return wd
}

func changeDir(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		log.Fatal(err)
	}
}

func getDir() []string {
	fname := []string{}
	finfo, err := os.ReadDir(getwd())
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range finfo {
		fname = append(fname, f.Name())
	}
	return fname
}

func sendFile(path string, conn net.Conn) {
	//只读方式打开文件
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	//读文件内容
	buf := make([]byte, 4*1024)
	for {
		num, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("file.Read err = ", err)
			}

			return
		}

		//发送内容
		conn.Write(buf[:num])
	}
}
