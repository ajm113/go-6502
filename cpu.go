package main

import "fmt"

type (
	Byte rune
	Word uint16

	CPU struct {
		PC Byte
		SP Byte // program counter, stack pointer

		A, X, Y Byte // registers

		C, // carry flag
		Z, // zero flag
		I, // interrupt disable
		D, // decimal mode
		B, // break command
		V, // overflow flag
		N Byte // negative flag

		Cycles uint32 // keeps track of cycles ran
	}
)

const (
	CPUInitProgramCounter Byte = 0xFFFC
	CPUInitStackPointer   Byte = 0x0100
)

func (c *CPU) Reset(mem *Mem) {
	// reset our flags
	c.C = 0
	c.Z = 0
	c.I = 0
	c.D = 0
	c.B = 0
	c.V = 0
	c.N = 0

	// reset program counter and stack pointer
	c.PC = CPUInitProgramCounter
	c.SP = CPUInitStackPointer

	// reset registers
	c.A = 0
	c.X = 0
	c.Y = 0

	// reset emulator specific stats.
	c.Cycles = 0

	mem.Initialize()
}

func (c *CPU) Print() {
	fmt.Println("======================================================")
	fmt.Printf("cpu status: pc: %x sp: %x cycles: %d\n", c.PC, c.SP, c.Cycles)
	fmt.Printf("cpu flags: c: %x z: %x i: %x d: %x b: %x v: %x n: %x\n", c.C, c.Z, c.I, c.D, c.B, c.V, c.N)
	fmt.Printf("cpu registers: a: %x x: %x y: %x\n", c.A, c.X, c.Y)
	fmt.Println("======================================================")
}

func (c *CPU) ldaSetStatus() {
	if c.A == 0 {
		c.Z = 1
	}
	if (c.A & 0b10000000) > 0 {
		c.N = 1
	}
}

func (c *CPU) Execute(cycles uint32, mem *Mem) error {
	startingCycles := cycles

	defer func() {
		c.Cycles = startingCycles - cycles
	}()

	for cycles > 0 {
		inst := mem.NextByte(&c.PC, &cycles)

		switch inst {
		case Ins_LDA_IM:
			value := mem.NextByte(&c.PC, &cycles)
			c.A = value
			c.ldaSetStatus()
		case Ins_LDA_ZP:
			zeroPageAddress := mem.NextByte(&c.PC, &cycles)

			b, err := mem.ReadByte(zeroPageAddress, &cycles)
			if err != nil {
				return err
			}
			c.A = *b

			c.ldaSetStatus()
		case Ins_LDA_ZPX:
			zeroPageAddress := mem.NextByte(&c.PC, &cycles)

			zeroPageAddress += c.X
			cycles--
			b, err := mem.ReadByte(zeroPageAddress, &cycles)
			if err != nil {
				return err
			}
			c.A = *b

			c.ldaSetStatus()
		case Ins_LDA_JSR:
			jumpAddr := mem.NextWord(&c.PC, &cycles)
			mem.WriteWord(c.PC-1, &c.SP, &cycles)
			c.SP++
			c.PC = Byte(jumpAddr)
			cycles -= 2
		default:
			// revert the program counter and cycle # for debugging purposes and return em back after continuing..
			c.PC--
			cycles++
			return fmt.Errorf("instruction not handled: %x pc: %x cycle: %x", inst, c.PC, cycles)
		}
	}

	return nil
}
