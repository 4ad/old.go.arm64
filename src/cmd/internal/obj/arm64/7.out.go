// cmd/7c/7.out.h  from Vita Nuova.
// https://code.google.com/p/ken-cc/source/browse/src/cmd/7c/7.out.h
//
// 	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
// 	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
// 	Portions Copyright © 1997-1999 Vita Nuova Limited
// 	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
// 	Portions Copyright © 2004,2006 Bruce Ellis
// 	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
// 	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
// 	Portions Copyright © 2009 The Go Authors.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package arm64

const (
	NSNAME = 8
	NSYM   = 50
	NREG   = 32
	NFREG  = 32
)

const (
	REGMIN   = 7
	REGRT1   = 16
	REGRT2   = 17
	REGPR    = 18
	REGMAX   = 25
	REGENV   = 26
	REGTMP   = 27
	REGG     = 28
	REGFP    = 29
	REGLINK  = 30
	REGSP    = 31
	REGZERO  = 31
	FREGRET  = 0
	FREGMIN  = 7
	FREGMAX  = 26
	FREGEXT  = 26
	FREGZERO = 28
	FREGHALF = 29
	FREGONE  = 30
	FREGTWO  = 31
)

/*
 * GENERAL:
 *
 * compiler allocates R4 up as temps
 * compiler allocates register variables R7-R27
 * compiler allocates external registers R29 down
 *
 * compiler allocates register variables F7-F26
 * compiler allocates external registers F26 down
 */

/* compiler allocates register variables F0 up */
/* compiler allocates external registers F15 down */
const (
	BIG = 2048 - 8
)

const (
	LABEL   = 1 << 0
	LEAF    = 1 << 1
	FLOAT   = 1 << 2
	BRANCH  = 1 << 3
	LOAD    = 1 << 4
	FCMP    = 1 << 5
	SYNC    = 1 << 6
	LIST    = 1 << 7
	FOLL    = 1 << 8
	NOSCHED = 1 << 9
)

const (
	C_NONE = iota
	C_REG
	C_PAIR
	C_RSP
	C_SHIFT
	C_EXTREG
	C_FREG
	C_SPR
	C_COND
	C_ZCON
	C_ADDCON0
	C_ADDCON
	C_MOVCON
	C_BITCON
	C_ABCON
	C_MBCON
	C_LCON
	C_FCON
	C_VCON
	C_VCONADDR
	C_AACON
	C_LACON
	C_AECON
	C_SBRA
	C_LBRA
	C_NPAUTO
	C_NSAUTO
	C_PSAUTO
	C_PPAUTO
	C_UAUTO4K
	C_UAUTO8K
	C_UAUTO16K
	C_UAUTO32K
	C_UAUTO64K
	C_LAUTO
	C_SEXT1
	C_SEXT2
	C_SEXT4
	C_SEXT8
	C_SEXT16
	C_LEXT
	C_NPOREG
	C_NSOREG
	C_ZOREG
	C_PSOREG
	C_PPOREG
	C_UOREG4K
	C_UOREG8K
	C_UOREG16K
	C_UOREG32K
	C_UOREG64K
	C_LOREG
	C_ADDR
	C_ROFF
	C_XPOST
	C_XPRE
	C_VREG
	C_GOK
	C_NCLASS
)

