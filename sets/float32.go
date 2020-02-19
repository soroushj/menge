package sets

import (
	"fmt"
	"strings"
)

// Float32Set represents a set of float32 elements
type Float32Set map[float32]struct{}

// Add adds zero or more elements to the set
func (s Float32Set) Add(elems ...float32) {
	for _, e := range elems {
		s[e] = struct{}{}
	}
}

// Remove removes zero or more elements from the set
func (s Float32Set) Remove(elems ...float32) {
	for _, e := range elems {
		delete(s, e)
	}
}

// Empty empties the set
func (s Float32Set) Empty() {
	for e := range s {
		delete(s, e)
	}
}

// Has indicates whether the set has an element
func (s Float32Set) Has(elem float32) bool {
	_, ok := s[elem]
	return ok
}

// Size returns the size of the set
func (s Float32Set) Size() int {
	return len(s)
}

// IsEmpty indicates whether the set is empty
func (s Float32Set) IsEmpty() bool {
	return len(s) == 0
}

// AsSlice returns an equivalent slice with no specific order of the elements
func (s Float32Set) AsSlice() []float32 {
	a := make([]float32, len(s))
	i := 0
	for e := range s {
		a[i] = e
		i++
	}
	return a
}

// String returns a string representation of the set
func (s Float32Set) String() string {
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
func (s Float32Set) Equals(t Float32Set) bool {
	if len(s) != len(t) {
		return false
	}
	r := make(Float32Set, len(s))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		delete(r, e)
	}
	return len(r) == 0
}

// Union returns the union of s and t
func (s Float32Set) Union(t Float32Set) Float32Set {
	r := make(Float32Set, len(s)+len(t))
	for e := range s {
		r[e] = struct{}{}
	}
	for e := range t {
		r[e] = struct{}{}
	}
	return r
}

// Intersection returns the intersection of s and t
func (s Float32Set) Intersection(t Float32Set) Float32Set {
	var small, large Float32Set
	if len(s) <= len(t) {
		small, large = s, t
	} else {
		small, large = t, s
	}
	r := make(Float32Set, len(small))
	for e := range small {
		if large.Has(e) {
			r[e] = struct{}{}
		}
	}
	return r
}

// Difference returns the difference of s and t, i.e., s - t
func (s Float32Set) Difference(t Float32Set) Float32Set {
	r := make(Float32Set, len(s))
	for e := range s {
		if !t.Has(e) {
			r[e] = struct{}{}
		}
	}
	return r
}

// IsSubsetOf indicates whether s is a subset of t
func (s Float32Set) IsSubsetOf(t Float32Set) bool {
	for e := range s {
		if !t.Has(e) {
			return false
		}
	}
	return true
}

// IsDisjointFrom indicates whether s and t are disjoint
func (s Float32Set) IsDisjointFrom(t Float32Set) bool {
	var small, large Float32Set
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

// NewFloat32Set returns a new Float32Set containing zero or more elements
func NewFloat32Set(elems ...float32) Float32Set {
	s := make(Float32Set, len(elems))
	s.Add(elems...)
	return s
}
