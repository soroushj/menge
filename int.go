package menge

import (
	"fmt"
	"strings"
)

// IntSet represents a set of int elements.
type IntSet map[int]struct{}

// Add adds zero or more elements to the set.
func (s IntSet) Add(elems ...int) {
	for _, e := range elems {
		s[e] = struct{}{}
	}
}

// Remove removes zero or more elements from the set.
func (s IntSet) Remove(elems ...int) {
	for _, e := range elems {
		delete(s, e)
	}
}

// Empty empties the set.
func (s IntSet) Empty() {
	for e := range s {
		delete(s, e)
	}
}

// Has indicates whether the set has an element.
func (s IntSet) Has(elem int) bool {
	_, ok := s[elem]
	return ok
}

// Size returns the size of the set.
func (s IntSet) Size() int {
	return len(s)
}

// IsEmpty indicates whether the set is empty.
func (s IntSet) IsEmpty() bool {
	return len(s) == 0
}

// Clone returns a clone of the set.
func (s IntSet) Clone() IntSet {
	c := make(IntSet, len(s))
	for e := range s {
		c[e] = struct{}{}
	}
	return c
}

// AsSlice returns an equivalent slice with no specific order of the elements.
func (s IntSet) AsSlice() []int {
	a := make([]int, len(s))
	i := 0
	for e := range s {
		a[i] = e
		i++
	}
	return a
}

// String returns a string representation of the set.
func (s IntSet) String() string {
	b := &strings.Builder{}
	b.Grow(len(s) * 8)
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

// Equals indicates whether s and t are equal.
func (s IntSet) Equals(t IntSet) bool {
	if len(s) != len(t) {
		return false
	}
	for e := range s {
		if _, ok := t[e]; !ok {
			return false
		}
	}
	return true
}

// Union returns the union of s and t.
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

// Intersection returns the intersection of s and t.
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

// Difference returns the difference of s and t, i.e., s - t.
func (s IntSet) Difference(t IntSet) IntSet {
	r := make(IntSet, len(s))
	for e := range s {
		if _, ok := t[e]; !ok {
			r[e] = struct{}{}
		}
	}
	return r
}

// IsSubsetOf indicates whether s is a subset of t.
func (s IntSet) IsSubsetOf(t IntSet) bool {
	for e := range s {
		if _, ok := t[e]; !ok {
			return false
		}
	}
	return true
}

// IsProperSubsetOf indicates whether s is a proper subset of t.
func (s IntSet) IsProperSubsetOf(t IntSet) bool {
	for e := range s {
		if _, ok := t[e]; !ok {
			return false
		}
	}
	return len(s) != len(t)
}

// IsSupersetOf indicates whether s is a superset of t.
func (s IntSet) IsSupersetOf(t IntSet) bool {
	for e := range t {
		if _, ok := s[e]; !ok {
			return false
		}
	}
	return true
}

// IsProperSupersetOf indicates whether s is a proper superset of t.
func (s IntSet) IsProperSupersetOf(t IntSet) bool {
	for e := range t {
		if _, ok := s[e]; !ok {
			return false
		}
	}
	return len(s) != len(t)
}

// IsDisjointFrom indicates whether s and t are disjoint.
func (s IntSet) IsDisjointFrom(t IntSet) bool {
	var small, large IntSet
	if len(s) <= len(t) {
		small, large = s, t
	} else {
		small, large = t, s
	}
	for e := range small {
		if _, ok := large[e]; ok {
			return false
		}
	}
	return true
}

// NewIntSet returns a new IntSet containing zero or more elements.
func NewIntSet(elems ...int) IntSet {
	s := make(IntSet, len(elems))
	s.Add(elems...)
	return s
}
