package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
)

func main() {

	var jsonData = []byte(`{"name": "morpheus","job": "leader"}`)

	c, err := NewSocket("tcp", "127.0.0.1:9000")
	if err != nil {
		log.Fatalln("Failed opening connection")
	}

	fmt.Fprintln(c, bytes.NewBuffer(jsonData))
}

func NewSocket(n, a string) (conn net.Conn, err error) {
	conn, err = net.Dial(n, a)
	if err != nil {
		return nil, err
	}
	return

}
