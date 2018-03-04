// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// +build amd64 amd64p32 386

package SIMD_test

import (
	"reflect"
	"testing"

	"github.com/RossMerr/Caudex.SIMD"
)

func Test_Addx86(t *testing.T) {
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
			if got := SIMD.Addx86(tt.args.X1, tt.args.X2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Addx86() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Subtractx86(t *testing.T) {
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
			if got := SIMD.Subtractx86(tt.args.X1, tt.args.X2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subtractx86() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Multiplyx86(t *testing.T) {
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
			if got := SIMD.Multiplyx86(tt.args.X1, tt.args.X2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multiplyx86() = %v, want %v", got, tt.want)
			}
		})
	}
}
