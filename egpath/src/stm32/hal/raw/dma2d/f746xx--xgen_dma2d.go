// +build f746xx

package dma2d

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"mmio"
	"unsafe"

	"stm32/o/f746xx/mmap"
)

type DMA2D_Periph struct {
	CR      CR
	ISR     ISR
	IFCR    IFCR
	FGMAR   FGMAR
	FGOR    FGOR
	BGMAR   BGMAR
	BGOR    BGOR
	FGPFCCR FGPFCCR
	FGCOLR  FGCOLR
	BGPFCCR BGPFCCR
	BGCOLR  BGCOLR
	FGCMAR  FGCMAR
	BGCMAR  BGCMAR
	OPFCCR  OPFCCR
	OCOLR   OCOLR
	OMAR    OMAR
	OOR     OOR
	NLR     NLR
	LWR     LWR
	AMTCR   AMTCR
	_       [236]uint32
	FGCLUT  [256]FGCLUT
	BGCLUT  [256]BGCLUT
}

func (p *DMA2D_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var DMA2D = (*DMA2D_Periph)(unsafe.Pointer(uintptr(mmap.DMA2D_BASE)))

type CR_Bits uint32

type CR struct{ mmio.U32 }

func (r *CR) Bits(mask CR_Bits) CR_Bits { return CR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CR) StoreBits(mask, b CR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CR) SetBits(mask CR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CR) ClearBits(mask CR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CR) Load() CR_Bits             { return CR_Bits(r.U32.Load()) }
func (r *CR) Store(b CR_Bits)           { r.U32.Store(uint32(b)) }

type CR_Mask struct{ mmio.UM32 }

func (rm CR_Mask) Load() CR_Bits   { return CR_Bits(rm.UM32.Load()) }
func (rm CR_Mask) Store(b CR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) START() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(START)}}
}

func (p *DMA2D_Periph) SUSP() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(SUSP)}}
}

func (p *DMA2D_Periph) ABORT() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(ABORT)}}
}

func (p *DMA2D_Periph) TEIE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(TEIE)}}
}

func (p *DMA2D_Periph) TCIE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(TCIE)}}
}

func (p *DMA2D_Periph) TWIE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(TWIE)}}
}

func (p *DMA2D_Periph) CAEIE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(CAEIE)}}
}

func (p *DMA2D_Periph) CTCIE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(CTCIE)}}
}

func (p *DMA2D_Periph) CEIE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(CEIE)}}
}

func (p *DMA2D_Periph) MODE() CR_Mask {
	return CR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(MODE)}}
}

type ISR_Bits uint32

type ISR struct{ mmio.U32 }

func (r *ISR) Bits(mask ISR_Bits) ISR_Bits { return ISR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ISR) StoreBits(mask, b ISR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ISR) SetBits(mask ISR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ISR) ClearBits(mask ISR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ISR) Load() ISR_Bits              { return ISR_Bits(r.U32.Load()) }
func (r *ISR) Store(b ISR_Bits)            { r.U32.Store(uint32(b)) }

type ISR_Mask struct{ mmio.UM32 }

func (rm ISR_Mask) Load() ISR_Bits   { return ISR_Bits(rm.UM32.Load()) }
func (rm ISR_Mask) Store(b ISR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) TEIF() ISR_Mask {
	return ISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(TEIF)}}
}

func (p *DMA2D_Periph) TCIF() ISR_Mask {
	return ISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(TCIF)}}
}

func (p *DMA2D_Periph) TWIF() ISR_Mask {
	return ISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(TWIF)}}
}

func (p *DMA2D_Periph) CAEIF() ISR_Mask {
	return ISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(CAEIF)}}
}

func (p *DMA2D_Periph) CTCIF() ISR_Mask {
	return ISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(CTCIF)}}
}

func (p *DMA2D_Periph) CEIF() ISR_Mask {
	return ISR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(CEIF)}}
}

type IFCR_Bits uint32

type IFCR struct{ mmio.U32 }

func (r *IFCR) Bits(mask IFCR_Bits) IFCR_Bits { return IFCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *IFCR) StoreBits(mask, b IFCR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IFCR) SetBits(mask IFCR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *IFCR) ClearBits(mask IFCR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *IFCR) Load() IFCR_Bits               { return IFCR_Bits(r.U32.Load()) }
func (r *IFCR) Store(b IFCR_Bits)             { r.U32.Store(uint32(b)) }

