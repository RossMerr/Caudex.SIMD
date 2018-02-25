// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// +build !amd64,!amd64p32,!386,!arm,!s390x,!arm64

package float32x4

var Add = addGeneric

var Multiply = multiplyGeneric

var Subtract = subtractGeneric
