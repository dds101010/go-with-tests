package integers

import "testing"

func TestAdderFunc(t *testing.T) {
	t.Run("Regular Sum", func(t *testing.T) {
		want := 42
		got := Add(40, 2)

		if want != got {
			t.Errorf("got: %d, want: %d", got, want)
		}
	})
}