type IFCR_Mask struct{ mmio.UM32 }

func (rm IFCR_Mask) Load() IFCR_Bits   { return IFCR_Bits(rm.UM32.Load()) }
func (rm IFCR_Mask) Store(b IFCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) CTEIF() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(CTEIF)}}
}

func (p *DMA2D_Periph) CTCIF() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(CTCIF)}}
}

func (p *DMA2D_Periph) CTWIF() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(CTWIF)}}
}

func (p *DMA2D_Periph) CAECIF() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(CAECIF)}}
}

func (p *DMA2D_Periph) CCTCIF() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(CCTCIF)}}
}

func (p *DMA2D_Periph) CCEIF() IFCR_Mask {
	return IFCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(CCEIF)}}
}

type FGMAR_Bits uint32

type FGMAR struct{ mmio.U32 }

func (r *FGMAR) Bits(mask FGMAR_Bits) FGMAR_Bits { return FGMAR_Bits(r.U32.Bits(uint32(mask))) }
func (r *FGMAR) StoreBits(mask, b FGMAR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *FGMAR) SetBits(mask FGMAR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *FGMAR) ClearBits(mask FGMAR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *FGMAR) Load() FGMAR_Bits                { return FGMAR_Bits(r.U32.Load()) }
func (r *FGMAR) Store(b FGMAR_Bits)              { r.U32.Store(uint32(b)) }

type FGMAR_Mask struct{ mmio.UM32 }

func (rm FGMAR_Mask) Load() FGMAR_Bits   { return FGMAR_Bits(rm.UM32.Load()) }
func (rm FGMAR_Mask) Store(b FGMAR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) MA() FGMAR_Mask {
	return FGMAR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(MA)}}
}

type FGOR_Bits uint32

type FGOR struct{ mmio.U32 }

func (r *FGOR) Bits(mask FGOR_Bits) FGOR_Bits { return FGOR_Bits(r.U32.Bits(uint32(mask))) }
func (r *FGOR) StoreBits(mask, b FGOR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *FGOR) SetBits(mask FGOR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *FGOR) ClearBits(mask FGOR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *FGOR) Load() FGOR_Bits               { return FGOR_Bits(r.U32.Load()) }
func (r *FGOR) Store(b FGOR_Bits)             { r.U32.Store(uint32(b)) }

type FGOR_Mask struct{ mmio.UM32 }

func (rm FGOR_Mask) Load() FGOR_Bits   { return FGOR_Bits(rm.UM32.Load()) }
func (rm FGOR_Mask) Store(b FGOR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) LO() FGOR_Mask {
	return FGOR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(LO)}}
}

type BGMAR_Bits uint32

type BGMAR struct{ mmio.U32 }

func (r *BGMAR) Bits(mask BGMAR_Bits) BGMAR_Bits { return BGMAR_Bits(r.U32.Bits(uint32(mask))) }
func (r *BGMAR) StoreBits(mask, b BGMAR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BGMAR) SetBits(mask BGMAR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *BGMAR) ClearBits(mask BGMAR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *BGMAR) Load() BGMAR_Bits                { return BGMAR_Bits(r.U32.Load()) }
func (r *BGMAR) Store(b BGMAR_Bits)              { r.U32.Store(uint32(b)) }

type BGMAR_Mask struct{ mmio.UM32 }

func (rm BGMAR_Mask) Load() BGMAR_Bits   { return BGMAR_Bits(rm.UM32.Load()) }
func (rm BGMAR_Mask) Store(b BGMAR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) MA() BGMAR_Mask {
	return BGMAR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(MA)}}
}

type BGOR_Bits uint32

type BGOR struct{ mmio.U32 }

func (r *BGOR) Bits(mask BGOR_Bits) BGOR_Bits { return BGOR_Bits(r.U32.Bits(uint32(mask))) }
func (r *BGOR) StoreBits(mask, b BGOR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BGOR) SetBits(mask BGOR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *BGOR) ClearBits(mask BGOR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *BGOR) Load() BGOR_Bits               { return BGOR_Bits(r.U32.Load()) }
func (r *BGOR) Store(b BGOR_Bits)             { r.U32.Store(uint32(b)) }

type BGOR_Mask struct{ mmio.UM32 }

