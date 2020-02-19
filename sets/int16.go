package sets

import (
	"fmt"
	"strings"
)

// Int16Set represents a set of int16 elements
type Int16Set map[int16]struct{}

// Add adds zero or more elements to the set
func (s Int16Set) Add(elems ...int16) {
	for _, e := range elems {
		s[e] = struct{}{}
	}
}

// Remove removes zero or more elements from the set
func (s Int16Set) Remove(elems ...int16) {
	for _, e := range elems {
		delete(s, e)
	}
}

// Empty empties the set
func (s Int16Set) Empty() {
	for e := range s {
		delete(s, e)
	}
}

// Has indicates whether the set has an element
func (s Int16Set) Has(elem int16) bool {
	_, ok := s[elem]
	return ok
}

// Size returns the size of the set
func (s Int16Set) Size() int {
	return len(s)
}

// IsEmpty indicates whether the set is empty
func (s Int16Set) IsEmpty() bool {
	return len(s) == 0
}

// AsSlice returns an equivalent slice with no specific order of the elements
func (s Int16Set) AsSlice() []int16 {
	a := make([]int16, len(s))
	i := 0
	for e := range s {
		a[i] = e
		i++
	}
	return a
}

// String returns a string representation of the set
func (s Int16Set) String() string {
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
func (s Int16Set) Equals(t Int16Set) bool {
	if len(s) != len(t) {
		return false
	}
	r := make(Int16Set, len(s))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		delete(r, e)
	}
	return len(r) == 0
}

// Union returns the union of s and t
func (s Int16Set) Union(t Int16Set) Int16Set {
	r := make(Int16Set, len(s)+len(t))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		r[e] = struct{}{}
	}
	return r
}

// Intersection returns the intersection of s and t
func (s Int16Set) Intersection(t Int16Set) Int16Set {
	var small, large Int16Set
	if len(s) <= len(t) {
		small, large = s, t
	} else {
		small, large = t, s
	}
	r := make(Int16Set, len(small))
	for e := range small {
		if large.Has(e) {
			r[e] = struct{}{}
		}
	}
	return r
}

// Difference returns the difference of s and t, i.e., s - t
func (s Int16Set) Difference(t Int16Set) Int16Set {
	r := make(Int16Set, len(s))
	for e := range s {
		if !t.Has(e) {
			r[e] = struct{}{}
		}
	}
	return r
}

// IsSubsetOf indicates whether s is a subset of t
func (s Int16Set) IsSubsetOf(t Int16Set) bool {
	for e := range s {
		if !t.Has(e) {
			return false
		}
	}
	return true
}

// IsDisjointFrom indicates whether s and t are disjoint
func (s Int16Set) IsDisjointFrom(t Int16Set) bool {
	var small, large Int16Set
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

// NewInt16Set returns a new Int16Set containing zero or more elements
func NewInt16Set(elems ...int16) Int16Set {
	s := make(Int16Set, len(elems))
	s.Add(elems...)
	return s
}
