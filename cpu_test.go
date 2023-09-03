package main

import (
	"testing"
)

func TestCPUReset(t *testing.T) {
	cpu := &CPU{}
	mem := &Mem{}
	cpu.Reset(mem)

	// maybe a little redudant, but let's tripple check to make sure
	for _, b := range mem.Data {
		if b != 0 {
			t.Fatal("all bytes from reset aren't 0")
			t.Fail()
		}
	}

	if cpu.PC != CPUInitProgramCounter {
		t.Fatalf("expected PC to be default %x but got %x", CPUInitProgramCounter, cpu.PC)
		t.Fail()
	}

	if cpu.SP != CPUInitStackPointer {
		t.Fatalf("expected SP to be default %x but got %x", CPUInitStackPointer, cpu.SP)
		t.Fail()
	}
}

// TestCPUExecute Tests a simple program. This test SHOULD NOT be used to test different opcodes.
func TestCPUExecute(t *testing.T) {
	cpu := &CPU{}
	mem := &Mem{}
	cpu.Reset(mem)

	// simple inline program
	mem.Data[0xFFFC] = Ins_LDA_JSR
	mem.Data[0xFFFD] = 0x42
	mem.Data[0xFFFE] = 0x42
	mem.Data[0x4242] = Ins_LDA_IM
	mem.Data[0x4243] = 0x84

	err := cpu.Execute(9, mem)

	if err != nil {
		t.Fatalf("unexpected error running Execute: %s", err)
		t.Fail()
	}

	if cpu.A != 0x84 {
		t.Fatalf("expected A register to be 0x84 but got: %x", cpu.A)
		t.Fail()
	}
}

func TestCPUExecuteInvalidOpsCode(t *testing.T) {
	cpu := &CPU{}
	mem := &Mem{}
	cpu.Reset(mem)

	// simple inline program
	mem.Data[0xFFFC] = 0x0A8

	err := cpu.Execute(1, mem)

	if err.Error() != "instruction not handled: a8 pc: fffc cycle: 1" {
		t.Fatalf("expected instruction not handled error: %s", err)
		t.Fail()
	}
}