const (
	AXXX = iota
	AADC
	AADCS
	AADCSW
	AADCW
	AADD
	AADDS
	AADDSW
	AADDW
	AADR
	AADRP
	AAND
	AANDS
	AANDSW
	AANDW
	AASR
	AASRW
	AAT
	AB
	ABFI
	ABFIW
	ABFM
	ABFMW
	ABFXIL
	ABFXILW
	ABIC
	ABICS
	ABICSW
	ABICW
	ABL
	ABRK
	ACBNZ
	ACBNZW
	ACBZ
	ACBZW
	ACCMN
	ACCMNW
	ACCMP
	ACCMPW
	ACINC
	ACINCW
	ACINV
	ACINVW
	ACLREX
	ACLS
	ACLSW
	ACLZ
	ACLZW
	ACMN
	ACMNW
	ACMP
	ACMPW
	ACNEG
	ACNEGW
	ACRC32B
	ACRC32CB
	ACRC32CH
	ACRC32CW
	ACRC32CX
	ACRC32H
	ACRC32W
	ACRC32X
	ACSEL
	ACSELW
	ACSET
	ACSETM
	ACSETMW
	ACSETW
	ACSINC
	ACSINCW
	ACSINV
	ACSINVW
	ACSNEG
	ACSNEGW
	ADC
	ADCPS1
	ADCPS2
	ADCPS3
	ADMB
	ADRPS
	ADSB
	AEON
	AEONW
	AEOR
	AEORW
	AERET
	AEXTR
	AEXTRW
	AHINT
	AHLT
	AHVC
	AIC
	AISB
	ALDAR
	ALDARB
	ALDARH
	ALDARW
	ALDAXP
	ALDAXPW
	ALDAXR
	ALDAXRB
	ALDAXRH
	ALDAXRW
	ALDP
	ALDXR
	ALDXRB
	ALDXRH
	ALDXRW
	ALDXP
	ALDXPW
	ALSL
	ALSLW
	ALSR
	ALSRW
	AMADD
	AMADDW
	AMNEG
	AMNEGW
	AMOVK
	AMOVKW
	AMOVN
	AMOVNW
	AMOVZ
	AMOVZW
	AMRS
	AMSR
	AMSUB
	AMSUBW
	AMUL
	AMULW
	AMVN
	AMVNW
	ANEG
	ANEGS
	ANEGSW
	ANEGW
	ANGC
	ANGCS
	ANGCSW
	ANGCW
	ANOP
	AORN
	AORNW
	AORR
	AORRW
	APRFM
	APRFUM
	ARBIT
	ARBITW
	AREM
	AREMW
	ARET
	AREV
	AREV16
	AREV16W
	AREV32
	AREVW
	AROR
	ARORW
	ASBC
	ASBCS
	ASBCSW
	ASBCW
	ASBFIZ
	ASBFIZW
	ASBFM
	ASBFMW
	ASBFX
	ASBFXW
	ASDIV
	ASDIVW
	ASEV
	ASEVL
	ASMADDL
	ASMC
	ASMNEGL
	ASMSUBL
	ASMULH
	ASMULL
	ASTXR
	ASTXRB
	ASTXRH
	ASTXP
	ASTXPW
	ASTXRW
	ASTLP
	ASTLPW
	ASTLR
	ASTLRB
	ASTLRH
	ASTLRW
	ASTLXP
	ASTLXPW
	ASTLXR
	ASTLXRB
	ASTLXRH
	ASTLXRW
	ASTP
	ASUB
	ASUBS
	ASUBSW
	ASUBW
	ASVC
	ASXTB
	ASXTBW
	ASXTH
	ASXTHW
	ASXTW
	ASYS
	ASYSL
	ATBNZ
	ATBZ
	ATLBI
	ATST
	ATSTW
	AUBFIZ
	AUBFIZW
	AUBFM
	AUBFMW
	AUBFX
	AUBFXW
	AUDIV
	AUDIVW
	AUMADDL
	AUMNEGL
	AUMSUBL
	AUMULH
	AUMULL
	AUREM
	AUREMW
	AUXTB
	AUXTH
	AUXTW
	AUXTBW
	AUXTHW
	AWFE
	AWFI
	AYIELD
	AMOVB
	AMOVBU
	AMOVH
	AMOVHU
	AMOVW
	AMOVWU
	AMOV
	AMOVNP
	AMOVNPW
	AMOVP
	AMOVPD
	AMOVPQ
	AMOVPS
	AMOVPSW
	AMOVPW
	ABEQ
	ABNE
	ABCS
	ABHS
	ABCC
	ABLO
	ABMI
	ABPL
	ABVS
	ABVC
	ABHI
	ABLS
	ABGE
	ABLT
	ABGT
	ABLE
	AFABSD
	AFABSS
	AFADDD
	AFADDS
	AFCCMPD
	AFCCMPED
	AFCCMPS
	AFCCMPES
	AFCMPD
	AFCMPED
	AFCMPES
	AFCMPS
	AFCVTSD
	AFCVTDS
	AFCVTZSD
	AFCVTZSDW
	AFCVTZSS
	AFCVTZSSW
	AFCVTZUD
	AFCVTZUDW
	AFCVTZUS
	AFCVTZUSW
	AFDIVD
	AFDIVS
	AFMOVD
	AFMOVS
	AFMULD
	AFMULS
	AFNEGD
	AFNEGS
	AFSQRTD
	AFSQRTS
	AFSUBD
	AFSUBS
	ASCVTFD
	ASCVTFS
	ASCVTFWD
	ASCVTFWS
	AUCVTFD
	AUCVTFS
	AUCVTFWD
	AUCVTFWS
	ATEXT
	ADATA
	AGLOBL
	AHISTORY
	ANAME
	AWORD
	ADYNT
	AINIT
	ABCASE
	ACASE
	ADWORD
	ASIGNAME
	AGOK
	ARETURN
	AEND
	AFCSELS
	AFCSELD
	AFMAXS
	AFMINS
	AFMAXD
	AFMIND
	AFMAXNMS
	AFMAXNMD
	AFNMULS
	AFNMULD
	AFRINTNS
	AFRINTND
	AFRINTPS
	AFRINTPD
	AFRINTMS
	AFRINTMD
	AFRINTZS
	AFRINTZD
	AFRINTAS
	AFRINTAD
	AFRINTXS
	AFRINTXD
	AFRINTIS
	AFRINTID
	AFMADDS
	AFMADDD
	AFMSUBS
	AFMSUBD
	AFNMADDS
	AFNMADDD
	AFNMSUBS
	AFNMSUBD
	AFMINNMS
	AFMINNMD
	AFCVTDH
	AFCVTHS
	AFCVTHD
	AFCVTSH
	AAESD
	AAESE
	AAESIMC
	AAESMC
	ASHA1C
	ASHA1H
	ASHA1M
	ASHA1P
	ASHA1SU0
	ASHA1SU1
	ASHA256H
	ASHA256H2
	ASHA256SU0
	ASHA256SU1
	AUNDEF
	AUSEFIELD
	ATYPE
	AFUNCDATA
	APCDATA
	ACHECKNIL
	AVARDEF
	AVARKILL
	ADUFFCOPY
	ADUFFZERO
	ALAST
)

