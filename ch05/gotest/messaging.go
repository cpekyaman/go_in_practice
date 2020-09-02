package gotest

import "fmt"

type Message struct {
	Text string
}

func (msg *Message) Send(email string, subject string, body string) error {
	fmt.Printf("Sending %s to %s wıth subject %s\n", body, email, subject)
	return nil
}
