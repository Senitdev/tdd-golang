package rectangle

import "testing"

func TestRectangle(t *testing.T) {
	t.Helper()
	result := Rectangle{6.0, 2.5}
	got := CalculeRectangle(result)
	want := 17.0
	if got != want {
		t.Errorf("got %f , want %f", got, want)
	}
	t.Run("Test calcule Area", func(t *testing.T) {
		got := CalculeArea(result)
		want := 15.00
		if got != want {
			t.Errorf("got %f , want %f", got, want)
		}
	})
}
