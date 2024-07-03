package main

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"strings"
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
				decodeAndPrintData(tmp)

				fmt.Println("Length:", n)
				if err != nil {
					fmt.Println(err)
					c.Close()
					return
				}
				str := string(tmp[:])
				fmt.Println("Incoming data:", str)

				var jsonData map[string]interface{}
				if err := json.Unmarshal(tmp[:n], &jsonData); err == nil {
					// Successfully decoded JSON data
					fmt.Println("Decoded JSON data:", jsonData)
				} else {
					// Print the raw data if it's not JSON
					fmt.Println("Raw data:", str)
				}

				dst := make([]byte, hex.DecodedLen(len(str)))
				no, error1 := hex.Decode(dst, tmp)
				fmt.Printf("%s\n", dst[:no])
				if error1 != nil {
					fmt.Println("Hex Decoding Error:", error1.Error())
				}
			}
			// Shut down the connection.
			// c.Close()
		}(conn)
	}
}

func decodeAndPrintData(data []byte) {
	str := string(data)

	b := fmt.Sprintf("%x", str)
	fmt.Printf("Decoded bytes %v", []byte(b))

	fmt.Println("Incoming data:", str)
	fmt.Printf("%x", str)

	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err == nil {
		fmt.Println("Decoded JSON data:", jsonData)
	} else {
		fmt.Println("Failed to decode JSON:", err)
	}

	if decodedHex, err := hex.DecodeString(strings.TrimSpace(str)); err == nil {
		fmt.Println("Decoded hex data:", string(decodedHex))
	} else {
		fmt.Println("Failed to decode hex:", err)
	}

	if decodedBase64, err := base64.StdEncoding.DecodeString(strings.TrimSpace(str)); err == nil {
		fmt.Println("Decoded base64 data:", string(decodedBase64))
	} else {
		fmt.Println("Failed to decode base64:", err)
	}

	fmt.Println("Raw data:", str)
}
