package gotest

import "testing"

type MockMessage struct {
	subject, body, email string
}

func (msg *MockMessage) Send(email string, subject string, body string) error {
	msg.subject = subject
	msg.body = body
	msg.email = email

	return nil
}

func TestAlert(t *testing.T) {
	msgr := new(MockMessage)
	body := "this is garbage"

	Alert(msgr, body)

	if msgr.body != body {
		t.Errorf("Expected body to be %s but got %s\n", body, msgr.body)
	}
}
