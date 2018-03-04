// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// +build amd64 amd64p32 386

package SIMD

// go:noescape

// Addx86 Add Packed Double-Precision Floating-Point Values
func Addx86(x, y [2]float64) [2]float64

// go:noescape

// Subtractx86 Subtract Packed Double-Precision Floating-Point Values
func Subtractx86(x, y [2]float64) [2]float64

// go:noescape

// Multiplyx86 Multiply Packed Double-Precision Floating-Point Values
func Multiplyx86(x, y [2]float64) [2]float64

var Add = Addx86

var Multiply = Multiplyx86

var Subtract = Subtractx86
