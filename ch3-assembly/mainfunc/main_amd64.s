#include "textflag.h"
TEXT ·main(SB), $16-0
	MOVQ ·helloworld+0(SB), AX
	MOVQ ·helloworld+8(SB), BX
	MOVQ AX, 0(SP)
	MOVQ AX, 0(SP)
	CALL ·myprint(SB)
	RET
