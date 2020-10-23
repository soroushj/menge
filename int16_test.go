package menge

import (
	"testing"
)

func TestNewInt16Set(t *testing.T) {
	cases := []struct {
		arg  []int16
		want Int16Set
	}{
		{[]int16{}, Int16Set{}},
		{[]int16{1, 1}, Int16Set{1: struct{}{}}},
		{[]int16{1, 2}, Int16Set{1: struct{}{}, 2: struct{}{}}},
	}
	for _, c := range cases {
		got := NewInt16Set(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt16Set_Add(t *testing.T) {
	cases := []struct {
		set  Int16Set
		arg  []int16
		want Int16Set
	}{
		{NewInt16Set(), []int16{}, NewInt16Set()},
		{NewInt16Set(), []int16{1, 1}, NewInt16Set(1)},
		{NewInt16Set(), []int16{1, 2}, NewInt16Set(1, 2)},
		{NewInt16Set(1), []int16{}, NewInt16Set(1)},
		{NewInt16Set(1), []int16{1, 1}, NewInt16Set(1)},
		{NewInt16Set(1), []int16{2, 3}, NewInt16Set(1, 2, 3)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Add(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt16Set_Remove(t *testing.T) {
	cases := []struct {
		set  Int16Set
		arg  []int16
		want Int16Set
	}{
		{NewInt16Set(), []int16{}, NewInt16Set()},
		{NewInt16Set(1), []int16{1, 1}, NewInt16Set()},
		{NewInt16Set(1, 2), []int16{1, 2}, NewInt16Set()},
		{NewInt16Set(1), []int16{}, NewInt16Set(1)},
		{NewInt16Set(1), []int16{1, 1}, NewInt16Set()},
		{NewInt16Set(1, 2), []int16{3}, NewInt16Set(1, 2)},
		{NewInt16Set(1, 2, 3), []int16{2, 3}, NewInt16Set(1)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Remove(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt16Set_Empty(t *testing.T) {
	cases := []struct {
		set  Int16Set
		want Int16Set
	}{
		{NewInt16Set(), NewInt16Set()},
		{NewInt16Set(1, 2), NewInt16Set()},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Empty()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt16Set_Has(t *testing.T) {
	cases := []struct {
		set  Int16Set
		arg  int16
		want bool
	}{
		{NewInt16Set(), 1, false},
		{NewInt16Set(2), 1, false},
		{NewInt16Set(1), 1, true},
		{NewInt16Set(1, 2), 1, true},
	}
	for _, c := range cases {
		got := c.set.Has(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt16Set_Size(t *testing.T) {
	cases := []struct {
		set  Int16Set
		want int
	}{
		{NewInt16Set(), 0},
		{NewInt16Set(1, 2), 2},
	}
	for _, c := range cases {
		got := c.set.Size()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt16Set_IsEmpty(t *testing.T) {
	cases := []struct {
		set  Int16Set
		want bool
	}{
		{NewInt16Set(), true},
		{NewInt16Set(1, 2), false},
	}
	for _, c := range cases {
		got := c.set.IsEmpty()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt16Set_Clone(t *testing.T) {
	cases := []struct {
		set  Int16Set
		want Int16Set
	}{
		{NewInt16Set(), NewInt16Set()},
		{NewInt16Set(1, 2), NewInt16Set(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt16Set_AsSlice(t *testing.T) {
	cases := []struct {
		set  Int16Set
		want []int16
	}{
		{NewInt16Set(), []int16{}},
		{NewInt16Set(1, 2), []int16{1, 2}},
	}
	for _, c := range cases {
		got := c.set.AsSlice()
		if len(got) != len(c.want) || !NewInt16Set(got...).Equals(NewInt16Set(c.want...)) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt16Set_String(t *testing.T) {
	cases := []struct {
		set  Int16Set
		want []string
	}{
		{NewInt16Set(), []string{"{}"}},
		{NewInt16Set(1), []string{"{1}"}},
		{NewInt16Set(1, 2), []string{"{1 2}", "{2 1}"}},
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

func TestInt16Set_Equals(t *testing.T) {
	cases := []struct {
		set  Int16Set
		arg  Int16Set
		want bool
	}{
		{NewInt16Set(), NewInt16Set(), true},
		{NewInt16Set(1, 2), NewInt16Set(2, 1), true},
		{NewInt16Set(1, 2), NewInt16Set(1), false},
		{NewInt16Set(1), NewInt16Set(1, 2), false},
		{NewInt16Set(1), NewInt16Set(2), false},
	}
	for _, c := range cases {
		got := c.set.Equals(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt16Set_Union(t *testing.T) {
	cases := []struct {
		set  Int16Set
		arg  Int16Set
		want Int16Set
	}{
		{NewInt16Set(), NewInt16Set(), NewInt16Set()},
		{NewInt16Set(1), NewInt16Set(1), NewInt16Set(1)},
		{NewInt16Set(1), NewInt16Set(2), NewInt16Set(1, 2)},
		{NewInt16Set(1), NewInt16Set(1, 2), NewInt16Set(1, 2)},
		{NewInt16Set(1, 2), NewInt16Set(1), NewInt16Set(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Union(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt16Set_Intersection(t *testing.T) {
	cases := []struct {
		set  Int16Set
		arg  Int16Set
		want Int16Set
	}{
		{NewInt16Set(), NewInt16Set(), NewInt16Set()},
		{NewInt16Set(1), NewInt16Set(1), NewInt16Set(1)},
		{NewInt16Set(1), NewInt16Set(2), NewInt16Set()},
		{NewInt16Set(1), NewInt16Set(1, 2), NewInt16Set(1)},
		{NewInt16Set(1, 2), NewInt16Set(1), NewInt16Set(1)},
	}
	for _, c := range cases {
		got := c.set.Intersection(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt16Set_Difference(t *testing.T) {
	cases := []struct {
		set  Int16Set
		arg  Int16Set
		want Int16Set
	}{
		{NewInt16Set(), NewInt16Set(), NewInt16Set()},
		{NewInt16Set(1), NewInt16Set(1), NewInt16Set()},
		{NewInt16Set(1), NewInt16Set(2), NewInt16Set(1)},
		{NewInt16Set(1), NewInt16Set(1, 2), NewInt16Set()},
		{NewInt16Set(1, 2), NewInt16Set(1), NewInt16Set(2)},
	}
	for _, c := range cases {
		got := c.set.Difference(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt16Set_IsSubsetOf(t *testing.T) {
	cases := []struct {
		set  Int16Set
		arg  Int16Set
		want bool
	}{
		{NewInt16Set(), NewInt16Set(), true},
		{NewInt16Set(1), NewInt16Set(1), true},
		{NewInt16Set(1), NewInt16Set(1, 2), true},
		{NewInt16Set(1, 2), NewInt16Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt16Set_IsProperSubsetOf(t *testing.T) {
	cases := []struct {
		set  Int16Set
		arg  Int16Set
		want bool
	}{
		{NewInt16Set(), NewInt16Set(), false},
		{NewInt16Set(1), NewInt16Set(1), false},
		{NewInt16Set(1), NewInt16Set(1, 2), true},
		{NewInt16Set(1, 2), NewInt16Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsProperSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt16Set_IsSupersetOf(t *testing.T) {
	cases := []struct {
		set  Int16Set
		arg  Int16Set
		want bool
	}{
		{NewInt16Set(), NewInt16Set(), true},
		{NewInt16Set(1), NewInt16Set(1), true},
		{NewInt16Set(1), NewInt16Set(1, 2), false},
		{NewInt16Set(1, 2), NewInt16Set(1), true},
	}
	for _, c := range cases {
		got := c.set.IsSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt16Set_IsProperSupersetOf(t *testing.T) {
	cases := []struct {
		set  Int16Set
		arg  Int16Set
		want bool
	}{
		{NewInt16Set(), NewInt16Set(), false},
		{NewInt16Set(1), NewInt16Set(1), false},
		{NewInt16Set(1), NewInt16Set(1, 2), false},
		{NewInt16Set(1, 2), NewInt16Set(1), true},
	}
	for _, c := range cases {
		got := c.set.IsProperSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt16Set_IsDisjointFrom(t *testing.T) {
	cases := []struct {
		set  Int16Set
		arg  Int16Set
		want bool
	}{
		{NewInt16Set(), NewInt16Set(), true},
		{NewInt16Set(1), NewInt16Set(1), false},
		{NewInt16Set(1), NewInt16Set(2, 3), true},
		{NewInt16Set(1, 2), NewInt16Set(3), true},
		{NewInt16Set(1), NewInt16Set(1, 2), false},
		{NewInt16Set(1, 2), NewInt16Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsDisjointFrom(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
