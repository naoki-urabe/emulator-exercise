package main

const (
	EAX = iota
	ECX
	EDX
	EBX
	ESP
	EBP
	ESI
	EDI
	REGISTERS_COUNT
)

type Emulator struct {
	registers [REGISTERS_COUNT]uint32
	eflags    uint32
	memory    []uint32
	eip       uint32
}

func createEmu(size uint, eip uint32, esp uint32) *Emulator {
	emu := &Emulator{}
	emu.eip = eip
	emu.registers[ESP] = esp
	return emu
}
