package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

//type channel struct {
//	var name string
//	var list *client
//}

//type client struct {
//	var username string
//	var nickname string
//	var current_chanel channel
//	var ioport net.Conn
//}

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
			var acc string
			for {
				buf := make([]byte, 1)
				if _, err := io.ReadFull(c, buf); err != nil {
					break
				}
				if string(buf) == "\n" {
					ihandler(acc, c)
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

func strstart(str, comp string) (b bool) {
	if str == comp {
		return comp == str[:len(comp) - 1]
	}
	return false
}

func ihandler(inp string, c net.Conn) {
	fmt.Printf("%s\n", inp)
	switch {
	case strings.Contains(inp, "USER"):
		fmt.Printf("USER cmd found\n")
	case strings.Contains(inp, "PASS"):
		fmt.Printf("PASS cmd found\n")
	case strings.Contains(inp, "NICK"):
		fmt.Printf("NICK cmd found\n")
	default:
		fmt.Printf("%s\n", inp)
	}
}
