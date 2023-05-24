package menge_test

import (
	"testing"

	"github.com/soroushj/menge"
)

func TestNewUIntSet(t *testing.T) {
	cases := []struct {
		arg  []uint
		want menge.UIntSet
	}{
		{[]uint{}, menge.UIntSet{}},
		{[]uint{1, 1}, menge.UIntSet{1: struct{}{}}},
		{[]uint{1, 2}, menge.UIntSet{1: struct{}{}, 2: struct{}{}}},
	}
	for _, c := range cases {
		got := menge.NewUIntSet(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntSet_Add(t *testing.T) {
	cases := []struct {
		set  menge.UIntSet
		arg  []uint
		want menge.UIntSet
	}{
		{menge.NewUIntSet(), []uint{}, menge.NewUIntSet()},
		{menge.NewUIntSet(), []uint{1, 1}, menge.NewUIntSet(1)},
		{menge.NewUIntSet(), []uint{1, 2}, menge.NewUIntSet(1, 2)},
		{menge.NewUIntSet(1), []uint{}, menge.NewUIntSet(1)},
		{menge.NewUIntSet(1), []uint{1, 1}, menge.NewUIntSet(1)},
		{menge.NewUIntSet(1), []uint{2, 3}, menge.NewUIntSet(1, 2, 3)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Add(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntSet_Remove(t *testing.T) {
	cases := []struct {
		set  menge.UIntSet
		arg  []uint
		want menge.UIntSet
	}{
		{menge.NewUIntSet(), []uint{}, menge.NewUIntSet()},
		{menge.NewUIntSet(1), []uint{1, 1}, menge.NewUIntSet()},
		{menge.NewUIntSet(1, 2), []uint{1, 2}, menge.NewUIntSet()},
		{menge.NewUIntSet(1), []uint{}, menge.NewUIntSet(1)},
		{menge.NewUIntSet(1), []uint{1, 1}, menge.NewUIntSet()},
		{menge.NewUIntSet(1, 2), []uint{3}, menge.NewUIntSet(1, 2)},
		{menge.NewUIntSet(1, 2, 3), []uint{2, 3}, menge.NewUIntSet(1)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Remove(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntSet_Empty(t *testing.T) {
	cases := []struct {
		set  menge.UIntSet
		want menge.UIntSet
	}{
		{menge.NewUIntSet(), menge.NewUIntSet()},
		{menge.NewUIntSet(1, 2), menge.NewUIntSet()},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Empty()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntSet_Has(t *testing.T) {
	cases := []struct {
		set  menge.UIntSet
		arg  uint
		want bool
	}{
		{menge.NewUIntSet(), 1, false},
		{menge.NewUIntSet(2), 1, false},
		{menge.NewUIntSet(1), 1, true},
		{menge.NewUIntSet(1, 2), 1, true},
	}
	for _, c := range cases {
		got := c.set.Has(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntSet_Size(t *testing.T) {
	cases := []struct {
		set  menge.UIntSet
		want int
	}{
		{menge.NewUIntSet(), 0},
		{menge.NewUIntSet(1, 2), 2},
	}
	for _, c := range cases {
		got := c.set.Size()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntSet_IsEmpty(t *testing.T) {
	cases := []struct {
		set  menge.UIntSet
		want bool
	}{
		{menge.NewUIntSet(), true},
		{menge.NewUIntSet(1, 2), false},
	}
	for _, c := range cases {
		got := c.set.IsEmpty()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntSet_Clone(t *testing.T) {
	cases := []struct {
		set  menge.UIntSet
		want menge.UIntSet
	}{
		{menge.NewUIntSet(), menge.NewUIntSet()},
		{menge.NewUIntSet(1, 2), menge.NewUIntSet(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntSet_AsSlice(t *testing.T) {
	cases := []struct {
		set  menge.UIntSet
		want []uint
	}{
		{menge.NewUIntSet(), []uint{}},
		{menge.NewUIntSet(1, 2), []uint{1, 2}},
	}
	for _, c := range cases {
		got := c.set.AsSlice()
		if len(got) != len(c.want) || !menge.NewUIntSet(got...).Equals(menge.NewUIntSet(c.want...)) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntSet_String(t *testing.T) {
	cases := []struct {
		set  menge.UIntSet
		want []string
	}{
		{menge.NewUIntSet(), []string{"{}"}},
		{menge.NewUIntSet(1), []string{"{1}"}},
		{menge.NewUIntSet(1, 2), []string{"{1 2}", "{2 1}"}},
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

func TestUIntSet_Equals(t *testing.T) {
	cases := []struct {
		set  menge.UIntSet
		arg  menge.UIntSet
		want bool
	}{
		{menge.NewUIntSet(), menge.NewUIntSet(), true},
		{menge.NewUIntSet(1, 2), menge.NewUIntSet(2, 1), true},
		{menge.NewUIntSet(1, 2), menge.NewUIntSet(1), false},
		{menge.NewUIntSet(1), menge.NewUIntSet(1, 2), false},
		{menge.NewUIntSet(1), menge.NewUIntSet(2), false},
	}
	for _, c := range cases {
		got := c.set.Equals(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntSet_Union(t *testing.T) {
	cases := []struct {
		set  menge.UIntSet
		arg  menge.UIntSet
		want menge.UIntSet
	}{
		{menge.NewUIntSet(), menge.NewUIntSet(), menge.NewUIntSet()},
		{menge.NewUIntSet(1), menge.NewUIntSet(1), menge.NewUIntSet(1)},
		{menge.NewUIntSet(1), menge.NewUIntSet(2), menge.NewUIntSet(1, 2)},
		{menge.NewUIntSet(1), menge.NewUIntSet(1, 2), menge.NewUIntSet(1, 2)},
		{menge.NewUIntSet(1, 2), menge.NewUIntSet(1), menge.NewUIntSet(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Union(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntSet_Intersection(t *testing.T) {
	cases := []struct {
		set  menge.UIntSet
		arg  menge.UIntSet
		want menge.UIntSet
	}{
		{menge.NewUIntSet(), menge.NewUIntSet(), menge.NewUIntSet()},
		{menge.NewUIntSet(1), menge.NewUIntSet(1), menge.NewUIntSet(1)},
		{menge.NewUIntSet(1), menge.NewUIntSet(2), menge.NewUIntSet()},
		{menge.NewUIntSet(1), menge.NewUIntSet(1, 2), menge.NewUIntSet(1)},
		{menge.NewUIntSet(1, 2), menge.NewUIntSet(1), menge.NewUIntSet(1)},
	}
	for _, c := range cases {
		got := c.set.Intersection(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntSet_Difference(t *testing.T) {
	cases := []struct {
		set  menge.UIntSet
		arg  menge.UIntSet
		want menge.UIntSet
	}{
		{menge.NewUIntSet(), menge.NewUIntSet(), menge.NewUIntSet()},
		{menge.NewUIntSet(1), menge.NewUIntSet(1), menge.NewUIntSet()},
		{menge.NewUIntSet(1), menge.NewUIntSet(2), menge.NewUIntSet(1)},
		{menge.NewUIntSet(1), menge.NewUIntSet(1, 2), menge.NewUIntSet()},
		{menge.NewUIntSet(1, 2), menge.NewUIntSet(1), menge.NewUIntSet(2)},
	}
	for _, c := range cases {
		got := c.set.Difference(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntSet_IsSubsetOf(t *testing.T) {
	cases := []struct {
		set  menge.UIntSet
		arg  menge.UIntSet
		want bool
	}{
		{menge.NewUIntSet(), menge.NewUIntSet(), true},
		{menge.NewUIntSet(1), menge.NewUIntSet(1), true},
		{menge.NewUIntSet(1), menge.NewUIntSet(1, 2), true},
		{menge.NewUIntSet(1, 2), menge.NewUIntSet(1), false},
	}
	for _, c := range cases {
		got := c.set.IsSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntSet_IsProperSubsetOf(t *testing.T) {
	cases := []struct {
		set  menge.UIntSet
		arg  menge.UIntSet
		want bool
	}{
		{menge.NewUIntSet(), menge.NewUIntSet(), false},
		{menge.NewUIntSet(1), menge.NewUIntSet(1), false},
		{menge.NewUIntSet(1), menge.NewUIntSet(1, 2), true},
		{menge.NewUIntSet(1, 2), menge.NewUIntSet(1), false},
	}
	for _, c := range cases {
		got := c.set.IsProperSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntSet_IsSupersetOf(t *testing.T) {
	cases := []struct {
		set  menge.UIntSet
		arg  menge.UIntSet
		want bool
	}{
		{menge.NewUIntSet(), menge.NewUIntSet(), true},
		{menge.NewUIntSet(1), menge.NewUIntSet(1), true},
		{menge.NewUIntSet(1), menge.NewUIntSet(1, 2), false},
		{menge.NewUIntSet(1, 2), menge.NewUIntSet(1), true},
	}
	for _, c := range cases {
		got := c.set.IsSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntSet_IsProperSupersetOf(t *testing.T) {
	cases := []struct {
		set  menge.UIntSet
		arg  menge.UIntSet
		want bool
	}{
		{menge.NewUIntSet(), menge.NewUIntSet(), false},
		{menge.NewUIntSet(1), menge.NewUIntSet(1), false},
		{menge.NewUIntSet(1), menge.NewUIntSet(1, 2), false},
		{menge.NewUIntSet(1, 2), menge.NewUIntSet(1), true},
	}
	for _, c := range cases {
		got := c.set.IsProperSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntSet_IsDisjointFrom(t *testing.T) {
	cases := []struct {
		set  menge.UIntSet
		arg  menge.UIntSet
		want bool
	}{
		{menge.NewUIntSet(), menge.NewUIntSet(), true},
		{menge.NewUIntSet(1), menge.NewUIntSet(1), false},
		{menge.NewUIntSet(1), menge.NewUIntSet(2, 3), true},
		{menge.NewUIntSet(1, 2), menge.NewUIntSet(3), true},
		{menge.NewUIntSet(1), menge.NewUIntSet(1, 2), false},
		{menge.NewUIntSet(1, 2), menge.NewUIntSet(1), false},
	}
	for _, c := range cases {
		got := c.set.IsDisjointFrom(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
