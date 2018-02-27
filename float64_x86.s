// Copyright (c) 2018 Ross Merrigan
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

// +build 386 amd64 amd64p32

#include "textflag.h"

// MOVOU is MOVDQU.

// func addFloat64(x, y []float64) []float64
TEXT ·addFloat64(SB), NOSPLIT, $0-48	
    MOVQ a+0(FP), SI
	MOVQ b+24(FP), DI
	MOVOU (SI), X
	MOVOU (DI), y
	ADDPD y, x
    MOVOU x, (SI)
	MOVOU y, (DI)
	RET

// func subtractFloat64(x, y []float64) []float64
TEXT ·subtractFloat64(SB), NOSPLIT, $0-48	
    MOVQ a+0(FP), SI
	MOVQ b+24(FP), DI
	MOVOU (SI), X
	MOVOU (DI), y
	SUBPD y, x
    MOVOU x, (SI)
	MOVOU y, (DI)
	RET	

// func multiplyFloat64(x, y []float64) []float64
TEXT ·multiplyFloat64(SB), NOSPLIT, $0-48	
    MOVQ a+0(FP), SI
	MOVQ b+24(FP), DI
	MOVOU (SI), X
	MOVOU (DI), y
	MULPD y, x
    MOVOU x, (SI)
	MOVOU y, (DI)
	RET		