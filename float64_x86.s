// Copyright (c) 2018 Ross Merrigan
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// +build 386 amd64 amd64p32

#include "textflag.h"

// MOVOU is MOVDQU.

// func addfloat64(X1, X2 [2]float64) [2]float64
TEXT ·Addfloat64(SB), NOSPLIT, $0
	// Stack Pointer to MMX register 0 then 1
	MOVUPS x+0(FP), X0
	MOVUPS y+16(FP), X1
	// Add Single-Precision Floating-Point values on the MMX register 0 and 1, save to 0
	ADDPS X1, X0
	// MMX register 0 to Stack Pointer
	MOVUPS X0, ret+32(FP)
	// Return
	RET

