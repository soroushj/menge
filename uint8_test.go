package menge_test

import (
	"testing"

	"github.com/soroushj/menge"
)

func TestNewUInt8Set(t *testing.T) {
	cases := []struct {
		arg  []uint8
		want menge.UInt8Set
	}{
		{[]uint8{}, menge.UInt8Set{}},
		{[]uint8{1, 1}, menge.UInt8Set{1: struct{}{}}},
		{[]uint8{1, 2}, menge.UInt8Set{1: struct{}{}, 2: struct{}{}}},
	}
	for _, c := range cases {
		got := menge.NewUInt8Set(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt8Set_Add(t *testing.T) {
	cases := []struct {
		set  menge.UInt8Set
		arg  []uint8
		want menge.UInt8Set
	}{
		{menge.NewUInt8Set(), []uint8{}, menge.NewUInt8Set()},
		{menge.NewUInt8Set(), []uint8{1, 1}, menge.NewUInt8Set(1)},
		{menge.NewUInt8Set(), []uint8{1, 2}, menge.NewUInt8Set(1, 2)},
		{menge.NewUInt8Set(1), []uint8{}, menge.NewUInt8Set(1)},
		{menge.NewUInt8Set(1), []uint8{1, 1}, menge.NewUInt8Set(1)},
		{menge.NewUInt8Set(1), []uint8{2, 3}, menge.NewUInt8Set(1, 2, 3)},
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
		set  menge.UInt8Set
		arg  []uint8
		want menge.UInt8Set
	}{
		{menge.NewUInt8Set(), []uint8{}, menge.NewUInt8Set()},
		{menge.NewUInt8Set(1), []uint8{1, 1}, menge.NewUInt8Set()},
		{menge.NewUInt8Set(1, 2), []uint8{1, 2}, menge.NewUInt8Set()},
		{menge.NewUInt8Set(1), []uint8{}, menge.NewUInt8Set(1)},
		{menge.NewUInt8Set(1), []uint8{1, 1}, menge.NewUInt8Set()},
		{menge.NewUInt8Set(1, 2), []uint8{3}, menge.NewUInt8Set(1, 2)},
		{menge.NewUInt8Set(1, 2, 3), []uint8{2, 3}, menge.NewUInt8Set(1)},
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
		set  menge.UInt8Set
		want menge.UInt8Set
	}{
		{menge.NewUInt8Set(), menge.NewUInt8Set()},
		{menge.NewUInt8Set(1, 2), menge.NewUInt8Set()},
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
		set  menge.UInt8Set
		arg  uint8
		want bool
	}{
		{menge.NewUInt8Set(), 1, false},
		{menge.NewUInt8Set(2), 1, false},
		{menge.NewUInt8Set(1), 1, true},
		{menge.NewUInt8Set(1, 2), 1, true},
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
		set  menge.UInt8Set
		want int
	}{
		{menge.NewUInt8Set(), 0},
		{menge.NewUInt8Set(1, 2), 2},
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
		set  menge.UInt8Set
		want bool
	}{
		{menge.NewUInt8Set(), true},
		{menge.NewUInt8Set(1, 2), false},
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
		set  menge.UInt8Set
		want menge.UInt8Set
	}{
		{menge.NewUInt8Set(), menge.NewUInt8Set()},
		{menge.NewUInt8Set(1, 2), menge.NewUInt8Set(1, 2)},
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
		set  menge.UInt8Set
		want []uint8
	}{
		{menge.NewUInt8Set(), []uint8{}},
		{menge.NewUInt8Set(1, 2), []uint8{1, 2}},
	}
	for _, c := range cases {
		got := c.set.AsSlice()
		if len(got) != len(c.want) || !menge.NewUInt8Set(got...).Equals(menge.NewUInt8Set(c.want...)) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUInt8Set_String(t *testing.T) {
	cases := []struct {
		set  menge.UInt8Set
		want []string
	}{
		{menge.NewUInt8Set(), []string{"{}"}},
		{menge.NewUInt8Set(1), []string{"{1}"}},
		{menge.NewUInt8Set(1, 2), []string{"{1 2}", "{2 1}"}},
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
		set  menge.UInt8Set
		arg  menge.UInt8Set
		want bool
	}{
		{menge.NewUInt8Set(), menge.NewUInt8Set(), true},
		{menge.NewUInt8Set(1, 2), menge.NewUInt8Set(2, 1), true},
		{menge.NewUInt8Set(1, 2), menge.NewUInt8Set(1), false},
		{menge.NewUInt8Set(1), menge.NewUInt8Set(1, 2), false},
		{menge.NewUInt8Set(1), menge.NewUInt8Set(2), false},
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
		set  menge.UInt8Set
		arg  menge.UInt8Set
		want menge.UInt8Set
	}{
		{menge.NewUInt8Set(), menge.NewUInt8Set(), menge.NewUInt8Set()},
		{menge.NewUInt8Set(1), menge.NewUInt8Set(1), menge.NewUInt8Set(1)},
		{menge.NewUInt8Set(1), menge.NewUInt8Set(2), menge.NewUInt8Set(1, 2)},
		{menge.NewUInt8Set(1), menge.NewUInt8Set(1, 2), menge.NewUInt8Set(1, 2)},
		{menge.NewUInt8Set(1, 2), menge.NewUInt8Set(1), menge.NewUInt8Set(1, 2)},
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
		set  menge.UInt8Set
		arg  menge.UInt8Set
		want menge.UInt8Set
	}{
		{menge.NewUInt8Set(), menge.NewUInt8Set(), menge.NewUInt8Set()},
		{menge.NewUInt8Set(1), menge.NewUInt8Set(1), menge.NewUInt8Set(1)},
		{menge.NewUInt8Set(1), menge.NewUInt8Set(2), menge.NewUInt8Set()},
		{menge.NewUInt8Set(1), menge.NewUInt8Set(1, 2), menge.NewUInt8Set(1)},
		{menge.NewUInt8Set(1, 2), menge.NewUInt8Set(1), menge.NewUInt8Set(1)},
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
		set  menge.UInt8Set
		arg  menge.UInt8Set
		want menge.UInt8Set
	}{
		{menge.NewUInt8Set(), menge.NewUInt8Set(), menge.NewUInt8Set()},
		{menge.NewUInt8Set(1), menge.NewUInt8Set(1), menge.NewUInt8Set()},
		{menge.NewUInt8Set(1), menge.NewUInt8Set(2), menge.NewUInt8Set(1)},
		{menge.NewUInt8Set(1), menge.NewUInt8Set(1, 2), menge.NewUInt8Set()},
		{menge.NewUInt8Set(1, 2), menge.NewUInt8Set(1), menge.NewUInt8Set(2)},
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
		set  menge.UInt8Set
		arg  menge.UInt8Set
		want bool
	}{
		{menge.NewUInt8Set(), menge.NewUInt8Set(), true},
		{menge.NewUInt8Set(1), menge.NewUInt8Set(1), true},
		{menge.NewUInt8Set(1), menge.NewUInt8Set(1, 2), true},
		{menge.NewUInt8Set(1, 2), menge.NewUInt8Set(1), false},
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
		set  menge.UInt8Set
		arg  menge.UInt8Set
		want bool
	}{
		{menge.NewUInt8Set(), menge.NewUInt8Set(), false},
		{menge.NewUInt8Set(1), menge.NewUInt8Set(1), false},
		{menge.NewUInt8Set(1), menge.NewUInt8Set(1, 2), true},
		{menge.NewUInt8Set(1, 2), menge.NewUInt8Set(1), false},
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
		set  menge.UInt8Set
		arg  menge.UInt8Set
		want bool
	}{
		{menge.NewUInt8Set(), menge.NewUInt8Set(), true},
		{menge.NewUInt8Set(1), menge.NewUInt8Set(1), true},
		{menge.NewUInt8Set(1), menge.NewUInt8Set(1, 2), false},
		{menge.NewUInt8Set(1, 2), menge.NewUInt8Set(1), true},
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
		set  menge.UInt8Set
		arg  menge.UInt8Set
		want bool
	}{
		{menge.NewUInt8Set(), menge.NewUInt8Set(), false},
		{menge.NewUInt8Set(1), menge.NewUInt8Set(1), false},
		{menge.NewUInt8Set(1), menge.NewUInt8Set(1, 2), false},
		{menge.NewUInt8Set(1, 2), menge.NewUInt8Set(1), true},
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
		set  menge.UInt8Set
		arg  menge.UInt8Set
		want bool
	}{
		{menge.NewUInt8Set(), menge.NewUInt8Set(), true},
		{menge.NewUInt8Set(1), menge.NewUInt8Set(1), false},
		{menge.NewUInt8Set(1), menge.NewUInt8Set(2, 3), true},
		{menge.NewUInt8Set(1, 2), menge.NewUInt8Set(3), true},
		{menge.NewUInt8Set(1), menge.NewUInt8Set(1, 2), false},
		{menge.NewUInt8Set(1, 2), menge.NewUInt8Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsDisjointFrom(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
