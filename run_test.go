package main

import "testing"

/*
	- It needs to be in a file with a name like xxx_test.go
	- The test function must start with the word Test
	- The test function takes one argument only t *testing.T
*/

func TestHello(t *testing.T) {
	// testing.TB - interface that *testing.T and *testing.B both satisfy
	assertCorrectMessage := func(t testing.TB, got, want string) {
		// Для корректной нумерации строк
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elena", "French")
		want := "Bonjour, Elena"
		assertCorrectMessage(t, got, want)
	})
}
