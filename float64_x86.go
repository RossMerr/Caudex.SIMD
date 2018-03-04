// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// +build amd64 amd64p32 386

package float64SIMD

// go:noescape

// Addfloat64 Add Packed Double-Precision Floating-Point Values
func Addfloat64(x, y [2]float64) [2]float64

// go:noescape

// SubtractFloat64 Subtract Packed Double-Precision Floating-Point Values
func SubtractFloat64(x, y [2]float64) [2]float64

// go:noescape

// MultiplyFloat64 Multiply Packed Double-Precision Floating-Point Values
func MultiplyFloat64(x, y [2]float64) [2]float64

var Add = Addfloat64

var Multiply = MultiplyFloat64

var Subtract = SubtractFloat64
