package main

import "testing"

/*
	- It needs to be in a file with a name like xxx_test.go
	- The test function must start with the word Test
	- The test function takes one argument only t *testing.T
*/

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
