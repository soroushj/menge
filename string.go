// Package menge implements efficient sets of all basic types.
package menge

import (
	"fmt"
	"strings"
)

// StringSet represents a set of string elements
type StringSet map[string]struct{}

// Add adds zero or more elements to the set
func (s StringSet) Add(elems ...string) {
	for _, e := range elems {
		s[e] = struct{}{}
	}
}

// Remove removes zero or more elements from the set
func (s StringSet) Remove(elems ...string) {
	for _, e := range elems {
		delete(s, e)
	}
}

// Empty empties the set
func (s StringSet) Empty() {
	for e := range s {
		delete(s, e)
	}
}

// Has indicates whether the set has an element
func (s StringSet) Has(elem string) bool {
	_, ok := s[elem]
	return ok
}

// Size returns the size of the set
func (s StringSet) Size() int {
	return len(s)
}

// IsEmpty indicates whether the set is empty
func (s StringSet) IsEmpty() bool {
	return len(s) == 0
}

// AsSlice returns an equivalent slice with no specific order of the elements
func (s StringSet) AsSlice() []string {
	a := make([]string, len(s))
	i := 0
	for e := range s {
		a[i] = e
		i++
	}
	return a
}

// String returns a string representation of the set
func (s StringSet) String() string {
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

// Equals indicates whether s and t are equal
func (s StringSet) Equals(t StringSet) bool {
	if len(s) != len(t) {
		return false
	}
	r := make(StringSet, len(s))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		delete(r, e)
	}
	return len(r) == 0
}

// Union returns the union of s and t
func (s StringSet) Union(t StringSet) StringSet {
	r := make(StringSet, len(s)+len(t))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		r[e] = struct{}{}
	}
	return r
}

// Intersection returns the intersection of s and t
func (s StringSet) Intersection(t StringSet) StringSet {
	var small, large StringSet
	if len(s) <= len(t) {
		small, large = s, t
	} else {
		small, large = t, s
	}
	r := make(StringSet, len(small))
	for e := range small {
		if large.Has(e) {
			r[e] = struct{}{}
		}
	}
	return r
}

// Difference returns the difference of s and t, i.e., s - t
func (s StringSet) Difference(t StringSet) StringSet {
	r := make(StringSet, len(s))
	for e := range s {
		if !t.Has(e) {
			r[e] = struct{}{}
		}
	}
	return r
}

// IsSubsetOf indicates whether s is a subset of t
func (s StringSet) IsSubsetOf(t StringSet) bool {
	for e := range s {
		if !t.Has(e) {
			return false
		}
	}
	return true
}

// IsDisjointFrom indicates whether s and t are disjoint
func (s StringSet) IsDisjointFrom(t StringSet) bool {
	var small, large StringSet
	if len(s) <= len(t) {
		small, large = s, t
	} else {
		small, large = t, s
	}
	for e := range small {
		if large.Has(e) {
			return false
		}
	}
	return true
}

// NewStringSet returns a new StringSet containing zero or more elements
func NewStringSet(elems ...string) StringSet {
	s := make(StringSet, len(elems))
	s.Add(elems...)
	return s
}
