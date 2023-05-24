package menge_test

import (
	"testing"

	"github.com/soroushj/menge"
)

func TestNewInt32Set(t *testing.T) {
	cases := []struct {
		arg  []int32
		want menge.Int32Set
	}{
		{[]int32{}, menge.Int32Set{}},
		{[]int32{1, 1}, menge.Int32Set{1: struct{}{}}},
		{[]int32{1, 2}, menge.Int32Set{1: struct{}{}, 2: struct{}{}}},
	}
	for _, c := range cases {
		got := menge.NewInt32Set(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt32Set_Add(t *testing.T) {
	cases := []struct {
		set  menge.Int32Set
		arg  []int32
		want menge.Int32Set
	}{
		{menge.NewInt32Set(), []int32{}, menge.NewInt32Set()},
		{menge.NewInt32Set(), []int32{1, 1}, menge.NewInt32Set(1)},
		{menge.NewInt32Set(), []int32{1, 2}, menge.NewInt32Set(1, 2)},
		{menge.NewInt32Set(1), []int32{}, menge.NewInt32Set(1)},
		{menge.NewInt32Set(1), []int32{1, 1}, menge.NewInt32Set(1)},
		{menge.NewInt32Set(1), []int32{2, 3}, menge.NewInt32Set(1, 2, 3)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Add(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt32Set_Remove(t *testing.T) {
	cases := []struct {
		set  menge.Int32Set
		arg  []int32
		want menge.Int32Set
	}{
		{menge.NewInt32Set(), []int32{}, menge.NewInt32Set()},
		{menge.NewInt32Set(1), []int32{1, 1}, menge.NewInt32Set()},
		{menge.NewInt32Set(1, 2), []int32{1, 2}, menge.NewInt32Set()},
		{menge.NewInt32Set(1), []int32{}, menge.NewInt32Set(1)},
		{menge.NewInt32Set(1), []int32{1, 1}, menge.NewInt32Set()},
		{menge.NewInt32Set(1, 2), []int32{3}, menge.NewInt32Set(1, 2)},
		{menge.NewInt32Set(1, 2, 3), []int32{2, 3}, menge.NewInt32Set(1)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Remove(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt32Set_Empty(t *testing.T) {
	cases := []struct {
		set  menge.Int32Set
		want menge.Int32Set
	}{
		{menge.NewInt32Set(), menge.NewInt32Set()},
		{menge.NewInt32Set(1, 2), menge.NewInt32Set()},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Empty()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt32Set_Has(t *testing.T) {
	cases := []struct {
		set  menge.Int32Set
		arg  int32
		want bool
	}{
		{menge.NewInt32Set(), 1, false},
		{menge.NewInt32Set(2), 1, false},
		{menge.NewInt32Set(1), 1, true},
		{menge.NewInt32Set(1, 2), 1, true},
	}
	for _, c := range cases {
		got := c.set.Has(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt32Set_Size(t *testing.T) {
	cases := []struct {
		set  menge.Int32Set
		want int
	}{
		{menge.NewInt32Set(), 0},
		{menge.NewInt32Set(1, 2), 2},
	}
	for _, c := range cases {
		got := c.set.Size()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt32Set_IsEmpty(t *testing.T) {
	cases := []struct {
		set  menge.Int32Set
		want bool
	}{
		{menge.NewInt32Set(), true},
		{menge.NewInt32Set(1, 2), false},
	}
	for _, c := range cases {
		got := c.set.IsEmpty()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt32Set_Clone(t *testing.T) {
	cases := []struct {
		set  menge.Int32Set
		want menge.Int32Set
	}{
		{menge.NewInt32Set(), menge.NewInt32Set()},
		{menge.NewInt32Set(1, 2), menge.NewInt32Set(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt32Set_AsSlice(t *testing.T) {
	cases := []struct {
		set  menge.Int32Set
		want []int32
	}{
		{menge.NewInt32Set(), []int32{}},
		{menge.NewInt32Set(1, 2), []int32{1, 2}},
	}
	for _, c := range cases {
		got := c.set.AsSlice()
		if len(got) != len(c.want) || !menge.NewInt32Set(got...).Equals(menge.NewInt32Set(c.want...)) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt32Set_String(t *testing.T) {
	cases := []struct {
		set  menge.Int32Set
		want []string
	}{
		{menge.NewInt32Set(), []string{"{}"}},
		{menge.NewInt32Set(1), []string{"{1}"}},
		{menge.NewInt32Set(1, 2), []string{"{1 2}", "{2 1}"}},
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

func TestInt32Set_Equals(t *testing.T) {
	cases := []struct {
		set  menge.Int32Set
		arg  menge.Int32Set
		want bool
	}{
		{menge.NewInt32Set(), menge.NewInt32Set(), true},
		{menge.NewInt32Set(1, 2), menge.NewInt32Set(2, 1), true},
		{menge.NewInt32Set(1, 2), menge.NewInt32Set(1), false},
		{menge.NewInt32Set(1), menge.NewInt32Set(1, 2), false},
		{menge.NewInt32Set(1), menge.NewInt32Set(2), false},
	}
	for _, c := range cases {
		got := c.set.Equals(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt32Set_Union(t *testing.T) {
	cases := []struct {
		set  menge.Int32Set
		arg  menge.Int32Set
		want menge.Int32Set
	}{
		{menge.NewInt32Set(), menge.NewInt32Set(), menge.NewInt32Set()},
		{menge.NewInt32Set(1), menge.NewInt32Set(1), menge.NewInt32Set(1)},
		{menge.NewInt32Set(1), menge.NewInt32Set(2), menge.NewInt32Set(1, 2)},
		{menge.NewInt32Set(1), menge.NewInt32Set(1, 2), menge.NewInt32Set(1, 2)},
		{menge.NewInt32Set(1, 2), menge.NewInt32Set(1), menge.NewInt32Set(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Union(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt32Set_Intersection(t *testing.T) {
	cases := []struct {
		set  menge.Int32Set
		arg  menge.Int32Set
		want menge.Int32Set
	}{
		{menge.NewInt32Set(), menge.NewInt32Set(), menge.NewInt32Set()},
		{menge.NewInt32Set(1), menge.NewInt32Set(1), menge.NewInt32Set(1)},
		{menge.NewInt32Set(1), menge.NewInt32Set(2), menge.NewInt32Set()},
		{menge.NewInt32Set(1), menge.NewInt32Set(1, 2), menge.NewInt32Set(1)},
		{menge.NewInt32Set(1, 2), menge.NewInt32Set(1), menge.NewInt32Set(1)},
	}
	for _, c := range cases {
		got := c.set.Intersection(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt32Set_Difference(t *testing.T) {
	cases := []struct {
		set  menge.Int32Set
		arg  menge.Int32Set
		want menge.Int32Set
	}{
		{menge.NewInt32Set(), menge.NewInt32Set(), menge.NewInt32Set()},
		{menge.NewInt32Set(1), menge.NewInt32Set(1), menge.NewInt32Set()},
		{menge.NewInt32Set(1), menge.NewInt32Set(2), menge.NewInt32Set(1)},
		{menge.NewInt32Set(1), menge.NewInt32Set(1, 2), menge.NewInt32Set()},
		{menge.NewInt32Set(1, 2), menge.NewInt32Set(1), menge.NewInt32Set(2)},
	}
	for _, c := range cases {
		got := c.set.Difference(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt32Set_IsSubsetOf(t *testing.T) {
	cases := []struct {
		set  menge.Int32Set
		arg  menge.Int32Set
		want bool
	}{
		{menge.NewInt32Set(), menge.NewInt32Set(), true},
		{menge.NewInt32Set(1), menge.NewInt32Set(1), true},
		{menge.NewInt32Set(1), menge.NewInt32Set(1, 2), true},
		{menge.NewInt32Set(1, 2), menge.NewInt32Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt32Set_IsProperSubsetOf(t *testing.T) {
	cases := []struct {
		set  menge.Int32Set
		arg  menge.Int32Set
		want bool
	}{
		{menge.NewInt32Set(), menge.NewInt32Set(), false},
		{menge.NewInt32Set(1), menge.NewInt32Set(1), false},
		{menge.NewInt32Set(1), menge.NewInt32Set(1, 2), true},
		{menge.NewInt32Set(1, 2), menge.NewInt32Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsProperSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt32Set_IsSupersetOf(t *testing.T) {
	cases := []struct {
		set  menge.Int32Set
		arg  menge.Int32Set
		want bool
	}{
		{menge.NewInt32Set(), menge.NewInt32Set(), true},
		{menge.NewInt32Set(1), menge.NewInt32Set(1), true},
		{menge.NewInt32Set(1), menge.NewInt32Set(1, 2), false},
		{menge.NewInt32Set(1, 2), menge.NewInt32Set(1), true},
	}
	for _, c := range cases {
		got := c.set.IsSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt32Set_IsProperSupersetOf(t *testing.T) {
	cases := []struct {
		set  menge.Int32Set
		arg  menge.Int32Set
		want bool
	}{
		{menge.NewInt32Set(), menge.NewInt32Set(), false},
		{menge.NewInt32Set(1), menge.NewInt32Set(1), false},
		{menge.NewInt32Set(1), menge.NewInt32Set(1, 2), false},
		{menge.NewInt32Set(1, 2), menge.NewInt32Set(1), true},
	}
	for _, c := range cases {
		got := c.set.IsProperSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt32Set_IsDisjointFrom(t *testing.T) {
	cases := []struct {
		set  menge.Int32Set
		arg  menge.Int32Set
		want bool
	}{
		{menge.NewInt32Set(), menge.NewInt32Set(), true},
		{menge.NewInt32Set(1), menge.NewInt32Set(1), false},
		{menge.NewInt32Set(1), menge.NewInt32Set(2, 3), true},
		{menge.NewInt32Set(1, 2), menge.NewInt32Set(3), true},
		{menge.NewInt32Set(1), menge.NewInt32Set(1, 2), false},
		{menge.NewInt32Set(1, 2), menge.NewInt32Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsDisjointFrom(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
