// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

TEXT _rt0_arm64_linux(SB),NOSPLIT,$-8
	// reference the dummy symbols so that they end up in the binary.
	MOV	$dummydata(SB), R11
	MOV	$dummyrodata(SB), R11
	MOV	$dummynoptrdata(SB), R11
	MOV	$dummybss(SB), R11
	MOV	$dummynoptrbss(SB), R11
	MOV	0(SP), R0	// argc
	ADD	$8, SP, R1	// argv
	BL	main(SB)

TEXT main(SB),NOSPLIT,$-8
	MOV	$runtime·rt0_go(SB), R2
	BL	(R2)
exit:
	MOV $0, R0
	MOV	$94, R8	// sys_exit
	SVC
	B	exit

// Linker has a bug, and we need non-zero length symbols in
// these sections.

// .data
DATA dummydata(SB)/4, $2
GLOBL dummydata(SB), 0, $4

// .rodata
DATA dummyrodata(SB)/4, $1
GLOBL dummyrodata(SB), RODATA, $4

// .noptrdata
DATA dummynoptrdata(SB)/4, $3
GLOBL dummynoptrdata(SB), NOPTR, $4

// .bss
GLOBL dummybss(SB), 0, $4

// .noptrbss
GLOBL dummynoptrbss(SB), NOPTR, $4
