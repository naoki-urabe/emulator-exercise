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
	memory    []uint8
	eip       uint32
}

func createEmu(size uint, eip uint32, esp uint32) *Emulator {
	emu := &Emulator{}
	emu.memory = make([]uint8, size)
	emu.eip = eip
	emu.registers[ESP] = esp
	return emu
}

func (emu *Emulator) getCode8(index int) uint8 {
	return emu.memory[int(emu.eip)+index]
}
