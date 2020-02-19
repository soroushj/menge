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

func (s IntSet) Clear() {
	for e := range s {
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
	b := &strings.Builder{}
	b.Grow(len(s) * 4)
	fmt.Fprint(b, "{")
	first := true
	for e := range s {
		if first {
			first = false
			fmt.Fprintf(b, "%v", e)
		} else {
			fmt.Fprintf(b, " %v", e)
		}
	}
	fmt.Fprint(b, "}")
	return b.String()
}

func (s IntSet) Equals(t IntSet) bool {
	if len(s) != len(t) {
		return false
	}
	r := make(IntSet, len(s))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		delete(r, e)
	}
	return len(r) == 0
}

func (s IntSet) Union(t IntSet) IntSet {
	r := make(IntSet, len(s)+len(t))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		r[e] = struct{}{}
	}
	return r
}

func (s IntSet) Intersection(t IntSet) IntSet {
	var small, large IntSet
	if len(s) <= len(t) {
		small, large = s, t
	} else {
		small, large = t, s
	}
	r := make(IntSet, len(small))
	for e := range small {
		if _, ok := large[e]; ok {
			r[e] = struct{}{}
		}
	}
	return r
}

func (s IntSet) Disjoint(t IntSet) bool {
	return s.Intersection(t).IsEmpty()
}

func NewIntSet(elems ...int) IntSet {
	s := make(IntSet, len(elems))
	for _, e := range elems {
		s[e] = struct{}{}
	}
	return s
}
