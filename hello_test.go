package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}
	t.Run("saying hello to other people", func(t *testing.T) {
		got := Hello("Phet", "")
		want := "Hello Phet"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in Thai", func(t *testing.T) {
		got := Hello("Phet", "Thai")
		want := "สวัสดี Phet"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to the world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello world"
		assertCorrectMessage(t, got, want)
	})
}
