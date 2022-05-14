package main

import (
	"fmt"
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

type Emulator struct {
	registers [REGISTERS_COUNT]uint32
	eflags    uint32
	memory    []uint8
	eip       uint32
}

func CreateEmu(size uint, eip uint32, esp uint32) *Emulator {
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

func (emu *Emulator) get_sign_code32(index int) int32 {
	return int32(emu.get_code32(index))
}
