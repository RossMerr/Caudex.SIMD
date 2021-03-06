// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package SIMD

// AddGeneric Add Packed Double-Precision Floating-Point Values
func AddGeneric(x, y [2]float64) [2]float64 {
	results := [2]float64{}
	for i := 0; i < 2; i++ {
		results[i] = x[i] + y[i]
	}
	return results
}

// MultiplyGeneric Multiply Packed Double-Precision Floating-Point Values
func MultiplyGeneric(x, y [2]float64) [2]float64 {
	results := [2]float64{}
	for i := 0; i < 2; i++ {
		results[i] = x[i] * y[i]
	}
	return results
}

// SubtractGeneric Subtract Packed Double-Precision Floating-Point Values
func SubtractGeneric(x, y [2]float64) [2]float64 {
	results := [2]float64{}
	for i := 0; i < 2; i++ {
		results[i] = x[i] - y[i]
	}
	return results
}

// PackedDoubleFloat take two Packed Double-Precision Floating-Point Values and does a arithmetic operation
type PackedDoubleFloat = func(x, y [2]float64) [2]float64

// Add Packed Double-Precision Floating-Point Values
var Add PackedDoubleFloat

// Multiply Packed Double-Precision Floating-Point Values
var Multiply PackedDoubleFloat

// Subtract Packed Double-Precision Floating-Point Values
var Subtract PackedDoubleFloat
