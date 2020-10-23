package menge

import (
	"testing"
)

func TestNewInt32Set(t *testing.T) {
	cases := []struct {
		arg  []int32
		want Int32Set
	}{
		{[]int32{}, Int32Set{}},
		{[]int32{1, 1}, Int32Set{1: struct{}{}}},
		{[]int32{1, 2}, Int32Set{1: struct{}{}, 2: struct{}{}}},
	}
	for _, c := range cases {
		got := NewInt32Set(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt32Set_Add(t *testing.T) {
	cases := []struct {
		set  Int32Set
		arg  []int32
		want Int32Set
	}{
		{NewInt32Set(), []int32{}, NewInt32Set()},
		{NewInt32Set(), []int32{1, 1}, NewInt32Set(1)},
		{NewInt32Set(), []int32{1, 2}, NewInt32Set(1, 2)},
		{NewInt32Set(1), []int32{}, NewInt32Set(1)},
		{NewInt32Set(1), []int32{1, 1}, NewInt32Set(1)},
		{NewInt32Set(1), []int32{2, 3}, NewInt32Set(1, 2, 3)},
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
		set  Int32Set
		arg  []int32
		want Int32Set
	}{
		{NewInt32Set(), []int32{}, NewInt32Set()},
		{NewInt32Set(1), []int32{1, 1}, NewInt32Set()},
		{NewInt32Set(1, 2), []int32{1, 2}, NewInt32Set()},
		{NewInt32Set(1), []int32{}, NewInt32Set(1)},
		{NewInt32Set(1), []int32{1, 1}, NewInt32Set()},
		{NewInt32Set(1, 2), []int32{3}, NewInt32Set(1, 2)},
		{NewInt32Set(1, 2, 3), []int32{2, 3}, NewInt32Set(1)},
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
		set  Int32Set
		want Int32Set
	}{
		{NewInt32Set(), NewInt32Set()},
		{NewInt32Set(1, 2), NewInt32Set()},
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
		set  Int32Set
		arg  int32
		want bool
	}{
		{NewInt32Set(), 1, false},
		{NewInt32Set(2), 1, false},
		{NewInt32Set(1), 1, true},
		{NewInt32Set(1, 2), 1, true},
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
		set  Int32Set
		want int
	}{
		{NewInt32Set(), 0},
		{NewInt32Set(1, 2), 2},
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
		set  Int32Set
		want bool
	}{
		{NewInt32Set(), true},
		{NewInt32Set(1, 2), false},
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
		set  Int32Set
		want Int32Set
	}{
		{NewInt32Set(), NewInt32Set()},
		{NewInt32Set(1, 2), NewInt32Set(1, 2)},
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
		set  Int32Set
		want []int32
	}{
		{NewInt32Set(), []int32{}},
		{NewInt32Set(1, 2), []int32{1, 2}},
	}
	for _, c := range cases {
		got := c.set.AsSlice()
		if len(got) != len(c.want) || !NewInt32Set(got...).Equals(NewInt32Set(c.want...)) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestInt32Set_String(t *testing.T) {
	cases := []struct {
		set  Int32Set
		want []string
	}{
		{NewInt32Set(), []string{"{}"}},
		{NewInt32Set(1), []string{"{1}"}},
		{NewInt32Set(1, 2), []string{"{1 2}", "{2 1}"}},
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
		set  Int32Set
		arg  Int32Set
		want bool
	}{
		{NewInt32Set(), NewInt32Set(), true},
		{NewInt32Set(1, 2), NewInt32Set(2, 1), true},
		{NewInt32Set(1, 2), NewInt32Set(1), false},
		{NewInt32Set(1), NewInt32Set(1, 2), false},
		{NewInt32Set(1), NewInt32Set(2), false},
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
		set  Int32Set
		arg  Int32Set
		want Int32Set
	}{
		{NewInt32Set(), NewInt32Set(), NewInt32Set()},
		{NewInt32Set(1), NewInt32Set(1), NewInt32Set(1)},
		{NewInt32Set(1), NewInt32Set(2), NewInt32Set(1, 2)},
		{NewInt32Set(1), NewInt32Set(1, 2), NewInt32Set(1, 2)},
		{NewInt32Set(1, 2), NewInt32Set(1), NewInt32Set(1, 2)},
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
		set  Int32Set
		arg  Int32Set
		want Int32Set
	}{
		{NewInt32Set(), NewInt32Set(), NewInt32Set()},
		{NewInt32Set(1), NewInt32Set(1), NewInt32Set(1)},
		{NewInt32Set(1), NewInt32Set(2), NewInt32Set()},
		{NewInt32Set(1), NewInt32Set(1, 2), NewInt32Set(1)},
		{NewInt32Set(1, 2), NewInt32Set(1), NewInt32Set(1)},
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
		set  Int32Set
		arg  Int32Set
		want Int32Set
	}{
		{NewInt32Set(), NewInt32Set(), NewInt32Set()},
		{NewInt32Set(1), NewInt32Set(1), NewInt32Set()},
		{NewInt32Set(1), NewInt32Set(2), NewInt32Set(1)},
		{NewInt32Set(1), NewInt32Set(1, 2), NewInt32Set()},
		{NewInt32Set(1, 2), NewInt32Set(1), NewInt32Set(2)},
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
		set  Int32Set
		arg  Int32Set
		want bool
	}{
		{NewInt32Set(), NewInt32Set(), true},
		{NewInt32Set(1), NewInt32Set(1), true},
		{NewInt32Set(1), NewInt32Set(1, 2), true},
		{NewInt32Set(1, 2), NewInt32Set(1), false},
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
		set  Int32Set
		arg  Int32Set
		want bool
	}{
		{NewInt32Set(), NewInt32Set(), false},
		{NewInt32Set(1), NewInt32Set(1), false},
		{NewInt32Set(1), NewInt32Set(1, 2), true},
		{NewInt32Set(1, 2), NewInt32Set(1), false},
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
		set  Int32Set
		arg  Int32Set
		want bool
	}{
		{NewInt32Set(), NewInt32Set(), true},
		{NewInt32Set(1), NewInt32Set(1), true},
		{NewInt32Set(1), NewInt32Set(1, 2), false},
		{NewInt32Set(1, 2), NewInt32Set(1), true},
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
		set  Int32Set
		arg  Int32Set
		want bool
	}{
		{NewInt32Set(), NewInt32Set(), false},
		{NewInt32Set(1), NewInt32Set(1), false},
		{NewInt32Set(1), NewInt32Set(1, 2), false},
		{NewInt32Set(1, 2), NewInt32Set(1), true},
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
		set  Int32Set
		arg  Int32Set
		want bool
	}{
		{NewInt32Set(), NewInt32Set(), true},
		{NewInt32Set(1), NewInt32Set(1), false},
		{NewInt32Set(1), NewInt32Set(2, 3), true},
		{NewInt32Set(1, 2), NewInt32Set(3), true},
		{NewInt32Set(1), NewInt32Set(1, 2), false},
		{NewInt32Set(1, 2), NewInt32Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsDisjointFrom(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
