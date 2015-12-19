package rcc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"
)

type Ctrl struct {
	CR        CR
	ICSCR     ICSCR
	CFGR      CFGR
	CIR       CIR
	AHBRSTR   AHBRSTR
	APB2RSTR  APB2RSTR
	APB1RSTR  APB1RSTR
	AHBENR    AHBENR
	APB2ENR   APB2ENR
	APB1ENR   APB1ENR
	AHBLPENR  AHBLPENR
	APB2LPENR APB2LPENR
	APB1LPENR APB1LPENR
	CSR       CSR
}

var RCC = (*Ctrl)(unsafe.Pointer(uintptr(0x40023800)))

type CR_Bits uint32

func (b CR_Bits) Field(mask CR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR_Bits) J(v int) CR_Bits {
	return CR_Bits(bits.Make32(v, uint32(mask)))
}

type CR struct{ r mmio.U32 }

func (r *CR) Bits(mask CR_Bits) CR_Bits { return CR_Bits(r.r.Bits(uint32(mask))) }
func (r *CR) StoreBits(mask, b CR_Bits) { r.r.StoreBits(uint32(mask), uint32(b)) }
func (r *CR) SetBits(mask CR_Bits)      { r.r.SetBits(uint32(mask)) }
func (r *CR) ClearBits(mask CR_Bits)    { r.r.ClearBits(uint32(mask)) }
func (r *CR) Load() CR_Bits             { return CR_Bits(r.r.Load()) }
func (r *CR) Store(b CR_Bits)           { r.r.Store(uint32(b)) }

type ICSCR_Bits uint32

func (b ICSCR_Bits) Field(mask ICSCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ICSCR_Bits) J(v int) ICSCR_Bits {
	return ICSCR_Bits(bits.Make32(v, uint32(mask)))
}

type ICSCR struct{ r mmio.U32 }

func (r *ICSCR) Bits(mask ICSCR_Bits) ICSCR_Bits { return ICSCR_Bits(r.r.Bits(uint32(mask))) }
func (r *ICSCR) StoreBits(mask, b ICSCR_Bits)    { r.r.StoreBits(uint32(mask), uint32(b)) }
func (r *ICSCR) SetBits(mask ICSCR_Bits)         { r.r.SetBits(uint32(mask)) }
func (r *ICSCR) ClearBits(mask ICSCR_Bits)       { r.r.ClearBits(uint32(mask)) }
func (r *ICSCR) Load() ICSCR_Bits                { return ICSCR_Bits(r.r.Load()) }
func (r *ICSCR) Store(b ICSCR_Bits)              { r.r.Store(uint32(b)) }

type CFGR_Bits uint32

func (b CFGR_Bits) Field(mask CFGR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFGR_Bits) J(v int) CFGR_Bits {
	return CFGR_Bits(bits.Make32(v, uint32(mask)))
}

type CFGR struct{ r mmio.U32 }

func (r *CFGR) Bits(mask CFGR_Bits) CFGR_Bits { return CFGR_Bits(r.r.Bits(uint32(mask))) }
func (r *CFGR) StoreBits(mask, b CFGR_Bits)   { r.r.StoreBits(uint32(mask), uint32(b)) }
func (r *CFGR) SetBits(mask CFGR_Bits)        { r.r.SetBits(uint32(mask)) }
func (r *CFGR) ClearBits(mask CFGR_Bits)      { r.r.ClearBits(uint32(mask)) }
func (r *CFGR) Load() CFGR_Bits               { return CFGR_Bits(r.r.Load()) }
func (r *CFGR) Store(b CFGR_Bits)             { r.r.Store(uint32(b)) }

type CIR_Bits uint32

func (b CIR_Bits) Field(mask CIR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CIR_Bits) J(v int) CIR_Bits {
	return CIR_Bits(bits.Make32(v, uint32(mask)))
}

type CIR struct{ r mmio.U32 }

func (r *CIR) Bits(mask CIR_Bits) CIR_Bits { return CIR_Bits(r.r.Bits(uint32(mask))) }
func (r *CIR) StoreBits(mask, b CIR_Bits)  { r.r.StoreBits(uint32(mask), uint32(b)) }
func (r *CIR) SetBits(mask CIR_Bits)       { r.r.SetBits(uint32(mask)) }
func (r *CIR) ClearBits(mask CIR_Bits)     { r.r.ClearBits(uint32(mask)) }
func (r *CIR) Load() CIR_Bits              { return CIR_Bits(r.r.Load()) }
func (r *CIR) Store(b CIR_Bits)            { r.r.Store(uint32(b)) }

type AHBRSTR_Bits uint32

