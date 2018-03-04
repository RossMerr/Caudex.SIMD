// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package float64SIMD_test

import (
	"reflect"
	"testing"

	"github.com/RossMerr/Caudex.SIMD"
)

func Test_addGeneric(t *testing.T) {
	type args struct {
		X1 [4]float32
		X2 [4]float32
	}
	tests := []struct {
		name string
		args args
		want [4]float32
	}{
		{
			args: args{
				X1: [4]float32{1, 2, 3, 4},
				X2: [4]float32{1, 2, 3, 4},
			},
			want: [4]float32{2, 4, 6, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := float64SIMD.Addfloat64(tt.args.X1, tt.args.X2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addGeneric() = %v, want %v", got, tt.want)
			}
		})
	}
}
