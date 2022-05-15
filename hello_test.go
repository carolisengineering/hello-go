package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "hello and welcome!" 

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestIMadeThis(t *testing.T) {
	got := IMadeThis()
	want := "i made this http server!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}