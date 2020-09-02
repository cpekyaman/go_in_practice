package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	pb "github.com/cpekyaman/go_in_practice/ch10/proto"
	"github.com/golang/protobuf/proto"
)

func main() {
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Println("Could not get response")
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Could not read body: %s\n", err)
		return
	}
	var usr pb.User
	err = proto.Unmarshal(buf, &usr)
	if err != nil {
		fmt.Printf("Could not unmarshal user: %s\n", err)
		return
	}
	fmt.Println(usr.Name)
	fmt.Println(usr.Id)
	fmt.Println(usr.Email)
}
