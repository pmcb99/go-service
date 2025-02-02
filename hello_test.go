package main

import "testing"

func TestHelloRecipient(t *testing.T) {
	t.Run("saying hello to paul", func(t *testing.T) {
		got := HelloRecipient("Paul")
		want := "Hello, Paul"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("generic hello if no recip", func(t *testing.T) {
		got := HelloRecipient("")
		want := "Hello stranger..."

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
