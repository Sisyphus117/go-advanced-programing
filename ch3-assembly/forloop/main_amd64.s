#include "textflag.h"
TEXT ·Forloop(SB), NOSPLIT, $0-32
	MOVQ v0+0(FP), AX
	MOVQ cnt+8(FP), CX
	MOVQ step+16(FP), BX

	MOVQ $0, DX

LOOP:
	CMPQ DX, CX
	JGE  END
	ADDQ $1, DX
	ADDQ BX, AX
	JMP  LOOP

END:
	MOVQ AX, ret+24(FP)
	RET

