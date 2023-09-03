package main

func main() {
	cpu := &CPU{}
	mem := &Mem{}
	cpu.Reset(mem)

	// simple inline program
	mem.Data[0xFFFC] = Ins_LDA_JSR
	mem.Data[0xFFFD] = 0x42
	mem.Data[0xFFFE] = 0x42
	mem.Data[0x42] = Ins_LDA_IM
	mem.Data[0x43] = 0x84

	cpu.Execute(9, mem)
	cpu.Print()
	return
}
