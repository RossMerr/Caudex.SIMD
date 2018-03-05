// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// +build arm64

package SIMD

import "github.com/RossMerr/Caudex.SIMD/internal/cpu"

// go:noescape

// AddARM64 Add Packed Double-Precision Floating-Point Values
func AddARM64(x, y [2]float64) [2]float64

// go:noescape

// SubtractARM64 Subtract Packed Double-Precision Floating-Point Values
func SubtractARM64(x, y [2]float64) [2]float64

// go:noescape

// MultiplyARM64 Multiply Packed Double-Precision Floating-Point Values
func MultiplyARM64(x, y [2]float64) [2]float64

func init() {
	if cpu.ARM64.HasASIMD {
		Add = AddARM64
	} else {
		Add = AddGeneric
	}

	if cpu.ARM64.HasASIMD {
		Multiply = MultiplyARM64
	} else {
		Multiply = MultiplyGeneric
	}

	if cpu.ARM64.HasASIMD {
		Subtract = SubtractARM64
	} else {
		Subtract = SubtractGeneric
	}
}
