package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

// 广播器
type client chan<- string

var (
	entering   = make(chan client)
	leaving    = make(chan client)
	messages   = make(chan string)
	clientname = make(chan string)
)

func broadcaster() {
	clients := make(map[client]string)
	for {
		select {
		case msg := <-messages:
			//把所有接受的消息广播给所有的客户
			//发送消息通道
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			for _, name := range clients {
				cli <- name + " online"
			}
			clients[cli] = <-clientname
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	send := make(chan struct{})
	exit := make(chan struct{})
	go clienWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch
	clientname <- who
	go func() {
		for {
			select {
			case <-time.After(10 * time.Second):
				leaving <- ch
				messages <- who + " timeout offline"
				conn.Close()
				return
			case <-exit:
				leaving <- ch
				messages <- who + " has left"
				conn.Close()
				return
			case <-send:

			}
		}
	}()
	input := bufio.NewScanner(conn)
	for input.Scan() {
		send <- struct{}{}
		messages <- who + ": " + input.Text()
	}
	exit <- struct{}{}
	//注意，忽略inoput.Err()中可能的错误

	/* leaving <- ch
	messages <- who + " has left"
	conn.Close() */
}

func clienWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
