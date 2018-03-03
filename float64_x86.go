// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// +build amd64 amd64p32 386

package float64SIMD

// go:noescape

// Addfloat64 Add Packed Double-Precision Floating-Point Values
func Addfloat64(X1, X2 float64) float64

// go:noescape

// subtractFloat64 Subtract Packed Double-Precision Floating-Point Values
func subtractFloat64(x [4]float64, y [4]float64) [4]float64

// go:noescape

// multiplyFloat64 Multiply Packed Double-Precision Floating-Point Values
func multiplyFloat64(x [4]float64, y [4]float64) [4]float64

var Add = Addfloat64

var Multiply = multiplyFloat64

var Subtract = subtractFloat64