func (rm BGOR_Mask) Load() BGOR_Bits   { return BGOR_Bits(rm.UM32.Load()) }
func (rm BGOR_Mask) Store(b BGOR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) LO() BGOR_Mask {
	return BGOR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint32(LO)}}
}

type FGPFCCR_Bits uint32

type FGPFCCR struct{ mmio.U32 }

func (r *FGPFCCR) Bits(mask FGPFCCR_Bits) FGPFCCR_Bits { return FGPFCCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *FGPFCCR) StoreBits(mask, b FGPFCCR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *FGPFCCR) SetBits(mask FGPFCCR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *FGPFCCR) ClearBits(mask FGPFCCR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *FGPFCCR) Load() FGPFCCR_Bits                  { return FGPFCCR_Bits(r.U32.Load()) }
func (r *FGPFCCR) Store(b FGPFCCR_Bits)                { r.U32.Store(uint32(b)) }

type FGPFCCR_Mask struct{ mmio.UM32 }

func (rm FGPFCCR_Mask) Load() FGPFCCR_Bits   { return FGPFCCR_Bits(rm.UM32.Load()) }
func (rm FGPFCCR_Mask) Store(b FGPFCCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) CM() FGPFCCR_Mask {
	return FGPFCCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28)), uint32(CM)}}
}

func (p *DMA2D_Periph) CCM() FGPFCCR_Mask {
	return FGPFCCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28)), uint32(CCM)}}
}

func (p *DMA2D_Periph) START() FGPFCCR_Mask {
	return FGPFCCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28)), uint32(START)}}
}

func (p *DMA2D_Periph) CS() FGPFCCR_Mask {
	return FGPFCCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28)), uint32(CS)}}
}

func (p *DMA2D_Periph) AM() FGPFCCR_Mask {
	return FGPFCCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28)), uint32(AM)}}
}

func (p *DMA2D_Periph) ALPHA() FGPFCCR_Mask {
	return FGPFCCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28)), uint32(ALPHA)}}
}

type FGCOLR_Bits uint32

type FGCOLR struct{ mmio.U32 }

func (r *FGCOLR) Bits(mask FGCOLR_Bits) FGCOLR_Bits { return FGCOLR_Bits(r.U32.Bits(uint32(mask))) }
func (r *FGCOLR) StoreBits(mask, b FGCOLR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *FGCOLR) SetBits(mask FGCOLR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *FGCOLR) ClearBits(mask FGCOLR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *FGCOLR) Load() FGCOLR_Bits                 { return FGCOLR_Bits(r.U32.Load()) }
func (r *FGCOLR) Store(b FGCOLR_Bits)               { r.U32.Store(uint32(b)) }

type FGCOLR_Mask struct{ mmio.UM32 }

func (rm FGCOLR_Mask) Load() FGCOLR_Bits   { return FGCOLR_Bits(rm.UM32.Load()) }
func (rm FGCOLR_Mask) Store(b FGCOLR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) BLUE() FGCOLR_Mask {
	return FGCOLR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 32)), uint32(BLUE)}}
}

func (p *DMA2D_Periph) GREEN() FGCOLR_Mask {
	return FGCOLR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 32)), uint32(GREEN)}}
}

func (p *DMA2D_Periph) RED() FGCOLR_Mask {
	return FGCOLR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 32)), uint32(RED)}}
}

type BGPFCCR_Bits uint32

type BGPFCCR struct{ mmio.U32 }

func (r *BGPFCCR) Bits(mask BGPFCCR_Bits) BGPFCCR_Bits { return BGPFCCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *BGPFCCR) StoreBits(mask, b BGPFCCR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BGPFCCR) SetBits(mask BGPFCCR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *BGPFCCR) ClearBits(mask BGPFCCR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *BGPFCCR) Load() BGPFCCR_Bits                  { return BGPFCCR_Bits(r.U32.Load()) }
func (r *BGPFCCR) Store(b BGPFCCR_Bits)                { r.U32.Store(uint32(b)) }

type BGPFCCR_Mask struct{ mmio.UM32 }

func (rm BGPFCCR_Mask) Load() BGPFCCR_Bits   { return BGPFCCR_Bits(rm.UM32.Load()) }
func (rm BGPFCCR_Mask) Store(b BGPFCCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) CM() BGPFCCR_Mask {
	return BGPFCCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 36)), uint32(CM)}}
}

func (p *DMA2D_Periph) CCM() BGPFCCR_Mask {
	return BGPFCCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 36)), uint32(CCM)}}
}

