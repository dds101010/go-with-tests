package main

import "testing"

func TestHelloFunc (t *testing.T) {
	assertEqualString := func (t testing.TB, got, want string) {
		// this method call excludes this method from stack trace in case test fails
		t.Helper()
		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	}

	t.Run("Test without name", func (t *testing.T) {
		want := "Hello, World"
		got := Hello("")

		assertEqualString(t, got, want)
	})
	
	t.Run("Test with name as John", func (t *testing.T) {
		want := "Hello, World"
		got := Hello("John")
		
		assertEqualString(t, got, want)
	})
}