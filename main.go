package main

import "fmt"

func main() {
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

	if err == nil {
		fmt.Println("execution completed successfully")
	} else {
		fmt.Printf("!!! unexpected exception occurred: %s !!!", err)
	}

	cpu.Print()
	return
}
