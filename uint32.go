// Package menge implements efficient sets of all basic types.
package menge

import (
	"fmt"
	"strings"
)

// UInt32Set represents a set of uint32 elements
type UInt32Set map[uint32]struct{}

// Add adds zero or more elements to the set
func (s UInt32Set) Add(elems ...uint32) {
	for _, e := range elems {
		s[e] = struct{}{}
	}
}

// Remove removes zero or more elements from the set
func (s UInt32Set) Remove(elems ...uint32) {
	for _, e := range elems {
		delete(s, e)
	}
}

// Empty empties the set
func (s UInt32Set) Empty() {
	for e := range s {
		delete(s, e)
	}
}

// Has indicates whether the set has an element
func (s UInt32Set) Has(elem uint32) bool {
	_, ok := s[elem]
	return ok
}

// Size returns the size of the set
func (s UInt32Set) Size() int {
	return len(s)
}

// IsEmpty indicates whether the set is empty
func (s UInt32Set) IsEmpty() bool {
	return len(s) == 0
}

// AsSlice returns an equivalent slice with no specific order of the elements
func (s UInt32Set) AsSlice() []uint32 {
	a := make([]uint32, len(s))
	i := 0
	for e := range s {
		a[i] = e
		i++
	}
	return a
}

// String returns a string representation of the set
func (s UInt32Set) String() string {
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
func (s UInt32Set) Equals(t UInt32Set) bool {
	if len(s) != len(t) {
		return false
	}
	r := make(UInt32Set, len(s))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		delete(r, e)
	}
	return len(r) == 0
}

// Union returns the union of s and t
func (s UInt32Set) Union(t UInt32Set) UInt32Set {
	r := make(UInt32Set, len(s)+len(t))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		r[e] = struct{}{}
	}
	return r
}

// Intersection returns the intersection of s and t
func (s UInt32Set) Intersection(t UInt32Set) UInt32Set {
	var small, large UInt32Set
	if len(s) <= len(t) {
		small, large = s, t
	} else {
		small, large = t, s
	}
	r := make(UInt32Set, len(small))
	for e := range small {
		if large.Has(e) {
			r[e] = struct{}{}
		}
	}
	return r
}

// Difference returns the difference of s and t, i.e., s - t
func (s UInt32Set) Difference(t UInt32Set) UInt32Set {
	r := make(UInt32Set, len(s))
	for e := range s {
		if !t.Has(e) {
			r[e] = struct{}{}
		}
	}
	return r
}

// IsSubsetOf indicates whether s is a subset of t
func (s UInt32Set) IsSubsetOf(t UInt32Set) bool {
	for e := range s {
		if !t.Has(e) {
			return false
		}
	}
	return true
}

// IsDisjointFrom indicates whether s and t are disjoint
func (s UInt32Set) IsDisjointFrom(t UInt32Set) bool {
	var small, large UInt32Set
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

// NewUInt32Set returns a new UInt32Set containing zero or more elements
func NewUInt32Set(elems ...uint32) UInt32Set {
	s := make(UInt32Set, len(elems))
	s.Add(elems...)
	return s
}
