package trolly

import "testing"

func TestProbablyContains(t *testing.T) {

}

func testProbablyContains(t *testing.T, length, n, iters int) {

}

func TestCardinality(t *testing.T) {
	for i := 0; i < 22; i++ {
		s := make(Set, i)
		l := s.Cardinality()
		if i != l {
			t.Errorf("Expected cardinality %v; got %v", i, l)
		}
	}
}
