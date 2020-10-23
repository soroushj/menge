package menge

import (
	"testing"
)

func TestNewUInt8Set(t *testing.T) {
	cases := []struct {
		arg  []uint8
		want UInt8Set
	}{
		{[]uint8{}, UInt8Set{}},
		{[]uint8{1, 1}, UInt8Set{1: struct{}{}}},
		{[]uint8{1, 2}, UInt8Set{1: struct{}{}, 2: struct{}{}}},
	}
	for _, c := range cases {
		got := NewUInt8Set(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt8Set_Add(t *testing.T) {
	cases := []struct {
		set  UInt8Set
		arg  []uint8
		want UInt8Set
	}{
		{NewUInt8Set(), []uint8{}, NewUInt8Set()},
		{NewUInt8Set(), []uint8{1, 1}, NewUInt8Set(1)},
		{NewUInt8Set(), []uint8{1, 2}, NewUInt8Set(1, 2)},
		{NewUInt8Set(1), []uint8{}, NewUInt8Set(1)},
		{NewUInt8Set(1), []uint8{1, 1}, NewUInt8Set(1)},
		{NewUInt8Set(1), []uint8{2, 3}, NewUInt8Set(1, 2, 3)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Add(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt8Set_Remove(t *testing.T) {
	cases := []struct {
		set  UInt8Set
		arg  []uint8
		want UInt8Set
	}{
		{NewUInt8Set(), []uint8{}, NewUInt8Set()},
		{NewUInt8Set(1), []uint8{1, 1}, NewUInt8Set()},
		{NewUInt8Set(1, 2), []uint8{1, 2}, NewUInt8Set()},
		{NewUInt8Set(1), []uint8{}, NewUInt8Set(1)},
		{NewUInt8Set(1), []uint8{1, 1}, NewUInt8Set()},
		{NewUInt8Set(1, 2), []uint8{3}, NewUInt8Set(1, 2)},
		{NewUInt8Set(1, 2, 3), []uint8{2, 3}, NewUInt8Set(1)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Remove(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt8Set_Empty(t *testing.T) {
	cases := []struct {
		set  UInt8Set
		want UInt8Set
	}{
		{NewUInt8Set(), NewUInt8Set()},
		{NewUInt8Set(1, 2), NewUInt8Set()},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Empty()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt8Set_Has(t *testing.T) {
	cases := []struct {
		set  UInt8Set
		arg  uint8
		want bool
	}{
		{NewUInt8Set(), 1, false},
		{NewUInt8Set(2), 1, false},
		{NewUInt8Set(1), 1, true},
		{NewUInt8Set(1, 2), 1, true},
	}
	for _, c := range cases {
		got := c.set.Has(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt8Set_Size(t *testing.T) {
	cases := []struct {
		set  UInt8Set
		want int
	}{
		{NewUInt8Set(), 0},
		{NewUInt8Set(1, 2), 2},
	}
	for _, c := range cases {
		got := c.set.Size()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt8Set_IsEmpty(t *testing.T) {
	cases := []struct {
		set  UInt8Set
		want bool
	}{
		{NewUInt8Set(), true},
		{NewUInt8Set(1, 2), false},
	}
	for _, c := range cases {
		got := c.set.IsEmpty()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt8Set_Clone(t *testing.T) {
	cases := []struct {
		set  UInt8Set
		want UInt8Set
	}{
		{NewUInt8Set(), NewUInt8Set()},
		{NewUInt8Set(1, 2), NewUInt8Set(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt8Set_AsSlice(t *testing.T) {
	cases := []struct {
		set  UInt8Set
		want []uint8
	}{
		{NewUInt8Set(), []uint8{}},
		{NewUInt8Set(1, 2), []uint8{1, 2}},
	}
	for _, c := range cases {
		got := c.set.AsSlice()
		if len(got) != len(c.want) || !NewUInt8Set(got...).Equals(NewUInt8Set(c.want...)) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt8Set_String(t *testing.T) {
	cases := []struct {
		set  UInt8Set
		want []string
	}{
		{NewUInt8Set(), []string{"{}"}},
		{NewUInt8Set(1), []string{"{1}"}},
		{NewUInt8Set(1, 2), []string{"{1 2}", "{2 1}"}},
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

func TestUInt8Set_Equals(t *testing.T) {
	cases := []struct {
		set  UInt8Set
		arg  UInt8Set
		want bool
	}{
		{NewUInt8Set(), NewUInt8Set(), true},
		{NewUInt8Set(1, 2), NewUInt8Set(2, 1), true},
		{NewUInt8Set(1, 2), NewUInt8Set(1), false},
		{NewUInt8Set(1), NewUInt8Set(1, 2), false},
		{NewUInt8Set(1), NewUInt8Set(2), false},
	}
	for _, c := range cases {
		got := c.set.Equals(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt8Set_Union(t *testing.T) {
	cases := []struct {
		set  UInt8Set
		arg  UInt8Set
		want UInt8Set
	}{
		{NewUInt8Set(), NewUInt8Set(), NewUInt8Set()},
		{NewUInt8Set(1), NewUInt8Set(1), NewUInt8Set(1)},
		{NewUInt8Set(1), NewUInt8Set(2), NewUInt8Set(1, 2)},
		{NewUInt8Set(1), NewUInt8Set(1, 2), NewUInt8Set(1, 2)},
		{NewUInt8Set(1, 2), NewUInt8Set(1), NewUInt8Set(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Union(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt8Set_Intersection(t *testing.T) {
	cases := []struct {
		set  UInt8Set
		arg  UInt8Set
		want UInt8Set
	}{
		{NewUInt8Set(), NewUInt8Set(), NewUInt8Set()},
		{NewUInt8Set(1), NewUInt8Set(1), NewUInt8Set(1)},
		{NewUInt8Set(1), NewUInt8Set(2), NewUInt8Set()},
		{NewUInt8Set(1), NewUInt8Set(1, 2), NewUInt8Set(1)},
		{NewUInt8Set(1, 2), NewUInt8Set(1), NewUInt8Set(1)},
	}
	for _, c := range cases {
		got := c.set.Intersection(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt8Set_Difference(t *testing.T) {
	cases := []struct {
		set  UInt8Set
		arg  UInt8Set
		want UInt8Set
	}{
		{NewUInt8Set(), NewUInt8Set(), NewUInt8Set()},
		{NewUInt8Set(1), NewUInt8Set(1), NewUInt8Set()},
		{NewUInt8Set(1), NewUInt8Set(2), NewUInt8Set(1)},
		{NewUInt8Set(1), NewUInt8Set(1, 2), NewUInt8Set()},
		{NewUInt8Set(1, 2), NewUInt8Set(1), NewUInt8Set(2)},
	}
	for _, c := range cases {
		got := c.set.Difference(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt8Set_IsSubsetOf(t *testing.T) {
	cases := []struct {
		set  UInt8Set
		arg  UInt8Set
		want bool
	}{
		{NewUInt8Set(), NewUInt8Set(), true},
		{NewUInt8Set(1), NewUInt8Set(1), true},
		{NewUInt8Set(1), NewUInt8Set(1, 2), true},
		{NewUInt8Set(1, 2), NewUInt8Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt8Set_IsProperSubsetOf(t *testing.T) {
	cases := []struct {
		set  UInt8Set
		arg  UInt8Set
		want bool
	}{
		{NewUInt8Set(), NewUInt8Set(), false},
		{NewUInt8Set(1), NewUInt8Set(1), false},
		{NewUInt8Set(1), NewUInt8Set(1, 2), true},
		{NewUInt8Set(1, 2), NewUInt8Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsProperSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt8Set_IsSupersetOf(t *testing.T) {
	cases := []struct {
		set  UInt8Set
		arg  UInt8Set
		want bool
	}{
		{NewUInt8Set(), NewUInt8Set(), true},
		{NewUInt8Set(1), NewUInt8Set(1), true},
		{NewUInt8Set(1), NewUInt8Set(1, 2), false},
		{NewUInt8Set(1, 2), NewUInt8Set(1), true},
	}
	for _, c := range cases {
		got := c.set.IsSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt8Set_IsProperSupersetOf(t *testing.T) {
	cases := []struct {
		set  UInt8Set
		arg  UInt8Set
		want bool
	}{
		{NewUInt8Set(), NewUInt8Set(), false},
		{NewUInt8Set(1), NewUInt8Set(1), false},
		{NewUInt8Set(1), NewUInt8Set(1, 2), false},
		{NewUInt8Set(1, 2), NewUInt8Set(1), true},
	}
	for _, c := range cases {
		got := c.set.IsProperSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt8Set_IsDisjointFrom(t *testing.T) {
	cases := []struct {
		set  UInt8Set
		arg  UInt8Set
		want bool
	}{
		{NewUInt8Set(), NewUInt8Set(), true},
		{NewUInt8Set(1), NewUInt8Set(1), false},
		{NewUInt8Set(1), NewUInt8Set(2, 3), true},
		{NewUInt8Set(1, 2), NewUInt8Set(3), true},
		{NewUInt8Set(1), NewUInt8Set(1, 2), false},
		{NewUInt8Set(1, 2), NewUInt8Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsDisjointFrom(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