func (p *DMA2D_Periph) START() BGPFCCR_Mask {
	return BGPFCCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 36)), uint32(START)}}
}

func (p *DMA2D_Periph) CS() BGPFCCR_Mask {
	return BGPFCCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 36)), uint32(CS)}}
}

func (p *DMA2D_Periph) AM() BGPFCCR_Mask {
	return BGPFCCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 36)), uint32(AM)}}
}

func (p *DMA2D_Periph) ALPHA() BGPFCCR_Mask {
	return BGPFCCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 36)), uint32(ALPHA)}}
}

type BGCOLR_Bits uint32

type BGCOLR struct{ mmio.U32 }

func (r *BGCOLR) Bits(mask BGCOLR_Bits) BGCOLR_Bits { return BGCOLR_Bits(r.U32.Bits(uint32(mask))) }
func (r *BGCOLR) StoreBits(mask, b BGCOLR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BGCOLR) SetBits(mask BGCOLR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *BGCOLR) ClearBits(mask BGCOLR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *BGCOLR) Load() BGCOLR_Bits                 { return BGCOLR_Bits(r.U32.Load()) }
func (r *BGCOLR) Store(b BGCOLR_Bits)               { r.U32.Store(uint32(b)) }

type BGCOLR_Mask struct{ mmio.UM32 }

func (rm BGCOLR_Mask) Load() BGCOLR_Bits   { return BGCOLR_Bits(rm.UM32.Load()) }
func (rm BGCOLR_Mask) Store(b BGCOLR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) BLUE() BGCOLR_Mask {
	return BGCOLR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 40)), uint32(BLUE)}}
}

func (p *DMA2D_Periph) GREEN() BGCOLR_Mask {
	return BGCOLR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 40)), uint32(GREEN)}}
}

func (p *DMA2D_Periph) RED() BGCOLR_Mask {
	return BGCOLR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 40)), uint32(RED)}}
}

type FGCMAR_Bits uint32

type FGCMAR struct{ mmio.U32 }

func (r *FGCMAR) Bits(mask FGCMAR_Bits) FGCMAR_Bits { return FGCMAR_Bits(r.U32.Bits(uint32(mask))) }
func (r *FGCMAR) StoreBits(mask, b FGCMAR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *FGCMAR) SetBits(mask FGCMAR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *FGCMAR) ClearBits(mask FGCMAR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *FGCMAR) Load() FGCMAR_Bits                 { return FGCMAR_Bits(r.U32.Load()) }
func (r *FGCMAR) Store(b FGCMAR_Bits)               { r.U32.Store(uint32(b)) }

type FGCMAR_Mask struct{ mmio.UM32 }

func (rm FGCMAR_Mask) Load() FGCMAR_Bits   { return FGCMAR_Bits(rm.UM32.Load()) }
func (rm FGCMAR_Mask) Store(b FGCMAR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) MA() FGCMAR_Mask {
	return FGCMAR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 44)), uint32(MA)}}
}

type BGCMAR_Bits uint32

type BGCMAR struct{ mmio.U32 }

func (r *BGCMAR) Bits(mask BGCMAR_Bits) BGCMAR_Bits { return BGCMAR_Bits(r.U32.Bits(uint32(mask))) }
func (r *BGCMAR) StoreBits(mask, b BGCMAR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BGCMAR) SetBits(mask BGCMAR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *BGCMAR) ClearBits(mask BGCMAR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *BGCMAR) Load() BGCMAR_Bits                 { return BGCMAR_Bits(r.U32.Load()) }
func (r *BGCMAR) Store(b BGCMAR_Bits)               { r.U32.Store(uint32(b)) }

type BGCMAR_Mask struct{ mmio.UM32 }

func (rm BGCMAR_Mask) Load() BGCMAR_Bits   { return BGCMAR_Bits(rm.UM32.Load()) }
func (rm BGCMAR_Mask) Store(b BGCMAR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) MA() BGCMAR_Mask {
	return BGCMAR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 48)), uint32(MA)}}
}

type OPFCCR_Bits uint32

type OPFCCR struct{ mmio.U32 }

func (r *OPFCCR) Bits(mask OPFCCR_Bits) OPFCCR_Bits { return OPFCCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *OPFCCR) StoreBits(mask, b OPFCCR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OPFCCR) SetBits(mask OPFCCR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *OPFCCR) ClearBits(mask OPFCCR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *OPFCCR) Load() OPFCCR_Bits                 { return OPFCCR_Bits(r.U32.Load()) }
func (r *OPFCCR) Store(b OPFCCR_Bits)               { r.U32.Store(uint32(b)) }

