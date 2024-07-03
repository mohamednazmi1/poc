package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("Listening to port 2000")

	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		fmt.Println("Listen Error:", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Accept Error:", err)
		}

		go func(c net.Conn) {
			for {
				tmp := make([]byte, 4096)
				n, err := c.Read(tmp)
				if err != nil {
					fmt.Println(err)
					c.Close()
					return
				}
				data := fmt.Sprintf("%x", tmp[:n])
				fmt.Printf("%s, Data: %s", data, time.Now().UTC().Format("2006-01-02 15:04:05 MST"))
			}
		}(conn)
	}
}
