// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// +build amd64 amd64p32 386

package float64SIMD

import "github.com/RossMerr/intrinsic/sse2"

var Add = sse2.ADDSDm64float64

var Multiply = sse2.MULSDm64float64

var Subtract = sse2.SUBSDm64float64
