// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package float32x4

import SIMD "github.com/RossMerr/Caudex.SIMD"

var useSIMD = SIMD.X86.HasSIMD

//go:noescape
func AddAMD64(a, b *Float32x4) *Float32x4

func Add(a, b *Float32x4) *Float32x4 {
	if !useSIMD {
		return AddGeneric(a, b)
	} else {
		return AddAMD64(a, b)
	}
}

//go:noescape
func MultiplyAMD64(a, b *Float32x4) *Float32x4

func Multiply(a, b *Float32x4) *Float32x4 {
	if !useSIMD {
		return MultiplyGeneric(a, b)
	} else {
		return MultiplyAMD64(a, b)
	}
}

//go:noescape
func SubtractAMD64(a, b *Float32x4) *Float32x4

func Subtract(a, b *Float32x4) *Float32x4 {
	if !useSIMD {
		return SubtractGeneric(a, b)
	} else {
		return SubtractAMD64(a, b)
	}
}
