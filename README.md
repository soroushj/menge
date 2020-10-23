# Menge: Sets for Go

[![PkgGoDev](https://pkg.go.dev/badge/github.com/soroushj/menge)](https://pkg.go.dev/github.com/soroushj/menge)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/soroushj/menge)](https://goreportcard.com/report/github.com/soroushj/menge)
[![Build Status](https://travis-ci.org/soroushj/menge.svg?branch=master)](https://travis-ci.org/soroushj/menge)
[![codecov](https://codecov.io/gh/soroushj/menge/branch/master/graph/badge.svg)](https://codecov.io/gh/soroushj/menge)

## Overview

Menge implements efficient sets of all basic types:

- String: `StringSet`
- Integer: `IntSet`, `Int8Set`, `Int16Set`, `Int32Set`, `Int64Set`
- Unsigned integer: `UIntSet`, `UInt8Set`, `UInt16Set`, `UInt32Set`, `UInt64Set`, `UIntPtrSet`
- Float: `Float32Set`, `Float64Set`
- Complex: `Complex64Set`, `Complex128Set`

Note that there are no set types for `byte` and `rune`, as they are just aliases for `uint8` and `int32`.
Also, there is no practical use to a set of type `bool`.

## Concurrency

Menge sets use Go maps as their underlying data structure.
As a result, these sets are [not safe for concurrent use](https://golang.org/doc/faq#atomic_maps).
You can protect them with [sync.RWMutex](https://pkg.go.dev/sync#RWMutex).
Read the *Concurrency* section of [this article](https://blog.golang.org/go-maps-in-action) for more details.

## Example

You can run this example [on the Go Playground](https://play.golang.org/p/ZbD_0DGcHWM).

```go
package main

import (
	"fmt"
	"github.com/soroushj/menge"
)

func main() {
	s := menge.NewIntSet(1, 2, 3)
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
	a := menge.NewIntSet(1)
	b := menge.NewIntSet(1, 2)
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
```