type OPFCCR_Mask struct{ mmio.UM32 }

func (rm OPFCCR_Mask) Load() OPFCCR_Bits   { return OPFCCR_Bits(rm.UM32.Load()) }
func (rm OPFCCR_Mask) Store(b OPFCCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) CM() OPFCCR_Mask {
	return OPFCCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(CM)}}
}

type OCOLR_Bits uint32

type OCOLR struct{ mmio.U32 }

func (r *OCOLR) Bits(mask OCOLR_Bits) OCOLR_Bits { return OCOLR_Bits(r.U32.Bits(uint32(mask))) }
func (r *OCOLR) StoreBits(mask, b OCOLR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OCOLR) SetBits(mask OCOLR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *OCOLR) ClearBits(mask OCOLR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *OCOLR) Load() OCOLR_Bits                { return OCOLR_Bits(r.U32.Load()) }
func (r *OCOLR) Store(b OCOLR_Bits)              { r.U32.Store(uint32(b)) }

type OCOLR_Mask struct{ mmio.UM32 }

func (rm OCOLR_Mask) Load() OCOLR_Bits   { return OCOLR_Bits(rm.UM32.Load()) }
func (rm OCOLR_Mask) Store(b OCOLR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) BLUE_1() OCOLR_Mask {
	return OCOLR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 56)), uint32(BLUE_1)}}
}

func (p *DMA2D_Periph) GREEN_1() OCOLR_Mask {
	return OCOLR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 56)), uint32(GREEN_1)}}
}

func (p *DMA2D_Periph) RED_1() OCOLR_Mask {
	return OCOLR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 56)), uint32(RED_1)}}
}

func (p *DMA2D_Periph) ALPHA_1() OCOLR_Mask {
	return OCOLR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 56)), uint32(ALPHA_1)}}
}

type OMAR_Bits uint32

type OMAR struct{ mmio.U32 }

func (r *OMAR) Bits(mask OMAR_Bits) OMAR_Bits { return OMAR_Bits(r.U32.Bits(uint32(mask))) }
func (r *OMAR) StoreBits(mask, b OMAR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OMAR) SetBits(mask OMAR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *OMAR) ClearBits(mask OMAR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *OMAR) Load() OMAR_Bits               { return OMAR_Bits(r.U32.Load()) }
func (r *OMAR) Store(b OMAR_Bits)             { r.U32.Store(uint32(b)) }

type OMAR_Mask struct{ mmio.UM32 }

func (rm OMAR_Mask) Load() OMAR_Bits   { return OMAR_Bits(rm.UM32.Load()) }
func (rm OMAR_Mask) Store(b OMAR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) MA() OMAR_Mask {
	return OMAR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 60)), uint32(MA)}}
}

type OOR_Bits uint32

type OOR struct{ mmio.U32 }

func (r *OOR) Bits(mask OOR_Bits) OOR_Bits { return OOR_Bits(r.U32.Bits(uint32(mask))) }
func (r *OOR) StoreBits(mask, b OOR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OOR) SetBits(mask OOR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *OOR) ClearBits(mask OOR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *OOR) Load() OOR_Bits              { return OOR_Bits(r.U32.Load()) }
func (r *OOR) Store(b OOR_Bits)            { r.U32.Store(uint32(b)) }

type OOR_Mask struct{ mmio.UM32 }

func (rm OOR_Mask) Load() OOR_Bits   { return OOR_Bits(rm.UM32.Load()) }
func (rm OOR_Mask) Store(b OOR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) LO() OOR_Mask {
	return OOR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 64)), uint32(LO)}}
}

type NLR_Bits uint32

type NLR struct{ mmio.U32 }

func (r *NLR) Bits(mask NLR_Bits) NLR_Bits { return NLR_Bits(r.U32.Bits(uint32(mask))) }
func (r *NLR) StoreBits(mask, b NLR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *NLR) SetBits(mask NLR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *NLR) ClearBits(mask NLR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *NLR) Load() NLR_Bits              { return NLR_Bits(r.U32.Load()) }
func (r *NLR) Store(b NLR_Bits)            { r.U32.Store(uint32(b)) }

