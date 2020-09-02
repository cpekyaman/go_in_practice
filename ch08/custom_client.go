package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	cc := &http.Client{Timeout: 2 * time.Second}
	req, err := http.NewRequest("GET", "http://goinpracticebook.com", nil)
	if err != nil {
		fmt.Printf("Can't create request : %s\n", err)
		return
	}

	res, err := cc.Do(req)
	if err != nil {
		fmt.Printf("Can't send request: %s\n", err)
		handleNetworkError(err)
		return
	}
	defer res.Body.Close()

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Could not read body: %s\n", err)
		handleNetworkError(err)
		return
	}
	fmt.Printf("%s", buf)
}

func handleNetworkError(err error) {
	switch err := err.(type) {
	case *url.Error:
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			fmt.Printf("Timeout error occured: %s\n", err)
			return
		}
	case net.Error:
		if err.Timeout() {
			fmt.Printf("Timeout error occured: %s\n", err)
			return
		}
	case *net.OpError:
		if err.Timeout() {
			fmt.Printf("Timeout error occured %s\n", err)
			return
		}
	}

	errText := "closed network connection"
	if err != nil && strings.Contains(err.Error(), errText) {
		fmt.Printf("Timeout error occured %s\n", err)
		return
	}
}
