// Copyright (c) 2018 Ross Merrigan
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// +build 386 amd64 amd64p32

#include "textflag.h"

// MOVOU is MOVDQU.

// func addfloat64(X1, X2 [2]float64) [2]float64
TEXT ·Addfloat64(SB), NOSPLIT, $0-48
	// Stack Pointer to MMX register 0
	MOVSD   8(SP), X0
	// Stack Pointer to MMX register 1
	MOVSD   16(SP), X1
	// Add Double-Precision Floating-Point values on the MMX register 0 and 1, save to 0
	ADDSD   X1, X0
	// MMX register 0 to Stack Pointer
	MOVSD X0, ret+24(SP)
	// Return
	RET
