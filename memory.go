package main

const MaxMemory = 1024 * 64

type (
	Mem struct {
		Data [MaxMemory]Byte
	}
)

func (m *Mem) Initalize() {
	for i := 0; i < MaxMemory; i++ {
		m.Data[i] = 0
	}
}

func (m *Mem) NextByte(pc *Word, cycle *uint32) Byte {
	x := m.Data[*pc]
	*pc++
	*cycle--

	return x
}

func (m *Mem) NextWord(pc *Word, cycle *uint32) Word {
	x := Word(m.Data[*pc])
	*pc++

	x |= (Word(m.Data[*pc] << 8))

	*cycle -= 2

	// TODO: Handle big/little endian

	return x
}

func (m *Mem) ReadByte(address Word, cycle *uint32) Byte {
	x := m.Data[address]
	*cycle--

	return x
}
