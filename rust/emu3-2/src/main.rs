mod emulator;
mod instruction;

use std::env;
use std::fs::File;
use std::io::Read;
use emulator::*;
use std::io::BufReader;
use instruction::*;

fn load_machine_code(emu: &mut Emulator, index: i64, b: Vec<u8>) {
    for i in 0..b.len() {
        emu.memory[(index as usize)+i]=b[i];
    }
}

fn dump_registers(emu: &Emulator) {
    for i in 0..(REGISTERS::REGISTERS_COUNT as usize) {
        println!("{:x}", emu.registers[i]);
    }
}

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() > 2 {
        println!("px86 filename");
        return;
    }
    let path = &args[1];
    let mut file = File::open(path)
        .expect("can't open file\n");
    let mut reader = BufReader::new(file);
    let mut buffer = Vec::new();
    reader.read_to_end(&mut buffer);
    let size = 1<<20;
    let mut emu = emulator::create_emu(size,0x7c00,0x7c00);
    load_machine_code(&mut emu, 0x7c00, buffer);
    let mut instructions: Instruction =  [None; INSTRUCTIONS_COUNT];
    init_instructions(&mut instructions);
    while emu.eip < (size as u32) {
        let mut code = emu.get_code8(0);
        println!("code:{:x} eip:{:x}", code, emu.eip);
        match instructions[code as usize] {
            Some(inst) => inst(&mut emu),
            None => {
                println!("Not implemented");
            }
        }
        if emu.eip == 0x00 {
            println!("end of program");
            break;
        }
    }
    dump_registers(&emu);
}