func (b AHBRSTR_Bits) Field(mask AHBRSTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AHBRSTR_Bits) J(v int) AHBRSTR_Bits {
	return AHBRSTR_Bits(bits.Make32(v, uint32(mask)))
}

type AHBRSTR struct{ r mmio.U32 }

func (r *AHBRSTR) Bits(mask AHBRSTR_Bits) AHBRSTR_Bits { return AHBRSTR_Bits(r.r.Bits(uint32(mask))) }
func (r *AHBRSTR) StoreBits(mask, b AHBRSTR_Bits)      { r.r.StoreBits(uint32(mask), uint32(b)) }
func (r *AHBRSTR) SetBits(mask AHBRSTR_Bits)           { r.r.SetBits(uint32(mask)) }
func (r *AHBRSTR) ClearBits(mask AHBRSTR_Bits)         { r.r.ClearBits(uint32(mask)) }
func (r *AHBRSTR) Load() AHBRSTR_Bits                  { return AHBRSTR_Bits(r.r.Load()) }
func (r *AHBRSTR) Store(b AHBRSTR_Bits)                { r.r.Store(uint32(b)) }

type APB2RSTR_Bits uint32

func (b APB2RSTR_Bits) Field(mask APB2RSTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB2RSTR_Bits) J(v int) APB2RSTR_Bits {
	return APB2RSTR_Bits(bits.Make32(v, uint32(mask)))
}

type APB2RSTR struct{ r mmio.U32 }

func (r *APB2RSTR) Bits(mask APB2RSTR_Bits) APB2RSTR_Bits {
	return APB2RSTR_Bits(r.r.Bits(uint32(mask)))
}
func (r *APB2RSTR) StoreBits(mask, b APB2RSTR_Bits) { r.r.StoreBits(uint32(mask), uint32(b)) }
func (r *APB2RSTR) SetBits(mask APB2RSTR_Bits)      { r.r.SetBits(uint32(mask)) }
func (r *APB2RSTR) ClearBits(mask APB2RSTR_Bits)    { r.r.ClearBits(uint32(mask)) }
func (r *APB2RSTR) Load() APB2RSTR_Bits             { return APB2RSTR_Bits(r.r.Load()) }
func (r *APB2RSTR) Store(b APB2RSTR_Bits)           { r.r.Store(uint32(b)) }

type APB1RSTR_Bits uint32

func (b APB1RSTR_Bits) Field(mask APB1RSTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB1RSTR_Bits) J(v int) APB1RSTR_Bits {
	return APB1RSTR_Bits(bits.Make32(v, uint32(mask)))
}

type APB1RSTR struct{ r mmio.U32 }

func (r *APB1RSTR) Bits(mask APB1RSTR_Bits) APB1RSTR_Bits {
	return APB1RSTR_Bits(r.r.Bits(uint32(mask)))
}
func (r *APB1RSTR) StoreBits(mask, b APB1RSTR_Bits) { r.r.StoreBits(uint32(mask), uint32(b)) }
func (r *APB1RSTR) SetBits(mask APB1RSTR_Bits)      { r.r.SetBits(uint32(mask)) }
func (r *APB1RSTR) ClearBits(mask APB1RSTR_Bits)    { r.r.ClearBits(uint32(mask)) }
func (r *APB1RSTR) Load() APB1RSTR_Bits             { return APB1RSTR_Bits(r.r.Load()) }
func (r *APB1RSTR) Store(b APB1RSTR_Bits)           { r.r.Store(uint32(b)) }

type AHBENR_Bits uint32

func (b AHBENR_Bits) Field(mask AHBENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AHBENR_Bits) J(v int) AHBENR_Bits {
	return AHBENR_Bits(bits.Make32(v, uint32(mask)))
}

type AHBENR struct{ r mmio.U32 }

func (r *AHBENR) Bits(mask AHBENR_Bits) AHBENR_Bits { return AHBENR_Bits(r.r.Bits(uint32(mask))) }
func (r *AHBENR) StoreBits(mask, b AHBENR_Bits)     { r.r.StoreBits(uint32(mask), uint32(b)) }
func (r *AHBENR) SetBits(mask AHBENR_Bits)          { r.r.SetBits(uint32(mask)) }
func (r *AHBENR) ClearBits(mask AHBENR_Bits)        { r.r.ClearBits(uint32(mask)) }
func (r *AHBENR) Load() AHBENR_Bits                 { return AHBENR_Bits(r.r.Load()) }
func (r *AHBENR) Store(b AHBENR_Bits)               { r.r.Store(uint32(b)) }

type APB2ENR_Bits uint32

