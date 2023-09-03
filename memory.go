package main

import (
	"errors"
)

const MaxMemory = 1024 * 64

type (
	Mem struct {
		Data [MaxMemory]Byte
	}
)

func (m *Mem) Initialize() {
	for i := 0; i < MaxMemory; i++ {
		m.Data[i] = 0
	}
}

func (m *Mem) NextByte(pc *Byte, cycle *uint32) Byte {
	x := m.Data[*pc]
	*pc++
	*cycle--

	return x
}

func (m *Mem) NextWord(pc *Byte, cycle *uint32) Word {
	x := m.Data[*pc]
	*pc++

	x |= m.Data[*pc] << 8
	*pc++
	*cycle -= 2

	return Word(x)
}

func (m *Mem) WriteWord(data Byte, address *Byte, cycle *uint32) {
	m.Data[*address] = data & 0xFF
	*address++
	m.Data[*address+1] = (data >> 8)
	*address++
	*cycle -= 2
}

func (m *Mem) ReadByte(address Byte, cycle *uint32) (*Byte, error) {
	if address < 0 || address > MaxMemory {
		return nil, errors.New("out of range")
	}

	x := m.Data[address]
	*cycle--

	return &x, nil
}

// TODO: Use me when needed
// func isLittleEndian() bool {
// 	buf := [2]byte{}
// 	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)

// 	switch buf {
// 	case [2]byte{0xCD, 0xAB}:
// 		return true
// 	case [2]byte{0xAB, 0xCD}:
// 		return false
// 	default:
// 		panic("Could not determine native endianness.")
// 	}
// }
