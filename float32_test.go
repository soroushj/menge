package menge_test

import (
	"math"
	"testing"

	"github.com/soroushj/menge"
)

func TestNewFloat32Set(t *testing.T) {
	cases := []struct {
		arg  []float32
		want menge.Float32Set
	}{
		{[]float32{}, menge.Float32Set{}},
		{[]float32{1, 1}, menge.Float32Set{1: struct{}{}}},
		{[]float32{1, 2}, menge.Float32Set{1: struct{}{}, 2: struct{}{}}},
		{[]float32{float32(math.NaN())}, menge.Float32Set{}},
	}
	for _, c := range cases {
		got := menge.NewFloat32Set(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat32Set_Add(t *testing.T) {
	cases := []struct {
		set  menge.Float32Set
		arg  []float32
		want menge.Float32Set
	}{
		{menge.NewFloat32Set(), []float32{}, menge.NewFloat32Set()},
		{menge.NewFloat32Set(), []float32{1, 1}, menge.NewFloat32Set(1)},
		{menge.NewFloat32Set(), []float32{1, 2}, menge.NewFloat32Set(1, 2)},
		{menge.NewFloat32Set(1), []float32{}, menge.NewFloat32Set(1)},
		{menge.NewFloat32Set(1), []float32{1, 1}, menge.NewFloat32Set(1)},
		{menge.NewFloat32Set(1), []float32{2, 3}, menge.NewFloat32Set(1, 2, 3)},
		{menge.NewFloat32Set(), []float32{float32(math.NaN())}, menge.NewFloat32Set()},
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
		set  menge.Float32Set
		arg  []float32
		want menge.Float32Set
	}{
		{menge.NewFloat32Set(), []float32{}, menge.NewFloat32Set()},
		{menge.NewFloat32Set(1), []float32{1, 1}, menge.NewFloat32Set()},
		{menge.NewFloat32Set(1, 2), []float32{1, 2}, menge.NewFloat32Set()},
		{menge.NewFloat32Set(1), []float32{}, menge.NewFloat32Set(1)},
		{menge.NewFloat32Set(1), []float32{1, 1}, menge.NewFloat32Set()},
		{menge.NewFloat32Set(1, 2), []float32{3}, menge.NewFloat32Set(1, 2)},
		{menge.NewFloat32Set(1, 2, 3), []float32{2, 3}, menge.NewFloat32Set(1)},
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
		set  menge.Float32Set
		want menge.Float32Set
	}{
		{menge.NewFloat32Set(), menge.NewFloat32Set()},
		{menge.NewFloat32Set(1, 2), menge.NewFloat32Set()},
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
		set  menge.Float32Set
		arg  float32
		want bool
	}{
		{menge.NewFloat32Set(), 1, false},
		{menge.NewFloat32Set(2), 1, false},
		{menge.NewFloat32Set(1), 1, true},
		{menge.NewFloat32Set(1, 2), 1, true},
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
		set  menge.Float32Set
		want int
	}{
		{menge.NewFloat32Set(), 0},
		{menge.NewFloat32Set(1, 2), 2},
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
		set  menge.Float32Set
		want bool
	}{
		{menge.NewFloat32Set(), true},
		{menge.NewFloat32Set(1, 2), false},
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
		set  menge.Float32Set
		want menge.Float32Set
	}{
		{menge.NewFloat32Set(), menge.NewFloat32Set()},
		{menge.NewFloat32Set(1, 2), menge.NewFloat32Set(1, 2)},
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
		set  menge.Float32Set
		want []float32
	}{
		{menge.NewFloat32Set(), []float32{}},
		{menge.NewFloat32Set(1, 2), []float32{1, 2}},
	}
	for _, c := range cases {
		got := c.set.AsSlice()
		if len(got) != len(c.want) || !menge.NewFloat32Set(got...).Equals(menge.NewFloat32Set(c.want...)) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat32Set_String(t *testing.T) {
	cases := []struct {
		set  menge.Float32Set
		want []string
	}{
		{menge.NewFloat32Set(), []string{"{}"}},
		{menge.NewFloat32Set(1), []string{"{1}"}},
		{menge.NewFloat32Set(1, 2), []string{"{1 2}", "{2 1}"}},
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
		set  menge.Float32Set
		arg  menge.Float32Set
		want bool
	}{
		{menge.NewFloat32Set(), menge.NewFloat32Set(), true},
		{menge.NewFloat32Set(1, 2), menge.NewFloat32Set(2, 1), true},
		{menge.NewFloat32Set(1, 2), menge.NewFloat32Set(1), false},
		{menge.NewFloat32Set(1), menge.NewFloat32Set(1, 2), false},
		{menge.NewFloat32Set(1), menge.NewFloat32Set(2), false},
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
		set  menge.Float32Set
		arg  menge.Float32Set
		want menge.Float32Set
	}{
		{menge.NewFloat32Set(), menge.NewFloat32Set(), menge.NewFloat32Set()},
		{menge.NewFloat32Set(1), menge.NewFloat32Set(1), menge.NewFloat32Set(1)},
		{menge.NewFloat32Set(1), menge.NewFloat32Set(2), menge.NewFloat32Set(1, 2)},
		{menge.NewFloat32Set(1), menge.NewFloat32Set(1, 2), menge.NewFloat32Set(1, 2)},
		{menge.NewFloat32Set(1, 2), menge.NewFloat32Set(1), menge.NewFloat32Set(1, 2)},
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
		set  menge.Float32Set
		arg  menge.Float32Set
		want menge.Float32Set
	}{
		{menge.NewFloat32Set(), menge.NewFloat32Set(), menge.NewFloat32Set()},
		{menge.NewFloat32Set(1), menge.NewFloat32Set(1), menge.NewFloat32Set(1)},
		{menge.NewFloat32Set(1), menge.NewFloat32Set(2), menge.NewFloat32Set()},
		{menge.NewFloat32Set(1), menge.NewFloat32Set(1, 2), menge.NewFloat32Set(1)},
		{menge.NewFloat32Set(1, 2), menge.NewFloat32Set(1), menge.NewFloat32Set(1)},
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
		set  menge.Float32Set
		arg  menge.Float32Set
		want menge.Float32Set
	}{
		{menge.NewFloat32Set(), menge.NewFloat32Set(), menge.NewFloat32Set()},
		{menge.NewFloat32Set(1), menge.NewFloat32Set(1), menge.NewFloat32Set()},
		{menge.NewFloat32Set(1), menge.NewFloat32Set(2), menge.NewFloat32Set(1)},
		{menge.NewFloat32Set(1), menge.NewFloat32Set(1, 2), menge.NewFloat32Set()},
		{menge.NewFloat32Set(1, 2), menge.NewFloat32Set(1), menge.NewFloat32Set(2)},
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
		set  menge.Float32Set
		arg  menge.Float32Set
		want bool
	}{
		{menge.NewFloat32Set(), menge.NewFloat32Set(), true},
		{menge.NewFloat32Set(1), menge.NewFloat32Set(1), true},
		{menge.NewFloat32Set(1), menge.NewFloat32Set(1, 2), true},
		{menge.NewFloat32Set(1, 2), menge.NewFloat32Set(1), false},
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
		set  menge.Float32Set
		arg  menge.Float32Set
		want bool
	}{
		{menge.NewFloat32Set(), menge.NewFloat32Set(), false},
		{menge.NewFloat32Set(1), menge.NewFloat32Set(1), false},
		{menge.NewFloat32Set(1), menge.NewFloat32Set(1, 2), true},
		{menge.NewFloat32Set(1, 2), menge.NewFloat32Set(1), false},
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
		set  menge.Float32Set
		arg  menge.Float32Set
		want bool
	}{
		{menge.NewFloat32Set(), menge.NewFloat32Set(), true},
		{menge.NewFloat32Set(1), menge.NewFloat32Set(1), true},
		{menge.NewFloat32Set(1), menge.NewFloat32Set(1, 2), false},
		{menge.NewFloat32Set(1, 2), menge.NewFloat32Set(1), true},
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
		set  menge.Float32Set
		arg  menge.Float32Set
		want bool
	}{
		{menge.NewFloat32Set(), menge.NewFloat32Set(), false},
		{menge.NewFloat32Set(1), menge.NewFloat32Set(1), false},
		{menge.NewFloat32Set(1), menge.NewFloat32Set(1, 2), false},
		{menge.NewFloat32Set(1, 2), menge.NewFloat32Set(1), true},
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
		set  menge.Float32Set
		arg  menge.Float32Set
		want bool
	}{
		{menge.NewFloat32Set(), menge.NewFloat32Set(), true},
		{menge.NewFloat32Set(1), menge.NewFloat32Set(1), false},
		{menge.NewFloat32Set(1), menge.NewFloat32Set(2, 3), true},
		{menge.NewFloat32Set(1, 2), menge.NewFloat32Set(3), true},
		{menge.NewFloat32Set(1), menge.NewFloat32Set(1, 2), false},
		{menge.NewFloat32Set(1, 2), menge.NewFloat32Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsDisjointFrom(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
