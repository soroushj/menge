package menge

import (
	"testing"
)

func TestNewUInt16Set(t *testing.T) {
	cases := []struct {
		arg  []uint16
		want UInt16Set
	}{
		{[]uint16{}, UInt16Set{}},
		{[]uint16{1, 1}, UInt16Set{1: struct{}{}}},
		{[]uint16{1, 2}, UInt16Set{1: struct{}{}, 2: struct{}{}}},
	}
	for _, c := range cases {
		got := NewUInt16Set(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt16Set_Add(t *testing.T) {
	cases := []struct {
		set  UInt16Set
		arg  []uint16
		want UInt16Set
	}{
		{NewUInt16Set(), []uint16{}, NewUInt16Set()},
		{NewUInt16Set(), []uint16{1, 1}, NewUInt16Set(1)},
		{NewUInt16Set(), []uint16{1, 2}, NewUInt16Set(1, 2)},
		{NewUInt16Set(1), []uint16{}, NewUInt16Set(1)},
		{NewUInt16Set(1), []uint16{1, 1}, NewUInt16Set(1)},
		{NewUInt16Set(1), []uint16{2, 3}, NewUInt16Set(1, 2, 3)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Add(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt16Set_Remove(t *testing.T) {
	cases := []struct {
		set  UInt16Set
		arg  []uint16
		want UInt16Set
	}{
		{NewUInt16Set(), []uint16{}, NewUInt16Set()},
		{NewUInt16Set(1), []uint16{1, 1}, NewUInt16Set()},
		{NewUInt16Set(1, 2), []uint16{1, 2}, NewUInt16Set()},
		{NewUInt16Set(1), []uint16{}, NewUInt16Set(1)},
		{NewUInt16Set(1), []uint16{1, 1}, NewUInt16Set()},
		{NewUInt16Set(1, 2), []uint16{3}, NewUInt16Set(1, 2)},
		{NewUInt16Set(1, 2, 3), []uint16{2, 3}, NewUInt16Set(1)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Remove(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt16Set_Empty(t *testing.T) {
	cases := []struct {
		set  UInt16Set
		want UInt16Set
	}{
		{NewUInt16Set(), NewUInt16Set()},
		{NewUInt16Set(1, 2), NewUInt16Set()},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Empty()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt16Set_Has(t *testing.T) {
	cases := []struct {
		set  UInt16Set
		arg  uint16
		want bool
	}{
		{NewUInt16Set(), 1, false},
		{NewUInt16Set(2), 1, false},
		{NewUInt16Set(1), 1, true},
		{NewUInt16Set(1, 2), 1, true},
	}
	for _, c := range cases {
		got := c.set.Has(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt16Set_Size(t *testing.T) {
	cases := []struct {
		set  UInt16Set
		want int
	}{
		{NewUInt16Set(), 0},
		{NewUInt16Set(1, 2), 2},
	}
	for _, c := range cases {
		got := c.set.Size()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt16Set_IsEmpty(t *testing.T) {
	cases := []struct {
		set  UInt16Set
		want bool
	}{
		{NewUInt16Set(), true},
		{NewUInt16Set(1, 2), false},
	}
	for _, c := range cases {
		got := c.set.IsEmpty()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt16Set_Clone(t *testing.T) {
	cases := []struct {
		set  UInt16Set
		want UInt16Set
	}{
		{NewUInt16Set(), NewUInt16Set()},
		{NewUInt16Set(1, 2), NewUInt16Set(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt16Set_AsSlice(t *testing.T) {
	cases := []struct {
		set  UInt16Set
		want []uint16
	}{
		{NewUInt16Set(), []uint16{}},
		{NewUInt16Set(1, 2), []uint16{1, 2}},
	}
	for _, c := range cases {
		got := c.set.AsSlice()
		if len(got) != len(c.want) || !NewUInt16Set(got...).Equals(NewUInt16Set(c.want...)) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt16Set_String(t *testing.T) {
	cases := []struct {
		set  UInt16Set
		want []string
	}{
		{NewUInt16Set(), []string{"{}"}},
		{NewUInt16Set(1), []string{"{1}"}},
		{NewUInt16Set(1, 2), []string{"{1 2}", "{2 1}"}},
	}
	contains := func(ss []string, s string) bool {
		for _, v := range ss {
			if v == s {
				return true
			}
		}
		return false
	}
	for _, c := range cases {
		got := c.set.String()
		if !contains(c.want, got) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt16Set_Equals(t *testing.T) {
	cases := []struct {
		set  UInt16Set
		arg  UInt16Set
		want bool
	}{
		{NewUInt16Set(), NewUInt16Set(), true},
		{NewUInt16Set(1, 2), NewUInt16Set(2, 1), true},
		{NewUInt16Set(1, 2), NewUInt16Set(1), false},
		{NewUInt16Set(1), NewUInt16Set(1, 2), false},
		{NewUInt16Set(1), NewUInt16Set(2), false},
	}
	for _, c := range cases {
		got := c.set.Equals(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt16Set_Union(t *testing.T) {
	cases := []struct {
		set  UInt16Set
		arg  UInt16Set
		want UInt16Set
	}{
		{NewUInt16Set(), NewUInt16Set(), NewUInt16Set()},
		{NewUInt16Set(1), NewUInt16Set(1), NewUInt16Set(1)},
		{NewUInt16Set(1), NewUInt16Set(2), NewUInt16Set(1, 2)},
		{NewUInt16Set(1), NewUInt16Set(1, 2), NewUInt16Set(1, 2)},
		{NewUInt16Set(1, 2), NewUInt16Set(1), NewUInt16Set(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Union(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt16Set_Intersection(t *testing.T) {
	cases := []struct {
		set  UInt16Set
		arg  UInt16Set
		want UInt16Set
	}{
		{NewUInt16Set(), NewUInt16Set(), NewUInt16Set()},
		{NewUInt16Set(1), NewUInt16Set(1), NewUInt16Set(1)},
		{NewUInt16Set(1), NewUInt16Set(2), NewUInt16Set()},
		{NewUInt16Set(1), NewUInt16Set(1, 2), NewUInt16Set(1)},
		{NewUInt16Set(1, 2), NewUInt16Set(1), NewUInt16Set(1)},
	}
	for _, c := range cases {
		got := c.set.Intersection(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt16Set_Difference(t *testing.T) {
	cases := []struct {
		set  UInt16Set
		arg  UInt16Set
		want UInt16Set
	}{
		{NewUInt16Set(), NewUInt16Set(), NewUInt16Set()},
		{NewUInt16Set(1), NewUInt16Set(1), NewUInt16Set()},
		{NewUInt16Set(1), NewUInt16Set(2), NewUInt16Set(1)},
		{NewUInt16Set(1), NewUInt16Set(1, 2), NewUInt16Set()},
		{NewUInt16Set(1, 2), NewUInt16Set(1), NewUInt16Set(2)},
	}
	for _, c := range cases {
		got := c.set.Difference(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt16Set_IsSubsetOf(t *testing.T) {
	cases := []struct {
		set  UInt16Set
		arg  UInt16Set
		want bool
	}{
		{NewUInt16Set(), NewUInt16Set(), true},
		{NewUInt16Set(1), NewUInt16Set(1), true},
		{NewUInt16Set(1), NewUInt16Set(1, 2), true},
		{NewUInt16Set(1, 2), NewUInt16Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt16Set_IsProperSubsetOf(t *testing.T) {
	cases := []struct {
		set  UInt16Set
		arg  UInt16Set
		want bool
	}{
		{NewUInt16Set(), NewUInt16Set(), false},
		{NewUInt16Set(1), NewUInt16Set(1), false},
		{NewUInt16Set(1), NewUInt16Set(1, 2), true},
		{NewUInt16Set(1, 2), NewUInt16Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsProperSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt16Set_IsSupersetOf(t *testing.T) {
	cases := []struct {
		set  UInt16Set
		arg  UInt16Set
		want bool
	}{
		{NewUInt16Set(), NewUInt16Set(), true},
		{NewUInt16Set(1), NewUInt16Set(1), true},
		{NewUInt16Set(1), NewUInt16Set(1, 2), false},
		{NewUInt16Set(1, 2), NewUInt16Set(1), true},
	}
	for _, c := range cases {
		got := c.set.IsSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt16Set_IsProperSupersetOf(t *testing.T) {
	cases := []struct {
		set  UInt16Set
		arg  UInt16Set
		want bool
	}{
		{NewUInt16Set(), NewUInt16Set(), false},
		{NewUInt16Set(1), NewUInt16Set(1), false},
		{NewUInt16Set(1), NewUInt16Set(1, 2), false},
		{NewUInt16Set(1, 2), NewUInt16Set(1), true},
	}
	for _, c := range cases {
		got := c.set.IsProperSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt16Set_IsDisjointFrom(t *testing.T) {
	cases := []struct {
		set  UInt16Set
		arg  UInt16Set
		want bool
	}{
		{NewUInt16Set(), NewUInt16Set(), true},
		{NewUInt16Set(1), NewUInt16Set(1), false},
		{NewUInt16Set(1), NewUInt16Set(2, 3), true},
		{NewUInt16Set(1, 2), NewUInt16Set(3), true},
		{NewUInt16Set(1), NewUInt16Set(1, 2), false},
		{NewUInt16Set(1, 2), NewUInt16Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsDisjointFrom(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
