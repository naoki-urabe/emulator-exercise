TARGET = px86
OBJS = main.o emulator_function.o instruction.o modrm.c
Z_TOOLS = ../z_tools

CC = gcc
CFLAGS += -Wall

.PHONY: all
all :
	make $(TARGET)

%.o : %.c Makefile
	$(CC) $(CFLAGS) -c $<

$(TARGET) : $(OBJS) Makefile
	$(CC) -o $@ $(OBJS)
