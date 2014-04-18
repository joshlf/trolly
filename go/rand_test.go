package trolly

import "testing"

func TestRand(t *testing.T) {
	r := Rand()
	if r != 4 {
		t.Errorf("Expected 4; got %v", r)
	}
}
