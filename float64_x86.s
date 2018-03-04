// Copyright (c) 2018 Ross Merrigan
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// +build 386 amd64 amd64p32

#include "textflag.h"

// func Addx86(x, y [2]float64) [2]float64
TEXT ·Addx86(SB), NOSPLIT, $0
	// Stack Pointer to MMX register 0 then 1
	MOVUPD x+0(FP), X0
	MOVUPD y+16(FP), X1
	// Add Double-Precision Floating-Point values on the MMX register 0 and 1, save to 0
	ADDPD X1, X0
	// MMX register 0 to Stack Pointer
	MOVUPD X0, ret+32(FP)
	// Return
	RET

// func Subtractx86(x, y [2]float64) [2]float64
TEXT ·Subtractx86(SB), NOSPLIT, $0
	// Stack Pointer to MMX register 0 then 1
	MOVUPD x+0(FP), X0
	MOVUPD y+16(FP), X1
	// Subtract Double-Precision Floating-Point values on the MMX register 0 and 1, save to 0
	SUBPD X1, X0
	// MMX register 0 to Stack Pointer
	MOVUPD X0, ret+32(FP)
	// Return
	RET

// func Multiplyx86(x, y [2]float64) [2]float64
TEXT ·Multiplyx86(SB), NOSPLIT, $0
	// Stack Pointer to MMX register 0 then 1
	MOVUPD x+0(FP), X0
	MOVUPD y+16(FP), X1
	// Multiply  Double-Precision Floating-Point values on the MMX register 0 and 1, save to 0
	MULPD X1, X0
	// MMX register 0 to Stack Pointer
	MOVUPD X0, ret+32(FP)
	// Return
	RET
