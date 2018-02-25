// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package float32x4

import SIMD "github.com/RossMerr/Caudex.SIMD"

var useSIMD = SIMD.X64.HasSIMD

//go:noescape
func addAMD64(a, b *Float32x4) *Float32x4

//go:noescape
func subtractAMD64(a, b *Float32x4) *Float32x4

//go:noescape
func multiplyAMD64(a, b *Float32x4) *Float32x4

func add(a, b *Float32x4) *Float32x4 {
	if !useSIMD {
		return addGeneric(a, b)
	} else {
		return addAMD64(a, b)
	}
}

func multiply(a, b *Float32x4) *Float32x4 {
	if !useSIMD {
		return multiplyGeneric(a, b)
	} else {
		return multiplyAMD64(a, b)
	}
}

func subtract(a, b *Float32x4) *Float32x4 {
	if !useSIMD {
		return subtractGeneric(a, b)
	} else {
		return subtractAMD64(a, b)
	}
}
