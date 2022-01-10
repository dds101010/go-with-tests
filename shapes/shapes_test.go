package shapes

import "testing"

// In Go interface resolution is implicit. If the type you pass in matches
// what the interface is asking for, it will compile.
func checkArea(t testing.TB, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()
	if got != want {
		// Use of g will print a more precise decimal number in the error message
		// The %#v format string will print out our struct with the values in its field
		t.Errorf("%#v got %g want %g", shape, got, want)
	}
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		// The f is for our float64 and the .2 means print 2 decimal places
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	// anonymous struct
	tests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, tt := range tests {
		checkArea(t, tt.shape, tt.want)
	}
}
