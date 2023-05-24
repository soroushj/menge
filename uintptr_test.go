package menge_test

import (
	"testing"

	"github.com/soroushj/menge"
)

func TestNewUIntPtrSet(t *testing.T) {
	cases := []struct {
		arg  []uintptr
		want menge.UIntPtrSet
	}{
		{[]uintptr{}, menge.UIntPtrSet{}},
		{[]uintptr{1, 1}, menge.UIntPtrSet{1: struct{}{}}},
		{[]uintptr{1, 2}, menge.UIntPtrSet{1: struct{}{}, 2: struct{}{}}},
	}
	for _, c := range cases {
		got := menge.NewUIntPtrSet(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntPtrSet_Add(t *testing.T) {
	cases := []struct {
		set  menge.UIntPtrSet
		arg  []uintptr
		want menge.UIntPtrSet
	}{
		{menge.NewUIntPtrSet(), []uintptr{}, menge.NewUIntPtrSet()},
		{menge.NewUIntPtrSet(), []uintptr{1, 1}, menge.NewUIntPtrSet(1)},
		{menge.NewUIntPtrSet(), []uintptr{1, 2}, menge.NewUIntPtrSet(1, 2)},
		{menge.NewUIntPtrSet(1), []uintptr{}, menge.NewUIntPtrSet(1)},
		{menge.NewUIntPtrSet(1), []uintptr{1, 1}, menge.NewUIntPtrSet(1)},
		{menge.NewUIntPtrSet(1), []uintptr{2, 3}, menge.NewUIntPtrSet(1, 2, 3)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Add(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntPtrSet_Remove(t *testing.T) {
	cases := []struct {
		set  menge.UIntPtrSet
		arg  []uintptr
		want menge.UIntPtrSet
	}{
		{menge.NewUIntPtrSet(), []uintptr{}, menge.NewUIntPtrSet()},
		{menge.NewUIntPtrSet(1), []uintptr{1, 1}, menge.NewUIntPtrSet()},
		{menge.NewUIntPtrSet(1, 2), []uintptr{1, 2}, menge.NewUIntPtrSet()},
		{menge.NewUIntPtrSet(1), []uintptr{}, menge.NewUIntPtrSet(1)},
		{menge.NewUIntPtrSet(1), []uintptr{1, 1}, menge.NewUIntPtrSet()},
		{menge.NewUIntPtrSet(1, 2), []uintptr{3}, menge.NewUIntPtrSet(1, 2)},
		{menge.NewUIntPtrSet(1, 2, 3), []uintptr{2, 3}, menge.NewUIntPtrSet(1)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Remove(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntPtrSet_Empty(t *testing.T) {
	cases := []struct {
		set  menge.UIntPtrSet
		want menge.UIntPtrSet
	}{
		{menge.NewUIntPtrSet(), menge.NewUIntPtrSet()},
		{menge.NewUIntPtrSet(1, 2), menge.NewUIntPtrSet()},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Empty()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntPtrSet_Has(t *testing.T) {
	cases := []struct {
		set  menge.UIntPtrSet
		arg  uintptr
		want bool
	}{
		{menge.NewUIntPtrSet(), 1, false},
		{menge.NewUIntPtrSet(2), 1, false},
		{menge.NewUIntPtrSet(1), 1, true},
		{menge.NewUIntPtrSet(1, 2), 1, true},
	}
	for _, c := range cases {
		got := c.set.Has(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntPtrSet_Size(t *testing.T) {
	cases := []struct {
		set  menge.UIntPtrSet
		want int
	}{
		{menge.NewUIntPtrSet(), 0},
		{menge.NewUIntPtrSet(1, 2), 2},
	}
	for _, c := range cases {
		got := c.set.Size()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntPtrSet_IsEmpty(t *testing.T) {
	cases := []struct {
		set  menge.UIntPtrSet
		want bool
	}{
		{menge.NewUIntPtrSet(), true},
		{menge.NewUIntPtrSet(1, 2), false},
	}
	for _, c := range cases {
		got := c.set.IsEmpty()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntPtrSet_Clone(t *testing.T) {
	cases := []struct {
		set  menge.UIntPtrSet
		want menge.UIntPtrSet
	}{
		{menge.NewUIntPtrSet(), menge.NewUIntPtrSet()},
		{menge.NewUIntPtrSet(1, 2), menge.NewUIntPtrSet(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntPtrSet_AsSlice(t *testing.T) {
	cases := []struct {
		set  menge.UIntPtrSet
		want []uintptr
	}{
		{menge.NewUIntPtrSet(), []uintptr{}},
		{menge.NewUIntPtrSet(1, 2), []uintptr{1, 2}},
	}
	for _, c := range cases {
		got := c.set.AsSlice()
		if len(got) != len(c.want) || !menge.NewUIntPtrSet(got...).Equals(menge.NewUIntPtrSet(c.want...)) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntPtrSet_String(t *testing.T) {
	cases := []struct {
		set  menge.UIntPtrSet
		want []string
	}{
		{menge.NewUIntPtrSet(), []string{"{}"}},
		{menge.NewUIntPtrSet(1), []string{"{1}"}},
		{menge.NewUIntPtrSet(1, 2), []string{"{1 2}", "{2 1}"}},
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

func TestUIntPtrSet_Equals(t *testing.T) {
	cases := []struct {
		set  menge.UIntPtrSet
		arg  menge.UIntPtrSet
		want bool
	}{
		{menge.NewUIntPtrSet(), menge.NewUIntPtrSet(), true},
		{menge.NewUIntPtrSet(1, 2), menge.NewUIntPtrSet(2, 1), true},
		{menge.NewUIntPtrSet(1, 2), menge.NewUIntPtrSet(1), false},
		{menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(1, 2), false},
		{menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(2), false},
	}
	for _, c := range cases {
		got := c.set.Equals(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntPtrSet_Union(t *testing.T) {
	cases := []struct {
		set  menge.UIntPtrSet
		arg  menge.UIntPtrSet
		want menge.UIntPtrSet
	}{
		{menge.NewUIntPtrSet(), menge.NewUIntPtrSet(), menge.NewUIntPtrSet()},
		{menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(1)},
		{menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(2), menge.NewUIntPtrSet(1, 2)},
		{menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(1, 2), menge.NewUIntPtrSet(1, 2)},
		{menge.NewUIntPtrSet(1, 2), menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Union(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntPtrSet_Intersection(t *testing.T) {
	cases := []struct {
		set  menge.UIntPtrSet
		arg  menge.UIntPtrSet
		want menge.UIntPtrSet
	}{
		{menge.NewUIntPtrSet(), menge.NewUIntPtrSet(), menge.NewUIntPtrSet()},
		{menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(1)},
		{menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(2), menge.NewUIntPtrSet()},
		{menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(1, 2), menge.NewUIntPtrSet(1)},
		{menge.NewUIntPtrSet(1, 2), menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(1)},
	}
	for _, c := range cases {
		got := c.set.Intersection(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntPtrSet_Difference(t *testing.T) {
	cases := []struct {
		set  menge.UIntPtrSet
		arg  menge.UIntPtrSet
		want menge.UIntPtrSet
	}{
		{menge.NewUIntPtrSet(), menge.NewUIntPtrSet(), menge.NewUIntPtrSet()},
		{menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(1), menge.NewUIntPtrSet()},
		{menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(2), menge.NewUIntPtrSet(1)},
		{menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(1, 2), menge.NewUIntPtrSet()},
		{menge.NewUIntPtrSet(1, 2), menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(2)},
	}
	for _, c := range cases {
		got := c.set.Difference(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntPtrSet_IsSubsetOf(t *testing.T) {
	cases := []struct {
		set  menge.UIntPtrSet
		arg  menge.UIntPtrSet
		want bool
	}{
		{menge.NewUIntPtrSet(), menge.NewUIntPtrSet(), true},
		{menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(1), true},
		{menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(1, 2), true},
		{menge.NewUIntPtrSet(1, 2), menge.NewUIntPtrSet(1), false},
	}
	for _, c := range cases {
		got := c.set.IsSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntPtrSet_IsProperSubsetOf(t *testing.T) {
	cases := []struct {
		set  menge.UIntPtrSet
		arg  menge.UIntPtrSet
		want bool
	}{
		{menge.NewUIntPtrSet(), menge.NewUIntPtrSet(), false},
		{menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(1), false},
		{menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(1, 2), true},
		{menge.NewUIntPtrSet(1, 2), menge.NewUIntPtrSet(1), false},
	}
	for _, c := range cases {
		got := c.set.IsProperSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntPtrSet_IsSupersetOf(t *testing.T) {
	cases := []struct {
		set  menge.UIntPtrSet
		arg  menge.UIntPtrSet
		want bool
	}{
		{menge.NewUIntPtrSet(), menge.NewUIntPtrSet(), true},
		{menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(1), true},
		{menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(1, 2), false},
		{menge.NewUIntPtrSet(1, 2), menge.NewUIntPtrSet(1), true},
	}
	for _, c := range cases {
		got := c.set.IsSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntPtrSet_IsProperSupersetOf(t *testing.T) {
	cases := []struct {
		set  menge.UIntPtrSet
		arg  menge.UIntPtrSet
		want bool
	}{
		{menge.NewUIntPtrSet(), menge.NewUIntPtrSet(), false},
		{menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(1), false},
		{menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(1, 2), false},
		{menge.NewUIntPtrSet(1, 2), menge.NewUIntPtrSet(1), true},
	}
	for _, c := range cases {
		got := c.set.IsProperSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntPtrSet_IsDisjointFrom(t *testing.T) {
	cases := []struct {
		set  menge.UIntPtrSet
		arg  menge.UIntPtrSet
		want bool
	}{
		{menge.NewUIntPtrSet(), menge.NewUIntPtrSet(), true},
		{menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(1), false},
		{menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(2, 3), true},
		{menge.NewUIntPtrSet(1, 2), menge.NewUIntPtrSet(3), true},
		{menge.NewUIntPtrSet(1), menge.NewUIntPtrSet(1, 2), false},
		{menge.NewUIntPtrSet(1, 2), menge.NewUIntPtrSet(1), false},
	}
	for _, c := range cases {
		got := c.set.IsDisjointFrom(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
