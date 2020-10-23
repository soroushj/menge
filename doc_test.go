package menge

import (
	"fmt"
)

func Example() {
	s := NewIntSet(1, 2, 3)
	fmt.Println("Set:", s)
	for e := range s {
		fmt.Println("Element:", e)
	}
	fmt.Println("Remove 1, 2; add 3, 4")
	s.Remove(1, 2)
	s.Add(3, 4)
	fmt.Println("Has 1?", s.Has(1))
	fmt.Println("Size:", s.Size())
	c := s.Clone()
	fmt.Printf("Clone: %v (%T)\n", c, c)
	l := s.AsSlice()
	fmt.Printf("Slice: %v (%T)\n", l, l)
	fmt.Println("Empty")
	s.Empty()
	fmt.Println("Is empty?", s.IsEmpty())
	a := NewIntSet(1)
	b := NewIntSet(1, 2)
	fmt.Printf("Does %v equal %v? %v\n", a, b, a.Equals(b))
	fmt.Printf("Is %v a subset of %v? %v\n", a, b, a.IsSubsetOf(b))
	fmt.Printf("Is %v a proper subset of %v? %v\n", a, b, a.IsProperSubsetOf(b))
	fmt.Printf("Is %v a superset of %v? %v\n", a, b, a.IsSupersetOf(b))
	fmt.Printf("Is %v a proper superset of %v? %v\n", a, b, a.IsProperSupersetOf(b))
	fmt.Printf("Are %v and %v disjoint? %v\n", a, b, a.IsDisjointFrom(b))
	fmt.Printf("%v ⋃ %v = %v\n", a, b, a.Union(b))
	fmt.Printf("%v ⋂ %v = %v\n", a, b, a.Intersection(b))
	fmt.Printf("%v - %v = %v\n", a, b, a.Difference(b))
}
