package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people in 'Spanish'", func(t *testing.T) {
		got := Hello("Eduardo", "Spanish")
		want := "Hola, Eduardo"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in 'English'", func(t *testing.T) {
		got := Hello("Eduardo", "English")
		want := "Hello, Eduardo"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people in 'French'", func(t *testing.T) {
		got := Hello("Eduardo", "French")
		want := "Bonjour, Eduardo"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
