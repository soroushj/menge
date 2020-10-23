package menge

import (
	"testing"
)

func TestNewComplex128Set(t *testing.T) {
	cases := []struct {
		arg  []complex128
		want Complex128Set
	}{
		{[]complex128{}, Complex128Set{}},
		{[]complex128{1, 1}, Complex128Set{1: struct{}{}}},
		{[]complex128{1, 2}, Complex128Set{1: struct{}{}, 2: struct{}{}}},
	}
	for _, c := range cases {
		got := NewComplex128Set(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex128Set_Add(t *testing.T) {
	cases := []struct {
		set  Complex128Set
		arg  []complex128
		want Complex128Set
	}{
		{NewComplex128Set(), []complex128{}, NewComplex128Set()},
		{NewComplex128Set(), []complex128{1, 1}, NewComplex128Set(1)},
		{NewComplex128Set(), []complex128{1, 2}, NewComplex128Set(1, 2)},
		{NewComplex128Set(1), []complex128{}, NewComplex128Set(1)},
		{NewComplex128Set(1), []complex128{1, 1}, NewComplex128Set(1)},
		{NewComplex128Set(1), []complex128{2, 3}, NewComplex128Set(1, 2, 3)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Add(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex128Set_Remove(t *testing.T) {
	cases := []struct {
		set  Complex128Set
		arg  []complex128
		want Complex128Set
	}{
		{NewComplex128Set(), []complex128{}, NewComplex128Set()},
		{NewComplex128Set(1), []complex128{1, 1}, NewComplex128Set()},
		{NewComplex128Set(1, 2), []complex128{1, 2}, NewComplex128Set()},
		{NewComplex128Set(1), []complex128{}, NewComplex128Set(1)},
		{NewComplex128Set(1), []complex128{1, 1}, NewComplex128Set()},
		{NewComplex128Set(1, 2), []complex128{3}, NewComplex128Set(1, 2)},
		{NewComplex128Set(1, 2, 3), []complex128{2, 3}, NewComplex128Set(1)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Remove(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex128Set_Empty(t *testing.T) {
	cases := []struct {
		set  Complex128Set
		want Complex128Set
	}{
		{NewComplex128Set(), NewComplex128Set()},
		{NewComplex128Set(1, 2), NewComplex128Set()},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Empty()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex128Set_Has(t *testing.T) {
	cases := []struct {
		set  Complex128Set
		arg  complex128
		want bool
	}{
		{NewComplex128Set(), 1, false},
		{NewComplex128Set(2), 1, false},
		{NewComplex128Set(1), 1, true},
		{NewComplex128Set(1, 2), 1, true},
	}
	for _, c := range cases {
		got := c.set.Has(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex128Set_Size(t *testing.T) {
	cases := []struct {
		set  Complex128Set
		want int
	}{
		{NewComplex128Set(), 0},
		{NewComplex128Set(1, 2), 2},
	}
	for _, c := range cases {
		got := c.set.Size()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex128Set_IsEmpty(t *testing.T) {
	cases := []struct {
		set  Complex128Set
		want bool
	}{
		{NewComplex128Set(), true},
		{NewComplex128Set(1, 2), false},
	}
	for _, c := range cases {
		got := c.set.IsEmpty()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex128Set_Clone(t *testing.T) {
	cases := []struct {
		set  Complex128Set
		want Complex128Set
	}{
		{NewComplex128Set(), NewComplex128Set()},
		{NewComplex128Set(1, 2), NewComplex128Set(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex128Set_AsSlice(t *testing.T) {
	cases := []struct {
		set  Complex128Set
		want []complex128
	}{
		{NewComplex128Set(), []complex128{}},
		{NewComplex128Set(1, 2), []complex128{1, 2}},
	}
	for _, c := range cases {
		got := c.set.AsSlice()
		if len(got) != len(c.want) || !NewComplex128Set(got...).Equals(NewComplex128Set(c.want...)) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex128Set_String(t *testing.T) {
	cases := []struct {
		set  Complex128Set
		want []string
	}{
		{NewComplex128Set(), []string{"{}"}},
		{NewComplex128Set(1), []string{"{(1+0i)}"}},
		{NewComplex128Set(1, 2), []string{"{(1+0i) (2+0i)}", "{(2+0i) (1+0i)}"}},
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

func TestComplex128Set_Equals(t *testing.T) {
	cases := []struct {
		set  Complex128Set
		arg  Complex128Set
		want bool
	}{
		{NewComplex128Set(), NewComplex128Set(), true},
		{NewComplex128Set(1, 2), NewComplex128Set(2, 1), true},
		{NewComplex128Set(1, 2), NewComplex128Set(1), false},
		{NewComplex128Set(1), NewComplex128Set(1, 2), false},
		{NewComplex128Set(1), NewComplex128Set(2), false},
	}
	for _, c := range cases {
		got := c.set.Equals(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex128Set_Union(t *testing.T) {
	cases := []struct {
		set  Complex128Set
		arg  Complex128Set
		want Complex128Set
	}{
		{NewComplex128Set(), NewComplex128Set(), NewComplex128Set()},
		{NewComplex128Set(1), NewComplex128Set(1), NewComplex128Set(1)},
		{NewComplex128Set(1), NewComplex128Set(2), NewComplex128Set(1, 2)},
		{NewComplex128Set(1), NewComplex128Set(1, 2), NewComplex128Set(1, 2)},
		{NewComplex128Set(1, 2), NewComplex128Set(1), NewComplex128Set(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Union(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex128Set_Intersection(t *testing.T) {
	cases := []struct {
		set  Complex128Set
		arg  Complex128Set
		want Complex128Set
	}{
		{NewComplex128Set(), NewComplex128Set(), NewComplex128Set()},
		{NewComplex128Set(1), NewComplex128Set(1), NewComplex128Set(1)},
		{NewComplex128Set(1), NewComplex128Set(2), NewComplex128Set()},
		{NewComplex128Set(1), NewComplex128Set(1, 2), NewComplex128Set(1)},
		{NewComplex128Set(1, 2), NewComplex128Set(1), NewComplex128Set(1)},
	}
	for _, c := range cases {
		got := c.set.Intersection(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex128Set_Difference(t *testing.T) {
	cases := []struct {
		set  Complex128Set
		arg  Complex128Set
		want Complex128Set
	}{
		{NewComplex128Set(), NewComplex128Set(), NewComplex128Set()},
		{NewComplex128Set(1), NewComplex128Set(1), NewComplex128Set()},
		{NewComplex128Set(1), NewComplex128Set(2), NewComplex128Set(1)},
		{NewComplex128Set(1), NewComplex128Set(1, 2), NewComplex128Set()},
		{NewComplex128Set(1, 2), NewComplex128Set(1), NewComplex128Set(2)},
	}
	for _, c := range cases {
		got := c.set.Difference(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex128Set_IsSubsetOf(t *testing.T) {
	cases := []struct {
		set  Complex128Set
		arg  Complex128Set
		want bool
	}{
		{NewComplex128Set(), NewComplex128Set(), true},
		{NewComplex128Set(1), NewComplex128Set(1), true},
		{NewComplex128Set(1), NewComplex128Set(1, 2), true},
		{NewComplex128Set(1, 2), NewComplex128Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex128Set_IsProperSubsetOf(t *testing.T) {
	cases := []struct {
		set  Complex128Set
		arg  Complex128Set
		want bool
	}{
		{NewComplex128Set(), NewComplex128Set(), false},
		{NewComplex128Set(1), NewComplex128Set(1), false},
		{NewComplex128Set(1), NewComplex128Set(1, 2), true},
		{NewComplex128Set(1, 2), NewComplex128Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsProperSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex128Set_IsSupersetOf(t *testing.T) {
	cases := []struct {
		set  Complex128Set
		arg  Complex128Set
		want bool
	}{
		{NewComplex128Set(), NewComplex128Set(), true},
		{NewComplex128Set(1), NewComplex128Set(1), true},
		{NewComplex128Set(1), NewComplex128Set(1, 2), false},
		{NewComplex128Set(1, 2), NewComplex128Set(1), true},
	}
	for _, c := range cases {
		got := c.set.IsSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex128Set_IsProperSupersetOf(t *testing.T) {
	cases := []struct {
		set  Complex128Set
		arg  Complex128Set
		want bool
	}{
		{NewComplex128Set(), NewComplex128Set(), false},
		{NewComplex128Set(1), NewComplex128Set(1), false},
		{NewComplex128Set(1), NewComplex128Set(1, 2), false},
		{NewComplex128Set(1, 2), NewComplex128Set(1), true},
	}
	for _, c := range cases {
		got := c.set.IsProperSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex128Set_IsDisjointFrom(t *testing.T) {
	cases := []struct {
		set  Complex128Set
		arg  Complex128Set
		want bool
	}{
		{NewComplex128Set(), NewComplex128Set(), true},
		{NewComplex128Set(1), NewComplex128Set(1), false},
		{NewComplex128Set(1), NewComplex128Set(2, 3), true},
		{NewComplex128Set(1, 2), NewComplex128Set(3), true},
		{NewComplex128Set(1), NewComplex128Set(1, 2), false},
		{NewComplex128Set(1, 2), NewComplex128Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsDisjointFrom(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
