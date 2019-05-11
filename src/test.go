package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// Listen on TCP port 2000 on all available unicast and
	// anycast IP addresses of the local system.
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			// Echo all incoming data.
			//io.Copy(c, c)
			var acc string
			for {
				buf := make([]byte, 1)
				if _, err := io.ReadFull(c, buf); err != nil {
					log.Fatal(err)
				}
				if string(buf) == "\n" {
					fmt.Printf("%s\n", acc)
					acc = ""
				} else {
					acc += string(buf)
				}
				io.WriteString(c, string(buf))
			}
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}
