package main

import (
	"net/http"

	upb "github.com/cpekyaman/go_in_practice/ch10/proto"
	proto "github.com/golang/protobuf/proto"
)

func main() {
	http.HandleFunc("/", handleProto)
	http.ListenAndServe(":8080", nil)

}

func handleProto(rw http.ResponseWriter, req *http.Request) {
	u := &upb.User{
		Name:  "John Doe",
		Id:    100,
		Email: "john.doe@lost.com",
	}

	body, err := proto.Marshal(u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "x-protobuf")
	rw.Write(body)
}
