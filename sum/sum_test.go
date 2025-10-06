package sum

import "testing"

func TestSum(t *testing.T) {
	numer := [5]int{1, 2, 3, 5, 7}
	got := Sum(numer)
	want := 18
	if got != want {
		t.Errorf("got %d , want %d", got, want)
	}
}
