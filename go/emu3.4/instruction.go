package main

type instruction_func func(emu *Emulator)

var instructions [256]instruction_func

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

func near_jump(emu *Emulator) {
	var diff int32
	diff = emu.get_sign_code32(1)
	emu.eip += uint32(diff + 5)
}

func InitInstructions() {
	for i := 0; i < 8; i++ {
		instructions[0xB8+i] = mov_r32_imm32
	}
	instructions[0xE9] = near_jump
	instructions[0xEB] = short_jump
}
