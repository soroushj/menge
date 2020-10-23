package menge

import (
	"testing"
)

func TestNewComplex64Set(t *testing.T) {
	cases := []struct {
		arg  []complex64
		want Complex64Set
	}{
		{[]complex64{}, Complex64Set{}},
		{[]complex64{1, 1}, Complex64Set{1: struct{}{}}},
		{[]complex64{1, 2}, Complex64Set{1: struct{}{}, 2: struct{}{}}},
	}
	for _, c := range cases {
		got := NewComplex64Set(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex64Set_Add(t *testing.T) {
	cases := []struct {
		set  Complex64Set
		arg  []complex64
		want Complex64Set
	}{
		{NewComplex64Set(), []complex64{}, NewComplex64Set()},
		{NewComplex64Set(), []complex64{1, 1}, NewComplex64Set(1)},
		{NewComplex64Set(), []complex64{1, 2}, NewComplex64Set(1, 2)},
		{NewComplex64Set(1), []complex64{}, NewComplex64Set(1)},
		{NewComplex64Set(1), []complex64{1, 1}, NewComplex64Set(1)},
		{NewComplex64Set(1), []complex64{2, 3}, NewComplex64Set(1, 2, 3)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Add(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex64Set_Remove(t *testing.T) {
	cases := []struct {
		set  Complex64Set
		arg  []complex64
		want Complex64Set
	}{
		{NewComplex64Set(), []complex64{}, NewComplex64Set()},
		{NewComplex64Set(1), []complex64{1, 1}, NewComplex64Set()},
		{NewComplex64Set(1, 2), []complex64{1, 2}, NewComplex64Set()},
		{NewComplex64Set(1), []complex64{}, NewComplex64Set(1)},
		{NewComplex64Set(1), []complex64{1, 1}, NewComplex64Set()},
		{NewComplex64Set(1, 2), []complex64{3}, NewComplex64Set(1, 2)},
		{NewComplex64Set(1, 2, 3), []complex64{2, 3}, NewComplex64Set(1)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Remove(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex64Set_Empty(t *testing.T) {
	cases := []struct {
		set  Complex64Set
		want Complex64Set
	}{
		{NewComplex64Set(), NewComplex64Set()},
		{NewComplex64Set(1, 2), NewComplex64Set()},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Empty()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex64Set_Has(t *testing.T) {
	cases := []struct {
		set  Complex64Set
		arg  complex64
		want bool
	}{
		{NewComplex64Set(), 1, false},
		{NewComplex64Set(2), 1, false},
		{NewComplex64Set(1), 1, true},
		{NewComplex64Set(1, 2), 1, true},
	}
	for _, c := range cases {
		got := c.set.Has(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex64Set_Size(t *testing.T) {
	cases := []struct {
		set  Complex64Set
		want int
	}{
		{NewComplex64Set(), 0},
		{NewComplex64Set(1, 2), 2},
	}
	for _, c := range cases {
		got := c.set.Size()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex64Set_IsEmpty(t *testing.T) {
	cases := []struct {
		set  Complex64Set
		want bool
	}{
		{NewComplex64Set(), true},
		{NewComplex64Set(1, 2), false},
	}
	for _, c := range cases {
		got := c.set.IsEmpty()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex64Set_Clone(t *testing.T) {
	cases := []struct {
		set  Complex64Set
		want Complex64Set
	}{
		{NewComplex64Set(), NewComplex64Set()},
		{NewComplex64Set(1, 2), NewComplex64Set(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex64Set_AsSlice(t *testing.T) {
	cases := []struct {
		set  Complex64Set
		want []complex64
	}{
		{NewComplex64Set(), []complex64{}},
		{NewComplex64Set(1, 2), []complex64{1, 2}},
	}
	for _, c := range cases {
		got := c.set.AsSlice()
		if len(got) != len(c.want) || !NewComplex64Set(got...).Equals(NewComplex64Set(c.want...)) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex64Set_String(t *testing.T) {
	cases := []struct {
		set  Complex64Set
		want []string
	}{
		{NewComplex64Set(), []string{"{}"}},
		{NewComplex64Set(1), []string{"{(1+0i)}"}},
		{NewComplex64Set(1, 2), []string{"{(1+0i) (2+0i)}", "{(2+0i) (1+0i)}"}},
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

func TestComplex64Set_Equals(t *testing.T) {
	cases := []struct {
		set  Complex64Set
		arg  Complex64Set
		want bool
	}{
		{NewComplex64Set(), NewComplex64Set(), true},
		{NewComplex64Set(1, 2), NewComplex64Set(2, 1), true},
		{NewComplex64Set(1, 2), NewComplex64Set(1), false},
		{NewComplex64Set(1), NewComplex64Set(1, 2), false},
		{NewComplex64Set(1), NewComplex64Set(2), false},
	}
	for _, c := range cases {
		got := c.set.Equals(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex64Set_Union(t *testing.T) {
	cases := []struct {
		set  Complex64Set
		arg  Complex64Set
		want Complex64Set
	}{
		{NewComplex64Set(), NewComplex64Set(), NewComplex64Set()},
		{NewComplex64Set(1), NewComplex64Set(1), NewComplex64Set(1)},
		{NewComplex64Set(1), NewComplex64Set(2), NewComplex64Set(1, 2)},
		{NewComplex64Set(1), NewComplex64Set(1, 2), NewComplex64Set(1, 2)},
		{NewComplex64Set(1, 2), NewComplex64Set(1), NewComplex64Set(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Union(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex64Set_Intersection(t *testing.T) {
	cases := []struct {
		set  Complex64Set
		arg  Complex64Set
		want Complex64Set
	}{
		{NewComplex64Set(), NewComplex64Set(), NewComplex64Set()},
		{NewComplex64Set(1), NewComplex64Set(1), NewComplex64Set(1)},
		{NewComplex64Set(1), NewComplex64Set(2), NewComplex64Set()},
		{NewComplex64Set(1), NewComplex64Set(1, 2), NewComplex64Set(1)},
		{NewComplex64Set(1, 2), NewComplex64Set(1), NewComplex64Set(1)},
	}
	for _, c := range cases {
		got := c.set.Intersection(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex64Set_Difference(t *testing.T) {
	cases := []struct {
		set  Complex64Set
		arg  Complex64Set
		want Complex64Set
	}{
		{NewComplex64Set(), NewComplex64Set(), NewComplex64Set()},
		{NewComplex64Set(1), NewComplex64Set(1), NewComplex64Set()},
		{NewComplex64Set(1), NewComplex64Set(2), NewComplex64Set(1)},
		{NewComplex64Set(1), NewComplex64Set(1, 2), NewComplex64Set()},
		{NewComplex64Set(1, 2), NewComplex64Set(1), NewComplex64Set(2)},
	}
	for _, c := range cases {
		got := c.set.Difference(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex64Set_IsSubsetOf(t *testing.T) {
	cases := []struct {
		set  Complex64Set
		arg  Complex64Set
		want bool
	}{
		{NewComplex64Set(), NewComplex64Set(), true},
		{NewComplex64Set(1), NewComplex64Set(1), true},
		{NewComplex64Set(1), NewComplex64Set(1, 2), true},
		{NewComplex64Set(1, 2), NewComplex64Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex64Set_IsProperSubsetOf(t *testing.T) {
	cases := []struct {
		set  Complex64Set
		arg  Complex64Set
		want bool
	}{
		{NewComplex64Set(), NewComplex64Set(), false},
		{NewComplex64Set(1), NewComplex64Set(1), false},
		{NewComplex64Set(1), NewComplex64Set(1, 2), true},
		{NewComplex64Set(1, 2), NewComplex64Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsProperSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex64Set_IsSupersetOf(t *testing.T) {
	cases := []struct {
		set  Complex64Set
		arg  Complex64Set
		want bool
	}{
		{NewComplex64Set(), NewComplex64Set(), true},
		{NewComplex64Set(1), NewComplex64Set(1), true},
		{NewComplex64Set(1), NewComplex64Set(1, 2), false},
		{NewComplex64Set(1, 2), NewComplex64Set(1), true},
	}
	for _, c := range cases {
		got := c.set.IsSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex64Set_IsProperSupersetOf(t *testing.T) {
	cases := []struct {
		set  Complex64Set
		arg  Complex64Set
		want bool
	}{
		{NewComplex64Set(), NewComplex64Set(), false},
		{NewComplex64Set(1), NewComplex64Set(1), false},
		{NewComplex64Set(1), NewComplex64Set(1, 2), false},
		{NewComplex64Set(1, 2), NewComplex64Set(1), true},
	}
	for _, c := range cases {
		got := c.set.IsProperSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex64Set_IsDisjointFrom(t *testing.T) {
	cases := []struct {
		set  Complex64Set
		arg  Complex64Set
		want bool
	}{
		{NewComplex64Set(), NewComplex64Set(), true},
		{NewComplex64Set(1), NewComplex64Set(1), false},
		{NewComplex64Set(1), NewComplex64Set(2, 3), true},
		{NewComplex64Set(1, 2), NewComplex64Set(3), true},
		{NewComplex64Set(1), NewComplex64Set(1, 2), false},
		{NewComplex64Set(1, 2), NewComplex64Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsDisjointFrom(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
