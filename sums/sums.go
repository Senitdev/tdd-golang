package sums

func Sums(number []int) int {
	sum := 0
	for _, i := range number {
		sum += i
	}
	return sum
}
