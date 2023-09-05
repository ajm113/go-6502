package main

import "testing"

func TestLDAIM(t *testing.T) {
	cpu := &CPU{}
	mem := &Mem{}
	cpu.Reset(mem)

	// simple inline program
	mem.Data[CPUInitProgramCounter] = Ins_LDA_IM
	mem.Data[CPUInitProgramCounter+1] = 0x85

	err := cpu.Execute(2, mem)

	// test execution
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		t.Fail()
	}

	if cpu.Cycles != 2 {
		t.Errorf("expected cycle ran of 2, but got: %d", cpu.Cycles)
		t.Fail()
	}

	// test registers...
	if cpu.A != 0x85 {
		t.Errorf("expected A 0x85 but got: %x", cpu.A)
		t.Fail()
	}

	if cpu.X != 0 {
		t.Errorf("expected X 0 but got: %x", cpu.X)
		t.Fail()
	}

	if cpu.Y != 0 {
		t.Errorf("expected Y 0 but got: %x", cpu.Y)
		t.Fail()
	}

	// Test LDA flags...
	if cpu.C != 0 {
		t.Error("expected C flag 0")
		t.Fail()
	}

	if cpu.Z != 0 {
		t.Error("expected Z flag 0")
		t.Fail()
	}

	if cpu.I != 0 {
		t.Error("expected I flag 0")
		t.Fail()
	}

	if cpu.D != 0 {
		t.Error("expected D flag 0")
		t.Fail()
	}

	if cpu.B != 0 {
		t.Error("expected D flag 0")
		t.Fail()
	}

	if cpu.V != 0 {
		t.Error("expected V flag 0")
		t.Fail()
	}

	if cpu.N != 1 {
		t.Error("expected N flag 1")
		t.Fail()
	}
}

func TestLDAIMSetAZero(t *testing.T) {
	cpu := &CPU{}
	mem := &Mem{}
	cpu.Reset(mem)

	// simple inline program
	mem.Data[CPUInitProgramCounter] = Ins_LDA_IM
	mem.Data[CPUInitProgramCounter+1] = 0

	err := cpu.Execute(2, mem)

	// test execution
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		t.Fail()
	}

	if cpu.Cycles != 2 {
		t.Errorf("expected cycle ran of 2, but got: %d", cpu.Cycles)
		t.Fail()
	}

	// test registers...
	if cpu.A != 0 {
		t.Errorf("expected A 0x85 but got: %x", cpu.A)
		t.Fail()
	}

	if cpu.X != 0 {
		t.Errorf("expected X 0 but got: %x", cpu.X)
		t.Fail()
	}

	if cpu.Y != 0 {
		t.Errorf("expected Y 0 but got: %x", cpu.Y)
		t.Fail()
	}

	// Test LDA flags...
	if cpu.C != 0 {
		t.Error("expected C flag 0")
		t.Fail()
	}

	if cpu.Z != 1 {
		t.Error("expected Z flag 1")
		t.Fail()
	}

	if cpu.I != 0 {
		t.Error("expected I flag 0")
		t.Fail()
	}

	if cpu.D != 0 {
		t.Error("expected D flag 0")
		t.Fail()
	}

	if cpu.B != 0 {
		t.Error("expected D flag 0")
		t.Fail()
	}

	if cpu.V != 0 {
		t.Error("expected V flag 0")
		t.Fail()
	}

	if cpu.N != 0 {
		t.Error("expected N flag 0")
		t.Fail()
	}
}

func TestLDAZP(t *testing.T) {
	cpu := &CPU{}
	mem := &Mem{}
	cpu.Reset(mem)

	// simple inline program
	mem.Data[CPUInitProgramCounter] = Ins_LDA_ZP
	mem.Data[CPUInitProgramCounter+1] = CPUInitProgramCounter - 2
	mem.Data[CPUInitProgramCounter-2] = 0x85

	err := cpu.Execute(3, mem)

	// test execution
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		t.Fail()
	}

	if cpu.Cycles != 3 {
		t.Errorf("expected cycle ran of 3, but got: %d", cpu.Cycles)
		t.Fail()
	}

	// test registers...
	if cpu.A != 0x85 {
		t.Errorf("expected A 0x85 but got: %x", cpu.A)
		t.Fail()
	}

	if cpu.X != 0 {
		t.Errorf("expected X 0 but got: %x", cpu.X)
		t.Fail()
	}

	if cpu.Y != 0 {
		t.Errorf("expected Y 0 but got: %x", cpu.Y)
		t.Fail()
	}

	// Test LDA flags...
	if cpu.C != 0 {
		t.Error("expected C flag 0")
		t.Fail()
	}

	if cpu.Z != 0 {
		t.Error("expected Z flag 0")
		t.Fail()
	}

	if cpu.I != 0 {
		t.Error("expected I flag 0")
		t.Fail()
	}

	if cpu.D != 0 {
		t.Error("expected D flag 0")
		t.Fail()
	}

	if cpu.B != 0 {
		t.Error("expected D flag 0")
		t.Fail()
	}

	if cpu.V != 0 {
		t.Error("expected V flag 0")
		t.Fail()
	}

	if cpu.N != 1 {
		t.Error("expected N flag 1")
		t.Fail()
	}
}
