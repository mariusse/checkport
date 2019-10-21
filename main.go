package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const timeout = time.Second * 5

func main() {
	URIs := os.Args[1:]
	if len(URIs) == 0 {
		fmt.Println("Usage: checkport <host:port>...")
		os.Exit(0)
	}

	fmt.Println("Timeout is set to ", timeout, " seconds.")
	for _, v := range URIs {
		if isPortOpen(v) {
			fmt.Printf("OK %s\n", v)
			continue
		}
		fmt.Printf("X %s\n", v)
	}
}

func isPortOpen(uri string) bool {
	_, err := net.DialTimeout("tcp", uri, timeout)
	if err == nil {
		return true
	}
	return false
}
