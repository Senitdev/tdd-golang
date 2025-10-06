package sums

import "testing"

func TestSums(t *testing.T) {
	numbers := []int{1, 2, 3}
	got := Sums(numbers)
	want := 6
	if got != want {
		t.Errorf("Excepted : got %d , want %d , slice : %v", got, want, numbers)
	}
	t.Run("Collection of 5 number", func(t *testing.T) {
		num := []int{1, 2, 7, 6, 4}
		got := Sums(num)
		want := 20
		if got != want {
			t.Errorf("got %d , want %d , slices %v", got, want, num)
		}
	})
	t.Run("Collection size not fixed", func(t *testing.T) {
		nums := []int{8, 10, 3, 2, 9, 1}
		got := Sums(nums)
		want := 33
		if got != want {
			t.Errorf("got %d , want %d , slices  len not fixed: %v", got, want, nums)
		}

	})

}
