use crate::emulator::Emulator;

// mod emulator;

pub const INSTRUCTIONS_COUNT: usize = 256;
pub type Instruction = [Option<fn(&mut Emulator)>; INSTRUCTIONS_COUNT];

pub fn mov_r32_imm32(emu: &mut Emulator) {
    let reg = emu.get_code8(0) - 0xB8;
    let value = emu.get_code32(1);
    emu.set_register32(reg as u32, value);
    emu.eip += 5;
}

pub fn short_jump(emu: &mut Emulator) {
    let value = emu.get_sign_code8(1);
    emu.eip = ((emu.eip as i8) + (value + 2)) as u32;
}

pub fn init_instructions(instructions: &mut Instruction) {
    for i in 0..8 {
        instructions[0xB8 + i] = Some(mov_r32_imm32);
    }
    instructions[0xEB] = Some(short_jump);
}