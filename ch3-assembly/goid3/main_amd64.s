#include "textflag.h"
TEXT ·getg(SB), NOSPLIT, $32-16
	MOVQ (TLS), AX
	MOVQ $type·runtime·g(SB), BX
	MOVQ AX, 8(SP)
	MOVQ BX, 0(SP)
	CALL runtime·convT2E(SB)
	MOVQ 16(SP), AX
	MOVQ 24(SP), BX

	MOVQ AX, ret+0(FP)
	MOVQ BX, ret+8(FP)
	RET
