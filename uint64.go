package menge

import (
	"fmt"
	"strings"
)

// UInt64Set represents a set of uint64 elements.
type UInt64Set map[uint64]struct{}

// Add adds zero or more elements to the set.
func (s UInt64Set) Add(elems ...uint64) {
	for _, e := range elems {
		s[e] = struct{}{}
	}
}

// Remove removes zero or more elements from the set.
func (s UInt64Set) Remove(elems ...uint64) {
	for _, e := range elems {
		delete(s, e)
	}
}

// Empty empties the set.
func (s UInt64Set) Empty() {
	for e := range s {
		delete(s, e)
	}
}

// Has indicates whether the set has an element.
func (s UInt64Set) Has(elem uint64) bool {
	_, ok := s[elem]
	return ok
}

// Size returns the size of the set.
func (s UInt64Set) Size() int {
	return len(s)
}

// IsEmpty indicates whether the set is empty.
func (s UInt64Set) IsEmpty() bool {
	return len(s) == 0
}

// AsSlice returns an equivalent slice with no specific order of the elements.
func (s UInt64Set) AsSlice() []uint64 {
	a := make([]uint64, len(s))
	i := 0
	for e := range s {
		a[i] = e
		i++
	}
	return a
}

// String returns a string representation of the set.
func (s UInt64Set) String() string {
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
func (s UInt64Set) Equals(t UInt64Set) bool {
	if len(s) != len(t) {
		return false
	}
	r := make(UInt64Set, len(s))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		delete(r, e)
	}
	return len(r) == 0
}

// Union returns the union of s and t.
func (s UInt64Set) Union(t UInt64Set) UInt64Set {
	r := make(UInt64Set, len(s)+len(t))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		r[e] = struct{}{}
	}
	return r
}

// Intersection returns the intersection of s and t.
func (s UInt64Set) Intersection(t UInt64Set) UInt64Set {
	var small, large UInt64Set
	if len(s) <= len(t) {
		small, large = s, t
	} else {
		small, large = t, s
	}
	r := make(UInt64Set, len(small))
	for e := range small {
		if large.Has(e) {
			r[e] = struct{}{}
		}
	}
	return r
}

// Difference returns the difference of s and t, i.e., s - t.
func (s UInt64Set) Difference(t UInt64Set) UInt64Set {
	r := make(UInt64Set, len(s))
	for e := range s {
		if !t.Has(e) {
			r[e] = struct{}{}
		}
	}
	return r
}

// IsSubsetOf indicates whether s is a subset of t.
func (s UInt64Set) IsSubsetOf(t UInt64Set) bool {
	for e := range s {
		if !t.Has(e) {
			return false
		}
	}
	return true
}

// IsDisjointFrom indicates whether s and t are disjoint.
func (s UInt64Set) IsDisjointFrom(t UInt64Set) bool {
	var small, large UInt64Set
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

// NewUInt64Set returns a new UInt64Set containing zero or more elements.
func NewUInt64Set(elems ...uint64) UInt64Set {
	s := make(UInt64Set, len(elems))
	s.Add(elems...)
	return s
}
