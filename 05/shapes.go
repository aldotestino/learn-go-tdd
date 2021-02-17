package shapes

import "math"

// Shape is an interface
type Shape interface {
	Area() float64
}

// Rectangle define all the properties of a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Area is a method to calculate the area
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle define all the properties of a circle
type Circle struct {
	Radius float64
}

// Area is a method to calculate the area
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Triangle define all the properties of a circle
type Triangle struct {
	Base   float64
	Height float64
}

// Area is a method to calculate the area
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

// Perimeter returns the perimeter of a rectangle
func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Width + rect.Height)
}
