package sum

//calcule somme d'un tableau
func Sum(p [5]int) int {
	result := 0
	for _, i := range p {
		result += i
	}
	return result
}
