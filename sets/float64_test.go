package sets

import "testing"

func TestFloat64SetBasics(t *testing.T) {
	// NewFloat64Set() -> {}
	s := NewFloat64Set()
	if s.Has(1) {
		t.Error("NewFloat64Set() Has(1) got: true - want: false - expected set: {}")
	}
	if z := s.Size(); z != 0 {
		t.Errorf("NewFloat64Set() Size() got: %v - want: 0 - expected set: {}", z)
	}
	if !s.IsEmpty() {
		t.Error("NewFloat64Set() IsEmpty() got: false - want: true - expected set: {}")
	}
	if a := s.AsSlice(); len(a) != 0 {
		t.Errorf("NewFloat64Set() AsSlice() got: %v - want: [] - expected set: {}", a)
	}
	if r := s.String(); r != "{}" {
		t.Errorf("NewFloat64Set() String() got: %v - want: {} - expected set: {}", r)
	}
	// NewFloat64Set() Add(1, 2, 1) -> {1 2}
	s.Add(1, 2, 1)
	if !s.Has(1) {
		t.Error("NewFloat64Set() Add(1, 2, 1) Has(1) got: false - want: true - expected set: {1 2}")
	}
	if !s.Has(2) {
		t.Error("NewFloat64Set() Add(1, 2, 1) Has(2) got: false - want: true - expected set: {1 2}")
	}
	if z := s.Size(); z != 2 {
		t.Errorf("NewFloat64Set() Add(1, 2, 1) Size() got: %v - want: 2 - expected set: {1 2}", z)
	}
	if s.IsEmpty() {
		t.Error("NewFloat64Set() Add(1, 2, 1) IsEmpty() got: true - want: false - expected set: {1 2}")
	}
	if a := s.AsSlice(); len(a) != 2 || !((a[0] == 1 && a[1] == 2) || (a[0] == 2 && a[1] == 1)) {
		t.Errorf("NewFloat64Set() Add(1, 2, 1) AsSlice() got: %v - want: [1 2] or [2 1] - expected set: {1 2}", a)
	}
	if r := s.String(); r != "{1 2}" && r != "{2 1}" {
		t.Errorf("NewFloat64Set() Add(1, 2, 1) String() got: %v - want: {1 2} or {2 1} - expected set: {1 2}", r)
	}
	// NewFloat64Set() Add(1, 2, 1) Remove(2, 2) -> {1}
	s.Remove(2, 2)
	if !s.Has(1) {
		t.Error("NewFloat64Set() Add(1, 2, 1) Remove(2, 2) Has(1) got: false - want: true - expected set: {1}")
	}
	if s.Has(2) {
		t.Error("NewFloat64Set() Add(1, 2, 1) Remove(2, 2) Has(2) got: true - want: false - expected set: {1}")
	}
	if z := s.Size(); z != 1 {
		t.Errorf("NewFloat64Set() Add(1, 2, 1) Remove(2, 2) Size() got: %v - want: 1 - expected set: {1}", z)
	}
	if s.IsEmpty() {
		t.Error("NewFloat64Set() Add(1, 2, 1) Remove(2, 2) IsEmpty() got: true - want: false - expected set: {1}")
	}
	if a := s.AsSlice(); len(a) != 1 || a[0] != 1 {
		t.Errorf("NewFloat64Set() Add(1, 2, 1) Remove(2, 2) AsSlice() got: %v - want: [1] - expected set: {1}", a)
	}
	if r := s.String(); r != "{1}" {
		t.Errorf("NewFloat64Set() Add(1, 2, 1) Remove(2, 2) String() got: %v - want: {1} - expected set: {1}", r)
	}
	// NewFloat64Set(2, 1, 2) -> {1 2}
	s = NewFloat64Set(2, 1, 2)
	if !s.Has(1) {
		t.Error("NewFloat64Set(2, 1, 2) Has(1) got: false - want: true - expected set: {1 2}")
	}
	if !s.Has(2) {
		t.Error("NewFloat64Set(2, 1, 2) Has(2) got: false - want: true - expected set: {1 2}")
	}
	if z := s.Size(); z != 2 {
		t.Errorf("NewFloat64Set(2, 1, 2) Size() got: %v - want: 2 - expected set: {1 2}", z)
	}
	if s.IsEmpty() {
		t.Error("NewFloat64Set(2, 1, 2) IsEmpty() got: true - want: false - expected set: {1 2}")
	}
	if a := s.AsSlice(); len(a) != 2 || !((a[0] == 1 && a[1] == 2) || (a[0] == 2 && a[1] == 1)) {
		t.Errorf("NewFloat64Set(2, 1, 2) AsSlice() got: %v - want: [1 2] or [2 1] - expected set: {1 2}", a)
	}
	if r := s.String(); r != "{1 2}" && r != "{2 1}" {
		t.Errorf("NewFloat64Set(2, 1, 2) String() got: %v - want: {1 2} or {2 1} - expected set: {1 2}", r)
	}
	// NewFloat64Set(2, 1, 2) Empty() -> {}
	s.Empty()
	if !s.IsEmpty() {
		t.Error("NewFloat64Set(2, 1, 2) Empty() IsEmpty() got: false - want: true - expected set: {}")
	}
	// NewFloat64Set(2, 1, 2) Empty() Empty() -> {}
	s.Empty()
	if !s.IsEmpty() {
		t.Error("NewFloat64Set(2, 1, 2) Empty() Empty() IsEmpty() got: false - want: true - expected set: {}")
	}
}

