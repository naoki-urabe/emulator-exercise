pub enum REGISTERS {
    EAX,
    ECX,
    EDX,
    EBX,
    ESP,
    EBP,
    ESI,
    EDI,
    REGISTERS_COUNT,
}

pub struct Emulator {
    pub registers: [u32; REGISTERS::REGISTERS_COUNT as usize],
    pub eflags:     u32,
    pub memory:     Vec<u8>,
    pub eip:        u32,
}

pub fn create_emu(size: usize, eip: u32, esp: u32) -> Emulator {
    let mut emu = Emulator {
        registers: Default::default(),
        eflags: 0,
        memory: vec![0;size],
        eip:    eip,
    };
    emu.registers[REGISTERS::ESP as usize] = esp;
    return emu
}

impl Emulator {
    pub fn get_code8(&mut self, index:usize) -> u8 {
        return self.memory[(self.eip as usize) + index]
    }
    pub fn get_sign_code8(&mut self, index:usize) -> i8 {
        return (self.memory[(self.eip as usize) + index] as i8);
    }
    pub fn get_code32(&mut self, index:usize) -> u32 {
        let mut value: u32 = 0;
        for i in 0..4 {
            value |= (self.get_code8(index+ i)as u32) << 8*i;
        }
        return value;
    }
    pub fn set_register32(&mut self, reg: u32, value: u32) {
        self.registers[reg as usize] = value;
    }
}