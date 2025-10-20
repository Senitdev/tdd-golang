package transaction

import (
	"testing"
)

func TestTransaction(t *testing.T) {
	wallet := Deposite(10)
	got := wallet
	want := 10
	if got != want {
		t.Errorf("got %d , want %d", got, want)
	}
	t.Run("Second depot", func(t *testing.T) {
		mt := Deposite(13)
		got := mt
		want := 23
		if got != want {
			t.Errorf("got %d , want %d", got, want)
		}
	})
	t.Run("Retrait sur le montant", func(t *testing.T) {
		mt2 := Retrait(15)
		got := mt2
		want := 8
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("Solde insuffisant", func(t *testing.T) {
		mt3 := Retrait(100)
		got := mt3
		want := 0
		if got != want {
			t.Errorf("A verifier")
		}
	})
}
