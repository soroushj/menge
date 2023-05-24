package menge_test

import (
	"testing"

	"github.com/soroushj/menge"
)

func TestNewUInt32Set(t *testing.T) {
	cases := []struct {
		arg  []uint32
		want menge.UInt32Set
	}{
		{[]uint32{}, menge.UInt32Set{}},
		{[]uint32{1, 1}, menge.UInt32Set{1: struct{}{}}},
		{[]uint32{1, 2}, menge.UInt32Set{1: struct{}{}, 2: struct{}{}}},
	}
	for _, c := range cases {
		got := menge.NewUInt32Set(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt32Set_Add(t *testing.T) {
	cases := []struct {
		set  menge.UInt32Set
		arg  []uint32
		want menge.UInt32Set
	}{
		{menge.NewUInt32Set(), []uint32{}, menge.NewUInt32Set()},
		{menge.NewUInt32Set(), []uint32{1, 1}, menge.NewUInt32Set(1)},
		{menge.NewUInt32Set(), []uint32{1, 2}, menge.NewUInt32Set(1, 2)},
		{menge.NewUInt32Set(1), []uint32{}, menge.NewUInt32Set(1)},
		{menge.NewUInt32Set(1), []uint32{1, 1}, menge.NewUInt32Set(1)},
		{menge.NewUInt32Set(1), []uint32{2, 3}, menge.NewUInt32Set(1, 2, 3)},
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
		set  menge.UInt32Set
		arg  []uint32
		want menge.UInt32Set
	}{
		{menge.NewUInt32Set(), []uint32{}, menge.NewUInt32Set()},
		{menge.NewUInt32Set(1), []uint32{1, 1}, menge.NewUInt32Set()},
		{menge.NewUInt32Set(1, 2), []uint32{1, 2}, menge.NewUInt32Set()},
		{menge.NewUInt32Set(1), []uint32{}, menge.NewUInt32Set(1)},
		{menge.NewUInt32Set(1), []uint32{1, 1}, menge.NewUInt32Set()},
		{menge.NewUInt32Set(1, 2), []uint32{3}, menge.NewUInt32Set(1, 2)},
		{menge.NewUInt32Set(1, 2, 3), []uint32{2, 3}, menge.NewUInt32Set(1)},
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
		set  menge.UInt32Set
		want menge.UInt32Set
	}{
		{menge.NewUInt32Set(), menge.NewUInt32Set()},
		{menge.NewUInt32Set(1, 2), menge.NewUInt32Set()},
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
		set  menge.UInt32Set
		arg  uint32
		want bool
	}{
		{menge.NewUInt32Set(), 1, false},
		{menge.NewUInt32Set(2), 1, false},
		{menge.NewUInt32Set(1), 1, true},
		{menge.NewUInt32Set(1, 2), 1, true},
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
		set  menge.UInt32Set
		want int
	}{
		{menge.NewUInt32Set(), 0},
		{menge.NewUInt32Set(1, 2), 2},
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
		set  menge.UInt32Set
		want bool
	}{
		{menge.NewUInt32Set(), true},
		{menge.NewUInt32Set(1, 2), false},
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
		set  menge.UInt32Set
		want menge.UInt32Set
	}{
		{menge.NewUInt32Set(), menge.NewUInt32Set()},
		{menge.NewUInt32Set(1, 2), menge.NewUInt32Set(1, 2)},
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
		set  menge.UInt32Set
		want []uint32
	}{
		{menge.NewUInt32Set(), []uint32{}},
		{menge.NewUInt32Set(1, 2), []uint32{1, 2}},
	}
	for _, c := range cases {
		got := c.set.AsSlice()
		if len(got) != len(c.want) || !menge.NewUInt32Set(got...).Equals(menge.NewUInt32Set(c.want...)) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt32Set_String(t *testing.T) {
	cases := []struct {
		set  menge.UInt32Set
		want []string
	}{
		{menge.NewUInt32Set(), []string{"{}"}},
		{menge.NewUInt32Set(1), []string{"{1}"}},
		{menge.NewUInt32Set(1, 2), []string{"{1 2}", "{2 1}"}},
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
		set  menge.UInt32Set
		arg  menge.UInt32Set
		want bool
	}{
		{menge.NewUInt32Set(), menge.NewUInt32Set(), true},
		{menge.NewUInt32Set(1, 2), menge.NewUInt32Set(2, 1), true},
		{menge.NewUInt32Set(1, 2), menge.NewUInt32Set(1), false},
		{menge.NewUInt32Set(1), menge.NewUInt32Set(1, 2), false},
		{menge.NewUInt32Set(1), menge.NewUInt32Set(2), false},
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
		set  menge.UInt32Set
		arg  menge.UInt32Set
		want menge.UInt32Set
	}{
		{menge.NewUInt32Set(), menge.NewUInt32Set(), menge.NewUInt32Set()},
		{menge.NewUInt32Set(1), menge.NewUInt32Set(1), menge.NewUInt32Set(1)},
		{menge.NewUInt32Set(1), menge.NewUInt32Set(2), menge.NewUInt32Set(1, 2)},
		{menge.NewUInt32Set(1), menge.NewUInt32Set(1, 2), menge.NewUInt32Set(1, 2)},
		{menge.NewUInt32Set(1, 2), menge.NewUInt32Set(1), menge.NewUInt32Set(1, 2)},
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
		set  menge.UInt32Set
		arg  menge.UInt32Set
		want menge.UInt32Set
	}{
		{menge.NewUInt32Set(), menge.NewUInt32Set(), menge.NewUInt32Set()},
		{menge.NewUInt32Set(1), menge.NewUInt32Set(1), menge.NewUInt32Set(1)},
		{menge.NewUInt32Set(1), menge.NewUInt32Set(2), menge.NewUInt32Set()},
		{menge.NewUInt32Set(1), menge.NewUInt32Set(1, 2), menge.NewUInt32Set(1)},
		{menge.NewUInt32Set(1, 2), menge.NewUInt32Set(1), menge.NewUInt32Set(1)},
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
		set  menge.UInt32Set
		arg  menge.UInt32Set
		want menge.UInt32Set
	}{
		{menge.NewUInt32Set(), menge.NewUInt32Set(), menge.NewUInt32Set()},
		{menge.NewUInt32Set(1), menge.NewUInt32Set(1), menge.NewUInt32Set()},
		{menge.NewUInt32Set(1), menge.NewUInt32Set(2), menge.NewUInt32Set(1)},
		{menge.NewUInt32Set(1), menge.NewUInt32Set(1, 2), menge.NewUInt32Set()},
		{menge.NewUInt32Set(1, 2), menge.NewUInt32Set(1), menge.NewUInt32Set(2)},
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
		set  menge.UInt32Set
		arg  menge.UInt32Set
		want bool
	}{
		{menge.NewUInt32Set(), menge.NewUInt32Set(), true},
		{menge.NewUInt32Set(1), menge.NewUInt32Set(1), true},
		{menge.NewUInt32Set(1), menge.NewUInt32Set(1, 2), true},
		{menge.NewUInt32Set(1, 2), menge.NewUInt32Set(1), false},
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
		set  menge.UInt32Set
		arg  menge.UInt32Set
		want bool
	}{
		{menge.NewUInt32Set(), menge.NewUInt32Set(), false},
		{menge.NewUInt32Set(1), menge.NewUInt32Set(1), false},
		{menge.NewUInt32Set(1), menge.NewUInt32Set(1, 2), true},
		{menge.NewUInt32Set(1, 2), menge.NewUInt32Set(1), false},
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
		set  menge.UInt32Set
		arg  menge.UInt32Set
		want bool
	}{
		{menge.NewUInt32Set(), menge.NewUInt32Set(), true},
		{menge.NewUInt32Set(1), menge.NewUInt32Set(1), true},
		{menge.NewUInt32Set(1), menge.NewUInt32Set(1, 2), false},
		{menge.NewUInt32Set(1, 2), menge.NewUInt32Set(1), true},
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
		set  menge.UInt32Set
		arg  menge.UInt32Set
		want bool
	}{
		{menge.NewUInt32Set(), menge.NewUInt32Set(), false},
		{menge.NewUInt32Set(1), menge.NewUInt32Set(1), false},
		{menge.NewUInt32Set(1), menge.NewUInt32Set(1, 2), false},
		{menge.NewUInt32Set(1, 2), menge.NewUInt32Set(1), true},
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
		set  menge.UInt32Set
		arg  menge.UInt32Set
		want bool
	}{
		{menge.NewUInt32Set(), menge.NewUInt32Set(), true},
		{menge.NewUInt32Set(1), menge.NewUInt32Set(1), false},
		{menge.NewUInt32Set(1), menge.NewUInt32Set(2, 3), true},
		{menge.NewUInt32Set(1, 2), menge.NewUInt32Set(3), true},
		{menge.NewUInt32Set(1), menge.NewUInt32Set(1, 2), false},
		{menge.NewUInt32Set(1, 2), menge.NewUInt32Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsDisjointFrom(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
