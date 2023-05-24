package menge_test

import (
	"testing"

	"github.com/soroushj/menge"
)

func TestNewInt64Set(t *testing.T) {
	cases := []struct {
		arg  []int64
		want menge.Int64Set
	}{
		{[]int64{}, menge.Int64Set{}},
		{[]int64{1, 1}, menge.Int64Set{1: struct{}{}}},
		{[]int64{1, 2}, menge.Int64Set{1: struct{}{}, 2: struct{}{}}},
	}
	for _, c := range cases {
		got := menge.NewInt64Set(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt64Set_Add(t *testing.T) {
	cases := []struct {
		set  menge.Int64Set
		arg  []int64
		want menge.Int64Set
	}{
		{menge.NewInt64Set(), []int64{}, menge.NewInt64Set()},
		{menge.NewInt64Set(), []int64{1, 1}, menge.NewInt64Set(1)},
		{menge.NewInt64Set(), []int64{1, 2}, menge.NewInt64Set(1, 2)},
		{menge.NewInt64Set(1), []int64{}, menge.NewInt64Set(1)},
		{menge.NewInt64Set(1), []int64{1, 1}, menge.NewInt64Set(1)},
		{menge.NewInt64Set(1), []int64{2, 3}, menge.NewInt64Set(1, 2, 3)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Add(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt64Set_Remove(t *testing.T) {
	cases := []struct {
		set  menge.Int64Set
		arg  []int64
		want menge.Int64Set
	}{
		{menge.NewInt64Set(), []int64{}, menge.NewInt64Set()},
		{menge.NewInt64Set(1), []int64{1, 1}, menge.NewInt64Set()},
		{menge.NewInt64Set(1, 2), []int64{1, 2}, menge.NewInt64Set()},
		{menge.NewInt64Set(1), []int64{}, menge.NewInt64Set(1)},
		{menge.NewInt64Set(1), []int64{1, 1}, menge.NewInt64Set()},
		{menge.NewInt64Set(1, 2), []int64{3}, menge.NewInt64Set(1, 2)},
		{menge.NewInt64Set(1, 2, 3), []int64{2, 3}, menge.NewInt64Set(1)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Remove(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt64Set_Empty(t *testing.T) {
	cases := []struct {
		set  menge.Int64Set
		want menge.Int64Set
	}{
		{menge.NewInt64Set(), menge.NewInt64Set()},
		{menge.NewInt64Set(1, 2), menge.NewInt64Set()},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Empty()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt64Set_Has(t *testing.T) {
	cases := []struct {
		set  menge.Int64Set
		arg  int64
		want bool
	}{
		{menge.NewInt64Set(), 1, false},
		{menge.NewInt64Set(2), 1, false},
		{menge.NewInt64Set(1), 1, true},
		{menge.NewInt64Set(1, 2), 1, true},
	}
	for _, c := range cases {
		got := c.set.Has(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt64Set_Size(t *testing.T) {
	cases := []struct {
		set  menge.Int64Set
		want int
	}{
		{menge.NewInt64Set(), 0},
		{menge.NewInt64Set(1, 2), 2},
	}
	for _, c := range cases {
		got := c.set.Size()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt64Set_IsEmpty(t *testing.T) {
	cases := []struct {
		set  menge.Int64Set
		want bool
	}{
		{menge.NewInt64Set(), true},
		{menge.NewInt64Set(1, 2), false},
	}
	for _, c := range cases {
		got := c.set.IsEmpty()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt64Set_Clone(t *testing.T) {
	cases := []struct {
		set  menge.Int64Set
		want menge.Int64Set
	}{
		{menge.NewInt64Set(), menge.NewInt64Set()},
		{menge.NewInt64Set(1, 2), menge.NewInt64Set(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt64Set_AsSlice(t *testing.T) {
	cases := []struct {
		set  menge.Int64Set
		want []int64
	}{
		{menge.NewInt64Set(), []int64{}},
		{menge.NewInt64Set(1, 2), []int64{1, 2}},
	}
	for _, c := range cases {
		got := c.set.AsSlice()
		if len(got) != len(c.want) || !menge.NewInt64Set(got...).Equals(menge.NewInt64Set(c.want...)) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt64Set_String(t *testing.T) {
	cases := []struct {
		set  menge.Int64Set
		want []string
	}{
		{menge.NewInt64Set(), []string{"{}"}},
		{menge.NewInt64Set(1), []string{"{1}"}},
		{menge.NewInt64Set(1, 2), []string{"{1 2}", "{2 1}"}},
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

func TestInt64Set_Equals(t *testing.T) {
	cases := []struct {
		set  menge.Int64Set
		arg  menge.Int64Set
		want bool
	}{
		{menge.NewInt64Set(), menge.NewInt64Set(), true},
		{menge.NewInt64Set(1, 2), menge.NewInt64Set(2, 1), true},
		{menge.NewInt64Set(1, 2), menge.NewInt64Set(1), false},
		{menge.NewInt64Set(1), menge.NewInt64Set(1, 2), false},
		{menge.NewInt64Set(1), menge.NewInt64Set(2), false},
	}
	for _, c := range cases {
		got := c.set.Equals(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt64Set_Union(t *testing.T) {
	cases := []struct {
		set  menge.Int64Set
		arg  menge.Int64Set
		want menge.Int64Set
	}{
		{menge.NewInt64Set(), menge.NewInt64Set(), menge.NewInt64Set()},
		{menge.NewInt64Set(1), menge.NewInt64Set(1), menge.NewInt64Set(1)},
		{menge.NewInt64Set(1), menge.NewInt64Set(2), menge.NewInt64Set(1, 2)},
		{menge.NewInt64Set(1), menge.NewInt64Set(1, 2), menge.NewInt64Set(1, 2)},
		{menge.NewInt64Set(1, 2), menge.NewInt64Set(1), menge.NewInt64Set(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Union(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt64Set_Intersection(t *testing.T) {
	cases := []struct {
		set  menge.Int64Set
		arg  menge.Int64Set
		want menge.Int64Set
	}{
		{menge.NewInt64Set(), menge.NewInt64Set(), menge.NewInt64Set()},
		{menge.NewInt64Set(1), menge.NewInt64Set(1), menge.NewInt64Set(1)},
		{menge.NewInt64Set(1), menge.NewInt64Set(2), menge.NewInt64Set()},
		{menge.NewInt64Set(1), menge.NewInt64Set(1, 2), menge.NewInt64Set(1)},
		{menge.NewInt64Set(1, 2), menge.NewInt64Set(1), menge.NewInt64Set(1)},
	}
	for _, c := range cases {
		got := c.set.Intersection(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt64Set_Difference(t *testing.T) {
	cases := []struct {
		set  menge.Int64Set
		arg  menge.Int64Set
		want menge.Int64Set
	}{
		{menge.NewInt64Set(), menge.NewInt64Set(), menge.NewInt64Set()},
		{menge.NewInt64Set(1), menge.NewInt64Set(1), menge.NewInt64Set()},
		{menge.NewInt64Set(1), menge.NewInt64Set(2), menge.NewInt64Set(1)},
		{menge.NewInt64Set(1), menge.NewInt64Set(1, 2), menge.NewInt64Set()},
		{menge.NewInt64Set(1, 2), menge.NewInt64Set(1), menge.NewInt64Set(2)},
	}
	for _, c := range cases {
		got := c.set.Difference(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt64Set_IsSubsetOf(t *testing.T) {
	cases := []struct {
		set  menge.Int64Set
		arg  menge.Int64Set
		want bool
	}{
		{menge.NewInt64Set(), menge.NewInt64Set(), true},
		{menge.NewInt64Set(1), menge.NewInt64Set(1), true},
		{menge.NewInt64Set(1), menge.NewInt64Set(1, 2), true},
		{menge.NewInt64Set(1, 2), menge.NewInt64Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt64Set_IsProperSubsetOf(t *testing.T) {
	cases := []struct {
		set  menge.Int64Set
		arg  menge.Int64Set
		want bool
	}{
		{menge.NewInt64Set(), menge.NewInt64Set(), false},
		{menge.NewInt64Set(1), menge.NewInt64Set(1), false},
		{menge.NewInt64Set(1), menge.NewInt64Set(1, 2), true},
		{menge.NewInt64Set(1, 2), menge.NewInt64Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsProperSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt64Set_IsSupersetOf(t *testing.T) {
	cases := []struct {
		set  menge.Int64Set
		arg  menge.Int64Set
		want bool
	}{
		{menge.NewInt64Set(), menge.NewInt64Set(), true},
		{menge.NewInt64Set(1), menge.NewInt64Set(1), true},
		{menge.NewInt64Set(1), menge.NewInt64Set(1, 2), false},
		{menge.NewInt64Set(1, 2), menge.NewInt64Set(1), true},
	}
	for _, c := range cases {
		got := c.set.IsSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt64Set_IsProperSupersetOf(t *testing.T) {
	cases := []struct {
		set  menge.Int64Set
		arg  menge.Int64Set
		want bool
	}{
		{menge.NewInt64Set(), menge.NewInt64Set(), false},
		{menge.NewInt64Set(1), menge.NewInt64Set(1), false},
		{menge.NewInt64Set(1), menge.NewInt64Set(1, 2), false},
		{menge.NewInt64Set(1, 2), menge.NewInt64Set(1), true},
	}
	for _, c := range cases {
		got := c.set.IsProperSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt64Set_IsDisjointFrom(t *testing.T) {
	cases := []struct {
		set  menge.Int64Set
		arg  menge.Int64Set
		want bool
	}{
		{menge.NewInt64Set(), menge.NewInt64Set(), true},
		{menge.NewInt64Set(1), menge.NewInt64Set(1), false},
		{menge.NewInt64Set(1), menge.NewInt64Set(2, 3), true},
		{menge.NewInt64Set(1, 2), menge.NewInt64Set(3), true},
		{menge.NewInt64Set(1), menge.NewInt64Set(1, 2), false},
		{menge.NewInt64Set(1, 2), menge.NewInt64Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsDisjointFrom(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
