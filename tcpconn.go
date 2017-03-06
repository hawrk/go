package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "usage:%s ip:port", os.Args[0])
		os.Exit(1)
	}
	address := os.Args[1]
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("error to Dial:", err)
		os.Exit(1)
	}

	//
	_, err = conn.Write([]byte("head/http/1.0\n\n"))
	if err != nil {
		fmt.Println("error to write", err.Error())
		os.Exit(1)
	}
}