func TestFloat64SetEquals(t *testing.T) {
	a := NewFloat64Set()
	b := NewFloat64Set()
	if !a.Equals(b) {
		t.Errorf("%v.Equals(%v) got: false - want: true", a, b)
	}
	a = NewFloat64Set(1, 2)
	b = NewFloat64Set(1, 2)
	if !a.Equals(b) {
		t.Errorf("%v.Equals(%v) got: false - want: true", a, b)
	}
	a = NewFloat64Set(1, 2)
	b = NewFloat64Set(2)
	if a.Equals(b) {
		t.Errorf("%v.Equals(%v) got: true - want: false", a, b)
	}
	a = NewFloat64Set(1)
	b = NewFloat64Set(2)
	if a.Equals(b) {
		t.Errorf("%v.Equals(%v) got: true - want: false", a, b)
	}
}

func TestFloat64SetUnion(t *testing.T) {
	a := NewFloat64Set()
	b := NewFloat64Set()
	w := NewFloat64Set()
	if g := a.Union(b); !g.Equals(w) {
		t.Errorf("%v.Union(%v) got: %v - want: %v", a, b, g, w)
	}
	a = NewFloat64Set(1, 2)
	b = NewFloat64Set(2, 3)
	w = NewFloat64Set(1, 2, 3)
	if g := a.Union(b); !g.Equals(w) {
		t.Errorf("%v.Union(%v) got: %v - want: %v", a, b, g, w)
	}
}

func TestFloat64SetIntersection(t *testing.T) {
	a := NewFloat64Set()
	b := NewFloat64Set()
	w := NewFloat64Set()
	if g := a.Intersection(b); !g.Equals(w) {
		t.Errorf("%v.Intersection(%v) got: %v - want: %v", a, b, g, w)
	}
	a = NewFloat64Set(1, 2, 3)
	b = NewFloat64Set(3, 4)
	w = NewFloat64Set(3)
	if g := a.Intersection(b); !g.Equals(w) {
		t.Errorf("%v.Intersection(%v) got: %v - want: %v", a, b, g, w)
	}
	if g := b.Intersection(a); !g.Equals(w) {
		t.Errorf("%v.Intersection(%v) got: %v - want: %v", b, a, g, w)
	}
}

func TestFloat64SetDifference(t *testing.T) {
	a := NewFloat64Set()
	b := NewFloat64Set()
	w := NewFloat64Set()
	if g := a.Difference(b); !g.Equals(w) {
		t.Errorf("%v.Difference(%v) got: %v - want: %v", a, b, g, w)
	}
	a = NewFloat64Set(1, 2)
	b = NewFloat64Set(2, 3)
	w = NewFloat64Set(1)
	if g := a.Difference(b); !g.Equals(w) {
		t.Errorf("%v.Difference(%v) got: %v - want: %v", a, b, g, w)
	}
}

func TestFloat64SetIsSubsetOf(t *testing.T) {
	a := NewFloat64Set()
	b := NewFloat64Set()
	if !a.IsSubsetOf(b) {
		t.Errorf("%v.IsSubsetOf(%v) got: false - want: true", a, b)
	}
	a = NewFloat64Set(1, 2)
	b = NewFloat64Set(1, 2, 3)
	if !a.IsSubsetOf(b) {
		t.Errorf("%v.IsSubsetOf(%v) got: false - want: true", a, b)
	}
	if b.IsSubsetOf(a) {
		t.Errorf("%v.IsSubsetOf(%v) got: true - want: false", b, a)
	}
}

func TestFloat64SetIsDisjointFrom(t *testing.T) {
	a := NewFloat64Set()
	b := NewFloat64Set()
	if !a.IsDisjointFrom(b) {
		t.Errorf("%v.IsDisjointFrom(%v) got: false - want: true", a, b)
	}
	a = NewFloat64Set(1, 2)
	b = NewFloat64Set(3, 4, 5)
	if !a.IsDisjointFrom(b) {
		t.Errorf("%v.IsDisjointFrom(%v) got: false - want: true", a, b)
	}
	if !b.IsDisjointFrom(a) {
		t.Errorf("%v.IsDisjointFrom(%v) got: false - want: true", b, a)
	}
	a = NewFloat64Set(1, 2)
	b = NewFloat64Set(2, 3, 4)
	if a.IsDisjointFrom(b) {
		t.Errorf("%v.IsDisjointFrom(%v) got: true - want: false", a, b)
	}
	if b.IsDisjointFrom(a) {
		t.Errorf("%v.IsDisjointFrom(%v) got: true - want: false", b, a)
	}
}
