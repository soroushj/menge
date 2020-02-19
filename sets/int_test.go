package sets

import "testing"

func TestIntSetBasics(t *testing.T) {
	// NewIntSet() -> {}
	s := NewIntSet()
	if s.Has(1) {
		t.Error("NewIntSet() Has(1) got: true - want: false - expected set: {}")
	}
	if z := s.Size(); z != 0 {
		t.Errorf("NewIntSet() Size() got: %v - want: 0 - expected set: {}", z)
	}
	if !s.IsEmpty() {
		t.Error("NewIntSet() IsEmpty() got: false - want: true - expected set: {}")
	}
	if a := s.AsSlice(); len(a) != 0 {
		t.Errorf("NewIntSet() AsSlice() got: %v - want: [] - expected set: {}", a)
	}
	if r := s.String(); r != "{}" {
		t.Errorf("NewIntSet() String() got: %v - want: {} - expected set: {}", r)
	}
	// NewIntSet() Add(1, 2, 1) -> {1 2}
	s.Add(1, 2, 1)
	if !s.Has(1) {
		t.Error("NewIntSet() Add(1, 2, 1) Has(1) got: false - want: true - expected set: {1 2}")
	}
	if !s.Has(2) {
		t.Error("NewIntSet() Add(1, 2, 1) Has(2) got: false - want: true - expected set: {1 2}")
	}
	if z := s.Size(); z != 2 {
		t.Errorf("NewIntSet() Add(1, 2, 1) Size() got: %v - want: 2 - expected set: {1 2}", z)
	}
	if s.IsEmpty() {
		t.Error("NewIntSet() Add(1, 2, 1) IsEmpty() got: true - want: false - expected set: {1 2}")
	}
	if a := s.AsSlice(); len(a) != 2 || !((a[0] == 1 && a[1] == 2) || (a[0] == 2 && a[1] == 1)) {
		t.Errorf("NewIntSet() Add(1, 2, 1) AsSlice() got: %v - want: [1 2] or [2 1] - expected set: {1 2}", a)
	}
	if r := s.String(); r != "{1 2}" && r != "{2 1}" {
		t.Errorf("NewIntSet() Add(1, 2, 1) String() got: %v - want: {1 2} or {2 1} - expected set: {1 2}", r)
	}
	// NewIntSet() Add(1, 2, 1) Remove(2, 2) -> {1}
	s.Remove(2, 2)
	if !s.Has(1) {
		t.Error("NewIntSet() Add(1, 2, 1) Remove(2, 2) Has(1) got: false - want: true - expected set: {1}")
	}
	if s.Has(2) {
		t.Error("NewIntSet() Add(1, 2, 1) Remove(2, 2) Has(2) got: true - want: false - expected set: {1}")
	}
	if z := s.Size(); z != 1 {
		t.Errorf("NewIntSet() Add(1, 2, 1) Remove(2, 2) Size() got: %v - want: 1 - expected set: {1}", z)
	}
	if s.IsEmpty() {
		t.Error("NewIntSet() Add(1, 2, 1) Remove(2, 2) IsEmpty() got: true - want: false - expected set: {1}")
	}
	if a := s.AsSlice(); len(a) != 1 || a[0] != 1 {
		t.Errorf("NewIntSet() Add(1, 2, 1) Remove(2, 2) AsSlice() got: %v - want: [1] - expected set: {1}", a)
	}
	if r := s.String(); r != "{1}" {
		t.Errorf("NewIntSet() Add(1, 2, 1) Remove(2, 2) String() got: %v - want: {1} - expected set: {1}", r)
	}
	// NewIntSet(2, 1, 2) -> {1 2}
	s = NewIntSet(2, 1, 2)
	if !s.Has(1) {
		t.Error("NewIntSet(2, 1, 2) Has(1) got: false - want: true - expected set: {1 2}")
	}
	if !s.Has(2) {
		t.Error("NewIntSet(2, 1, 2) Has(2) got: false - want: true - expected set: {1 2}")
	}
	if z := s.Size(); z != 2 {
		t.Errorf("NewIntSet(2, 1, 2) Size() got: %v - want: 2 - expected set: {1 2}", z)
	}
	if s.IsEmpty() {
		t.Error("NewIntSet(2, 1, 2) IsEmpty() got: true - want: false - expected set: {1 2}")
	}
	if a := s.AsSlice(); len(a) != 2 || !((a[0] == 1 && a[1] == 2) || (a[0] == 2 && a[1] == 1)) {
		t.Errorf("NewIntSet(2, 1, 2) AsSlice() got: %v - want: [1 2] or [2 1] - expected set: {1 2}", a)
	}
	if r := s.String(); r != "{1 2}" && r != "{2 1}" {
		t.Errorf("NewIntSet(2, 1, 2) String() got: %v - want: {1 2} or {2 1} - expected set: {1 2}", r)
	}
	// NewIntSet(2, 1, 2) Clear() -> {}
	s.Clear()
	if !s.IsEmpty() {
		t.Error("NewIntSet(2, 1, 2) Clear() IsEmpty() got: false - want: true - expected set: {}")
	}
	// NewIntSet(2, 1, 2) Clear() Clear() -> {}
	s.Clear()
	if !s.IsEmpty() {
		t.Error("NewIntSet(2, 1, 2) Clear() Clear() IsEmpty() got: false - want: true - expected set: {}")
	}
}

