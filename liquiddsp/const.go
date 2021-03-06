// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Wed, 24 Nov 2021 17:31:22 CST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package liquiddsp

/*
#cgo CFLAGS: -msse4.1 -Wall -fPIC -I. -I.. -I../include
#cgo LDFLAGS: -lm ${SRCDIR}/../libliquid.a
#include <inttypes.h>
#include <complex.h>
#include "liquid.h"
#include "liquid.internal.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// LiquidCrcNumSchemes as defined in include/liquid.h:1102
	LiquidCrcNumSchemes = 7
	// LiquidFecNumSchemes as defined in include/liquid.h:1164
	LiquidFecNumSchemes = 28
	// LiquidModemNumSchemes as defined in include/liquid.h:7261
	LiquidModemNumSchemes = 52
)

// CrcScheme as declared in include/liquid.h:1111
type CrcScheme int32

// CrcScheme enumeration from include/liquid.h:1111
const (
	LiquidCrcUnknown  CrcScheme = iota
	LiquidCrcNone     CrcScheme = 1
	LiquidCrcChecksum CrcScheme = 2
	LiquidCrc8        CrcScheme = 3
	LiquidCrc16       CrcScheme = 4
	LiquidCrc24       CrcScheme = 5
	LiquidCrc32       CrcScheme = 6
)

// FecScheme as declared in include/liquid.h:1202
type FecScheme int32

// FecScheme enumeration from include/liquid.h:1202
const (
	LiquidFecUnknown    FecScheme = iota
	LiquidFecNone       FecScheme = 1
	LiquidFecRep3       FecScheme = 2
	LiquidFecRep5       FecScheme = 3
	LiquidFecHamming74  FecScheme = 4
	LiquidFecHamming84  FecScheme = 5
	LiquidFecHamming128 FecScheme = 6
	LiquidFecGolay2412  FecScheme = 7
	LiquidFecSecded2216 FecScheme = 8
	LiquidFecSecded3932 FecScheme = 9
	LiquidFecSecded7264 FecScheme = 10
	LiquidFecConvV27    FecScheme = 11
	LiquidFecConvV29    FecScheme = 12
	LiquidFecConvV39    FecScheme = 13
	LiquidFecConvV615   FecScheme = 14
	LiquidFecConvV27p23 FecScheme = 15
	LiquidFecConvV27p34 FecScheme = 16
	LiquidFecConvV27p45 FecScheme = 17
	LiquidFecConvV27p56 FecScheme = 18
	LiquidFecConvV27p67 FecScheme = 19
	LiquidFecConvV27p78 FecScheme = 20
	LiquidFecConvV29p23 FecScheme = 21
	LiquidFecConvV29p34 FecScheme = 22
	LiquidFecConvV29p45 FecScheme = 23
	LiquidFecConvV29p56 FecScheme = 24
	LiquidFecConvV29p67 FecScheme = 25
	LiquidFecConvV29p78 FecScheme = 26
	LiquidFecRsM8       FecScheme = 27
)

// ModulationScheme as declared in include/liquid.h:7312
type ModulationScheme int32

// ModulationScheme enumeration from include/liquid.h:7312
const (
	LiquidModemUnknown   ModulationScheme = iota
	LiquidModemPsk2      ModulationScheme = 1
	LiquidModemPsk4      ModulationScheme = 2
	LiquidModemPsk8      ModulationScheme = 3
	LiquidModemPsk16     ModulationScheme = 4
	LiquidModemPsk32     ModulationScheme = 5
	LiquidModemPsk64     ModulationScheme = 6
	LiquidModemPsk128    ModulationScheme = 7
	LiquidModemPsk256    ModulationScheme = 8
	LiquidModemDpsk2     ModulationScheme = 9
	LiquidModemDpsk4     ModulationScheme = 10
	LiquidModemDpsk8     ModulationScheme = 11
	LiquidModemDpsk16    ModulationScheme = 12
	LiquidModemDpsk32    ModulationScheme = 13
	LiquidModemDpsk64    ModulationScheme = 14
	LiquidModemDpsk128   ModulationScheme = 15
	LiquidModemDpsk256   ModulationScheme = 16
	LiquidModemAsk2      ModulationScheme = 17
	LiquidModemAsk4      ModulationScheme = 18
	LiquidModemAsk8      ModulationScheme = 19
	LiquidModemAsk16     ModulationScheme = 20
	LiquidModemAsk32     ModulationScheme = 21
	LiquidModemAsk64     ModulationScheme = 22
	LiquidModemAsk128    ModulationScheme = 23
	LiquidModemAsk256    ModulationScheme = 24
	LiquidModemQam4      ModulationScheme = 25
	LiquidModemQam8      ModulationScheme = 26
	LiquidModemQam16     ModulationScheme = 27
	LiquidModemQam32     ModulationScheme = 28
	LiquidModemQam64     ModulationScheme = 29
	LiquidModemQam128    ModulationScheme = 30
	LiquidModemQam256    ModulationScheme = 31
	LiquidModemApsk4     ModulationScheme = 32
	LiquidModemApsk8     ModulationScheme = 33
	LiquidModemApsk16    ModulationScheme = 34
	LiquidModemApsk32    ModulationScheme = 35
	LiquidModemApsk64    ModulationScheme = 36
	LiquidModemApsk128   ModulationScheme = 37
	LiquidModemApsk256   ModulationScheme = 38
	LiquidModemBpsk      ModulationScheme = 39
	LiquidModemQpsk      ModulationScheme = 40
	LiquidModemOok       ModulationScheme = 41
	LiquidModemSqam32    ModulationScheme = 42
	LiquidModemSqam128   ModulationScheme = 43
	LiquidModemV29       ModulationScheme = 44
	LiquidModemArb16opt  ModulationScheme = 45
	LiquidModemArb32opt  ModulationScheme = 46
	LiquidModemArb64opt  ModulationScheme = 47
	LiquidModemArb128opt ModulationScheme = 48
	LiquidModemArb256opt ModulationScheme = 49
	LiquidModemArb64vt   ModulationScheme = 50
	LiquidModemArb       ModulationScheme = 51
)
