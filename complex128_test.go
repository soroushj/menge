package menge_test

import (
	"testing"

	"github.com/soroushj/menge"
)

func TestNewComplex128Set(t *testing.T) {
	cases := []struct {
		arg  []complex128
		want menge.Complex128Set
	}{
		{[]complex128{}, menge.Complex128Set{}},
		{[]complex128{1, 1}, menge.Complex128Set{1: struct{}{}}},
		{[]complex128{1, 2}, menge.Complex128Set{1: struct{}{}, 2: struct{}{}}},
	}
	for _, c := range cases {
		got := menge.NewComplex128Set(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex128Set_Add(t *testing.T) {
	cases := []struct {
		set  menge.Complex128Set
		arg  []complex128
		want menge.Complex128Set
	}{
		{menge.NewComplex128Set(), []complex128{}, menge.NewComplex128Set()},
		{menge.NewComplex128Set(), []complex128{1, 1}, menge.NewComplex128Set(1)},
		{menge.NewComplex128Set(), []complex128{1, 2}, menge.NewComplex128Set(1, 2)},
		{menge.NewComplex128Set(1), []complex128{}, menge.NewComplex128Set(1)},
		{menge.NewComplex128Set(1), []complex128{1, 1}, menge.NewComplex128Set(1)},
		{menge.NewComplex128Set(1), []complex128{2, 3}, menge.NewComplex128Set(1, 2, 3)},
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
		set  menge.Complex128Set
		arg  []complex128
		want menge.Complex128Set
	}{
		{menge.NewComplex128Set(), []complex128{}, menge.NewComplex128Set()},
		{menge.NewComplex128Set(1), []complex128{1, 1}, menge.NewComplex128Set()},
		{menge.NewComplex128Set(1, 2), []complex128{1, 2}, menge.NewComplex128Set()},
		{menge.NewComplex128Set(1), []complex128{}, menge.NewComplex128Set(1)},
		{menge.NewComplex128Set(1), []complex128{1, 1}, menge.NewComplex128Set()},
		{menge.NewComplex128Set(1, 2), []complex128{3}, menge.NewComplex128Set(1, 2)},
		{menge.NewComplex128Set(1, 2, 3), []complex128{2, 3}, menge.NewComplex128Set(1)},
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
		set  menge.Complex128Set
		want menge.Complex128Set
	}{
		{menge.NewComplex128Set(), menge.NewComplex128Set()},
		{menge.NewComplex128Set(1, 2), menge.NewComplex128Set()},
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
		set  menge.Complex128Set
		arg  complex128
		want bool
	}{
		{menge.NewComplex128Set(), 1, false},
		{menge.NewComplex128Set(2), 1, false},
		{menge.NewComplex128Set(1), 1, true},
		{menge.NewComplex128Set(1, 2), 1, true},
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
		set  menge.Complex128Set
		want int
	}{
		{menge.NewComplex128Set(), 0},
		{menge.NewComplex128Set(1, 2), 2},
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
		set  menge.Complex128Set
		want bool
	}{
		{menge.NewComplex128Set(), true},
		{menge.NewComplex128Set(1, 2), false},
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
		set  menge.Complex128Set
		want menge.Complex128Set
	}{
		{menge.NewComplex128Set(), menge.NewComplex128Set()},
		{menge.NewComplex128Set(1, 2), menge.NewComplex128Set(1, 2)},
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
		set  menge.Complex128Set
		want []complex128
	}{
		{menge.NewComplex128Set(), []complex128{}},
		{menge.NewComplex128Set(1, 2), []complex128{1, 2}},
	}
	for _, c := range cases {
		got := c.set.AsSlice()
		if len(got) != len(c.want) || !menge.NewComplex128Set(got...).Equals(menge.NewComplex128Set(c.want...)) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestComplex128Set_String(t *testing.T) {
	cases := []struct {
		set  menge.Complex128Set
		want []string
	}{
		{menge.NewComplex128Set(), []string{"{}"}},
		{menge.NewComplex128Set(1), []string{"{(1+0i)}"}},
		{menge.NewComplex128Set(1, 2), []string{"{(1+0i) (2+0i)}", "{(2+0i) (1+0i)}"}},
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
		set  menge.Complex128Set
		arg  menge.Complex128Set
		want bool
	}{
		{menge.NewComplex128Set(), menge.NewComplex128Set(), true},
		{menge.NewComplex128Set(1, 2), menge.NewComplex128Set(2, 1), true},
		{menge.NewComplex128Set(1, 2), menge.NewComplex128Set(1), false},
		{menge.NewComplex128Set(1), menge.NewComplex128Set(1, 2), false},
		{menge.NewComplex128Set(1), menge.NewComplex128Set(2), false},
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
		set  menge.Complex128Set
		arg  menge.Complex128Set
		want menge.Complex128Set
	}{
		{menge.NewComplex128Set(), menge.NewComplex128Set(), menge.NewComplex128Set()},
		{menge.NewComplex128Set(1), menge.NewComplex128Set(1), menge.NewComplex128Set(1)},
		{menge.NewComplex128Set(1), menge.NewComplex128Set(2), menge.NewComplex128Set(1, 2)},
		{menge.NewComplex128Set(1), menge.NewComplex128Set(1, 2), menge.NewComplex128Set(1, 2)},
		{menge.NewComplex128Set(1, 2), menge.NewComplex128Set(1), menge.NewComplex128Set(1, 2)},
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
		set  menge.Complex128Set
		arg  menge.Complex128Set
		want menge.Complex128Set
	}{
		{menge.NewComplex128Set(), menge.NewComplex128Set(), menge.NewComplex128Set()},
		{menge.NewComplex128Set(1), menge.NewComplex128Set(1), menge.NewComplex128Set(1)},
		{menge.NewComplex128Set(1), menge.NewComplex128Set(2), menge.NewComplex128Set()},
		{menge.NewComplex128Set(1), menge.NewComplex128Set(1, 2), menge.NewComplex128Set(1)},
		{menge.NewComplex128Set(1, 2), menge.NewComplex128Set(1), menge.NewComplex128Set(1)},
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
		set  menge.Complex128Set
		arg  menge.Complex128Set
		want menge.Complex128Set
	}{
		{menge.NewComplex128Set(), menge.NewComplex128Set(), menge.NewComplex128Set()},
		{menge.NewComplex128Set(1), menge.NewComplex128Set(1), menge.NewComplex128Set()},
		{menge.NewComplex128Set(1), menge.NewComplex128Set(2), menge.NewComplex128Set(1)},
		{menge.NewComplex128Set(1), menge.NewComplex128Set(1, 2), menge.NewComplex128Set()},
		{menge.NewComplex128Set(1, 2), menge.NewComplex128Set(1), menge.NewComplex128Set(2)},
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
		set  menge.Complex128Set
		arg  menge.Complex128Set
		want bool
	}{
		{menge.NewComplex128Set(), menge.NewComplex128Set(), true},
		{menge.NewComplex128Set(1), menge.NewComplex128Set(1), true},
		{menge.NewComplex128Set(1), menge.NewComplex128Set(1, 2), true},
		{menge.NewComplex128Set(1, 2), menge.NewComplex128Set(1), false},
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
		set  menge.Complex128Set
		arg  menge.Complex128Set
		want bool
	}{
		{menge.NewComplex128Set(), menge.NewComplex128Set(), false},
		{menge.NewComplex128Set(1), menge.NewComplex128Set(1), false},
		{menge.NewComplex128Set(1), menge.NewComplex128Set(1, 2), true},
		{menge.NewComplex128Set(1, 2), menge.NewComplex128Set(1), false},
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
		set  menge.Complex128Set
		arg  menge.Complex128Set
		want bool
	}{
		{menge.NewComplex128Set(), menge.NewComplex128Set(), true},
		{menge.NewComplex128Set(1), menge.NewComplex128Set(1), true},
		{menge.NewComplex128Set(1), menge.NewComplex128Set(1, 2), false},
		{menge.NewComplex128Set(1, 2), menge.NewComplex128Set(1), true},
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
		set  menge.Complex128Set
		arg  menge.Complex128Set
		want bool
	}{
		{menge.NewComplex128Set(), menge.NewComplex128Set(), false},
		{menge.NewComplex128Set(1), menge.NewComplex128Set(1), false},
		{menge.NewComplex128Set(1), menge.NewComplex128Set(1, 2), false},
		{menge.NewComplex128Set(1, 2), menge.NewComplex128Set(1), true},
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
		set  menge.Complex128Set
		arg  menge.Complex128Set
		want bool
	}{
		{menge.NewComplex128Set(), menge.NewComplex128Set(), true},
		{menge.NewComplex128Set(1), menge.NewComplex128Set(1), false},
		{menge.NewComplex128Set(1), menge.NewComplex128Set(2, 3), true},
		{menge.NewComplex128Set(1, 2), menge.NewComplex128Set(3), true},
		{menge.NewComplex128Set(1), menge.NewComplex128Set(1, 2), false},
		{menge.NewComplex128Set(1, 2), menge.NewComplex128Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsDisjointFrom(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
