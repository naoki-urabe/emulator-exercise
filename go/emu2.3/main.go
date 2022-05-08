package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	EAX             = iota
	ECX             = iota
	EDX             = iota
	EBX             = iota
	ESP             = iota
	EBP             = iota
	ESI             = iota
	EDI             = iota
	REGISTERS_COUNT = iota
)

var registers_name = []string{"EAX", "ECX", "EDX", "EBX", "ESP", "EBP", "ESI", "EDI"}

const MEMORY_SIZE = 1024 * 1024

type Emulator struct {
	registers [REGISTERS_COUNT]uint32
	eflags    uint32
	memory    []uint8
	eip       uint32
}

func create_emu(size uint, eip uint32, esp uint32) *Emulator {
	emu := &Emulator{}
	emu.memory = make([]uint8, size)
	emu.eip = eip
	emu.registers[ESP] = esp
	return emu
}

func dump_registers(emu *Emulator) {
	for i := 0; i < REGISTERS_COUNT; i++ {
		fmt.Printf("%s = %08x\n", registers_name[i], emu.registers[i])
	}
}

func (emu *Emulator) get_code8(index int) uint8 {
	return emu.memory[int(emu.eip)+index]
}

func (emu *Emulator) get_sign_code8(index int) int8 {
	return int8(emu.memory[int(emu.eip)+index])
}

func (emu *Emulator) get_code32(index int) uint32 {
	var ret uint32
	for i := 0; i < 4; i++ {
		ret |= uint32(emu.get_code8(index+i)) << (i * 8)
	}
	return ret
}

func mov_r32_imm32(emu *Emulator) {
	reg := emu.get_code8(0) - 0xB8
	value := emu.get_code32(1)
	emu.registers[reg] = value
	emu.eip += 5
}

func short_jump(emu *Emulator) {
	diff := emu.get_sign_code8(1)
	emu.eip += uint32(diff + 2)
}

type instruction_func func(emu *Emulator)

var instructions [256]instruction_func

func init_instructions() {
	for i := 0; i < 8; i++ {
		instructions[0xB8+i] = mov_r32_imm32
	}
	instructions[0xEB] = short_jump
}

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
	emu := create_emu(MEMORY_SIZE, 0x0000, 0x7c00)
	binary, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("%sファイルが開けません\n", os.Args[1])
	}
	if err := copyByIndex(emu.memory, binary, 0x0000); err != nil {
		log.Fatalln(err)
	}
	init_instructions()
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