type NLR_Mask struct{ mmio.UM32 }

func (rm NLR_Mask) Load() NLR_Bits   { return NLR_Bits(rm.UM32.Load()) }
func (rm NLR_Mask) Store(b NLR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) NL() NLR_Mask {
	return NLR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 68)), uint32(NL)}}
}

func (p *DMA2D_Periph) PL() NLR_Mask {
	return NLR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 68)), uint32(PL)}}
}

type LWR_Bits uint32

type LWR struct{ mmio.U32 }

func (r *LWR) Bits(mask LWR_Bits) LWR_Bits { return LWR_Bits(r.U32.Bits(uint32(mask))) }
func (r *LWR) StoreBits(mask, b LWR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *LWR) SetBits(mask LWR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *LWR) ClearBits(mask LWR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *LWR) Load() LWR_Bits              { return LWR_Bits(r.U32.Load()) }
func (r *LWR) Store(b LWR_Bits)            { r.U32.Store(uint32(b)) }

type LWR_Mask struct{ mmio.UM32 }

func (rm LWR_Mask) Load() LWR_Bits   { return LWR_Bits(rm.UM32.Load()) }
func (rm LWR_Mask) Store(b LWR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) LW() LWR_Mask {
	return LWR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 72)), uint32(LW)}}
}

type AMTCR_Bits uint32

type AMTCR struct{ mmio.U32 }

func (r *AMTCR) Bits(mask AMTCR_Bits) AMTCR_Bits { return AMTCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *AMTCR) StoreBits(mask, b AMTCR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *AMTCR) SetBits(mask AMTCR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *AMTCR) ClearBits(mask AMTCR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *AMTCR) Load() AMTCR_Bits                { return AMTCR_Bits(r.U32.Load()) }
func (r *AMTCR) Store(b AMTCR_Bits)              { r.U32.Store(uint32(b)) }

type AMTCR_Mask struct{ mmio.UM32 }

func (rm AMTCR_Mask) Load() AMTCR_Bits   { return AMTCR_Bits(rm.UM32.Load()) }
func (rm AMTCR_Mask) Store(b AMTCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *DMA2D_Periph) EN() AMTCR_Mask {
	return AMTCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 76)), uint32(EN)}}
}

func (p *DMA2D_Periph) DT() AMTCR_Mask {
	return AMTCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 76)), uint32(DT)}}
}

type FGCLUT_Bits uint32

type FGCLUT struct{ mmio.U32 }

func (r *FGCLUT) Bits(mask FGCLUT_Bits) FGCLUT_Bits { return FGCLUT_Bits(r.U32.Bits(uint32(mask))) }
func (r *FGCLUT) StoreBits(mask, b FGCLUT_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *FGCLUT) SetBits(mask FGCLUT_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *FGCLUT) ClearBits(mask FGCLUT_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *FGCLUT) Load() FGCLUT_Bits                 { return FGCLUT_Bits(r.U32.Load()) }
func (r *FGCLUT) Store(b FGCLUT_Bits)               { r.U32.Store(uint32(b)) }

type FGCLUT_Mask struct{ mmio.UM32 }

func (rm FGCLUT_Mask) Load() FGCLUT_Bits   { return FGCLUT_Bits(rm.UM32.Load()) }
func (rm FGCLUT_Mask) Store(b FGCLUT_Bits) { rm.UM32.Store(uint32(b)) }

type BGCLUT_Bits uint32

type BGCLUT struct{ mmio.U32 }

func (r *BGCLUT) Bits(mask BGCLUT_Bits) BGCLUT_Bits { return BGCLUT_Bits(r.U32.Bits(uint32(mask))) }
func (r *BGCLUT) StoreBits(mask, b BGCLUT_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BGCLUT) SetBits(mask BGCLUT_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *BGCLUT) ClearBits(mask BGCLUT_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *BGCLUT) Load() BGCLUT_Bits                 { return BGCLUT_Bits(r.U32.Load()) }
func (r *BGCLUT) Store(b BGCLUT_Bits)               { r.U32.Store(uint32(b)) }

type BGCLUT_Mask struct{ mmio.UM32 }

func (rm BGCLUT_Mask) Load() BGCLUT_Bits   { return BGCLUT_Bits(rm.UM32.Load()) }
func (rm BGCLUT_Mask) Store(b BGCLUT_Bits) { rm.UM32.Store(uint32(b)) }