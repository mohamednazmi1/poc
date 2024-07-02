package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Listening to port 2000")

	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		fmt.Println("Listen Error:", err)
	}
	defer l.Close()

	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Accept Error:", err)
		}

		go func(c net.Conn) {
			// io.Copy(c, c)

			for {
				tmp := make([]byte, 4096)
				n, err := c.Read(tmp)
				fmt.Println("Length:", n)
				if err != nil {
					fmt.Println(err)
					c.Close()
					return
				}
				str := string(tmp[:])
				fmt.Println("Incoming data:", str)
			}
			// Shut down the connection.
			// c.Close()
		}(conn)
	}
}
