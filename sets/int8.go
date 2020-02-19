package sets

import (
	"fmt"
	"strings"
)

// Int8Set represents a set of int8 elements
type Int8Set map[int8]struct{}

// Add adds zero or more elements to the set
func (s Int8Set) Add(elems ...int8) {
	for _, e := range elems {
		s[e] = struct{}{}
	}
}

// Remove removes zero or more elements from the set
func (s Int8Set) Remove(elems ...int8) {
	for _, e := range elems {
		delete(s, e)
	}
}

// Empty empties the set
func (s Int8Set) Empty() {
	for e := range s {
		delete(s, e)
	}
}

// Has indicates whether the set has an element
func (s Int8Set) Has(elem int8) bool {
	_, ok := s[elem]
	return ok
}

// Size returns the size of the set
func (s Int8Set) Size() int {
	return len(s)
}

// IsEmpty indicates whether the set is empty
func (s Int8Set) IsEmpty() bool {
	return len(s) == 0
}

// AsSlice returns an equivalent slice with no specific order of the elements
func (s Int8Set) AsSlice() []int8 {
	a := make([]int8, len(s))
	i := 0
	for e := range s {
		a[i] = e
		i++
	}
	return a
}

// String returns a string representation of the set
func (s Int8Set) String() string {
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
func (s Int8Set) Equals(t Int8Set) bool {
	if len(s) != len(t) {
		return false
	}
	r := make(Int8Set, len(s))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		delete(r, e)
	}
	return len(r) == 0
}

// Union returns the union of s and t
func (s Int8Set) Union(t Int8Set) Int8Set {
	r := make(Int8Set, len(s)+len(t))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		r[e] = struct{}{}
	}
	return r
}

// Intersection returns the intersection of s and t
func (s Int8Set) Intersection(t Int8Set) Int8Set {
	var small, large Int8Set
	if len(s) <= len(t) {
		small, large = s, t
	} else {
		small, large = t, s
	}
	r := make(Int8Set, len(small))
	for e := range small {
		if large.Has(e) {
			r[e] = struct{}{}
		}
	}
	return r
}

// Difference returns the difference of s and t, i.e., s - t
func (s Int8Set) Difference(t Int8Set) Int8Set {
	r := make(Int8Set, len(s))
	for e := range s {
		if !t.Has(e) {
			r[e] = struct{}{}
		}
	}
	return r
}

// IsSubsetOf indicates whether s is a subset of t
func (s Int8Set) IsSubsetOf(t Int8Set) bool {
	for e := range s {
		if !t.Has(e) {
			return false
		}
	}
	return true
}

// IsDisjointFrom indicates whether s and t are disjoint
func (s Int8Set) IsDisjointFrom(t Int8Set) bool {
	var small, large Int8Set
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

// NewInt8Set returns a new Int8Set containing zero or more elements
func NewInt8Set(elems ...int8) Int8Set {
	s := make(Int8Set, len(elems))
	s.Add(elems...)
	return s
}
