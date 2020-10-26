package menge

import (
	"math"
	"testing"
)

func TestNewFloat32Set(t *testing.T) {
	cases := []struct {
		arg  []float32
		want Float32Set
	}{
		{[]float32{}, Float32Set{}},
		{[]float32{1, 1}, Float32Set{1: struct{}{}}},
		{[]float32{1, 2}, Float32Set{1: struct{}{}, 2: struct{}{}}},
		{[]float32{float32(math.NaN())}, Float32Set{}},
	}
	for _, c := range cases {
		got := NewFloat32Set(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat32Set_Add(t *testing.T) {
	cases := []struct {
		set  Float32Set
		arg  []float32
		want Float32Set
	}{
		{NewFloat32Set(), []float32{}, NewFloat32Set()},
		{NewFloat32Set(), []float32{1, 1}, NewFloat32Set(1)},
		{NewFloat32Set(), []float32{1, 2}, NewFloat32Set(1, 2)},
		{NewFloat32Set(1), []float32{}, NewFloat32Set(1)},
		{NewFloat32Set(1), []float32{1, 1}, NewFloat32Set(1)},
		{NewFloat32Set(1), []float32{2, 3}, NewFloat32Set(1, 2, 3)},
		{NewFloat32Set(), []float32{float32(math.NaN())}, NewFloat32Set()},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Add(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat32Set_Remove(t *testing.T) {
	cases := []struct {
		set  Float32Set
		arg  []float32
		want Float32Set
	}{
		{NewFloat32Set(), []float32{}, NewFloat32Set()},
		{NewFloat32Set(1), []float32{1, 1}, NewFloat32Set()},
		{NewFloat32Set(1, 2), []float32{1, 2}, NewFloat32Set()},
		{NewFloat32Set(1), []float32{}, NewFloat32Set(1)},
		{NewFloat32Set(1), []float32{1, 1}, NewFloat32Set()},
		{NewFloat32Set(1, 2), []float32{3}, NewFloat32Set(1, 2)},
		{NewFloat32Set(1, 2, 3), []float32{2, 3}, NewFloat32Set(1)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Remove(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat32Set_Empty(t *testing.T) {
	cases := []struct {
		set  Float32Set
		want Float32Set
	}{
		{NewFloat32Set(), NewFloat32Set()},
		{NewFloat32Set(1, 2), NewFloat32Set()},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Empty()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat32Set_Has(t *testing.T) {
	cases := []struct {
		set  Float32Set
		arg  float32
		want bool
	}{
		{NewFloat32Set(), 1, false},
		{NewFloat32Set(2), 1, false},
		{NewFloat32Set(1), 1, true},
		{NewFloat32Set(1, 2), 1, true},
	}
	for _, c := range cases {
		got := c.set.Has(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat32Set_Size(t *testing.T) {
	cases := []struct {
		set  Float32Set
		want int
	}{
		{NewFloat32Set(), 0},
		{NewFloat32Set(1, 2), 2},
	}
	for _, c := range cases {
		got := c.set.Size()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat32Set_IsEmpty(t *testing.T) {
	cases := []struct {
		set  Float32Set
		want bool
	}{
		{NewFloat32Set(), true},
		{NewFloat32Set(1, 2), false},
	}
	for _, c := range cases {
		got := c.set.IsEmpty()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat32Set_Clone(t *testing.T) {
	cases := []struct {
		set  Float32Set
		want Float32Set
	}{
		{NewFloat32Set(), NewFloat32Set()},
		{NewFloat32Set(1, 2), NewFloat32Set(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat32Set_AsSlice(t *testing.T) {
	cases := []struct {
		set  Float32Set
		want []float32
	}{
		{NewFloat32Set(), []float32{}},
		{NewFloat32Set(1, 2), []float32{1, 2}},
	}
	for _, c := range cases {
		got := c.set.AsSlice()
		if len(got) != len(c.want) || !NewFloat32Set(got...).Equals(NewFloat32Set(c.want...)) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat32Set_String(t *testing.T) {
	cases := []struct {
		set  Float32Set
		want []string
	}{
		{NewFloat32Set(), []string{"{}"}},
		{NewFloat32Set(1), []string{"{1}"}},
		{NewFloat32Set(1, 2), []string{"{1 2}", "{2 1}"}},
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

func TestFloat32Set_Equals(t *testing.T) {
	cases := []struct {
		set  Float32Set
		arg  Float32Set
		want bool
	}{
		{NewFloat32Set(), NewFloat32Set(), true},
		{NewFloat32Set(1, 2), NewFloat32Set(2, 1), true},
		{NewFloat32Set(1, 2), NewFloat32Set(1), false},
		{NewFloat32Set(1), NewFloat32Set(1, 2), false},
		{NewFloat32Set(1), NewFloat32Set(2), false},
	}
	for _, c := range cases {
		got := c.set.Equals(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat32Set_Union(t *testing.T) {
	cases := []struct {
		set  Float32Set
		arg  Float32Set
		want Float32Set
	}{
		{NewFloat32Set(), NewFloat32Set(), NewFloat32Set()},
		{NewFloat32Set(1), NewFloat32Set(1), NewFloat32Set(1)},
		{NewFloat32Set(1), NewFloat32Set(2), NewFloat32Set(1, 2)},
		{NewFloat32Set(1), NewFloat32Set(1, 2), NewFloat32Set(1, 2)},
		{NewFloat32Set(1, 2), NewFloat32Set(1), NewFloat32Set(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Union(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat32Set_Intersection(t *testing.T) {
	cases := []struct {
		set  Float32Set
		arg  Float32Set
		want Float32Set
	}{
		{NewFloat32Set(), NewFloat32Set(), NewFloat32Set()},
		{NewFloat32Set(1), NewFloat32Set(1), NewFloat32Set(1)},
		{NewFloat32Set(1), NewFloat32Set(2), NewFloat32Set()},
		{NewFloat32Set(1), NewFloat32Set(1, 2), NewFloat32Set(1)},
		{NewFloat32Set(1, 2), NewFloat32Set(1), NewFloat32Set(1)},
	}
	for _, c := range cases {
		got := c.set.Intersection(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat32Set_Difference(t *testing.T) {
	cases := []struct {
		set  Float32Set
		arg  Float32Set
		want Float32Set
	}{
		{NewFloat32Set(), NewFloat32Set(), NewFloat32Set()},
		{NewFloat32Set(1), NewFloat32Set(1), NewFloat32Set()},
		{NewFloat32Set(1), NewFloat32Set(2), NewFloat32Set(1)},
		{NewFloat32Set(1), NewFloat32Set(1, 2), NewFloat32Set()},
		{NewFloat32Set(1, 2), NewFloat32Set(1), NewFloat32Set(2)},
	}
	for _, c := range cases {
		got := c.set.Difference(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat32Set_IsSubsetOf(t *testing.T) {
	cases := []struct {
		set  Float32Set
		arg  Float32Set
		want bool
	}{
		{NewFloat32Set(), NewFloat32Set(), true},
		{NewFloat32Set(1), NewFloat32Set(1), true},
		{NewFloat32Set(1), NewFloat32Set(1, 2), true},
		{NewFloat32Set(1, 2), NewFloat32Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat32Set_IsProperSubsetOf(t *testing.T) {
	cases := []struct {
		set  Float32Set
		arg  Float32Set
		want bool
	}{
		{NewFloat32Set(), NewFloat32Set(), false},
		{NewFloat32Set(1), NewFloat32Set(1), false},
		{NewFloat32Set(1), NewFloat32Set(1, 2), true},
		{NewFloat32Set(1, 2), NewFloat32Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsProperSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat32Set_IsSupersetOf(t *testing.T) {
	cases := []struct {
		set  Float32Set
		arg  Float32Set
		want bool
	}{
		{NewFloat32Set(), NewFloat32Set(), true},
		{NewFloat32Set(1), NewFloat32Set(1), true},
		{NewFloat32Set(1), NewFloat32Set(1, 2), false},
		{NewFloat32Set(1, 2), NewFloat32Set(1), true},
	}
	for _, c := range cases {
		got := c.set.IsSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat32Set_IsProperSupersetOf(t *testing.T) {
	cases := []struct {
		set  Float32Set
		arg  Float32Set
		want bool
	}{
		{NewFloat32Set(), NewFloat32Set(), false},
		{NewFloat32Set(1), NewFloat32Set(1), false},
		{NewFloat32Set(1), NewFloat32Set(1, 2), false},
		{NewFloat32Set(1, 2), NewFloat32Set(1), true},
	}
	for _, c := range cases {
		got := c.set.IsProperSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat32Set_IsDisjointFrom(t *testing.T) {
	cases := []struct {
		set  Float32Set
		arg  Float32Set
		want bool
	}{
		{NewFloat32Set(), NewFloat32Set(), true},
		{NewFloat32Set(1), NewFloat32Set(1), false},
		{NewFloat32Set(1), NewFloat32Set(2, 3), true},
		{NewFloat32Set(1, 2), NewFloat32Set(3), true},
		{NewFloat32Set(1), NewFloat32Set(1, 2), false},
		{NewFloat32Set(1, 2), NewFloat32Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsDisjointFrom(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
