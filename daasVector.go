package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {

	var jsonData = []byte(`{
		"name": "morpheus",
		"job": "leader"
	}`)

	c, err := NewSocket("tcp", "127.0.0.1:9000")
	if err != nil {
		log.Fatalln("Failed opening connection")
	}

	fmt.Fprintln(c, bytes.NewBuffer(jsonData))

	// wait for reply
	//message, _ := bufio.NewReader(c).ReadString('\n')
	//fmt.Print("Message from server: " + message)

}

// NewClient crates a new http client with the provided timeout.
// If redirects is true, the request will follow any redirects in the response
func NewClient(t time.Duration, redirects bool) (client *http.Client) {
	if redirects {
		client = &http.Client{
			Timeout: t * time.Millisecond,
		}
	} else {
		client = &http.Client{
			Timeout: t * time.Millisecond,
			// Do not Follow redirect
			CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
		}
	}
	return client
}

func NewRequest(url string, d []byte) (*http.Request, error) {
	// Create Request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(d))
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "scte35Mon/1.0")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return req, err
}

func NewSocket(n, a string) (conn net.Conn, err error) {
	conn, err = net.Dial(n, a)
	if err != nil {
		return nil, err
	}
	return

}
