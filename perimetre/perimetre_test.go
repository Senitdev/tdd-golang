package perimetre

import "testing"

func TestPerimetre(t *testing.T) {
	got := Perimetre(3, 7)
	want := 20.0
	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
	t.Run("Test périmetre avec valeur décimal", func(t *testing.T) {
		got := Perimetre(7.5, 8.4)
		want := 31.8
		if got != want {
			t.Errorf("got %f , want %f", got, want)
		}
	})
}
