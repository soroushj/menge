package menge

import (
	"fmt"
	"strings"
)

// Complex64Set represents a set of complex64 elements.
type Complex64Set map[complex64]struct{}

// Add adds zero or more elements to the set.
func (s Complex64Set) Add(elems ...complex64) {
	for _, e := range elems {
		s[e] = struct{}{}
	}
}

// Remove removes zero or more elements from the set.
func (s Complex64Set) Remove(elems ...complex64) {
	for _, e := range elems {
		delete(s, e)
	}
}

// Empty empties the set.
func (s Complex64Set) Empty() {
	for e := range s {
		delete(s, e)
	}
}

// Has indicates whether the set has an element.
func (s Complex64Set) Has(elem complex64) bool {
	_, ok := s[elem]
	return ok
}

// Size returns the size of the set.
func (s Complex64Set) Size() int {
	return len(s)
}

// IsEmpty indicates whether the set is empty.
func (s Complex64Set) IsEmpty() bool {
	return len(s) == 0
}

// AsSlice returns an equivalent slice with no specific order of the elements.
func (s Complex64Set) AsSlice() []complex64 {
	a := make([]complex64, len(s))
	i := 0
	for e := range s {
		a[i] = e
		i++
	}
	return a
}

// String returns a string representation of the set.
func (s Complex64Set) String() string {
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
func (s Complex64Set) Equals(t Complex64Set) bool {
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
func (s Complex64Set) Union(t Complex64Set) Complex64Set {
	r := make(Complex64Set, len(s)+len(t))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		r[e] = struct{}{}
	}
	return r
}

// Intersection returns the intersection of s and t.
func (s Complex64Set) Intersection(t Complex64Set) Complex64Set {
	var small, large Complex64Set
	if len(s) <= len(t) {
		small, large = s, t
	} else {
		small, large = t, s
	}
	r := make(Complex64Set, len(small))
	for e := range small {
		if _, ok := large[e]; ok {
			r[e] = struct{}{}
		}
	}
	return r
}

// Difference returns the difference of s and t, i.e., s - t.
func (s Complex64Set) Difference(t Complex64Set) Complex64Set {
	r := make(Complex64Set, len(s))
	for e := range s {
		if _, ok := t[e]; !ok {
			r[e] = struct{}{}
		}
	}
	return r
}

// IsSubsetOf indicates whether s is a subset of t.
func (s Complex64Set) IsSubsetOf(t Complex64Set) bool {
	for e := range s {
		if _, ok := t[e]; !ok {
			return false
		}
	}
	return true
}

// IsProperSubsetOf indicates whether s is a proper subset of t.
func (s Complex64Set) IsProperSubsetOf(t Complex64Set) bool {
	for e := range s {
		if _, ok := t[e]; !ok {
			return false
		}
	}
	return len(s) != len(t)
}

// IsSupersetOf indicates whether s is a superset of t.
func (s Complex64Set) IsSupersetOf(t Complex64Set) bool {
	for e := range t {
		if _, ok := s[e]; !ok {
			return false
		}
	}
	return true
}

// IsProperSupersetOf indicates whether s is a proper superset of t.
func (s Complex64Set) IsProperSupersetOf(t Complex64Set) bool {
	for e := range t {
		if _, ok := s[e]; !ok {
			return false
		}
	}
	return len(s) != len(t)
}

// IsDisjointFrom indicates whether s and t are disjoint.
func (s Complex64Set) IsDisjointFrom(t Complex64Set) bool {
	var small, large Complex64Set
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

// NewComplex64Set returns a new Complex64Set containing zero or more elements.
func NewComplex64Set(elems ...complex64) Complex64Set {
	s := make(Complex64Set, len(elems))
	s.Add(elems...)
	return s
}
