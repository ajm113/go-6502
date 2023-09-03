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

	if *b != Ins_LDA_IM {
		t.Errorf("bytes is not %x given: %x", Ins_LDA_IM, b)
		t.Fail()
	}

	if cycle != 0 {
		t.Errorf("cycle is not 0 given: %d", cycle)
		t.Fail()
	}
}

func TestNextBytes(t *testing.T) {
	var (
		pc    Byte   = CPUInitProgramCounter
		cycle uint32 = 2
	)

	mem := &Mem{}
	mem.Initialize()

	mem.Data[CPUInitProgramCounter] = Ins_LDA_IM
	mem.Data[CPUInitProgramCounter+1] = 0x84

	cmd := mem.NextByte(&pc, &cycle)

	if pc != CPUInitProgramCounter+1 {
		t.Errorf("expected program counter to increment + 1: %x", pc)
		t.Fail()
	}

	if cycle != 1 {
		t.Errorf("expected cycles to count down 1: %x", cycle)
		t.Fail()
	}

	if cmd != Ins_LDA_IM {
		t.Errorf("bytes is not %x given: %x", Ins_LDA_IM, cmd)
		t.Fail()
	}

	val := mem.NextByte(&pc, &cycle)
	if val != 0x84 {
		t.Errorf("bytes is not %x given: %x", 0x84, val)
		t.Fail()
	}

	if cycle != 0 {
		t.Errorf("cycle is not 0 given: %d", cycle)
		t.Fail()
	}
}

func TestWriteWord(t *testing.T) {
	var (
		pc        Byte   = CPUInitProgramCounter
		cycle     uint32 = 2
		testWrite Byte   = Ins_LDA_ZP
	)

	mem := &Mem{}
	mem.Initialize()

	mem.Data[CPUInitProgramCounter] = Ins_LDA_IM
	mem.Data[CPUInitProgramCounter+1] = 0x84

	mem.WriteWord(testWrite, &pc, &cycle)

	if pc != CPUInitProgramCounter+2 {
		t.Errorf("expected program counter to increment + 2: %x", pc)
		t.Fail()
	}

	if cycle != 0 {
		t.Errorf("expected cycles to count down 1: %x", cycle)
		t.Fail()
	}

	if mem.Data[CPUInitProgramCounter] != Ins_LDA_ZP {
		t.Errorf("expected memory address of %x: %x", CPUInitProgramCounter, mem.Data[CPUInitProgramCounter])
		t.Fail()
	}

	if mem.Data[CPUInitProgramCounter+1] != 0x84 {
		t.Errorf("expected memory address of %x: %x", CPUInitProgramCounter+1, mem.Data[CPUInitProgramCounter+1])
		t.Fail()
	}

}
