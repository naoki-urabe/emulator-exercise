TARGET = helloworld.bin
Z_TOOLS = ../z_tools

CC = gcc
LD = ld
AS = nasm
CFLAGS += -nostdlib -fno-asynchronous-unwind-tables \
	-I$(Z_TOOLS)/i386-elf-gcc/include -g -fno-stack-protector -m32 -march=i386 -fno-pie
LDFLAGS += --entry=func --oformat=binary -Ttext 0x7c00 -m elf_i386

.PHONY: all
all :
	make $(TARGET)

%.o : %.c Makefile
	$(CC) $(CFLAGS) -c $<

%.bin : %.o Makefile
	$(LD) $(LDFLAGS) -o $@ $<

%.bin : %.asm Makefile
	$(AS) -f bin -o $@ $<