func (b APB2ENR_Bits) Field(mask APB2ENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB2ENR_Bits) J(v int) APB2ENR_Bits {
	return APB2ENR_Bits(bits.Make32(v, uint32(mask)))
}

type APB2ENR struct{ r mmio.U32 }

func (r *APB2ENR) Bits(mask APB2ENR_Bits) APB2ENR_Bits { return APB2ENR_Bits(r.r.Bits(uint32(mask))) }
func (r *APB2ENR) StoreBits(mask, b APB2ENR_Bits)      { r.r.StoreBits(uint32(mask), uint32(b)) }
func (r *APB2ENR) SetBits(mask APB2ENR_Bits)           { r.r.SetBits(uint32(mask)) }
func (r *APB2ENR) ClearBits(mask APB2ENR_Bits)         { r.r.ClearBits(uint32(mask)) }
func (r *APB2ENR) Load() APB2ENR_Bits                  { return APB2ENR_Bits(r.r.Load()) }
func (r *APB2ENR) Store(b APB2ENR_Bits)                { r.r.Store(uint32(b)) }

type APB1ENR_Bits uint32

func (b APB1ENR_Bits) Field(mask APB1ENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB1ENR_Bits) J(v int) APB1ENR_Bits {
	return APB1ENR_Bits(bits.Make32(v, uint32(mask)))
}

type APB1ENR struct{ r mmio.U32 }

func (r *APB1ENR) Bits(mask APB1ENR_Bits) APB1ENR_Bits { return APB1ENR_Bits(r.r.Bits(uint32(mask))) }
func (r *APB1ENR) StoreBits(mask, b APB1ENR_Bits)      { r.r.StoreBits(uint32(mask), uint32(b)) }
func (r *APB1ENR) SetBits(mask APB1ENR_Bits)           { r.r.SetBits(uint32(mask)) }
func (r *APB1ENR) ClearBits(mask APB1ENR_Bits)         { r.r.ClearBits(uint32(mask)) }
func (r *APB1ENR) Load() APB1ENR_Bits                  { return APB1ENR_Bits(r.r.Load()) }
func (r *APB1ENR) Store(b APB1ENR_Bits)                { r.r.Store(uint32(b)) }

type AHBLPENR_Bits uint32

func (b AHBLPENR_Bits) Field(mask AHBLPENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AHBLPENR_Bits) J(v int) AHBLPENR_Bits {
	return AHBLPENR_Bits(bits.Make32(v, uint32(mask)))
}

type AHBLPENR struct{ r mmio.U32 }

func (r *AHBLPENR) Bits(mask AHBLPENR_Bits) AHBLPENR_Bits {
	return AHBLPENR_Bits(r.r.Bits(uint32(mask)))
}
func (r *AHBLPENR) StoreBits(mask, b AHBLPENR_Bits) { r.r.StoreBits(uint32(mask), uint32(b)) }
func (r *AHBLPENR) SetBits(mask AHBLPENR_Bits)      { r.r.SetBits(uint32(mask)) }
func (r *AHBLPENR) ClearBits(mask AHBLPENR_Bits)    { r.r.ClearBits(uint32(mask)) }
func (r *AHBLPENR) Load() AHBLPENR_Bits             { return AHBLPENR_Bits(r.r.Load()) }
func (r *AHBLPENR) Store(b AHBLPENR_Bits)           { r.r.Store(uint32(b)) }

type APB2LPENR_Bits uint32

func (b APB2LPENR_Bits) Field(mask APB2LPENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB2LPENR_Bits) J(v int) APB2LPENR_Bits {
	return APB2LPENR_Bits(bits.Make32(v, uint32(mask)))
}

type APB2LPENR struct{ r mmio.U32 }

func (r *APB2LPENR) Bits(mask APB2LPENR_Bits) APB2LPENR_Bits {
	return APB2LPENR_Bits(r.r.Bits(uint32(mask)))
}
func (r *APB2LPENR) StoreBits(mask, b APB2LPENR_Bits) { r.r.StoreBits(uint32(mask), uint32(b)) }
func (r *APB2LPENR) SetBits(mask APB2LPENR_Bits)      { r.r.SetBits(uint32(mask)) }
func (r *APB2LPENR) ClearBits(mask APB2LPENR_Bits)    { r.r.ClearBits(uint32(mask)) }
func (r *APB2LPENR) Load() APB2LPENR_Bits             { return APB2LPENR_Bits(r.r.Load()) }
func (r *APB2LPENR) Store(b APB2LPENR_Bits)           { r.r.Store(uint32(b)) }

type APB1LPENR_Bits uint32

