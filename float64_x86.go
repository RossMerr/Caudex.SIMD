// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// +build amd64 amd64p32 386

package float64SIMD

// go:noescape

// addfloat64 Add Packed Double-Precision Floating-Point Values
func addfloat64(x []float64, y []float64)

// go:noescape

// subtractFloat64 Subtract Packed Double-Precision Floating-Point Values
func subtractFloat64(x []float64, y []float64)

// go:noescape

// multiplyFloat64 Multiply Packed Double-Precision Floating-Point Values
func multiplyFloat64(x []float64, y []float64)

var Add = addfloat64

var Multiply = multiplyFloat64

var Subtract = subtractFloat64
