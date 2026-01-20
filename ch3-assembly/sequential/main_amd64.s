TEXT ·main(SB), $24-0
	MOVQ $0, a-8*2(SP)
	MOVQ $0, b-8*1(SP)

	MOVQ $10, AX
	MOVQ AX, a-8*2(SP)

	MOVQ AX, 0(SP)
	CALL ·myprint(SB)

	MOVQ a-8*2(SP), AX

	// MOVQ b-8*1(SP), BX

	MOVQ  AX, BX
	ADDQ  BX, BX
	IMULQ AX, BX
	MOVQ  BX, b-8*1(SP)

	MOVQ BX, 0(SP)
	CALL ·myprint(SB)
	RET
