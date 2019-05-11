package main

import (
	//"bufio"
	"fmt"
	//"os"
)

func strstart(str, comp string) (b bool) {
	return comp == str[:len(comp)]
}

func main() {
	//reader := bufio.NewReader(os.Stdin)
    //fmt.Print("Enter text: ")
    //text, _ := reader.ReadString('\n')
    //fmt.Println(text)

    var imp string
    fmt.Scan(&imp)

	switch {
	case strstart(imp, "PASS NICK USER"):
	case strstart(imp, "NICK"):
	case strstart(imp, "JOIN"):
	case strstart(imp, "PART"):
	case strstart(imp, "NAMES"):
	case strstart(imp, "LIST"):
	case strstart(imp, "PRIVMSG"):
	default:
	}

    //ln := ""
    //fmt.Sscanln("%v", ln)
    //fmt.Println(ln)
}
