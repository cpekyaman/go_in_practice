package gotest

import "testing"

func TestHello(t *testing.T) {
	s := Hello()
	if s != "Hello" {
		t.Errorf("Expected %s but got %s", "Hello", s)
	}
}
