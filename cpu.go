// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package SIMD

var X86 x86

func init() {
	X86.HasSIMD = true
}

type x86 struct {
	HasSIMD bool
}
