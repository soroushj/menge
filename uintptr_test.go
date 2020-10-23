package menge

import (
	"testing"
)

func TestNewUIntPtrSet(t *testing.T) {
	cases := []struct {
		arg  []uintptr
		want UIntPtrSet
	}{
		{[]uintptr{}, UIntPtrSet{}},
		{[]uintptr{1, 1}, UIntPtrSet{1: struct{}{}}},
		{[]uintptr{1, 2}, UIntPtrSet{1: struct{}{}, 2: struct{}{}}},
	}
	for _, c := range cases {
		got := NewUIntPtrSet(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntPtrSet_Add(t *testing.T) {
	cases := []struct {
		set  UIntPtrSet
		arg  []uintptr
		want UIntPtrSet
	}{
		{NewUIntPtrSet(), []uintptr{}, NewUIntPtrSet()},
		{NewUIntPtrSet(), []uintptr{1, 1}, NewUIntPtrSet(1)},
		{NewUIntPtrSet(), []uintptr{1, 2}, NewUIntPtrSet(1, 2)},
		{NewUIntPtrSet(1), []uintptr{}, NewUIntPtrSet(1)},
		{NewUIntPtrSet(1), []uintptr{1, 1}, NewUIntPtrSet(1)},
		{NewUIntPtrSet(1), []uintptr{2, 3}, NewUIntPtrSet(1, 2, 3)},
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
		set  UIntPtrSet
		arg  []uintptr
		want UIntPtrSet
	}{
		{NewUIntPtrSet(), []uintptr{}, NewUIntPtrSet()},
		{NewUIntPtrSet(1), []uintptr{1, 1}, NewUIntPtrSet()},
		{NewUIntPtrSet(1, 2), []uintptr{1, 2}, NewUIntPtrSet()},
		{NewUIntPtrSet(1), []uintptr{}, NewUIntPtrSet(1)},
		{NewUIntPtrSet(1), []uintptr{1, 1}, NewUIntPtrSet()},
		{NewUIntPtrSet(1, 2), []uintptr{3}, NewUIntPtrSet(1, 2)},
		{NewUIntPtrSet(1, 2, 3), []uintptr{2, 3}, NewUIntPtrSet(1)},
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
		set  UIntPtrSet
		want UIntPtrSet
	}{
		{NewUIntPtrSet(), NewUIntPtrSet()},
		{NewUIntPtrSet(1, 2), NewUIntPtrSet()},
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
		set  UIntPtrSet
		arg  uintptr
		want bool
	}{
		{NewUIntPtrSet(), 1, false},
		{NewUIntPtrSet(2), 1, false},
		{NewUIntPtrSet(1), 1, true},
		{NewUIntPtrSet(1, 2), 1, true},
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
		set  UIntPtrSet
		want int
	}{
		{NewUIntPtrSet(), 0},
		{NewUIntPtrSet(1, 2), 2},
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
		set  UIntPtrSet
		want bool
	}{
		{NewUIntPtrSet(), true},
		{NewUIntPtrSet(1, 2), false},
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
		set  UIntPtrSet
		want UIntPtrSet
	}{
		{NewUIntPtrSet(), NewUIntPtrSet()},
		{NewUIntPtrSet(1, 2), NewUIntPtrSet(1, 2)},
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
		set  UIntPtrSet
		want []uintptr
	}{
		{NewUIntPtrSet(), []uintptr{}},
		{NewUIntPtrSet(1, 2), []uintptr{1, 2}},
	}
	for _, c := range cases {
		got := c.set.AsSlice()
		if len(got) != len(c.want) || !NewUIntPtrSet(got...).Equals(NewUIntPtrSet(c.want...)) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestUIntPtrSet_String(t *testing.T) {
	cases := []struct {
		set  UIntPtrSet
		want []string
	}{
		{NewUIntPtrSet(), []string{"{}"}},
		{NewUIntPtrSet(1), []string{"{1}"}},
		{NewUIntPtrSet(1, 2), []string{"{1 2}", "{2 1}"}},
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
		set  UIntPtrSet
		arg  UIntPtrSet
		want bool
	}{
		{NewUIntPtrSet(), NewUIntPtrSet(), true},
		{NewUIntPtrSet(1, 2), NewUIntPtrSet(2, 1), true},
		{NewUIntPtrSet(1, 2), NewUIntPtrSet(1), false},
		{NewUIntPtrSet(1), NewUIntPtrSet(1, 2), false},
		{NewUIntPtrSet(1), NewUIntPtrSet(2), false},
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
		set  UIntPtrSet
		arg  UIntPtrSet
		want UIntPtrSet
	}{
		{NewUIntPtrSet(), NewUIntPtrSet(), NewUIntPtrSet()},
		{NewUIntPtrSet(1), NewUIntPtrSet(1), NewUIntPtrSet(1)},
		{NewUIntPtrSet(1), NewUIntPtrSet(2), NewUIntPtrSet(1, 2)},
		{NewUIntPtrSet(1), NewUIntPtrSet(1, 2), NewUIntPtrSet(1, 2)},
		{NewUIntPtrSet(1, 2), NewUIntPtrSet(1), NewUIntPtrSet(1, 2)},
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
		set  UIntPtrSet
		arg  UIntPtrSet
		want UIntPtrSet
	}{
		{NewUIntPtrSet(), NewUIntPtrSet(), NewUIntPtrSet()},
		{NewUIntPtrSet(1), NewUIntPtrSet(1), NewUIntPtrSet(1)},
		{NewUIntPtrSet(1), NewUIntPtrSet(2), NewUIntPtrSet()},
		{NewUIntPtrSet(1), NewUIntPtrSet(1, 2), NewUIntPtrSet(1)},
		{NewUIntPtrSet(1, 2), NewUIntPtrSet(1), NewUIntPtrSet(1)},
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
		set  UIntPtrSet
		arg  UIntPtrSet
		want UIntPtrSet
	}{
		{NewUIntPtrSet(), NewUIntPtrSet(), NewUIntPtrSet()},
		{NewUIntPtrSet(1), NewUIntPtrSet(1), NewUIntPtrSet()},
		{NewUIntPtrSet(1), NewUIntPtrSet(2), NewUIntPtrSet(1)},
		{NewUIntPtrSet(1), NewUIntPtrSet(1, 2), NewUIntPtrSet()},
		{NewUIntPtrSet(1, 2), NewUIntPtrSet(1), NewUIntPtrSet(2)},
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
		set  UIntPtrSet
		arg  UIntPtrSet
		want bool
	}{
		{NewUIntPtrSet(), NewUIntPtrSet(), true},
		{NewUIntPtrSet(1), NewUIntPtrSet(1), true},
		{NewUIntPtrSet(1), NewUIntPtrSet(1, 2), true},
		{NewUIntPtrSet(1, 2), NewUIntPtrSet(1), false},
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
		set  UIntPtrSet
		arg  UIntPtrSet
		want bool
	}{
		{NewUIntPtrSet(), NewUIntPtrSet(), false},
		{NewUIntPtrSet(1), NewUIntPtrSet(1), false},
		{NewUIntPtrSet(1), NewUIntPtrSet(1, 2), true},
		{NewUIntPtrSet(1, 2), NewUIntPtrSet(1), false},
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
		set  UIntPtrSet
		arg  UIntPtrSet
		want bool
	}{
		{NewUIntPtrSet(), NewUIntPtrSet(), true},
		{NewUIntPtrSet(1), NewUIntPtrSet(1), true},
		{NewUIntPtrSet(1), NewUIntPtrSet(1, 2), false},
		{NewUIntPtrSet(1, 2), NewUIntPtrSet(1), true},
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
		set  UIntPtrSet
		arg  UIntPtrSet
		want bool
	}{
		{NewUIntPtrSet(), NewUIntPtrSet(), false},
		{NewUIntPtrSet(1), NewUIntPtrSet(1), false},
		{NewUIntPtrSet(1), NewUIntPtrSet(1, 2), false},
		{NewUIntPtrSet(1, 2), NewUIntPtrSet(1), true},
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
		set  UIntPtrSet
		arg  UIntPtrSet
		want bool
	}{
		{NewUIntPtrSet(), NewUIntPtrSet(), true},
		{NewUIntPtrSet(1), NewUIntPtrSet(1), false},
		{NewUIntPtrSet(1), NewUIntPtrSet(2, 3), true},
		{NewUIntPtrSet(1, 2), NewUIntPtrSet(3), true},
		{NewUIntPtrSet(1), NewUIntPtrSet(1, 2), false},
		{NewUIntPtrSet(1, 2), NewUIntPtrSet(1), false},
	}
	for _, c := range cases {
		got := c.set.IsDisjointFrom(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
