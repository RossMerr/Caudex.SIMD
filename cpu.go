// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package SIMD

var X64 x64

func init() {
	X64.HasSIMD = true
}

type x64 struct {
	HasSIMD bool
}
