package test

import (
	"../hello"
	"testing"
)

func TestHello_1(t *testing.T) {
	t.Run("saying hello people", func(t *testing.T) {
		got := hello.Hello("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := hello.Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
