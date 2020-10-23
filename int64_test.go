package menge

import (
	"testing"
)

func TestNewInt64Set(t *testing.T) {
	cases := []struct {
		arg  []int64
		want Int64Set
	}{
		{[]int64{}, Int64Set{}},
		{[]int64{1, 1}, Int64Set{1: struct{}{}}},
		{[]int64{1, 2}, Int64Set{1: struct{}{}, 2: struct{}{}}},
	}
	for _, c := range cases {
		got := NewInt64Set(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt64Set_Add(t *testing.T) {
	cases := []struct {
		set  Int64Set
		arg  []int64
		want Int64Set
	}{
		{NewInt64Set(), []int64{}, NewInt64Set()},
		{NewInt64Set(), []int64{1, 1}, NewInt64Set(1)},
		{NewInt64Set(), []int64{1, 2}, NewInt64Set(1, 2)},
		{NewInt64Set(1), []int64{}, NewInt64Set(1)},
		{NewInt64Set(1), []int64{1, 1}, NewInt64Set(1)},
		{NewInt64Set(1), []int64{2, 3}, NewInt64Set(1, 2, 3)},
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
		set  Int64Set
		arg  []int64
		want Int64Set
	}{
		{NewInt64Set(), []int64{}, NewInt64Set()},
		{NewInt64Set(1), []int64{1, 1}, NewInt64Set()},
		{NewInt64Set(1, 2), []int64{1, 2}, NewInt64Set()},
		{NewInt64Set(1), []int64{}, NewInt64Set(1)},
		{NewInt64Set(1), []int64{1, 1}, NewInt64Set()},
		{NewInt64Set(1, 2), []int64{3}, NewInt64Set(1, 2)},
		{NewInt64Set(1, 2, 3), []int64{2, 3}, NewInt64Set(1)},
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
		set  Int64Set
		want Int64Set
	}{
		{NewInt64Set(), NewInt64Set()},
		{NewInt64Set(1, 2), NewInt64Set()},
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
		set  Int64Set
		arg  int64
		want bool
	}{
		{NewInt64Set(), 1, false},
		{NewInt64Set(2), 1, false},
		{NewInt64Set(1), 1, true},
		{NewInt64Set(1, 2), 1, true},
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
		set  Int64Set
		want int
	}{
		{NewInt64Set(), 0},
		{NewInt64Set(1, 2), 2},
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
		set  Int64Set
		want bool
	}{
		{NewInt64Set(), true},
		{NewInt64Set(1, 2), false},
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
		set  Int64Set
		want Int64Set
	}{
		{NewInt64Set(), NewInt64Set()},
		{NewInt64Set(1, 2), NewInt64Set(1, 2)},
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
		set  Int64Set
		want []int64
	}{
		{NewInt64Set(), []int64{}},
		{NewInt64Set(1, 2), []int64{1, 2}},
	}
	for _, c := range cases {
		got := c.set.AsSlice()
		if len(got) != len(c.want) || !NewInt64Set(got...).Equals(NewInt64Set(c.want...)) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt64Set_String(t *testing.T) {
	cases := []struct {
		set  Int64Set
		want []string
	}{
		{NewInt64Set(), []string{"{}"}},
		{NewInt64Set(1), []string{"{1}"}},
		{NewInt64Set(1, 2), []string{"{1 2}", "{2 1}"}},
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
		set  Int64Set
		arg  Int64Set
		want bool
	}{
		{NewInt64Set(), NewInt64Set(), true},
		{NewInt64Set(1, 2), NewInt64Set(2, 1), true},
		{NewInt64Set(1, 2), NewInt64Set(1), false},
		{NewInt64Set(1), NewInt64Set(1, 2), false},
		{NewInt64Set(1), NewInt64Set(2), false},
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
		set  Int64Set
		arg  Int64Set
		want Int64Set
	}{
		{NewInt64Set(), NewInt64Set(), NewInt64Set()},
		{NewInt64Set(1), NewInt64Set(1), NewInt64Set(1)},
		{NewInt64Set(1), NewInt64Set(2), NewInt64Set(1, 2)},
		{NewInt64Set(1), NewInt64Set(1, 2), NewInt64Set(1, 2)},
		{NewInt64Set(1, 2), NewInt64Set(1), NewInt64Set(1, 2)},
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
		set  Int64Set
		arg  Int64Set
		want Int64Set
	}{
		{NewInt64Set(), NewInt64Set(), NewInt64Set()},
		{NewInt64Set(1), NewInt64Set(1), NewInt64Set(1)},
		{NewInt64Set(1), NewInt64Set(2), NewInt64Set()},
		{NewInt64Set(1), NewInt64Set(1, 2), NewInt64Set(1)},
		{NewInt64Set(1, 2), NewInt64Set(1), NewInt64Set(1)},
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
		set  Int64Set
		arg  Int64Set
		want Int64Set
	}{
		{NewInt64Set(), NewInt64Set(), NewInt64Set()},
		{NewInt64Set(1), NewInt64Set(1), NewInt64Set()},
		{NewInt64Set(1), NewInt64Set(2), NewInt64Set(1)},
		{NewInt64Set(1), NewInt64Set(1, 2), NewInt64Set()},
		{NewInt64Set(1, 2), NewInt64Set(1), NewInt64Set(2)},
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
		set  Int64Set
		arg  Int64Set
		want bool
	}{
		{NewInt64Set(), NewInt64Set(), true},
		{NewInt64Set(1), NewInt64Set(1), true},
		{NewInt64Set(1), NewInt64Set(1, 2), true},
		{NewInt64Set(1, 2), NewInt64Set(1), false},
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
		set  Int64Set
		arg  Int64Set
		want bool
	}{
		{NewInt64Set(), NewInt64Set(), false},
		{NewInt64Set(1), NewInt64Set(1), false},
		{NewInt64Set(1), NewInt64Set(1, 2), true},
		{NewInt64Set(1, 2), NewInt64Set(1), false},
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
		set  Int64Set
		arg  Int64Set
		want bool
	}{
		{NewInt64Set(), NewInt64Set(), true},
		{NewInt64Set(1), NewInt64Set(1), true},
		{NewInt64Set(1), NewInt64Set(1, 2), false},
		{NewInt64Set(1, 2), NewInt64Set(1), true},
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
		set  Int64Set
		arg  Int64Set
		want bool
	}{
		{NewInt64Set(), NewInt64Set(), false},
		{NewInt64Set(1), NewInt64Set(1), false},
		{NewInt64Set(1), NewInt64Set(1, 2), false},
		{NewInt64Set(1, 2), NewInt64Set(1), true},
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
		set  Int64Set
		arg  Int64Set
		want bool
	}{
		{NewInt64Set(), NewInt64Set(), true},
		{NewInt64Set(1), NewInt64Set(1), false},
		{NewInt64Set(1), NewInt64Set(2, 3), true},
		{NewInt64Set(1, 2), NewInt64Set(3), true},
		{NewInt64Set(1), NewInt64Set(1, 2), false},
		{NewInt64Set(1, 2), NewInt64Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsDisjointFrom(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
