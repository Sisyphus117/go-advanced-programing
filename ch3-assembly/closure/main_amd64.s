#include "textflag.h"
TEXT ·ptrToFunc(SB), NOSPLIT, $0-16
	MOVQ a+0(FP), AX
	MOVQ AX, ret+8(FP)
	RET

TEXT ·asmFunTwiceClosureAddr(SB), NOSPLIT, $0-8
	LEAQ ·asmFunTwiceClosureBody(SB), AX
	MOVQ AX, ret+0(FP)
	RET

TEXT ·asmFunTwiceClosureBody(SB), NOSPLIT|NEEDCTXT, $0-8
	MOVQ 8(DX), AX
	ADDQ AX, AX
	MOVQ AX, 8(DX)
	MOVQ AX, ret+0(FP)
	RET
