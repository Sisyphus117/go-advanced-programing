#include "textflag.h"
TEXT ·If(SB), NOSPLIT, $0-32
	MOVQ ok+1*0(FP), CX
	MOVQ a+8*1(FP), AX
	MOVQ b+8*2(FP), BX

	CMPQ CX, $0
	JZ   L
	MOVQ AX, ret+24(FP)
	RET

L:
	MOVQ BX, ret+24(FP)
	RET