/* form offset parameter to SYS; special register number */
func SYSARG5(op0 int, op1 int, Cn int, Cm int, op2 int) int {

	return op0<<19 | op1<<16 | Cn<<12 | Cm<<8 | op2<<5
}

func SYSARG4(op1 int, Cn int, Cm int, op2 int) int {
	return SYSARG5(0, op1, Cn, Cm, op2)
}

/* type/name */
const (
	D_GOK = 0 + iota
	D_NONE
	D_EXTERN
	D_STATIC
	D_AUTO
	D_PARAM
	D_BRANCH
	D_OREG
	D_XPRE
	D_XPOST
	D_CONST
	D_DCONST
	D_FCONST
	D_SCONST
	D_REG
	D_SP
	D_FREG
	D_VREG
	D_SPR
	D_FILE
	D_OCONST
	D_FILE1
	D_SHIFT
	D_PAIR
	D_ADDR
	D_ADRP
	D_ADRLO
	D_EXTREG
	D_ROFF
	D_COND
	D_VLANE
	D_VSET
	D_LAST
	D_R0        = 0
	D_F0        = D_R0 + NREG
	D_DAIF      = 3<<19 | 3<<16 | 4<<12 | 2<<8 | 1<<5
	D_NZCV      = 3<<19 | 3<<16 | 4<<12 | 2<<8 | 0<<5
	D_FPSR      = 3<<19 | 3<<16 | 4<<12 | 4<<8 | 1<<5
	D_FPCR      = 3<<19 | 3<<16 | 4<<12 | 4<<8 | 0<<5
	D_SPSR_EL1  = 3<<19 | 0<<16 | 4<<12 | 0<<8 | 0<<5
	D_ELR_EL1   = 3<<19 | 0<<16 | 4<<12 | 0<<8 | 1<<5
	D_SPSR_EL2  = 3<<19 | 4<<16 | 4<<12 | 0<<8 | 0<<5
	D_ELR_EL2   = 3<<19 | 4<<16 | 4<<12 | 0<<8 | 1<<5
	D_CurrentEL = 3<<19 | 0<<16 | 4<<12 | 2<<8 | 2<<5
	D_SP_EL0    = 3<<19 | 0<<16 | 4<<12 | 1<<8 | 0<<5
	D_SPSel     = 3<<19 | 0<<16 | 4<<12 | 2<<8 | 0<<5
	D_DAIFSet   = 1<<30 | 0
	D_DAIFClr   = 1<<30 | 1
)

/*
 * this is the ranlib header
 */
var SYMDEF string
