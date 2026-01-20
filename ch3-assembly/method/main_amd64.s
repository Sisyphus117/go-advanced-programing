#include "textflag.h"
TEXT ·myInt·twice(SB), NOSPLIT, $0-16
	MOVQ a+0(FP), AX
	SHLQ $1, AX
	MOVQ AX, ret+8(FP)
	RET
