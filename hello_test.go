package main

import "testing"

func TestStringMatching(t *testing.T) {
	assertStringsMatch := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("testing hello function", func(t *testing.T) {
		got := Hello("english")
		want := "hello and welcome!"
		assertStringsMatch(t, got, want)
	})
	t.Run("tesing hello function in french", func(t *testing.T) {
		got := Hello("french")
		want := "bonjour et bienvenue!"
		assertStringsMatch(t, got, want)
	})
	t.Run("testing i made this function", func(t *testing.T) {
		got := IMadeThis()
		want := "i made this http server!"
		assertStringsMatch(t, got, want)
	})
}