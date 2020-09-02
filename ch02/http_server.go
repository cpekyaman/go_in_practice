package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/braintree/manners"
)

func main() {
	sch := make(chan os.Signal)
	signal.Notify(sch, os.Kill, os.Interrupt)
	go shutdownListener(sch)

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/bye", bye)
	http.HandleFunc("/", home)

	log.Fatal(manners.ListenAndServe(":9000", nil))
}

func shutdownListener(s <-chan os.Signal) {
	<-s
	manners.Close()
}

func home(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "Welcome %s\n")
}

func hello(resp http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	name := q.Get("name")
	if name == "" {
		name = "John Doe"
	}
	fmt.Fprintf(resp, "Hello %s\n", name)
}

func bye(resp http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	name := q.Get("name")
	if name == "" {
		name = "John Doe"
	}
	fmt.Fprintf(resp, "Bye %s\n", name)
}
