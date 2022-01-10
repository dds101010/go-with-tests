package shapes

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (r Circle) Area() float64 {
	return math.Pi * r.Radius * r.Radius
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

// func Area(r Rectangle) float64 {
// 	return r.Width * r.Height
// }
