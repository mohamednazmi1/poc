package main

import (
	"encoding/json"
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
				// fmt.Println("Incoming data:", str)
				var jsonData map[string]interface{}
				if err := json.Unmarshal(tmp[:n], &jsonData); err == nil {
					// Successfully decoded JSON data
					fmt.Println("Decoded JSON data:", jsonData)
				} else {
					// Print the raw data if it's not JSON
					fmt.Println("Raw data:", str)
				}
			}
			// Shut down the connection.
			// c.Close()
		}(conn)
	}
}
