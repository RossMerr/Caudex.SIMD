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
		X1 float64
		X2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{
				X1: 1,
				X2: 1,
			},
			want: 2,
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
