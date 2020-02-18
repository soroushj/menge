package sets

import "testing"

func TestIntSet(t *testing.T) {
	s := NewIntSet()
	if s.Has(1) {
		t.Error("Newly-created set has an element (1)")
	}
	if z := s.Size(); z != 0 {
		t.Errorf("Expected size to be 0, but got %v", z)
	}
	if !s.IsEmpty() {
		t.Error("IsEmpty() returned false for a newly-created set")
	}
	if a := s.AsSlice(); len(a) != 0 {
		t.Errorf("AsSlice() returned %v for an empty set", a)
	}
	if r := s.String(); r != "{}" {
		t.Errorf("String() returned %v for an empty set", r)
	}
	s.Add(1, 2, 1)
	if !s.Has(1) {
		t.Error("Set doesn't have an added element (1)")
	}
	if !s.Has(2) {
		t.Error("Set doesn't have an added element (2)")
	}
	if z := s.Size(); z != 2 {
		t.Errorf("Expected size to be 2, but got %v", z)
	}
	if s.IsEmpty() {
		t.Error("IsEmpty() returned true for a non-empty set")
	}
	if a := s.AsSlice(); len(a) != 2 || !((a[0] == 1 && a[1] == 2) || (a[0] == 2 && a[1] == 1)) {
		t.Errorf("AsSlice() returned %v for set {1 2}", a)
	}
	if r := s.String(); r != "{1 2}" && r != "{2 1}" {
		t.Errorf("String() returned %v for set {1 2}", r)
	}
	s.Remove(1, 2, 2)
	if s.Has(1) {
		t.Error("Set has a removed element (1)")
	}
	if s.Has(2) {
		t.Error("Set has a removed element (2)")
	}
	if z := s.Size(); z != 0 {
		t.Errorf("Expected size to be 0, but got %v", z)
	}
	if !s.IsEmpty() {
		t.Error("IsEmpty() returned false for an empty set")
	}
	s.Clear()
	if !s.IsEmpty() {
		t.Error("Cleared an empty set, but no longer is empty")
	}
	s.Add(1, 2, 3)
	s.Clear()
	if !s.IsEmpty() {
		t.Error("Cleared a set, but is not empty")
	}
}
