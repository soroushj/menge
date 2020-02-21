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
