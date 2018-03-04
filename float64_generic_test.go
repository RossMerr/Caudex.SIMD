// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package SIMD_test

import (
	"reflect"
	"testing"

	"github.com/RossMerr/Caudex.SIMD"
)

func Test_AddGeneric(t *testing.T) {
	type args struct {
		X1 [2]float64
		X2 [2]float64
	}
	tests := []struct {
		name string
		args args
		want [2]float64
	}{
		{
			args: args{
				X1: [2]float64{1, 2},
				X2: [2]float64{1, 2},
			},
			want: [2]float64{2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SIMD.AddGeneric(tt.args.X1, tt.args.X2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddGeneric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_SubtractGeneric(t *testing.T) {
	type args struct {
		X1 [2]float64
		X2 [2]float64
	}
	tests := []struct {
		name string
		args args
		want [2]float64
	}{
		{
			args: args{
				X1: [2]float64{2, 4},
				X2: [2]float64{1, 2},
			},
			want: [2]float64{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SIMD.SubtractGeneric(tt.args.X1, tt.args.X2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubtractGeneric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MultiplyGeneric(t *testing.T) {
	type args struct {
		X1 [2]float64
		X2 [2]float64
	}
	tests := []struct {
		name string
		args args
		want [2]float64
	}{
		{
			args: args{
				X1: [2]float64{2, 4},
				X2: [2]float64{1, 2},
			},
			want: [2]float64{2, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SIMD.MultiplyGeneric(tt.args.X1, tt.args.X2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MultiplyGeneric() = %v, want %v", got, tt.want)
			}
		})
	}
}
