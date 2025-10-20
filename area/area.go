package area

import "math"

type Rectangles struct {
	width  float64
	height float64
}

func CalRec(r Rectangles) float64 {
	return 2 * (r.height + r.width)
}
func CalculeArea(r Rectangles) float64 {
	return r.height * r.width
}

type Circle struct {
	raduis float64
}

func CalculeCercle(r Circle) float64 {
	return math.Pi * r.raduis * r.raduis
}
