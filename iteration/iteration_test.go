package iteration

import "testing"

func TestRepeatFunc(t *testing.T) {
	t.Run("Test Repeat function 1", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})

	t.Run("Test Repeat function 2", func(t *testing.T) {
		got := Repeat("a", 1)
		want := "a"

		if got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})
}
