package menge

import (
	"testing"
)

func TestNewStringSet(t *testing.T) {
	cases := []struct {
		arg  []string
		want StringSet
	}{
		{[]string{}, StringSet{}},
		{[]string{"1", "1"}, StringSet{"1": struct{}{}}},
		{[]string{"1", "2"}, StringSet{"1": struct{}{}, "2": struct{}{}}},
	}
	for _, c := range cases {
		got := NewStringSet(c.arg...)
		if !got.Equals(c.want) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestStringSet_Add(t *testing.T) {
	cases := []struct {
		set  StringSet
		arg  []string
		want StringSet
	}{
		{NewStringSet(), []string{}, NewStringSet()},
		{NewStringSet(), []string{"1", "1"}, NewStringSet("1")},
		{NewStringSet(), []string{"1", "2"}, NewStringSet("1", "2")},
		{NewStringSet("1"), []string{}, NewStringSet("1")},
		{NewStringSet("1"), []string{"1", "1"}, NewStringSet("1")},
		{NewStringSet("1"), []string{"2", "3"}, NewStringSet("1", "2", "3")},
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
		set  StringSet
		arg  []string
		want StringSet
	}{
		{NewStringSet(), []string{}, NewStringSet()},
		{NewStringSet("1"), []string{"1", "1"}, NewStringSet()},
		{NewStringSet("1", "2"), []string{"1", "2"}, NewStringSet()},
		{NewStringSet("1"), []string{}, NewStringSet("1")},
		{NewStringSet("1"), []string{"1", "1"}, NewStringSet()},
		{NewStringSet("1", "2"), []string{"3"}, NewStringSet("1", "2")},
		{NewStringSet("1", "2", "3"), []string{"2", "3"}, NewStringSet("1")},
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
		set  StringSet
		want StringSet
	}{
		{NewStringSet(), NewStringSet()},
		{NewStringSet("1", "2"), NewStringSet()},
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
		set  StringSet
		arg  string
		want bool
	}{
		{NewStringSet(), "1", false},
		{NewStringSet("2"), "1", false},
		{NewStringSet("1"), "1", true},
		{NewStringSet("1", "2"), "1", true},
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
		set  StringSet
		want int
	}{
		{NewStringSet(), 0},
		{NewStringSet("1", "2"), 2},
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
		set  StringSet
		want bool
	}{
		{NewStringSet(), true},
		{NewStringSet("1", "2"), false},
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
		set  StringSet
		want StringSet
	}{
		{NewStringSet(), NewStringSet()},
		{NewStringSet("1", "2"), NewStringSet("1", "2")},
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
		set  StringSet
		want []string
	}{
		{NewStringSet(), []string{}},
		{NewStringSet("1", "2"), []string{"1", "2"}},
	}
	for _, c := range cases {
		got := c.set.AsSlice()
		if len(got) != len(c.want) || !NewStringSet(got...).Equals(NewStringSet(c.want...)) {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}

func TestStringSet_String(t *testing.T) {
	cases := []struct {
		set  StringSet
		want []string
	}{
		{NewStringSet(), []string{"{}"}},
		{NewStringSet("1"), []string{"{1}"}},
		{NewStringSet("1", "2"), []string{"{1 2}", "{2 1}"}},
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
		set  StringSet
		arg  StringSet
		want bool
	}{
		{NewStringSet(), NewStringSet(), true},
		{NewStringSet("1", "2"), NewStringSet("2", "1"), true},
		{NewStringSet("1", "2"), NewStringSet("1"), false},
		{NewStringSet("1"), NewStringSet("1", "2"), false},
		{NewStringSet("1"), NewStringSet("2"), false},
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
		set  StringSet
		arg  StringSet
		want StringSet
	}{
		{NewStringSet(), NewStringSet(), NewStringSet()},
		{NewStringSet("1"), NewStringSet("1"), NewStringSet("1")},
		{NewStringSet("1"), NewStringSet("2"), NewStringSet("1", "2")},
		{NewStringSet("1"), NewStringSet("1", "2"), NewStringSet("1", "2")},
		{NewStringSet("1", "2"), NewStringSet("1"), NewStringSet("1", "2")},
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
		set  StringSet
		arg  StringSet
		want StringSet
	}{
		{NewStringSet(), NewStringSet(), NewStringSet()},
		{NewStringSet("1"), NewStringSet("1"), NewStringSet("1")},
		{NewStringSet("1"), NewStringSet("2"), NewStringSet()},
		{NewStringSet("1"), NewStringSet("1", "2"), NewStringSet("1")},
		{NewStringSet("1", "2"), NewStringSet("1"), NewStringSet("1")},
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
		set  StringSet
		arg  StringSet
		want StringSet
	}{
		{NewStringSet(), NewStringSet(), NewStringSet()},
		{NewStringSet("1"), NewStringSet("1"), NewStringSet()},
		{NewStringSet("1"), NewStringSet("2"), NewStringSet("1")},
		{NewStringSet("1"), NewStringSet("1", "2"), NewStringSet()},
		{NewStringSet("1", "2"), NewStringSet("1"), NewStringSet("2")},
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
		set  StringSet
		arg  StringSet
		want bool
	}{
		{NewStringSet(), NewStringSet(), true},
		{NewStringSet("1"), NewStringSet("1"), true},
		{NewStringSet("1"), NewStringSet("1", "2"), true},
		{NewStringSet("1", "2"), NewStringSet("1"), false},
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
		set  StringSet
		arg  StringSet
		want bool
	}{
		{NewStringSet(), NewStringSet(), false},
		{NewStringSet("1"), NewStringSet("1"), false},
		{NewStringSet("1"), NewStringSet("1", "2"), true},
		{NewStringSet("1", "2"), NewStringSet("1"), false},
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
		set  StringSet
		arg  StringSet
		want bool
	}{
		{NewStringSet(), NewStringSet(), true},
		{NewStringSet("1"), NewStringSet("1"), true},
		{NewStringSet("1"), NewStringSet("1", "2"), false},
		{NewStringSet("1", "2"), NewStringSet("1"), true},
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
		set  StringSet
		arg  StringSet
		want bool
	}{
		{NewStringSet(), NewStringSet(), false},
		{NewStringSet("1"), NewStringSet("1"), false},
		{NewStringSet("1"), NewStringSet("1", "2"), false},
		{NewStringSet("1", "2"), NewStringSet("1"), true},
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
		set  StringSet
		arg  StringSet
		want bool
	}{
		{NewStringSet(), NewStringSet(), true},
		{NewStringSet("1"), NewStringSet("1"), false},
		{NewStringSet("1"), NewStringSet("2", "3"), true},
		{NewStringSet("1", "2"), NewStringSet("3"), true},
		{NewStringSet("1"), NewStringSet("1", "2"), false},
		{NewStringSet("1", "2"), NewStringSet("1"), false},
	}
	for _, c := range cases {
		got := c.set.IsDisjointFrom(c.arg)
		if got != c.want {
			t.Errorf("case: %v got: %v", c, got)
		}
	}
}
