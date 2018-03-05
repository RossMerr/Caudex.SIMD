// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// +build !amd64,!amd64p32,!386,!arm64

package SIMD

func init() {
	Add = AddGeneric

	Multiply = MultiplyGeneric

	Subtract = SubtractGeneric
}
