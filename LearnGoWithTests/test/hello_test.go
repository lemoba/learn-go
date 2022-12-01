package test

import (
	"../hello"
	"testing"
)

func TestHello(t *testing.T) {
	got := hello.Hello("ranen")
	want := "Hello, ranen"

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
