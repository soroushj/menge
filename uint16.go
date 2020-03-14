package menge

import (
	"fmt"
	"strings"
)

// UInt16Set represents a set of uint16 elements.
type UInt16Set map[uint16]struct{}

// Add adds zero or more elements to the set.
func (s UInt16Set) Add(elems ...uint16) {
	for _, e := range elems {
		s[e] = struct{}{}
	}
}

// Remove removes zero or more elements from the set.
func (s UInt16Set) Remove(elems ...uint16) {
	for _, e := range elems {
		delete(s, e)
	}
}

// Empty empties the set.
func (s UInt16Set) Empty() {
	for e := range s {
		delete(s, e)
	}
}

// Has indicates whether the set has an element.
func (s UInt16Set) Has(elem uint16) bool {
	_, ok := s[elem]
	return ok
}

// Size returns the size of the set.
func (s UInt16Set) Size() int {
	return len(s)
}

// IsEmpty indicates whether the set is empty.
func (s UInt16Set) IsEmpty() bool {
	return len(s) == 0
}

// AsSlice returns an equivalent slice with no specific order of the elements.
func (s UInt16Set) AsSlice() []uint16 {
	a := make([]uint16, len(s))
	i := 0
	for e := range s {
		a[i] = e
		i++
	}
	return a
}

// String returns a string representation of the set.
func (s UInt16Set) String() string {
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
func (s UInt16Set) Equals(t UInt16Set) bool {
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
func (s UInt16Set) Union(t UInt16Set) UInt16Set {
	r := make(UInt16Set, len(s)+len(t))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		r[e] = struct{}{}
	}
	return r
}

// Intersection returns the intersection of s and t.
func (s UInt16Set) Intersection(t UInt16Set) UInt16Set {
	var small, large UInt16Set
	if len(s) <= len(t) {
		small, large = s, t
	} else {
		small, large = t, s
	}
	r := make(UInt16Set, len(small))
	for e := range small {
		if _, ok := large[e]; ok {
			r[e] = struct{}{}
		}
	}
	return r
}

// Difference returns the difference of s and t, i.e., s - t.
func (s UInt16Set) Difference(t UInt16Set) UInt16Set {
	r := make(UInt16Set, len(s))
	for e := range s {
		if _, ok := t[e]; !ok {
			r[e] = struct{}{}
		}
	}
	return r
}

// IsSubsetOf indicates whether s is a subset of t.
func (s UInt16Set) IsSubsetOf(t UInt16Set) bool {
	for e := range s {
		if _, ok := t[e]; !ok {
			return false
		}
	}
	return true
}

// IsDisjointFrom indicates whether s and t are disjoint.
func (s UInt16Set) IsDisjointFrom(t UInt16Set) bool {
	var small, large UInt16Set
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

// NewUInt16Set returns a new UInt16Set containing zero or more elements.
func NewUInt16Set(elems ...uint16) UInt16Set {
	s := make(UInt16Set, len(elems))
	s.Add(elems...)
	return s
}
