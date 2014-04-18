package trolly

import (
	"math"
	"math/rand"
)

type Set []interface{}

func (s Set) ProbablyContains(elem interface{}, n int) bool {
	l := len(s)
	for i := 0; i < n; i++ {
		if s[rand.Intn(l)] == elem {
			return true
		}
	}
	return false
}

func (s Set) Cardinality() int {
	return int(math.Log2(float64(len(s.subsets()))))
}

func (s Set) subsets() []Set {
	if len(s) == 0 {
		return []Set{Set{}}
	}

	subs := Set(s[1:]).subsets()
	l := len(subs)
	for i := 0; i < l; i++ {
		t := append(Set{}, subs[i]...)
		subs = append(subs, append(t, s[0]))
	}

	return subs
}
