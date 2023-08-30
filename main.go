package main

func main() {
	cpu := &CPU{}
	mem := &Mem{}
	cpu.Reset(mem)

	// simple inline program
	mem.Data[0xFFFC] = Ins_LDA_ZP
	mem.Data[0xFFFD] = 0x42
	mem.Data[0x0042] = 0x84

	cpu.Execute(3, mem)
	cpu.Print()
	return
}
