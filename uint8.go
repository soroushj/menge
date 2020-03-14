package menge

import (
	"fmt"
	"strings"
)

// UInt8Set represents a set of uint8 elements.
type UInt8Set map[uint8]struct{}

// Add adds zero or more elements to the set.
func (s UInt8Set) Add(elems ...uint8) {
	for _, e := range elems {
		s[e] = struct{}{}
	}
}

// Remove removes zero or more elements from the set.
func (s UInt8Set) Remove(elems ...uint8) {
	for _, e := range elems {
		delete(s, e)
	}
}

// Empty empties the set.
func (s UInt8Set) Empty() {
	for e := range s {
		delete(s, e)
	}
}

// Has indicates whether the set has an element.
func (s UInt8Set) Has(elem uint8) bool {
	_, ok := s[elem]
	return ok
}

// Size returns the size of the set.
func (s UInt8Set) Size() int {
	return len(s)
}

// IsEmpty indicates whether the set is empty.
func (s UInt8Set) IsEmpty() bool {
	return len(s) == 0
}

// AsSlice returns an equivalent slice with no specific order of the elements.
func (s UInt8Set) AsSlice() []uint8 {
	a := make([]uint8, len(s))
	i := 0
	for e := range s {
		a[i] = e
		i++
	}
	return a
}

// String returns a string representation of the set.
func (s UInt8Set) String() string {
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
func (s UInt8Set) Equals(t UInt8Set) bool {
	if len(s) != len(t) {
		return false
	}
	r := make(UInt8Set, len(s))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		delete(r, e)
	}
	return len(r) == 0
}

// Union returns the union of s and t.
func (s UInt8Set) Union(t UInt8Set) UInt8Set {
	r := make(UInt8Set, len(s)+len(t))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		r[e] = struct{}{}
	}
	return r
}

// Intersection returns the intersection of s and t.
func (s UInt8Set) Intersection(t UInt8Set) UInt8Set {
	var small, large UInt8Set
	if len(s) <= len(t) {
		small, large = s, t
	} else {
		small, large = t, s
	}
	r := make(UInt8Set, len(small))
	for e := range small {
		if _, ok := large[e]; ok {
			r[e] = struct{}{}
		}
	}
	return r
}

// Difference returns the difference of s and t, i.e., s - t.
func (s UInt8Set) Difference(t UInt8Set) UInt8Set {
	r := make(UInt8Set, len(s))
	for e := range s {
		if _, ok := t[e]; !ok {
			r[e] = struct{}{}
		}
	}
	return r
}

// IsSubsetOf indicates whether s is a subset of t.
func (s UInt8Set) IsSubsetOf(t UInt8Set) bool {
	for e := range s {
		if _, ok := t[e]; !ok {
			return false
		}
	}
	return true
}

// IsDisjointFrom indicates whether s and t are disjoint.
func (s UInt8Set) IsDisjointFrom(t UInt8Set) bool {
	var small, large UInt8Set
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

// NewUInt8Set returns a new UInt8Set containing zero or more elements.
func NewUInt8Set(elems ...uint8) UInt8Set {
	s := make(UInt8Set, len(elems))
	s.Add(elems...)
	return s
}
