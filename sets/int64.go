package sets

import (
	"fmt"
	"strings"
)

// Int64Set represents a set of int64 elements
type Int64Set map[int64]struct{}

// Add adds zero or more elements to the set
func (s Int64Set) Add(elems ...int64) {
	for _, e := range elems {
		s[e] = struct{}{}
	}
}

// Remove removes zero or more elements from the set
func (s Int64Set) Remove(elems ...int64) {
	for _, e := range elems {
		delete(s, e)
	}
}

// Empty empties the set
func (s Int64Set) Empty() {
	for e := range s {
		delete(s, e)
	}
}

// Has indicates whether the set has an element
func (s Int64Set) Has(elem int64) bool {
	_, ok := s[elem]
	return ok
}

// Size returns the size of the set
func (s Int64Set) Size() int {
	return len(s)
}

// IsEmpty indicates whether the set is empty
func (s Int64Set) IsEmpty() bool {
	return len(s) == 0
}

// AsSlice returns an equivalent slice with no specific order of the elements
func (s Int64Set) AsSlice() []int64 {
	a := make([]int64, len(s))
	i := 0
	for e := range s {
		a[i] = e
		i++
	}
	return a
}

// String returns a string representation of the set
func (s Int64Set) String() string {
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
func (s Int64Set) Equals(t Int64Set) bool {
	if len(s) != len(t) {
		return false
	}
	r := make(Int64Set, len(s))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		delete(r, e)
	}
	return len(r) == 0
}

// Union returns the union of s and t
func (s Int64Set) Union(t Int64Set) Int64Set {
	r := make(Int64Set, len(s)+len(t))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		r[e] = struct{}{}
	}
	return r
}

// Intersection returns the intersection of s and t
func (s Int64Set) Intersection(t Int64Set) Int64Set {
	var small, large Int64Set
	if len(s) <= len(t) {
		small, large = s, t
	} else {
		small, large = t, s
	}
	r := make(Int64Set, len(small))
	for e := range small {
		if large.Has(e) {
			r[e] = struct{}{}
		}
	}
	return r
}

// Difference returns the difference of s and t, i.e., s - t
func (s Int64Set) Difference(t Int64Set) Int64Set {
	r := make(Int64Set, len(s))
	for e := range s {
		if !t.Has(e) {
			r[e] = struct{}{}
		}
	}
	return r
}

// IsSubsetOf indicates whether s is a subset of t
func (s Int64Set) IsSubsetOf(t Int64Set) bool {
	for e := range s {
		if !t.Has(e) {
			return false
		}
	}
	return true
}

// IsDisjointFrom indicates whether s and t are disjoint
func (s Int64Set) IsDisjointFrom(t Int64Set) bool {
	var small, large Int64Set
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

// NewInt64Set returns a new Int64Set containing zero or more elements
func NewInt64Set(elems ...int64) Int64Set {
	s := make(Int64Set, len(elems))
	s.Add(elems...)
	return s
}
