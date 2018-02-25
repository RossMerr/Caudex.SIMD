// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package float32x4

// Float32x4 SIMD vector of length 4
type Float32x4 struct {
}

func NewFloat32x4(a, b, c, d float32) *Float32x4 {
	return &Float32x4{}
}

func addGeneric(a, b *Float32x4) *Float32x4 {
	return nil
}

func multiplyGeneric(a, b *Float32x4) *Float32x4 {
	return nil
}

func subtractGeneric(a, b *Float32x4) *Float32x4 {
	return nil
}
