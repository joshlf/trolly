package trolly

import "testing"

func TestCardinality(t *testing.T) {
	for i := 0; i < 22; i++ {
		s := make(Set, i)
		l := s.Cardinality()
		if i != l {
			t.Errorf("Expected cardinality %v; got %v", i, l)
		}
	}
}
