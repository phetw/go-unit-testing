package main

import "testing"

func TestAdd(t *testing.T) {
	got := Add(2, 5)
	want := 7

	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}
