package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hi", func(t *testing.T) {
		result := Hello("Chris", "English")
		expected := "Hello, Chris"
		assertCorrectMessage(t, expected, result)
	})
	t.Run("say hello world when no string is provided", func(t *testing.T) {
		expected := "Hello, World"
		result := Hello("", "English")
		assertCorrectMessage(t, expected, result)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Sabrina", "Spanish")
		want := "Hola, Sabrina"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Sabrina", "French")
		want := "Bonjour, Sabrina"
		assertCorrectMessage(t, got, want)

	})

}
func assertCorrectMessage(t testing.TB, result, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("got %q want % q", result, expected)
	}
}
