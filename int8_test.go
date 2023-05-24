package menge_test

import (
	"testing"

	"github.com/soroushj/menge"
)

func TestNewInt8Set(t *testing.T) {
	cases := []struct {
		arg  []int8
		want menge.Int8Set
	}{
		{[]int8{}, menge.Int8Set{}},
		{[]int8{1, 1}, menge.Int8Set{1: struct{}{}}},
		{[]int8{1, 2}, menge.Int8Set{1: struct{}{}, 2: struct{}{}}},
	}
	for _, c := range cases {
		got := menge.NewInt8Set(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt8Set_Add(t *testing.T) {
	cases := []struct {
		set  menge.Int8Set
		arg  []int8
		want menge.Int8Set
	}{
		{menge.NewInt8Set(), []int8{}, menge.NewInt8Set()},
		{menge.NewInt8Set(), []int8{1, 1}, menge.NewInt8Set(1)},
		{menge.NewInt8Set(), []int8{1, 2}, menge.NewInt8Set(1, 2)},
		{menge.NewInt8Set(1), []int8{}, menge.NewInt8Set(1)},
		{menge.NewInt8Set(1), []int8{1, 1}, menge.NewInt8Set(1)},
		{menge.NewInt8Set(1), []int8{2, 3}, menge.NewInt8Set(1, 2, 3)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Add(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt8Set_Remove(t *testing.T) {
	cases := []struct {
		set  menge.Int8Set
		arg  []int8
		want menge.Int8Set
	}{
		{menge.NewInt8Set(), []int8{}, menge.NewInt8Set()},
		{menge.NewInt8Set(1), []int8{1, 1}, menge.NewInt8Set()},
		{menge.NewInt8Set(1, 2), []int8{1, 2}, menge.NewInt8Set()},
		{menge.NewInt8Set(1), []int8{}, menge.NewInt8Set(1)},
		{menge.NewInt8Set(1), []int8{1, 1}, menge.NewInt8Set()},
		{menge.NewInt8Set(1, 2), []int8{3}, menge.NewInt8Set(1, 2)},
		{menge.NewInt8Set(1, 2, 3), []int8{2, 3}, menge.NewInt8Set(1)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Remove(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt8Set_Empty(t *testing.T) {
	cases := []struct {
		set  menge.Int8Set
		want menge.Int8Set
	}{
		{menge.NewInt8Set(), menge.NewInt8Set()},
		{menge.NewInt8Set(1, 2), menge.NewInt8Set()},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Empty()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt8Set_Has(t *testing.T) {
	cases := []struct {
		set  menge.Int8Set
		arg  int8
		want bool
	}{
		{menge.NewInt8Set(), 1, false},
		{menge.NewInt8Set(2), 1, false},
		{menge.NewInt8Set(1), 1, true},
		{menge.NewInt8Set(1, 2), 1, true},
	}
	for _, c := range cases {
		got := c.set.Has(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt8Set_Size(t *testing.T) {
	cases := []struct {
		set  menge.Int8Set
		want int
	}{
		{menge.NewInt8Set(), 0},
		{menge.NewInt8Set(1, 2), 2},
	}
	for _, c := range cases {
		got := c.set.Size()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt8Set_IsEmpty(t *testing.T) {
	cases := []struct {
		set  menge.Int8Set
		want bool
	}{
		{menge.NewInt8Set(), true},
		{menge.NewInt8Set(1, 2), false},
	}
	for _, c := range cases {
		got := c.set.IsEmpty()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt8Set_Clone(t *testing.T) {
	cases := []struct {
		set  menge.Int8Set
		want menge.Int8Set
	}{
		{menge.NewInt8Set(), menge.NewInt8Set()},
		{menge.NewInt8Set(1, 2), menge.NewInt8Set(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt8Set_AsSlice(t *testing.T) {
	cases := []struct {
		set  menge.Int8Set
		want []int8
	}{
		{menge.NewInt8Set(), []int8{}},
		{menge.NewInt8Set(1, 2), []int8{1, 2}},
	}
	for _, c := range cases {
		got := c.set.AsSlice()
		if len(got) != len(c.want) || !menge.NewInt8Set(got...).Equals(menge.NewInt8Set(c.want...)) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt8Set_String(t *testing.T) {
	cases := []struct {
		set  menge.Int8Set
		want []string
	}{
		{menge.NewInt8Set(), []string{"{}"}},
		{menge.NewInt8Set(1), []string{"{1}"}},
		{menge.NewInt8Set(1, 2), []string{"{1 2}", "{2 1}"}},
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

func TestInt8Set_Equals(t *testing.T) {
	cases := []struct {
		set  menge.Int8Set
		arg  menge.Int8Set
		want bool
	}{
		{menge.NewInt8Set(), menge.NewInt8Set(), true},
		{menge.NewInt8Set(1, 2), menge.NewInt8Set(2, 1), true},
		{menge.NewInt8Set(1, 2), menge.NewInt8Set(1), false},
		{menge.NewInt8Set(1), menge.NewInt8Set(1, 2), false},
		{menge.NewInt8Set(1), menge.NewInt8Set(2), false},
	}
	for _, c := range cases {
		got := c.set.Equals(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt8Set_Union(t *testing.T) {
	cases := []struct {
		set  menge.Int8Set
		arg  menge.Int8Set
		want menge.Int8Set
	}{
		{menge.NewInt8Set(), menge.NewInt8Set(), menge.NewInt8Set()},
		{menge.NewInt8Set(1), menge.NewInt8Set(1), menge.NewInt8Set(1)},
		{menge.NewInt8Set(1), menge.NewInt8Set(2), menge.NewInt8Set(1, 2)},
		{menge.NewInt8Set(1), menge.NewInt8Set(1, 2), menge.NewInt8Set(1, 2)},
		{menge.NewInt8Set(1, 2), menge.NewInt8Set(1), menge.NewInt8Set(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Union(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt8Set_Intersection(t *testing.T) {
	cases := []struct {
		set  menge.Int8Set
		arg  menge.Int8Set
		want menge.Int8Set
	}{
		{menge.NewInt8Set(), menge.NewInt8Set(), menge.NewInt8Set()},
		{menge.NewInt8Set(1), menge.NewInt8Set(1), menge.NewInt8Set(1)},
		{menge.NewInt8Set(1), menge.NewInt8Set(2), menge.NewInt8Set()},
		{menge.NewInt8Set(1), menge.NewInt8Set(1, 2), menge.NewInt8Set(1)},
		{menge.NewInt8Set(1, 2), menge.NewInt8Set(1), menge.NewInt8Set(1)},
	}
	for _, c := range cases {
		got := c.set.Intersection(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt8Set_Difference(t *testing.T) {
	cases := []struct {
		set  menge.Int8Set
		arg  menge.Int8Set
		want menge.Int8Set
	}{
		{menge.NewInt8Set(), menge.NewInt8Set(), menge.NewInt8Set()},
		{menge.NewInt8Set(1), menge.NewInt8Set(1), menge.NewInt8Set()},
		{menge.NewInt8Set(1), menge.NewInt8Set(2), menge.NewInt8Set(1)},
		{menge.NewInt8Set(1), menge.NewInt8Set(1, 2), menge.NewInt8Set()},
		{menge.NewInt8Set(1, 2), menge.NewInt8Set(1), menge.NewInt8Set(2)},
	}
	for _, c := range cases {
		got := c.set.Difference(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt8Set_IsSubsetOf(t *testing.T) {
	cases := []struct {
		set  menge.Int8Set
		arg  menge.Int8Set
		want bool
	}{
		{menge.NewInt8Set(), menge.NewInt8Set(), true},
		{menge.NewInt8Set(1), menge.NewInt8Set(1), true},
		{menge.NewInt8Set(1), menge.NewInt8Set(1, 2), true},
		{menge.NewInt8Set(1, 2), menge.NewInt8Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt8Set_IsProperSubsetOf(t *testing.T) {
	cases := []struct {
		set  menge.Int8Set
		arg  menge.Int8Set
		want bool
	}{
		{menge.NewInt8Set(), menge.NewInt8Set(), false},
		{menge.NewInt8Set(1), menge.NewInt8Set(1), false},
		{menge.NewInt8Set(1), menge.NewInt8Set(1, 2), true},
		{menge.NewInt8Set(1, 2), menge.NewInt8Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsProperSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt8Set_IsSupersetOf(t *testing.T) {
	cases := []struct {
		set  menge.Int8Set
		arg  menge.Int8Set
		want bool
	}{
		{menge.NewInt8Set(), menge.NewInt8Set(), true},
		{menge.NewInt8Set(1), menge.NewInt8Set(1), true},
		{menge.NewInt8Set(1), menge.NewInt8Set(1, 2), false},
		{menge.NewInt8Set(1, 2), menge.NewInt8Set(1), true},
	}
	for _, c := range cases {
		got := c.set.IsSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt8Set_IsProperSupersetOf(t *testing.T) {
	cases := []struct {
		set  menge.Int8Set
		arg  menge.Int8Set
		want bool
	}{
		{menge.NewInt8Set(), menge.NewInt8Set(), false},
		{menge.NewInt8Set(1), menge.NewInt8Set(1), false},
		{menge.NewInt8Set(1), menge.NewInt8Set(1, 2), false},
		{menge.NewInt8Set(1, 2), menge.NewInt8Set(1), true},
	}
	for _, c := range cases {
		got := c.set.IsProperSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt8Set_IsDisjointFrom(t *testing.T) {
	cases := []struct {
		set  menge.Int8Set
		arg  menge.Int8Set
		want bool
	}{
		{menge.NewInt8Set(), menge.NewInt8Set(), true},
		{menge.NewInt8Set(1), menge.NewInt8Set(1), false},
		{menge.NewInt8Set(1), menge.NewInt8Set(2, 3), true},
		{menge.NewInt8Set(1, 2), menge.NewInt8Set(3), true},
		{menge.NewInt8Set(1), menge.NewInt8Set(1, 2), false},
		{menge.NewInt8Set(1, 2), menge.NewInt8Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsDisjointFrom(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
