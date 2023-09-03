package main

import "testing"

func TestMemoryInit(t *testing.T) {
	m := &Mem{}
	m.Initialize()

	for _, b := range m.Data {
		if b != 0 {
			t.Fatal("all bytes from Initialize aren't 0")
		}
	}
}

func TestReadBytes(t *testing.T) {
	var (
		pc    Byte   = CPUInitProgramCounter
		cycle uint32 = 1
	)

	mem := &Mem{}
	mem.Initialize()

	mem.Data[CPUInitProgramCounter] = Ins_LDA_IM
	mem.Data[CPUInitProgramCounter+1] = 0x84

	b, err := mem.ReadByte(pc, &cycle)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		t.Fail()
	}

	if b != Ins_LDA_IM {
		t.Errorf("bytes is not %x given: %x", Ins_LDA_IM, b)
		t.Fail()
	}

	if cycle != 0 {
		t.Errorf("cycle is not 0 given: %d", cycle)
		t.Fail()
	}

}
