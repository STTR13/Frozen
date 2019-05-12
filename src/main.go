package main

const password string = nil

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

type client struct {
	var username string
	var hostname string
	var nickname string
	var current_chanel channel
	var ioport net.Conn
}

func main() {
	// Listen on TCP port 2000 on all available unicast and
	// anycast IP addresses of the local system.
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	client_list := make(chan []client, 0)
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
			var pw bool
			for {
				buf := make([]byte, 1)
				if _, err := io.ReadFull(c, buf); err != nil {
					break
				}
				if string(buf) == "\n" {
					if !pw { pw = pass(acc, c) }
					if pw { ihandler(acc, c) }
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

//func (cll chan []client) add(str string, c net.Conn)

func pass(str string, c net.Conn) bool {
	if strings.Contains(inp, "PASS") {
		t := strings.Split(inp, ' ')
		if len(t) == 2 && t[1][0:1] == ':' && t[1][1:] == password {
			return true
		} else {
			io.WriteString(c, "ERR_NEEDMOREPARAMS")
		}
	} else {
		if password == nil { return true }
		else { return false }
	}
	return false
}

func strstart(str, comp string) (b bool) {
	if len(str) >= len(comp) {
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
