package menge

import (
	"fmt"
	"strings"
)

// UIntPtrSet represents a set of uintptr elements.
type UIntPtrSet map[uintptr]struct{}

// Add adds zero or more elements to the set.
func (s UIntPtrSet) Add(elems ...uintptr) {
	for _, e := range elems {
		s[e] = struct{}{}
	}
}

// Remove removes zero or more elements from the set.
func (s UIntPtrSet) Remove(elems ...uintptr) {
	for _, e := range elems {
		delete(s, e)
	}
}

// Empty empties the set.
func (s UIntPtrSet) Empty() {
	for e := range s {
		delete(s, e)
	}
}

// Has indicates whether the set has an element.
func (s UIntPtrSet) Has(elem uintptr) bool {
	_, ok := s[elem]
	return ok
}

// Size returns the size of the set.
func (s UIntPtrSet) Size() int {
	return len(s)
}

// IsEmpty indicates whether the set is empty.
func (s UIntPtrSet) IsEmpty() bool {
	return len(s) == 0
}

// Clone returns a clone of the set.
func (s UIntPtrSet) Clone() UIntPtrSet {
	c := make(UIntPtrSet, len(s))
	for e := range s {
		c[e] = struct{}{}
	}
	return c
}

// AsSlice returns an equivalent slice with no specific order of the elements.
func (s UIntPtrSet) AsSlice() []uintptr {
	a := make([]uintptr, len(s))
	i := 0
	for e := range s {
		a[i] = e
		i++
	}
	return a
}

// String returns a string representation of the set.
func (s UIntPtrSet) String() string {
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
func (s UIntPtrSet) Equals(t UIntPtrSet) bool {
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
func (s UIntPtrSet) Union(t UIntPtrSet) UIntPtrSet {
	r := make(UIntPtrSet, len(s)+len(t))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		r[e] = struct{}{}
	}
	return r
}

// Intersection returns the intersection of s and t.
func (s UIntPtrSet) Intersection(t UIntPtrSet) UIntPtrSet {
	var small, large UIntPtrSet
	if len(s) <= len(t) {
		small, large = s, t
	} else {
		small, large = t, s
	}
	r := make(UIntPtrSet, len(small))
	for e := range small {
		if _, ok := large[e]; ok {
			r[e] = struct{}{}
		}
	}
	return r
}

// Difference returns the difference of s and t, i.e., s - t.
func (s UIntPtrSet) Difference(t UIntPtrSet) UIntPtrSet {
	r := make(UIntPtrSet, len(s))
	for e := range s {
		if _, ok := t[e]; !ok {
			r[e] = struct{}{}
		}
	}
	return r
}

// IsSubsetOf indicates whether s is a subset of t.
func (s UIntPtrSet) IsSubsetOf(t UIntPtrSet) bool {
	for e := range s {
		if _, ok := t[e]; !ok {
			return false
		}
	}
	return true
}

// IsProperSubsetOf indicates whether s is a proper subset of t.
func (s UIntPtrSet) IsProperSubsetOf(t UIntPtrSet) bool {
	for e := range s {
		if _, ok := t[e]; !ok {
			return false
		}
	}
	return len(s) != len(t)
}

// IsSupersetOf indicates whether s is a superset of t.
func (s UIntPtrSet) IsSupersetOf(t UIntPtrSet) bool {
	for e := range t {
		if _, ok := s[e]; !ok {
			return false
		}
	}
	return true
}

// IsProperSupersetOf indicates whether s is a proper superset of t.
func (s UIntPtrSet) IsProperSupersetOf(t UIntPtrSet) bool {
	for e := range t {
		if _, ok := s[e]; !ok {
			return false
		}
	}
	return len(s) != len(t)
}

// IsDisjointFrom indicates whether s and t are disjoint.
func (s UIntPtrSet) IsDisjointFrom(t UIntPtrSet) bool {
	var small, large UIntPtrSet
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

// NewUIntPtrSet returns a new UIntPtrSet containing zero or more elements.
func NewUIntPtrSet(elems ...uintptr) UIntPtrSet {
	s := make(UIntPtrSet, len(elems))
	s.Add(elems...)
	return s
}
