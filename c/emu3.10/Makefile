TARGET = px86
OBJS = main.o emulator_function.o instruction.o modrm.o
Z_TOOLS = ../ z_tools
CC = gcc
FORMAT = clang-format 
CFLAGS += -Wall
.PHONY : all 

all : 
	make $(TARGET) 
%.o : %.c Makefile
	$(CC) $(CFLAGS) -c $<
$(TARGET): $(OBJS) Makefile
	$(CC) $(CFLAGS) -o $@ $(OBJS)

asm-build:
	nasm -f bin -o call-test.bin call-test.asm
