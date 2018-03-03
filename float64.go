// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package float64SIMD

func AddGeneric(X1, X2 [2]float64) [2]float64 {
	results := [2]float64{}
	for i := 0; i < 2; i++ {
		results[i] = X1[i] + X2[i]
	}
	return results
}

func MultiplyGeneric(X1, X2 [4]float64) [4]float64 {
	results := [4]float64{}
	for i := 0; i < 4; i++ {
		results[i] = X1[i] * X2[i]
	}
	return results
}

func SubtractGeneric(X1, X2 [4]float64) [4]float64 {
	results := [4]float64{}
	for i := 0; i < 4; i++ {
		results[i] = X1[i] - X2[i]
	}
	return results
}
