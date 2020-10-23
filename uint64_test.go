package menge

import (
	"testing"
)

func TestNewUInt64Set(t *testing.T) {
	cases := []struct {
		arg  []uint64
		want UInt64Set
	}{
		{[]uint64{}, UInt64Set{}},
		{[]uint64{1, 1}, UInt64Set{1: struct{}{}}},
		{[]uint64{1, 2}, UInt64Set{1: struct{}{}, 2: struct{}{}}},
	}
	for _, c := range cases {
		got := NewUInt64Set(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt64Set_Add(t *testing.T) {
	cases := []struct {
		set  UInt64Set
		arg  []uint64
		want UInt64Set
	}{
		{NewUInt64Set(), []uint64{}, NewUInt64Set()},
		{NewUInt64Set(), []uint64{1, 1}, NewUInt64Set(1)},
		{NewUInt64Set(), []uint64{1, 2}, NewUInt64Set(1, 2)},
		{NewUInt64Set(1), []uint64{}, NewUInt64Set(1)},
		{NewUInt64Set(1), []uint64{1, 1}, NewUInt64Set(1)},
		{NewUInt64Set(1), []uint64{2, 3}, NewUInt64Set(1, 2, 3)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Add(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt64Set_Remove(t *testing.T) {
	cases := []struct {
		set  UInt64Set
		arg  []uint64
		want UInt64Set
	}{
		{NewUInt64Set(), []uint64{}, NewUInt64Set()},
		{NewUInt64Set(1), []uint64{1, 1}, NewUInt64Set()},
		{NewUInt64Set(1, 2), []uint64{1, 2}, NewUInt64Set()},
		{NewUInt64Set(1), []uint64{}, NewUInt64Set(1)},
		{NewUInt64Set(1), []uint64{1, 1}, NewUInt64Set()},
		{NewUInt64Set(1, 2), []uint64{3}, NewUInt64Set(1, 2)},
		{NewUInt64Set(1, 2, 3), []uint64{2, 3}, NewUInt64Set(1)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Remove(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt64Set_Empty(t *testing.T) {
	cases := []struct {
		set  UInt64Set
		want UInt64Set
	}{
		{NewUInt64Set(), NewUInt64Set()},
		{NewUInt64Set(1, 2), NewUInt64Set()},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Empty()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt64Set_Has(t *testing.T) {
	cases := []struct {
		set  UInt64Set
		arg  uint64
		want bool
	}{
		{NewUInt64Set(), 1, false},
		{NewUInt64Set(2), 1, false},
		{NewUInt64Set(1), 1, true},
		{NewUInt64Set(1, 2), 1, true},
	}
	for _, c := range cases {
		got := c.set.Has(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt64Set_Size(t *testing.T) {
	cases := []struct {
		set  UInt64Set
		want int
	}{
		{NewUInt64Set(), 0},
		{NewUInt64Set(1, 2), 2},
	}
	for _, c := range cases {
		got := c.set.Size()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt64Set_IsEmpty(t *testing.T) {
	cases := []struct {
		set  UInt64Set
		want bool
	}{
		{NewUInt64Set(), true},
		{NewUInt64Set(1, 2), false},
	}
	for _, c := range cases {
		got := c.set.IsEmpty()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt64Set_Clone(t *testing.T) {
	cases := []struct {
		set  UInt64Set
		want UInt64Set
	}{
		{NewUInt64Set(), NewUInt64Set()},
		{NewUInt64Set(1, 2), NewUInt64Set(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt64Set_AsSlice(t *testing.T) {
	cases := []struct {
		set  UInt64Set
		want []uint64
	}{
		{NewUInt64Set(), []uint64{}},
		{NewUInt64Set(1, 2), []uint64{1, 2}},
	}
	for _, c := range cases {
		got := c.set.AsSlice()
		if len(got) != len(c.want) || !NewUInt64Set(got...).Equals(NewUInt64Set(c.want...)) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt64Set_String(t *testing.T) {
	cases := []struct {
		set  UInt64Set
		want []string
	}{
		{NewUInt64Set(), []string{"{}"}},
		{NewUInt64Set(1), []string{"{1}"}},
		{NewUInt64Set(1, 2), []string{"{1 2}", "{2 1}"}},
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

func TestUInt64Set_Equals(t *testing.T) {
	cases := []struct {
		set  UInt64Set
		arg  UInt64Set
		want bool
	}{
		{NewUInt64Set(), NewUInt64Set(), true},
		{NewUInt64Set(1, 2), NewUInt64Set(2, 1), true},
		{NewUInt64Set(1, 2), NewUInt64Set(1), false},
		{NewUInt64Set(1), NewUInt64Set(1, 2), false},
		{NewUInt64Set(1), NewUInt64Set(2), false},
	}
	for _, c := range cases {
		got := c.set.Equals(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt64Set_Union(t *testing.T) {
	cases := []struct {
		set  UInt64Set
		arg  UInt64Set
		want UInt64Set
	}{
		{NewUInt64Set(), NewUInt64Set(), NewUInt64Set()},
		{NewUInt64Set(1), NewUInt64Set(1), NewUInt64Set(1)},
		{NewUInt64Set(1), NewUInt64Set(2), NewUInt64Set(1, 2)},
		{NewUInt64Set(1), NewUInt64Set(1, 2), NewUInt64Set(1, 2)},
		{NewUInt64Set(1, 2), NewUInt64Set(1), NewUInt64Set(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Union(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt64Set_Intersection(t *testing.T) {
	cases := []struct {
		set  UInt64Set
		arg  UInt64Set
		want UInt64Set
	}{
		{NewUInt64Set(), NewUInt64Set(), NewUInt64Set()},
		{NewUInt64Set(1), NewUInt64Set(1), NewUInt64Set(1)},
		{NewUInt64Set(1), NewUInt64Set(2), NewUInt64Set()},
		{NewUInt64Set(1), NewUInt64Set(1, 2), NewUInt64Set(1)},
		{NewUInt64Set(1, 2), NewUInt64Set(1), NewUInt64Set(1)},
	}
	for _, c := range cases {
		got := c.set.Intersection(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt64Set_Difference(t *testing.T) {
	cases := []struct {
		set  UInt64Set
		arg  UInt64Set
		want UInt64Set
	}{
		{NewUInt64Set(), NewUInt64Set(), NewUInt64Set()},
		{NewUInt64Set(1), NewUInt64Set(1), NewUInt64Set()},
		{NewUInt64Set(1), NewUInt64Set(2), NewUInt64Set(1)},
		{NewUInt64Set(1), NewUInt64Set(1, 2), NewUInt64Set()},
		{NewUInt64Set(1, 2), NewUInt64Set(1), NewUInt64Set(2)},
	}
	for _, c := range cases {
		got := c.set.Difference(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt64Set_IsSubsetOf(t *testing.T) {
	cases := []struct {
		set  UInt64Set
		arg  UInt64Set
		want bool
	}{
		{NewUInt64Set(), NewUInt64Set(), true},
		{NewUInt64Set(1), NewUInt64Set(1), true},
		{NewUInt64Set(1), NewUInt64Set(1, 2), true},
		{NewUInt64Set(1, 2), NewUInt64Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt64Set_IsProperSubsetOf(t *testing.T) {
	cases := []struct {
		set  UInt64Set
		arg  UInt64Set
		want bool
	}{
		{NewUInt64Set(), NewUInt64Set(), false},
		{NewUInt64Set(1), NewUInt64Set(1), false},
		{NewUInt64Set(1), NewUInt64Set(1, 2), true},
		{NewUInt64Set(1, 2), NewUInt64Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsProperSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt64Set_IsSupersetOf(t *testing.T) {
	cases := []struct {
		set  UInt64Set
		arg  UInt64Set
		want bool
	}{
		{NewUInt64Set(), NewUInt64Set(), true},
		{NewUInt64Set(1), NewUInt64Set(1), true},
		{NewUInt64Set(1), NewUInt64Set(1, 2), false},
		{NewUInt64Set(1, 2), NewUInt64Set(1), true},
	}
	for _, c := range cases {
		got := c.set.IsSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt64Set_IsProperSupersetOf(t *testing.T) {
	cases := []struct {
		set  UInt64Set
		arg  UInt64Set
		want bool
	}{
		{NewUInt64Set(), NewUInt64Set(), false},
		{NewUInt64Set(1), NewUInt64Set(1), false},
		{NewUInt64Set(1), NewUInt64Set(1, 2), false},
		{NewUInt64Set(1, 2), NewUInt64Set(1), true},
	}
	for _, c := range cases {
		got := c.set.IsProperSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt64Set_IsDisjointFrom(t *testing.T) {
	cases := []struct {
		set  UInt64Set
		arg  UInt64Set
		want bool
	}{
		{NewUInt64Set(), NewUInt64Set(), true},
		{NewUInt64Set(1), NewUInt64Set(1), false},
		{NewUInt64Set(1), NewUInt64Set(2, 3), true},
		{NewUInt64Set(1, 2), NewUInt64Set(3), true},
		{NewUInt64Set(1), NewUInt64Set(1, 2), false},
		{NewUInt64Set(1, 2), NewUInt64Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsDisjointFrom(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
