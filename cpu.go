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
	}
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
	c.PC = 0xFFFC
	c.SP = 0x0100

	// reset registers
	c.A = 0
	c.X = 0
	c.Y = 0

	mem.Initalize()
}

func (c *CPU) Print() {
	fmt.Println("======================================================")
	fmt.Printf("cpu status: pc: %x sp: %x\n", c.PC, c.SP)
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

func (c *CPU) Execute(cycles uint32, mem *Mem) {
	for cycles > 0 {
		inst := mem.NextByte(&c.PC, &cycles)

		switch inst {
		case Ins_LDA_IM:
			value := mem.NextByte(&c.PC, &cycles)
			c.A = value
			c.ldaSetStatus()
		case Ins_LDA_ZP:
			zeroPageAddress := mem.NextByte(&c.PC, &cycles)

			c.A = mem.ReadByte(zeroPageAddress, &cycles)

			c.ldaSetStatus()
		case Ins_LDA_ZPX:
			zeroPageAddress := mem.NextByte(&c.PC, &cycles)

			zeroPageAddress += c.X
			cycles--
			c.A = mem.ReadByte(zeroPageAddress, &cycles)

			c.ldaSetStatus()
		case Ins_LDA_JSR:
			jumpAddr := mem.NextWord(&c.PC, &cycles)
			mem.WriteWord(Word(c.PC-1), &c.SP, &cycles)
			c.SP++
			c.PC = Byte(jumpAddr)
			cycles -= 2
		default:
			// revert the program counter and cycle # for debugging purposes and return em back after continuing..
			c.PC--
			cycles++
			fmt.Printf("!!! instruction not handled: %x pc: %x cycle: %x !!!\n", inst, c.PC, cycles)
			cycles = 0
		}

	}
}
