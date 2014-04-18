package trolly

import (
	"fmt"
	"math"
	"testing"
)

func BenchmarkSetCardinalityLogarithmic(b *testing.B) {
	n := int(math.Log2(float64(b.N)))
	s := make(Set, n)
	fmt.Println(n, b.N)

	b.ResetTimer()
	s.Cardinality()
}
