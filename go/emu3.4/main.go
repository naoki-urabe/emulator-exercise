package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var registers_name = []string{"EAX", "ECX", "EDX", "EBX", "ESP", "EBP", "ESI", "EDI"}

const MEMORY_SIZE = 1024 * 1024

func copyByIndex(dst, src []uint8, index int) error {
	for i := 0; i < len(src); i++ {
		if index+len(src) > len(dst) {
			return fmt.Errorf("out of index")
		}
		dst[index+i] = src[i]
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: px86 filename")
	}
	emu := CreateEmu(MEMORY_SIZE, 0x7c00, 0x7c00)
	binary, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("%sファイルが開けません\n", os.Args[1])
	}
	if err := copyByIndex(emu.memory, binary, 0x7c00); err != nil {
		log.Fatalln(err)
	}
	InitInstructions()
	for emu.eip < MEMORY_SIZE {
		code := emu.get_code8(0)
		fmt.Printf("EIP = %X, Code = %02X\n", emu.eip, code)
		if instructions[code] == nil {
			fmt.Printf("\n\nNot Implemented: %X\n", code)
			break
		}
		instructions[code](emu)
		if emu.eip == 0x00 {
			fmt.Printf("\n\nend of program.\n\n")
			break
		}
	}
	dump_registers(emu)
}
