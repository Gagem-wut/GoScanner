package main

import (
	"github.com/fatih/color"
	"net"
	"strconv"
)

func main() {
	for i := 1; i <= 65535; i++ {
		var addrs = "127.0.0.1:" + strconv.Itoa(i)
		_, err := net.Dial("tcp", addrs)
		if err == nil {
			color.Green("Port: " + strconv.Itoa(i) + " Is open!")
		}
	}
}
