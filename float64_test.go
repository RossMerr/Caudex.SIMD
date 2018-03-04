package SIMD_test

import (
	"reflect"
	"testing"

	SIMD "github.com/RossMerr/Caudex.SIMD"
)

func Test_Add(t *testing.T) {
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
			if got := SIMD.Add(tt.args.X1, tt.args.X2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Subtract(t *testing.T) {
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
			if got := SIMD.Subtract(tt.args.X1, tt.args.X2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Multiply(t *testing.T) {
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
			if got := SIMD.Multiply(tt.args.X1, tt.args.X2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
