package menge

import (
	"fmt"
	"strings"
)

// Float64Set represents a set of float64 elements.
type Float64Set map[float64]struct{}

// Add adds zero or more elements to the set.
func (s Float64Set) Add(elems ...float64) {
	for _, e := range elems {
		s[e] = struct{}{}
	}
}

// Remove removes zero or more elements from the set.
func (s Float64Set) Remove(elems ...float64) {
	for _, e := range elems {
		delete(s, e)
	}
}

// Empty empties the set.
func (s Float64Set) Empty() {
	for e := range s {
		delete(s, e)
	}
}

// Has indicates whether the set has an element.
func (s Float64Set) Has(elem float64) bool {
	_, ok := s[elem]
	return ok
}

// Size returns the size of the set.
func (s Float64Set) Size() int {
	return len(s)
}

// IsEmpty indicates whether the set is empty.
func (s Float64Set) IsEmpty() bool {
	return len(s) == 0
}

// AsSlice returns an equivalent slice with no specific order of the elements.
func (s Float64Set) AsSlice() []float64 {
	a := make([]float64, len(s))
	i := 0
	for e := range s {
		a[i] = e
		i++
	}
	return a
}

// String returns a string representation of the set.
func (s Float64Set) String() string {
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

// Equals indicates whether s and t are equal.
func (s Float64Set) Equals(t Float64Set) bool {
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
func (s Float64Set) Union(t Float64Set) Float64Set {
	r := make(Float64Set, len(s)+len(t))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		r[e] = struct{}{}
	}
	return r
}

// Intersection returns the intersection of s and t.
func (s Float64Set) Intersection(t Float64Set) Float64Set {
	var small, large Float64Set
	if len(s) <= len(t) {
		small, large = s, t
	} else {
		small, large = t, s
	}
	r := make(Float64Set, len(small))
	for e := range small {
		if _, ok := large[e]; ok {
			r[e] = struct{}{}
		}
	}
	return r
}

// Difference returns the difference of s and t, i.e., s - t.
func (s Float64Set) Difference(t Float64Set) Float64Set {
	r := make(Float64Set, len(s))
	for e := range s {
		if _, ok := t[e]; !ok {
			r[e] = struct{}{}
		}
	}
	return r
}

// IsSubsetOf indicates whether s is a subset of t.
func (s Float64Set) IsSubsetOf(t Float64Set) bool {
	for e := range s {
		if _, ok := t[e]; !ok {
			return false
		}
	}
	return true
}

// IsDisjointFrom indicates whether s and t are disjoint.
func (s Float64Set) IsDisjointFrom(t Float64Set) bool {
	var small, large Float64Set
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

// NewFloat64Set returns a new Float64Set containing zero or more elements.
func NewFloat64Set(elems ...float64) Float64Set {
	s := make(Float64Set, len(elems))
	s.Add(elems...)
	return s
}
