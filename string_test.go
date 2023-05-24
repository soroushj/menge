package menge_test

import (
	"testing"

	"github.com/soroushj/menge"
)

func TestNewStringSet(t *testing.T) {
	cases := []struct {
		arg  []string
		want menge.StringSet
	}{
		{[]string{}, menge.StringSet{}},
		{[]string{"1", "1"}, menge.StringSet{"1": struct{}{}}},
		{[]string{"1", "2"}, menge.StringSet{"1": struct{}{}, "2": struct{}{}}},
	}
	for _, c := range cases {
		got := menge.NewStringSet(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestStringSet_Add(t *testing.T) {
	cases := []struct {
		set  menge.StringSet
		arg  []string
		want menge.StringSet
	}{
		{menge.NewStringSet(), []string{}, menge.NewStringSet()},
		{menge.NewStringSet(), []string{"1", "1"}, menge.NewStringSet("1")},
		{menge.NewStringSet(), []string{"1", "2"}, menge.NewStringSet("1", "2")},
		{menge.NewStringSet("1"), []string{}, menge.NewStringSet("1")},
		{menge.NewStringSet("1"), []string{"1", "1"}, menge.NewStringSet("1")},
		{menge.NewStringSet("1"), []string{"2", "3"}, menge.NewStringSet("1", "2", "3")},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Add(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestStringSet_Remove(t *testing.T) {
	cases := []struct {
		set  menge.StringSet
		arg  []string
		want menge.StringSet
	}{
		{menge.NewStringSet(), []string{}, menge.NewStringSet()},
		{menge.NewStringSet("1"), []string{"1", "1"}, menge.NewStringSet()},
		{menge.NewStringSet("1", "2"), []string{"1", "2"}, menge.NewStringSet()},
		{menge.NewStringSet("1"), []string{}, menge.NewStringSet("1")},
		{menge.NewStringSet("1"), []string{"1", "1"}, menge.NewStringSet()},
		{menge.NewStringSet("1", "2"), []string{"3"}, menge.NewStringSet("1", "2")},
		{menge.NewStringSet("1", "2", "3"), []string{"2", "3"}, menge.NewStringSet("1")},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Remove(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestStringSet_Empty(t *testing.T) {
	cases := []struct {
		set  menge.StringSet
		want menge.StringSet
	}{
		{menge.NewStringSet(), menge.NewStringSet()},
		{menge.NewStringSet("1", "2"), menge.NewStringSet()},
	}
	for _, c := range cases {
		got := c.set.Clone()
		got.Empty()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestStringSet_Has(t *testing.T) {
	cases := []struct {
		set  menge.StringSet
		arg  string
		want bool
	}{
		{menge.NewStringSet(), "1", false},
		{menge.NewStringSet("2"), "1", false},
		{menge.NewStringSet("1"), "1", true},
		{menge.NewStringSet("1", "2"), "1", true},
	}
	for _, c := range cases {
		got := c.set.Has(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestStringSet_Size(t *testing.T) {
	cases := []struct {
		set  menge.StringSet
		want int
	}{
		{menge.NewStringSet(), 0},
		{menge.NewStringSet("1", "2"), 2},
	}
	for _, c := range cases {
		got := c.set.Size()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestStringSet_IsEmpty(t *testing.T) {
	cases := []struct {
		set  menge.StringSet
		want bool
	}{
		{menge.NewStringSet(), true},
		{menge.NewStringSet("1", "2"), false},
	}
	for _, c := range cases {
		got := c.set.IsEmpty()
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestStringSet_Clone(t *testing.T) {
	cases := []struct {
		set  menge.StringSet
		want menge.StringSet
	}{
		{menge.NewStringSet(), menge.NewStringSet()},
		{menge.NewStringSet("1", "2"), menge.NewStringSet("1", "2")},
	}
	for _, c := range cases {
		got := c.set.Clone()
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestStringSet_AsSlice(t *testing.T) {
	cases := []struct {
		set  menge.StringSet
		want []string
	}{
		{menge.NewStringSet(), []string{}},
		{menge.NewStringSet("1", "2"), []string{"1", "2"}},
	}
	for _, c := range cases {
		got := c.set.AsSlice()
		if len(got) != len(c.want) || !menge.NewStringSet(got...).Equals(menge.NewStringSet(c.want...)) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestStringSet_String(t *testing.T) {
	cases := []struct {
		set  menge.StringSet
		want []string
	}{
		{menge.NewStringSet(), []string{"{}"}},
		{menge.NewStringSet("1"), []string{"{1}"}},
		{menge.NewStringSet("1", "2"), []string{"{1 2}", "{2 1}"}},
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

func TestStringSet_Equals(t *testing.T) {
	cases := []struct {
		set  menge.StringSet
		arg  menge.StringSet
		want bool
	}{
		{menge.NewStringSet(), menge.NewStringSet(), true},
		{menge.NewStringSet("1", "2"), menge.NewStringSet("2", "1"), true},
		{menge.NewStringSet("1", "2"), menge.NewStringSet("1"), false},
		{menge.NewStringSet("1"), menge.NewStringSet("1", "2"), false},
		{menge.NewStringSet("1"), menge.NewStringSet("2"), false},
	}
	for _, c := range cases {
		got := c.set.Equals(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestStringSet_Union(t *testing.T) {
	cases := []struct {
		set  menge.StringSet
		arg  menge.StringSet
		want menge.StringSet
	}{
		{menge.NewStringSet(), menge.NewStringSet(), menge.NewStringSet()},
		{menge.NewStringSet("1"), menge.NewStringSet("1"), menge.NewStringSet("1")},
		{menge.NewStringSet("1"), menge.NewStringSet("2"), menge.NewStringSet("1", "2")},
		{menge.NewStringSet("1"), menge.NewStringSet("1", "2"), menge.NewStringSet("1", "2")},
		{menge.NewStringSet("1", "2"), menge.NewStringSet("1"), menge.NewStringSet("1", "2")},
	}
	for _, c := range cases {
		got := c.set.Union(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestStringSet_Intersection(t *testing.T) {
	cases := []struct {
		set  menge.StringSet
		arg  menge.StringSet
		want menge.StringSet
	}{
		{menge.NewStringSet(), menge.NewStringSet(), menge.NewStringSet()},
		{menge.NewStringSet("1"), menge.NewStringSet("1"), menge.NewStringSet("1")},
		{menge.NewStringSet("1"), menge.NewStringSet("2"), menge.NewStringSet()},
		{menge.NewStringSet("1"), menge.NewStringSet("1", "2"), menge.NewStringSet("1")},
		{menge.NewStringSet("1", "2"), menge.NewStringSet("1"), menge.NewStringSet("1")},
	}
	for _, c := range cases {
		got := c.set.Intersection(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestStringSet_Difference(t *testing.T) {
	cases := []struct {
		set  menge.StringSet
		arg  menge.StringSet
		want menge.StringSet
	}{
		{menge.NewStringSet(), menge.NewStringSet(), menge.NewStringSet()},
		{menge.NewStringSet("1"), menge.NewStringSet("1"), menge.NewStringSet()},
		{menge.NewStringSet("1"), menge.NewStringSet("2"), menge.NewStringSet("1")},
		{menge.NewStringSet("1"), menge.NewStringSet("1", "2"), menge.NewStringSet()},
		{menge.NewStringSet("1", "2"), menge.NewStringSet("1"), menge.NewStringSet("2")},
	}
	for _, c := range cases {
		got := c.set.Difference(c.arg)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestStringSet_IsSubsetOf(t *testing.T) {
	cases := []struct {
		set  menge.StringSet
		arg  menge.StringSet
		want bool
	}{
		{menge.NewStringSet(), menge.NewStringSet(), true},
		{menge.NewStringSet("1"), menge.NewStringSet("1"), true},
		{menge.NewStringSet("1"), menge.NewStringSet("1", "2"), true},
		{menge.NewStringSet("1", "2"), menge.NewStringSet("1"), false},
	}
	for _, c := range cases {
		got := c.set.IsSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestStringSet_IsProperSubsetOf(t *testing.T) {
	cases := []struct {
		set  menge.StringSet
		arg  menge.StringSet
		want bool
	}{
		{menge.NewStringSet(), menge.NewStringSet(), false},
		{menge.NewStringSet("1"), menge.NewStringSet("1"), false},
		{menge.NewStringSet("1"), menge.NewStringSet("1", "2"), true},
		{menge.NewStringSet("1", "2"), menge.NewStringSet("1"), false},
	}
	for _, c := range cases {
		got := c.set.IsProperSubsetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestStringSet_IsSupersetOf(t *testing.T) {
	cases := []struct {
		set  menge.StringSet
		arg  menge.StringSet
		want bool
	}{
		{menge.NewStringSet(), menge.NewStringSet(), true},
		{menge.NewStringSet("1"), menge.NewStringSet("1"), true},
		{menge.NewStringSet("1"), menge.NewStringSet("1", "2"), false},
		{menge.NewStringSet("1", "2"), menge.NewStringSet("1"), true},
	}
	for _, c := range cases {
		got := c.set.IsSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestStringSet_IsProperSupersetOf(t *testing.T) {
	cases := []struct {
		set  menge.StringSet
		arg  menge.StringSet
		want bool
	}{
		{menge.NewStringSet(), menge.NewStringSet(), false},
		{menge.NewStringSet("1"), menge.NewStringSet("1"), false},
		{menge.NewStringSet("1"), menge.NewStringSet("1", "2"), false},
		{menge.NewStringSet("1", "2"), menge.NewStringSet("1"), true},
	}
	for _, c := range cases {
		got := c.set.IsProperSupersetOf(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestStringSet_IsDisjointFrom(t *testing.T) {
	cases := []struct {
		set  menge.StringSet
		arg  menge.StringSet
		want bool
	}{
		{menge.NewStringSet(), menge.NewStringSet(), true},
		{menge.NewStringSet("1"), menge.NewStringSet("1"), false},
		{menge.NewStringSet("1"), menge.NewStringSet("2", "3"), true},
		{menge.NewStringSet("1", "2"), menge.NewStringSet("3"), true},
		{menge.NewStringSet("1"), menge.NewStringSet("1", "2"), false},
		{menge.NewStringSet("1", "2"), menge.NewStringSet("1"), false},
	}
	for _, c := range cases {
		got := c.set.IsDisjointFrom(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
