package menge

import (
	"testing"
)

func TestNewUInt32Set(t *testing.T) {
	cases := []struct {
		arg  []uint32
		want UInt32Set
	}{
		{[]uint32{}, UInt32Set{}},
		{[]uint32{1, 1}, UInt32Set{1: struct{}{}}},
		{[]uint32{1, 2}, UInt32Set{1: struct{}{}, 2: struct{}{}}},
	}
	for _, c := range cases {
		got := NewUInt32Set(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt32Set_Add(t *testing.T) {
	cases := []struct {
		set  UInt32Set
		arg  []uint32
		want UInt32Set
	}{
		{NewUInt32Set(), []uint32{}, NewUInt32Set()},
		{NewUInt32Set(), []uint32{1, 1}, NewUInt32Set(1)},
		{NewUInt32Set(), []uint32{1, 2}, NewUInt32Set(1, 2)},
		{NewUInt32Set(1), []uint32{}, NewUInt32Set(1)},
		{NewUInt32Set(1), []uint32{1, 1}, NewUInt32Set(1)},
		{NewUInt32Set(1), []uint32{2, 3}, NewUInt32Set(1, 2, 3)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Add(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt32Set_Remove(t *testing.T) {
	cases := []struct {
		set  UInt32Set
		arg  []uint32
		want UInt32Set
	}{
		{NewUInt32Set(), []uint32{}, NewUInt32Set()},
		{NewUInt32Set(1), []uint32{1, 1}, NewUInt32Set()},
		{NewUInt32Set(1, 2), []uint32{1, 2}, NewUInt32Set()},
		{NewUInt32Set(1), []uint32{}, NewUInt32Set(1)},
		{NewUInt32Set(1), []uint32{1, 1}, NewUInt32Set()},
		{NewUInt32Set(1, 2), []uint32{3}, NewUInt32Set(1, 2)},
		{NewUInt32Set(1, 2, 3), []uint32{2, 3}, NewUInt32Set(1)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Remove(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt32Set_Empty(t *testing.T) {
	cases := []struct {
		set  UInt32Set
		want UInt32Set
	}{
		{NewUInt32Set(), NewUInt32Set()},
		{NewUInt32Set(1, 2), NewUInt32Set()},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Empty()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt32Set_Has(t *testing.T) {
	cases := []struct {
		set  UInt32Set
		arg  uint32
		want bool
	}{
		{NewUInt32Set(), 1, false},
		{NewUInt32Set(2), 1, false},
		{NewUInt32Set(1), 1, true},
		{NewUInt32Set(1, 2), 1, true},
	}
	for _, c := range cases {
		got := c.set.Has(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt32Set_Size(t *testing.T) {
	cases := []struct {
		set  UInt32Set
		want int
	}{
		{NewUInt32Set(), 0},
		{NewUInt32Set(1, 2), 2},
	}
	for _, c := range cases {
		got := c.set.Size()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt32Set_IsEmpty(t *testing.T) {
	cases := []struct {
		set  UInt32Set
		want bool
	}{
		{NewUInt32Set(), true},
		{NewUInt32Set(1, 2), false},
	}
	for _, c := range cases {
		got := c.set.IsEmpty()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt32Set_Clone(t *testing.T) {
	cases := []struct {
		set  UInt32Set
		want UInt32Set
	}{
		{NewUInt32Set(), NewUInt32Set()},
		{NewUInt32Set(1, 2), NewUInt32Set(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt32Set_AsSlice(t *testing.T) {
	cases := []struct {
		set  UInt32Set
		want []uint32
	}{
		{NewUInt32Set(), []uint32{}},
		{NewUInt32Set(1, 2), []uint32{1, 2}},
	}
	for _, c := range cases {
		got := c.set.AsSlice()
		if len(got) != len(c.want) || !NewUInt32Set(got...).Equals(NewUInt32Set(c.want...)) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt32Set_String(t *testing.T) {
	cases := []struct {
		set  UInt32Set
		want []string
	}{
		{NewUInt32Set(), []string{"{}"}},
		{NewUInt32Set(1), []string{"{1}"}},
		{NewUInt32Set(1, 2), []string{"{1 2}", "{2 1}"}},
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

func TestUInt32Set_Equals(t *testing.T) {
	cases := []struct {
		set  UInt32Set
		arg  UInt32Set
		want bool
	}{
		{NewUInt32Set(), NewUInt32Set(), true},
		{NewUInt32Set(1, 2), NewUInt32Set(2, 1), true},
		{NewUInt32Set(1, 2), NewUInt32Set(1), false},
		{NewUInt32Set(1), NewUInt32Set(1, 2), false},
		{NewUInt32Set(1), NewUInt32Set(2), false},
	}
	for _, c := range cases {
		got := c.set.Equals(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt32Set_Union(t *testing.T) {
	cases := []struct {
		set  UInt32Set
		arg  UInt32Set
		want UInt32Set
	}{
		{NewUInt32Set(), NewUInt32Set(), NewUInt32Set()},
		{NewUInt32Set(1), NewUInt32Set(1), NewUInt32Set(1)},
		{NewUInt32Set(1), NewUInt32Set(2), NewUInt32Set(1, 2)},
		{NewUInt32Set(1), NewUInt32Set(1, 2), NewUInt32Set(1, 2)},
		{NewUInt32Set(1, 2), NewUInt32Set(1), NewUInt32Set(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Union(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt32Set_Intersection(t *testing.T) {
	cases := []struct {
		set  UInt32Set
		arg  UInt32Set
		want UInt32Set
	}{
		{NewUInt32Set(), NewUInt32Set(), NewUInt32Set()},
		{NewUInt32Set(1), NewUInt32Set(1), NewUInt32Set(1)},
		{NewUInt32Set(1), NewUInt32Set(2), NewUInt32Set()},
		{NewUInt32Set(1), NewUInt32Set(1, 2), NewUInt32Set(1)},
		{NewUInt32Set(1, 2), NewUInt32Set(1), NewUInt32Set(1)},
	}
	for _, c := range cases {
		got := c.set.Intersection(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt32Set_Difference(t *testing.T) {
	cases := []struct {
		set  UInt32Set
		arg  UInt32Set
		want UInt32Set
	}{
		{NewUInt32Set(), NewUInt32Set(), NewUInt32Set()},
		{NewUInt32Set(1), NewUInt32Set(1), NewUInt32Set()},
		{NewUInt32Set(1), NewUInt32Set(2), NewUInt32Set(1)},
		{NewUInt32Set(1), NewUInt32Set(1, 2), NewUInt32Set()},
		{NewUInt32Set(1, 2), NewUInt32Set(1), NewUInt32Set(2)},
	}
	for _, c := range cases {
		got := c.set.Difference(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt32Set_IsSubsetOf(t *testing.T) {
	cases := []struct {
		set  UInt32Set
		arg  UInt32Set
		want bool
	}{
		{NewUInt32Set(), NewUInt32Set(), true},
		{NewUInt32Set(1), NewUInt32Set(1), true},
		{NewUInt32Set(1), NewUInt32Set(1, 2), true},
		{NewUInt32Set(1, 2), NewUInt32Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt32Set_IsProperSubsetOf(t *testing.T) {
	cases := []struct {
		set  UInt32Set
		arg  UInt32Set
		want bool
	}{
		{NewUInt32Set(), NewUInt32Set(), false},
		{NewUInt32Set(1), NewUInt32Set(1), false},
		{NewUInt32Set(1), NewUInt32Set(1, 2), true},
		{NewUInt32Set(1, 2), NewUInt32Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsProperSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt32Set_IsSupersetOf(t *testing.T) {
	cases := []struct {
		set  UInt32Set
		arg  UInt32Set
		want bool
	}{
		{NewUInt32Set(), NewUInt32Set(), true},
		{NewUInt32Set(1), NewUInt32Set(1), true},
		{NewUInt32Set(1), NewUInt32Set(1, 2), false},
		{NewUInt32Set(1, 2), NewUInt32Set(1), true},
	}
	for _, c := range cases {
		got := c.set.IsSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt32Set_IsProperSupersetOf(t *testing.T) {
	cases := []struct {
		set  UInt32Set
		arg  UInt32Set
		want bool
	}{
		{NewUInt32Set(), NewUInt32Set(), false},
		{NewUInt32Set(1), NewUInt32Set(1), false},
		{NewUInt32Set(1), NewUInt32Set(1, 2), false},
		{NewUInt32Set(1, 2), NewUInt32Set(1), true},
	}
	for _, c := range cases {
		got := c.set.IsProperSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt32Set_IsDisjointFrom(t *testing.T) {
	cases := []struct {
		set  UInt32Set
		arg  UInt32Set
		want bool
	}{
		{NewUInt32Set(), NewUInt32Set(), true},
		{NewUInt32Set(1), NewUInt32Set(1), false},
		{NewUInt32Set(1), NewUInt32Set(2, 3), true},
		{NewUInt32Set(1, 2), NewUInt32Set(3), true},
		{NewUInt32Set(1), NewUInt32Set(1, 2), false},
		{NewUInt32Set(1, 2), NewUInt32Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsDisjointFrom(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
