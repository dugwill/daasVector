package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	var jsonData = []byte(`{"name":"morpheus","job":"leader"}}`)

	c := NewClient(500*time.Millisecond, false)

	r, err := NewRequest("http://0.0.0.0:9000", jsonData)
	if err != nil {
		log.Fatalf("Failed creating new request")
	}

	response, error := c.Do(r)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
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
