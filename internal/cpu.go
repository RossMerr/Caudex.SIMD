// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package internal

var X86 x86
var ARM arm

func init() {
	X86.HasSSE2 = true
	X86.HasSSE3 = true
	X86.HasSSSE3 = true
	X86.HasSSE41 = true
	X86.HasSSE42 = true
	X86.HasAVX = true
	X86.HasAVX2 = true

	ARM.HasNEON = true
}

type x86 struct {
	// HasSSE2 Streaming SIMD Extensions 2
	HasSSE2 bool
	// HasSSE3 Streaming SIMD Extensions 3
	HasSSE3 bool
	// HasSSSE3 Supplemental Streaming SIMD Extensions 3
	HasSSSE3 bool
	// HasSSE41 Streaming SIMD Extensions 4.1
	HasSSE41 bool
	// HasSSE41 Streaming SIMD Extensions 4.2
	HasSSE42 bool
	// HasAVX Advanced Vector Extensions
	HasAVX bool
	// HasAVX2 Advanced Vector Extensions 2
	HasAVX2 bool
}

type arm struct {
	// HasNEON The Advanced SIMD extension
	HasNEON bool
}
