package sets

import (
	"fmt"
	"strings"
)

type IntSet map[int]struct{}

func (s IntSet) Add(elems ...int) {
	for _, e := range elems {
		s[e] = struct{}{}
	}
}

func (s IntSet) Remove(elems ...int) {
	for _, e := range elems {
		delete(s, e)
	}
}

func (s IntSet) Has(elem int) bool {
	_, ok := s[elem]
	return ok
}

func (s IntSet) Size() int {
	return len(s)
}

func (s IntSet) IsEmpty() bool {
	return len(s) == 0
}

func (s IntSet) Clear() {
	for e := range s {
		delete(s, e)
	}
}

func (s IntSet) AsSlice() []int {
	a := make([]int, len(s))
	i := 0
	for e := range s {
		a[i] = e
		i++
	}
	return a
}

func (s IntSet) String() string {
	if len(s) == 0 {
		return "{}"
	}
	b := &strings.Builder{}
	b.Grow(len(s) * 4)
	first := true
	for e := range s {
		if first {
			first = false
			fmt.Fprintf(b, "{%v", e)
		} else {
			fmt.Fprintf(b, " %v", e)
		}
	}
	fmt.Fprint(b, "}")
	return b.String()
}

func NewIntSet() IntSet {
	return make(IntSet)
}
