package menge

import (
	"fmt"
	"strings"
)

// Complex128Set represents a set of complex128 elements.
type Complex128Set map[complex128]struct{}

// Add adds zero or more elements to the set.
func (s Complex128Set) Add(elems ...complex128) {
	for _, e := range elems {
		s[e] = struct{}{}
	}
}

// Remove removes zero or more elements from the set.
func (s Complex128Set) Remove(elems ...complex128) {
	for _, e := range elems {
		delete(s, e)
	}
}

// Empty empties the set.
func (s Complex128Set) Empty() {
	for e := range s {
		delete(s, e)
	}
}

// Has indicates whether the set has an element.
func (s Complex128Set) Has(elem complex128) bool {
	_, ok := s[elem]
	return ok
}

// Size returns the size of the set.
func (s Complex128Set) Size() int {
	return len(s)
}

// IsEmpty indicates whether the set is empty.
func (s Complex128Set) IsEmpty() bool {
	return len(s) == 0
}

// AsSlice returns an equivalent slice with no specific order of the elements.
func (s Complex128Set) AsSlice() []complex128 {
	a := make([]complex128, len(s))
	i := 0
	for e := range s {
		a[i] = e
		i++
	}
	return a
}

// String returns a string representation of the set.
func (s Complex128Set) String() string {
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
func (s Complex128Set) Equals(t Complex128Set) bool {
	if len(s) != len(t) {
		return false
	}
	r := make(Complex128Set, len(s))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		delete(r, e)
	}
	return len(r) == 0
}

// Union returns the union of s and t.
func (s Complex128Set) Union(t Complex128Set) Complex128Set {
	r := make(Complex128Set, len(s)+len(t))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		r[e] = struct{}{}
	}
	return r
}

// Intersection returns the intersection of s and t.
func (s Complex128Set) Intersection(t Complex128Set) Complex128Set {
	var small, large Complex128Set
	if len(s) <= len(t) {
		small, large = s, t
	} else {
		small, large = t, s
	}
	r := make(Complex128Set, len(small))
	for e := range small {
		if _, ok := large[e]; ok {
			r[e] = struct{}{}
		}
	}
	return r
}

// Difference returns the difference of s and t, i.e., s - t.
func (s Complex128Set) Difference(t Complex128Set) Complex128Set {
	r := make(Complex128Set, len(s))
	for e := range s {
		if _, ok := t[e]; !ok {
			r[e] = struct{}{}
		}
	}
	return r
}

// IsSubsetOf indicates whether s is a subset of t.
func (s Complex128Set) IsSubsetOf(t Complex128Set) bool {
	for e := range s {
		if _, ok := t[e]; !ok {
			return false
		}
	}
	return true
}

// IsDisjointFrom indicates whether s and t are disjoint.
func (s Complex128Set) IsDisjointFrom(t Complex128Set) bool {
	var small, large Complex128Set
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

// NewComplex128Set returns a new Complex128Set containing zero or more elements.
func NewComplex128Set(elems ...complex128) Complex128Set {
	s := make(Complex128Set, len(elems))
	s.Add(elems...)
	return s
}
