package addition

import (
	"testing"
)

func TestAddition(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, but got %d", result)
	}
}
