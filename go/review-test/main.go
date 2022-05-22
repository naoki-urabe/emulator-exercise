package main

import (
	"fmt"
	"log"
	"os"
)

const MEMORY_SIZE = 2 << 20

func loadMachineCode(emu *Emulator, index int, b []uint8) error {
	if index+len(b) > MEMORY_SIZE {
		return fmt.Errorf("out of index")
	}
	for i := 0; i < len(b); i++ {
		emu.memory[index+i] = b[i]
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: px86 filename\n")
	}
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(f))
	emu := createEmu(MEMORY_SIZE, 0x0000, 0x7c00)
	loadMachineCode(emu, 0x0000, f)
	for emu.eip < MEMORY_SIZE {
		code := emu.getCode8(0)
		fmt.Printf("%x\n", code)
		if emu.eip == 0x00 {
			fmt.Println("end of program")
		}
	}
}
