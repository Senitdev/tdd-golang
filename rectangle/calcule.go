package rectangle

func CalculeRectangle(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}
func CalculeArea(rect Rectangle) float64 {
	return rect.height * rect.width
}
