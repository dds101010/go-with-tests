package main

import "testing"

func TestHelloFunc(t *testing.T) {
	assertEqualString := func(t testing.TB, got, want string) {
		// this method call excludes this method from stack trace in case test fails
		t.Helper()
		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	}

	t.Run("Test without name", func(t *testing.T) {
		want := "Hello, World"
		got := Hello("", "")

		assertEqualString(t, got, want)
	})

	t.Run("Test with name as John", func(t *testing.T) {
		want := "Hello, John"
		got := Hello("John", "")

		assertEqualString(t, got, want)
	})

	t.Run("Test with language as Spanish 1", func(t *testing.T) {
		want := "Hola, John"
		got := Hello("John", "Spanish")

		assertEqualString(t, got, want)
	})

	t.Run("Test with language as Spanish 2", func(t *testing.T) {
		want := "Hola, John"
		got := Hello("John", "Spanish")

		assertEqualString(t, got, want)
	})

	t.Run("Test with language as French", func(t *testing.T) {
		want := "Bonjour, Kate"
		got := Hello("Kate", "French")

		assertEqualString(t, got, want)
	})
}
