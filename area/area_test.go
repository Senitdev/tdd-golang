package area

import "testing"

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangles{2.0, 5.2}
		got := CalRec(rectangle)
		want := 14.4
		if got != want {
			t.Errorf("got %f, want %f", got, want)
		}
	})
	t.Run("Calcule surface", func(t *testing.T) {
		rec := Rectangles{3.2, 4.5}
		got := CalculeArea(rec)
		want := 14.4
		if got != want {
			t.Errorf("got %f , want %f", got, want)
		}
	})
	t.Run("Cercle", func(t *testing.T) {
		cercle := Circle{10}
		got := CalculeCercle(cercle)
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %f , want %f", got, want)
		}
	})

}