func (b APB1LPENR_Bits) Field(mask APB1LPENR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB1LPENR_Bits) J(v int) APB1LPENR_Bits {
	return APB1LPENR_Bits(bits.Make32(v, uint32(mask)))
}

type APB1LPENR struct{ r mmio.U32 }

func (r *APB1LPENR) Bits(mask APB1LPENR_Bits) APB1LPENR_Bits {
	return APB1LPENR_Bits(r.r.Bits(uint32(mask)))
}
func (r *APB1LPENR) StoreBits(mask, b APB1LPENR_Bits) { r.r.StoreBits(uint32(mask), uint32(b)) }
func (r *APB1LPENR) SetBits(mask APB1LPENR_Bits)      { r.r.SetBits(uint32(mask)) }
func (r *APB1LPENR) ClearBits(mask APB1LPENR_Bits)    { r.r.ClearBits(uint32(mask)) }
func (r *APB1LPENR) Load() APB1LPENR_Bits             { return APB1LPENR_Bits(r.r.Load()) }
func (r *APB1LPENR) Store(b APB1LPENR_Bits)           { r.r.Store(uint32(b)) }

type CSR_Bits uint32

func (b CSR_Bits) Field(mask CSR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSR_Bits) J(v int) CSR_Bits {
	return CSR_Bits(bits.Make32(v, uint32(mask)))
}

type CSR struct{ r mmio.U32 }

func (r *CSR) Bits(mask CSR_Bits) CSR_Bits { return CSR_Bits(r.r.Bits(uint32(mask))) }
func (r *CSR) StoreBits(mask, b CSR_Bits)  { r.r.StoreBits(uint32(mask), uint32(b)) }
func (r *CSR) SetBits(mask CSR_Bits)       { r.r.SetBits(uint32(mask)) }
func (r *CSR) ClearBits(mask CSR_Bits)     { r.r.ClearBits(uint32(mask)) }
func (r *CSR) Load() CSR_Bits              { return CSR_Bits(r.r.Load()) }
func (r *CSR) Store(b CSR_Bits)            { r.r.Store(uint32(b)) }

func (p *Ctrl) LSION() mmio.Bits32 {
	return mmio.Bits32{
		(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(LSION),
	}
}
func (p *Ctrl) LSIRDY() mmio.Bits32 {
	return mmio.Bits32{
		(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(LSIRDY),
	}
}
func (p *Ctrl) LSEON() mmio.Bits32 {
	return mmio.Bits32{
		(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(LSEON),
	}
}
func (p *Ctrl) LSERDY() mmio.Bits32 {
	return mmio.Bits32{
		(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(LSERDY),
	}
}
func (p *Ctrl) LSEBYP() mmio.Bits32 {
	return mmio.Bits32{
		(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(LSEBYP),
	}
}
func (p *Ctrl) LSECSSON() mmio.Bits32 {
	return mmio.Bits32{
		(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(LSECSSON),
	}
}
func (p *Ctrl) LSECSSD() mmio.Bits32 {
	return mmio.Bits32{
		(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(LSECSSD),
	}
}
func (p *Ctrl) RTCSEL() mmio.Bits32 {
	return mmio.Bits32{
		(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(RTCSEL),
	}
}
func (p *Ctrl) RTCEN() mmio.Bits32 {
	return mmio.Bits32{
		(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(RTCEN),
	}
}
func (p *Ctrl) RTCRST() mmio.Bits32 {
	return mmio.Bits32{
		(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(RTCRST),
	}
}
func (p *Ctrl) RMVF() mmio.Bits32 {
	return mmio.Bits32{
		(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(RMVF),
	}
}
func (p *Ctrl) OBLRSTF() mmio.Bits32 {
	return mmio.Bits32{
		(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(OBLRSTF),
	}
}
func (p *Ctrl) PINRSTF() mmio.Bits32 {
	return mmio.Bits32{
		(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(PINRSTF),
	}
}
func (p *Ctrl) PORRSTF() mmio.Bits32 {
	return mmio.Bits32{
		(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(PORRSTF),
	}
}
func (p *Ctrl) SFTRSTF() mmio.Bits32 {
	return mmio.Bits32{
		(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(SFTRSTF),
	}
}
func (p *Ctrl) IWDGRSTF() mmio.Bits32 {
	return mmio.Bits32{
		(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(IWDGRSTF),
	}
}
func (p *Ctrl) WWDGRSTF() mmio.Bits32 {
	return mmio.Bits32{
		(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(WWDGRSTF),
	}
}
func (p *Ctrl) LPWRRSTF() mmio.Bits32 {
	return mmio.Bits32{
		(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(LPWRRSTF),
	}
}
