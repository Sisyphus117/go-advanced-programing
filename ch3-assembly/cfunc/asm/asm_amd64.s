#include "textflag.h"
TEXT ·AsmCallCAdd(SB), NOSPLIT, $0
	MOVQ cfun+0(FP), AX
	MOVQ a+8(FP), DI
	MOVQ b+16(FP), SI
	CALL AX
	MOVQ AX, ret+24(FP)
	RET
