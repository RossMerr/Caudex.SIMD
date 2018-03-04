// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// +build amd64 amd64p32 386

package SIMD

import "github.com/RossMerr/Caudex.SIMD/internal/cpu"

// go:noescape

// Addx86 Add Packed Double-Precision Floating-Point Values
func Addx86(x, y [2]float64) [2]float64

// go:noescape

// Subtractx86 Subtract Packed Double-Precision Floating-Point Values
func Subtractx86(x, y [2]float64) [2]float64

// go:noescape

// Multiplyx86 Multiply Packed Double-Precision Floating-Point Values
func Multiplyx86(x, y [2]float64) [2]float64

func init() {
	if cpu.X86.HasSSE2 {
		Add = Addx86
	} else {
		Add = AddGeneric
	}

	if cpu.X86.HasSSE2 {
		Multiply = Multiplyx86
	} else {
		Multiply = MultiplyGeneric
	}

	if cpu.X86.HasSSE2 {
		Subtract = Subtractx86
	} else {
		Subtract = SubtractGeneric
	}
}

var Add PackedDoubleFloat

var Multiply PackedDoubleFloat

var Subtract PackedDoubleFloat
