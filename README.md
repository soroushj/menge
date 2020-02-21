# Menge: Sets for Go

Efficient Go implementation of sets for all basic types.

[![Go Report Card](https://goreportcard.com/badge/github.com/soroushj/menge)](https://goreportcard.com/report/github.com/soroushj/menge)
[![Build Status](https://travis-ci.org/soroushj/menge.svg?branch=master)](https://travis-ci.org/soroushj/menge)
[![codecov](https://codecov.io/gh/soroushj/menge/branch/master/graph/badge.svg)](https://codecov.io/gh/soroushj/menge)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/soroushj/menge?status.svg)](https://pkg.go.dev/github.com/soroushj/menge?tab=doc)

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
You can protect them with [sync.RWMutex](https://golang.org/pkg/sync/#RWMutex).
Read the *Concurrency* section of [this article](https://blog.golang.org/go-maps-in-action) for more details.

## Examples

### The Basics

You can run this example [on the Go Playground](https://play.golang.org/p/SGXwGO6Ysv0).

```go
s := menge.NewIntSet(1, 2, 3)
fmt.Println("Set:", s)
fmt.Println("Has 2?", s.Has(2))
s.Add(2)
s.Add(2, 3, 4)
s.Remove(1)
s.Remove(1, 2)
for e := range s {
	fmt.Println("Element:", e)
}
fmt.Println("Size:", s.Size())
var a []int = s.AsSlice()
fmt.Println("Slice:", a)
s.Empty()
fmt.Println("Is empty?", s.IsEmpty())
```

### Set Operations

You can run this example [on the Go Playground](https://play.golang.org/p/zdOKLFfrcAH).

```go
a := menge.NewIntSet(1, 2)
b := menge.NewIntSet(2, 3)
fmt.Printf("Does %v equal %v? %v\n", a, b, a.Equals(b))
fmt.Printf("Is %v a subset of %v? %v\n", a, b, a.IsSubsetOf(b))
fmt.Printf("Are %v and %v disjoint? %v\n", a, b, a.IsDisjointFrom(b))
fmt.Printf("%v ⋃ %v = %v\n", a, b, a.Union(b))
fmt.Printf("%v ⋂ %v = %v\n", a, b, a.Intersection(b))
fmt.Printf("%v - %v = %v\n", a, b, a.Difference(b))
```