func TestIntSetEquals(t *testing.T) {
	a := NewIntSet()
	b := NewIntSet()
	if !a.Equals(b) {
		t.Errorf("%v.Equals(%v) got: false - want: true", a, b)
	}
	a = NewIntSet(1, 2)
	b = NewIntSet(1, 2)
	if !a.Equals(b) {
		t.Errorf("%v.Equals(%v) got: false - want: true", a, b)
	}
	a = NewIntSet(1, 2)
	b = NewIntSet(2)
	if a.Equals(b) {
		t.Errorf("%v.Equals(%v) got: true - want: false", a, b)
	}
	a = NewIntSet(1)
	b = NewIntSet(2)
	if a.Equals(b) {
		t.Errorf("%v.Equals(%v) got: true - want: false", a, b)
	}
}

func TestIntSetUnion(t *testing.T) {
	a := NewIntSet()
	b := NewIntSet()
	w := NewIntSet()
	if g := a.Union(b); !g.Equals(w) {
		t.Errorf("%v.Union(%v) got: %v - want: %v", a, b, g, w)
	}
	a = NewIntSet(1, 2)
	b = NewIntSet(2, 3)
	w = NewIntSet(1, 2, 3)
	if g := a.Union(b); !g.Equals(w) {
		t.Errorf("%v.Union(%v) got: %v - want: %v", a, b, g, w)
	}
}

func TestIntSetIntersection(t *testing.T) {
	a := NewIntSet()
	b := NewIntSet()
	w := NewIntSet()
	if g := a.Intersection(b); !g.Equals(w) {
		t.Errorf("%v.Intersection(%v) got: %v - want: %v", a, b, g, w)
	}
	a = NewIntSet(1, 2, 3)
	b = NewIntSet(3, 4)
	w = NewIntSet(3)
	if g := a.Intersection(b); !g.Equals(w) {
		t.Errorf("%v.Intersection(%v) got: %v - want: %v", a, b, g, w)
	}
	if g := b.Intersection(a); !g.Equals(w) {
		t.Errorf("%v.Intersection(%v) got: %v - want: %v", b, a, g, w)
	}
}

func TestIntSetDisjoint(t *testing.T) {
	a := NewIntSet()
	b := NewIntSet()
	if !a.Disjoint(b) {
		t.Errorf("%v.Disjoint(%v) got: false - want: true", a, b)
	}
	a = NewIntSet(1, 2)
	b = NewIntSet(3, 4)
	if !a.Disjoint(b) {
		t.Errorf("%v.Disjoint(%v) got: false - want: true", a, b)
	}
	a = NewIntSet(1, 2)
	b = NewIntSet(2, 3)
	if a.Disjoint(b) {
		t.Errorf("%v.Disjoint(%v) got: true - want: false", a, b)
	}
}
