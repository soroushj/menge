package menge

import (
	"testing"
)

func TestNewIntSet(t *testing.T) {
	cases := []struct {
		arg  []int
		want IntSet
	}{
		{[]int{}, IntSet{}},
		{[]int{1, 1}, IntSet{1: struct{}{}}},
		{[]int{1, 2}, IntSet{1: struct{}{}, 2: struct{}{}}},
	}
	for _, c := range cases {
		got := NewIntSet(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestIntSet_Add(t *testing.T) {
	cases := []struct {
		set  IntSet
		arg  []int
		want IntSet
	}{
		{NewIntSet(), []int{}, NewIntSet()},
		{NewIntSet(), []int{1, 1}, NewIntSet(1)},
		{NewIntSet(), []int{1, 2}, NewIntSet(1, 2)},
		{NewIntSet(1), []int{}, NewIntSet(1)},
		{NewIntSet(1), []int{1, 1}, NewIntSet(1)},
		{NewIntSet(1), []int{2, 3}, NewIntSet(1, 2, 3)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Add(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestIntSet_Remove(t *testing.T) {
	cases := []struct {
		set  IntSet
		arg  []int
		want IntSet
	}{
		{NewIntSet(), []int{}, NewIntSet()},
		{NewIntSet(1), []int{1, 1}, NewIntSet()},
		{NewIntSet(1, 2), []int{1, 2}, NewIntSet()},
		{NewIntSet(1), []int{}, NewIntSet(1)},
		{NewIntSet(1), []int{1, 1}, NewIntSet()},
		{NewIntSet(1, 2), []int{3}, NewIntSet(1, 2)},
		{NewIntSet(1, 2, 3), []int{2, 3}, NewIntSet(1)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Remove(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestIntSet_Empty(t *testing.T) {
	cases := []struct {
		set  IntSet
		want IntSet
	}{
		{NewIntSet(), NewIntSet()},
		{NewIntSet(1, 2), NewIntSet()},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Empty()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestIntSet_Has(t *testing.T) {
	cases := []struct {
		set  IntSet
		arg  int
		want bool
	}{
		{NewIntSet(), 1, false},
		{NewIntSet(2), 1, false},
		{NewIntSet(1), 1, true},
		{NewIntSet(1, 2), 1, true},
	}
	for _, c := range cases {
		got := c.set.Has(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestIntSet_Size(t *testing.T) {
	cases := []struct {
		set  IntSet
		want int
	}{
		{NewIntSet(), 0},
		{NewIntSet(1, 2), 2},
	}
	for _, c := range cases {
		got := c.set.Size()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestIntSet_IsEmpty(t *testing.T) {
	cases := []struct {
		set  IntSet
		want bool
	}{
		{NewIntSet(), true},
		{NewIntSet(1, 2), false},
	}
	for _, c := range cases {
		got := c.set.IsEmpty()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestIntSet_Clone(t *testing.T) {
	cases := []struct {
		set  IntSet
		want IntSet
	}{
		{NewIntSet(), NewIntSet()},
		{NewIntSet(1, 2), NewIntSet(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestIntSet_AsSlice(t *testing.T) {
	cases := []struct {
		set  IntSet
		want []int
	}{
		{NewIntSet(), []int{}},
		{NewIntSet(1, 2), []int{1, 2}},
	}
	for _, c := range cases {
		got := c.set.AsSlice()
		if len(got) != len(c.want) || !NewIntSet(got...).Equals(NewIntSet(c.want...)) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestIntSet_String(t *testing.T) {
	cases := []struct {
		set  IntSet
		want []string
	}{
		{NewIntSet(), []string{"{}"}},
		{NewIntSet(1), []string{"{1}"}},
		{NewIntSet(1, 2), []string{"{1 2}", "{2 1}"}},
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

func TestIntSet_Equals(t *testing.T) {
	cases := []struct {
		set  IntSet
		arg  IntSet
		want bool
	}{
		{NewIntSet(), NewIntSet(), true},
		{NewIntSet(1, 2), NewIntSet(2, 1), true},
		{NewIntSet(1, 2), NewIntSet(1), false},
		{NewIntSet(1), NewIntSet(1, 2), false},
		{NewIntSet(1), NewIntSet(2), false},
	}
	for _, c := range cases {
		got := c.set.Equals(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestIntSet_Union(t *testing.T) {
	cases := []struct {
		set  IntSet
		arg  IntSet
		want IntSet
	}{
		{NewIntSet(), NewIntSet(), NewIntSet()},
		{NewIntSet(1), NewIntSet(1), NewIntSet(1)},
		{NewIntSet(1), NewIntSet(2), NewIntSet(1, 2)},
		{NewIntSet(1), NewIntSet(1, 2), NewIntSet(1, 2)},
		{NewIntSet(1, 2), NewIntSet(1), NewIntSet(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Union(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestIntSet_Intersection(t *testing.T) {
	cases := []struct {
		set  IntSet
		arg  IntSet
		want IntSet
	}{
		{NewIntSet(), NewIntSet(), NewIntSet()},
		{NewIntSet(1), NewIntSet(1), NewIntSet(1)},
		{NewIntSet(1), NewIntSet(2), NewIntSet()},
		{NewIntSet(1), NewIntSet(1, 2), NewIntSet(1)},
		{NewIntSet(1, 2), NewIntSet(1), NewIntSet(1)},
	}
	for _, c := range cases {
		got := c.set.Intersection(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestIntSet_Difference(t *testing.T) {
	cases := []struct {
		set  IntSet
		arg  IntSet
		want IntSet
	}{
		{NewIntSet(), NewIntSet(), NewIntSet()},
		{NewIntSet(1), NewIntSet(1), NewIntSet()},
		{NewIntSet(1), NewIntSet(2), NewIntSet(1)},
		{NewIntSet(1), NewIntSet(1, 2), NewIntSet()},
		{NewIntSet(1, 2), NewIntSet(1), NewIntSet(2)},
	}
	for _, c := range cases {
		got := c.set.Difference(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestIntSet_IsSubsetOf(t *testing.T) {
	cases := []struct {
		set  IntSet
		arg  IntSet
		want bool
	}{
		{NewIntSet(), NewIntSet(), true},
		{NewIntSet(1), NewIntSet(1), true},
		{NewIntSet(1), NewIntSet(1, 2), true},
		{NewIntSet(1, 2), NewIntSet(1), false},
	}
	for _, c := range cases {
		got := c.set.IsSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestIntSet_IsProperSubsetOf(t *testing.T) {
	cases := []struct {
		set  IntSet
		arg  IntSet
		want bool
	}{
		{NewIntSet(), NewIntSet(), false},
		{NewIntSet(1), NewIntSet(1), false},
		{NewIntSet(1), NewIntSet(1, 2), true},
		{NewIntSet(1, 2), NewIntSet(1), false},
	}
	for _, c := range cases {
		got := c.set.IsProperSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestIntSet_IsSupersetOf(t *testing.T) {
	cases := []struct {
		set  IntSet
		arg  IntSet
		want bool
	}{
		{NewIntSet(), NewIntSet(), true},
		{NewIntSet(1), NewIntSet(1), true},
		{NewIntSet(1), NewIntSet(1, 2), false},
		{NewIntSet(1, 2), NewIntSet(1), true},
	}
	for _, c := range cases {
		got := c.set.IsSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestIntSet_IsProperSupersetOf(t *testing.T) {
	cases := []struct {
		set  IntSet
		arg  IntSet
		want bool
	}{
		{NewIntSet(), NewIntSet(), false},
		{NewIntSet(1), NewIntSet(1), false},
		{NewIntSet(1), NewIntSet(1, 2), false},
		{NewIntSet(1, 2), NewIntSet(1), true},
	}
	for _, c := range cases {
		got := c.set.IsProperSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestIntSet_IsDisjointFrom(t *testing.T) {
	cases := []struct {
		set  IntSet
		arg  IntSet
		want bool
	}{
		{NewIntSet(), NewIntSet(), true},
		{NewIntSet(1), NewIntSet(1), false},
		{NewIntSet(1), NewIntSet(2, 3), true},
		{NewIntSet(1, 2), NewIntSet(3), true},
		{NewIntSet(1), NewIntSet(1, 2), false},
		{NewIntSet(1, 2), NewIntSet(1), false},
	}
	for _, c := range cases {
		got := c.set.IsDisjointFrom(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
