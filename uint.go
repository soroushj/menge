package menge

import (
	"fmt"
	"strings"
)

// UIntSet represents a set of uint elements.
type UIntSet map[uint]struct{}

// Add adds zero or more elements to the set.
func (s UIntSet) Add(elems ...uint) {
	for _, e := range elems {
		s[e] = struct{}{}
	}
}

// Remove removes zero or more elements from the set.
func (s UIntSet) Remove(elems ...uint) {
	for _, e := range elems {
		delete(s, e)
	}
}

// Empty empties the set.
func (s UIntSet) Empty() {
	for e := range s {
		delete(s, e)
	}
}

// Has indicates whether the set has an element.
func (s UIntSet) Has(elem uint) bool {
	_, ok := s[elem]
	return ok
}

// Size returns the size of the set.
func (s UIntSet) Size() int {
	return len(s)
}

// IsEmpty indicates whether the set is empty.
func (s UIntSet) IsEmpty() bool {
	return len(s) == 0
}

// Clone returns a clone of the set.
func (s UIntSet) Clone() UIntSet {
	c := make(UIntSet, len(s))
	for e := range s {
		c[e] = struct{}{}
	}
	return c
}

// AsSlice returns an equivalent slice with no specific order of the elements.
func (s UIntSet) AsSlice() []uint {
	a := make([]uint, len(s))
	i := 0
	for e := range s {
		a[i] = e
		i++
	}
	return a
}

// String returns a string representation of the set.
func (s UIntSet) String() string {
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
func (s UIntSet) Equals(t UIntSet) bool {
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
func (s UIntSet) Union(t UIntSet) UIntSet {
	r := make(UIntSet, len(s)+len(t))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		r[e] = struct{}{}
	}
	return r
}

// Intersection returns the intersection of s and t.
func (s UIntSet) Intersection(t UIntSet) UIntSet {
	var small, large UIntSet
	if len(s) <= len(t) {
		small, large = s, t
	} else {
		small, large = t, s
	}
	r := make(UIntSet, len(small))
	for e := range small {
		if _, ok := large[e]; ok {
			r[e] = struct{}{}
		}
	}
	return r
}

// Difference returns the difference of s and t, i.e., s - t.
func (s UIntSet) Difference(t UIntSet) UIntSet {
	r := make(UIntSet, len(s))
	for e := range s {
		if _, ok := t[e]; !ok {
			r[e] = struct{}{}
		}
	}
	return r
}

// IsSubsetOf indicates whether s is a subset of t.
func (s UIntSet) IsSubsetOf(t UIntSet) bool {
	for e := range s {
		if _, ok := t[e]; !ok {
			return false
		}
	}
	return true
}

// IsProperSubsetOf indicates whether s is a proper subset of t.
func (s UIntSet) IsProperSubsetOf(t UIntSet) bool {
	for e := range s {
		if _, ok := t[e]; !ok {
			return false
		}
	}
	return len(s) != len(t)
}

// IsSupersetOf indicates whether s is a superset of t.
func (s UIntSet) IsSupersetOf(t UIntSet) bool {
	for e := range t {
		if _, ok := s[e]; !ok {
			return false
		}
	}
	return true
}

// IsProperSupersetOf indicates whether s is a proper superset of t.
func (s UIntSet) IsProperSupersetOf(t UIntSet) bool {
	for e := range t {
		if _, ok := s[e]; !ok {
			return false
		}
	}
	return len(s) != len(t)
}

// IsDisjointFrom indicates whether s and t are disjoint.
func (s UIntSet) IsDisjointFrom(t UIntSet) bool {
	var small, large UIntSet
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

// NewUIntSet returns a new UIntSet containing zero or more elements.
func NewUIntSet(elems ...uint) UIntSet {
	s := make(UIntSet, len(elems))
	s.Add(elems...)
	return s
}
