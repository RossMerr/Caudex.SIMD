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
