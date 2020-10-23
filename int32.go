package menge

import (
	"fmt"
	"strings"
)

// Int32Set represents a set of int32 elements.
type Int32Set map[int32]struct{}

// Add adds zero or more elements to the set.
func (s Int32Set) Add(elems ...int32) {
	for _, e := range elems {
		s[e] = struct{}{}
	}
}

// Remove removes zero or more elements from the set.
func (s Int32Set) Remove(elems ...int32) {
	for _, e := range elems {
		delete(s, e)
	}
}

// Empty empties the set.
func (s Int32Set) Empty() {
	for e := range s {
		delete(s, e)
	}
}

// Has indicates whether the set has an element.
func (s Int32Set) Has(elem int32) bool {
	_, ok := s[elem]
	return ok
}

// Size returns the size of the set.
func (s Int32Set) Size() int {
	return len(s)
}

// IsEmpty indicates whether the set is empty.
func (s Int32Set) IsEmpty() bool {
	return len(s) == 0
}

// AsSlice returns an equivalent slice with no specific order of the elements.
func (s Int32Set) AsSlice() []int32 {
	a := make([]int32, len(s))
	i := 0
	for e := range s {
		a[i] = e
		i++
	}
	return a
}

// String returns a string representation of the set.
func (s Int32Set) String() string {
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
func (s Int32Set) Equals(t Int32Set) bool {
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
func (s Int32Set) Union(t Int32Set) Int32Set {
	r := make(Int32Set, len(s)+len(t))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		r[e] = struct{}{}
	}
	return r
}

// Intersection returns the intersection of s and t.
func (s Int32Set) Intersection(t Int32Set) Int32Set {
	var small, large Int32Set
	if len(s) <= len(t) {
		small, large = s, t
	} else {
		small, large = t, s
	}
	r := make(Int32Set, len(small))
	for e := range small {
		if _, ok := large[e]; ok {
			r[e] = struct{}{}
		}
	}
	return r
}

// Difference returns the difference of s and t, i.e., s - t.
func (s Int32Set) Difference(t Int32Set) Int32Set {
	r := make(Int32Set, len(s))
	for e := range s {
		if _, ok := t[e]; !ok {
			r[e] = struct{}{}
		}
	}
	return r
}

// IsSubsetOf indicates whether s is a subset of t.
func (s Int32Set) IsSubsetOf(t Int32Set) bool {
	for e := range s {
		if _, ok := t[e]; !ok {
			return false
		}
	}
	return true
}

// IsProperSubsetOf indicates whether s is a proper subset of t.
func (s Int32Set) IsProperSubsetOf(t Int32Set) bool {
	for e := range s {
		if _, ok := t[e]; !ok {
			return false
		}
	}
	return len(s) != len(t)
}

// IsSupersetOf indicates whether s is a superset of t.
func (s Int32Set) IsSupersetOf(t Int32Set) bool {
	for e := range t {
		if _, ok := s[e]; !ok {
			return false
		}
	}
	return true
}

// IsProperSupersetOf indicates whether s is a proper superset of t.
func (s Int32Set) IsProperSupersetOf(t Int32Set) bool {
	for e := range t {
		if _, ok := s[e]; !ok {
			return false
		}
	}
	return len(s) != len(t)
}

// IsDisjointFrom indicates whether s and t are disjoint.
func (s Int32Set) IsDisjointFrom(t Int32Set) bool {
	var small, large Int32Set
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

// NewInt32Set returns a new Int32Set containing zero or more elements.
func NewInt32Set(elems ...int32) Int32Set {
	s := make(Int32Set, len(elems))
	s.Add(elems...)
	return s
}
