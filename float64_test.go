package menge_test

import (
	"math"
	"testing"

	"github.com/soroushj/menge"
)

func TestNewFloat64Set(t *testing.T) {
	cases := []struct {
		arg  []float64
		want menge.Float64Set
	}{
		{[]float64{}, menge.Float64Set{}},
		{[]float64{1, 1}, menge.Float64Set{1: struct{}{}}},
		{[]float64{1, 2}, menge.Float64Set{1: struct{}{}, 2: struct{}{}}},
		{[]float64{math.NaN()}, menge.Float64Set{}},
	}
	for _, c := range cases {
		got := menge.NewFloat64Set(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat64Set_Add(t *testing.T) {
	cases := []struct {
		set  menge.Float64Set
		arg  []float64
		want menge.Float64Set
	}{
		{menge.NewFloat64Set(), []float64{}, menge.NewFloat64Set()},
		{menge.NewFloat64Set(), []float64{1, 1}, menge.NewFloat64Set(1)},
		{menge.NewFloat64Set(), []float64{1, 2}, menge.NewFloat64Set(1, 2)},
		{menge.NewFloat64Set(1), []float64{}, menge.NewFloat64Set(1)},
		{menge.NewFloat64Set(1), []float64{1, 1}, menge.NewFloat64Set(1)},
		{menge.NewFloat64Set(1), []float64{2, 3}, menge.NewFloat64Set(1, 2, 3)},
		{menge.NewFloat64Set(), []float64{math.NaN()}, menge.NewFloat64Set()},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Add(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat64Set_Remove(t *testing.T) {
	cases := []struct {
		set  menge.Float64Set
		arg  []float64
		want menge.Float64Set
	}{
		{menge.NewFloat64Set(), []float64{}, menge.NewFloat64Set()},
		{menge.NewFloat64Set(1), []float64{1, 1}, menge.NewFloat64Set()},
		{menge.NewFloat64Set(1, 2), []float64{1, 2}, menge.NewFloat64Set()},
		{menge.NewFloat64Set(1), []float64{}, menge.NewFloat64Set(1)},
		{menge.NewFloat64Set(1), []float64{1, 1}, menge.NewFloat64Set()},
		{menge.NewFloat64Set(1, 2), []float64{3}, menge.NewFloat64Set(1, 2)},
		{menge.NewFloat64Set(1, 2, 3), []float64{2, 3}, menge.NewFloat64Set(1)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Remove(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat64Set_Empty(t *testing.T) {
	cases := []struct {
		set  menge.Float64Set
		want menge.Float64Set
	}{
		{menge.NewFloat64Set(), menge.NewFloat64Set()},
		{menge.NewFloat64Set(1, 2), menge.NewFloat64Set()},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Empty()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat64Set_Has(t *testing.T) {
	cases := []struct {
		set  menge.Float64Set
		arg  float64
		want bool
	}{
		{menge.NewFloat64Set(), 1, false},
		{menge.NewFloat64Set(2), 1, false},
		{menge.NewFloat64Set(1), 1, true},
		{menge.NewFloat64Set(1, 2), 1, true},
	}
	for _, c := range cases {
		got := c.set.Has(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat64Set_Size(t *testing.T) {
	cases := []struct {
		set  menge.Float64Set
		want int
	}{
		{menge.NewFloat64Set(), 0},
		{menge.NewFloat64Set(1, 2), 2},
	}
	for _, c := range cases {
		got := c.set.Size()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat64Set_IsEmpty(t *testing.T) {
	cases := []struct {
		set  menge.Float64Set
		want bool
	}{
		{menge.NewFloat64Set(), true},
		{menge.NewFloat64Set(1, 2), false},
	}
	for _, c := range cases {
		got := c.set.IsEmpty()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat64Set_Clone(t *testing.T) {
	cases := []struct {
		set  menge.Float64Set
		want menge.Float64Set
	}{
		{menge.NewFloat64Set(), menge.NewFloat64Set()},
		{menge.NewFloat64Set(1, 2), menge.NewFloat64Set(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Clone()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat64Set_AsSlice(t *testing.T) {
	cases := []struct {
		set  menge.Float64Set
		want []float64
	}{
		{menge.NewFloat64Set(), []float64{}},
		{menge.NewFloat64Set(1, 2), []float64{1, 2}},
	}
	for _, c := range cases {
		got := c.set.AsSlice()
		if len(got) != len(c.want) || !menge.NewFloat64Set(got...).Equals(menge.NewFloat64Set(c.want...)) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat64Set_String(t *testing.T) {
	cases := []struct {
		set  menge.Float64Set
		want []string
	}{
		{menge.NewFloat64Set(), []string{"{}"}},
		{menge.NewFloat64Set(1), []string{"{1}"}},
		{menge.NewFloat64Set(1, 2), []string{"{1 2}", "{2 1}"}},
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

func TestFloat64Set_Equals(t *testing.T) {
	cases := []struct {
		set  menge.Float64Set
		arg  menge.Float64Set
		want bool
	}{
		{menge.NewFloat64Set(), menge.NewFloat64Set(), true},
		{menge.NewFloat64Set(1, 2), menge.NewFloat64Set(2, 1), true},
		{menge.NewFloat64Set(1, 2), menge.NewFloat64Set(1), false},
		{menge.NewFloat64Set(1), menge.NewFloat64Set(1, 2), false},
		{menge.NewFloat64Set(1), menge.NewFloat64Set(2), false},
	}
	for _, c := range cases {
		got := c.set.Equals(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat64Set_Union(t *testing.T) {
	cases := []struct {
		set  menge.Float64Set
		arg  menge.Float64Set
		want menge.Float64Set
	}{
		{menge.NewFloat64Set(), menge.NewFloat64Set(), menge.NewFloat64Set()},
		{menge.NewFloat64Set(1), menge.NewFloat64Set(1), menge.NewFloat64Set(1)},
		{menge.NewFloat64Set(1), menge.NewFloat64Set(2), menge.NewFloat64Set(1, 2)},
		{menge.NewFloat64Set(1), menge.NewFloat64Set(1, 2), menge.NewFloat64Set(1, 2)},
		{menge.NewFloat64Set(1, 2), menge.NewFloat64Set(1), menge.NewFloat64Set(1, 2)},
	}
	for _, c := range cases {
		got := c.set.Union(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat64Set_Intersection(t *testing.T) {
	cases := []struct {
		set  menge.Float64Set
		arg  menge.Float64Set
		want menge.Float64Set
	}{
		{menge.NewFloat64Set(), menge.NewFloat64Set(), menge.NewFloat64Set()},
		{menge.NewFloat64Set(1), menge.NewFloat64Set(1), menge.NewFloat64Set(1)},
		{menge.NewFloat64Set(1), menge.NewFloat64Set(2), menge.NewFloat64Set()},
		{menge.NewFloat64Set(1), menge.NewFloat64Set(1, 2), menge.NewFloat64Set(1)},
		{menge.NewFloat64Set(1, 2), menge.NewFloat64Set(1), menge.NewFloat64Set(1)},
	}
	for _, c := range cases {
		got := c.set.Intersection(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat64Set_Difference(t *testing.T) {
	cases := []struct {
		set  menge.Float64Set
		arg  menge.Float64Set
		want menge.Float64Set
	}{
		{menge.NewFloat64Set(), menge.NewFloat64Set(), menge.NewFloat64Set()},
		{menge.NewFloat64Set(1), menge.NewFloat64Set(1), menge.NewFloat64Set()},
		{menge.NewFloat64Set(1), menge.NewFloat64Set(2), menge.NewFloat64Set(1)},
		{menge.NewFloat64Set(1), menge.NewFloat64Set(1, 2), menge.NewFloat64Set()},
		{menge.NewFloat64Set(1, 2), menge.NewFloat64Set(1), menge.NewFloat64Set(2)},
	}
	for _, c := range cases {
		got := c.set.Difference(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat64Set_IsSubsetOf(t *testing.T) {
	cases := []struct {
		set  menge.Float64Set
		arg  menge.Float64Set
		want bool
	}{
		{menge.NewFloat64Set(), menge.NewFloat64Set(), true},
		{menge.NewFloat64Set(1), menge.NewFloat64Set(1), true},
		{menge.NewFloat64Set(1), menge.NewFloat64Set(1, 2), true},
		{menge.NewFloat64Set(1, 2), menge.NewFloat64Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat64Set_IsProperSubsetOf(t *testing.T) {
	cases := []struct {
		set  menge.Float64Set
		arg  menge.Float64Set
		want bool
	}{
		{menge.NewFloat64Set(), menge.NewFloat64Set(), false},
		{menge.NewFloat64Set(1), menge.NewFloat64Set(1), false},
		{menge.NewFloat64Set(1), menge.NewFloat64Set(1, 2), true},
		{menge.NewFloat64Set(1, 2), menge.NewFloat64Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsProperSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat64Set_IsSupersetOf(t *testing.T) {
	cases := []struct {
		set  menge.Float64Set
		arg  menge.Float64Set
		want bool
	}{
		{menge.NewFloat64Set(), menge.NewFloat64Set(), true},
		{menge.NewFloat64Set(1), menge.NewFloat64Set(1), true},
		{menge.NewFloat64Set(1), menge.NewFloat64Set(1, 2), false},
		{menge.NewFloat64Set(1, 2), menge.NewFloat64Set(1), true},
	}
	for _, c := range cases {
		got := c.set.IsSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat64Set_IsProperSupersetOf(t *testing.T) {
	cases := []struct {
		set  menge.Float64Set
		arg  menge.Float64Set
		want bool
	}{
		{menge.NewFloat64Set(), menge.NewFloat64Set(), false},
		{menge.NewFloat64Set(1), menge.NewFloat64Set(1), false},
		{menge.NewFloat64Set(1), menge.NewFloat64Set(1, 2), false},
		{menge.NewFloat64Set(1, 2), menge.NewFloat64Set(1), true},
	}
	for _, c := range cases {
		got := c.set.IsProperSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestFloat64Set_IsDisjointFrom(t *testing.T) {
	cases := []struct {
		set  menge.Float64Set
		arg  menge.Float64Set
		want bool
	}{
		{menge.NewFloat64Set(), menge.NewFloat64Set(), true},
		{menge.NewFloat64Set(1), menge.NewFloat64Set(1), false},
		{menge.NewFloat64Set(1), menge.NewFloat64Set(2, 3), true},
		{menge.NewFloat64Set(1, 2), menge.NewFloat64Set(3), true},
		{menge.NewFloat64Set(1), menge.NewFloat64Set(1, 2), false},
		{menge.NewFloat64Set(1, 2), menge.NewFloat64Set(1), false},
	}
	for _, c := range cases {
		got := c.set.IsDisjointFrom(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
