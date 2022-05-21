package main

import (
	"fmt"
	"log"
	"os"
)

const MEMORY_SIZE = 2 >> 20

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: px86 filename\n")
	}
	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(f))
	emu := createEmu(MEMORY_SIZE, 0x0000, 0x7c00)

	fmt.Println(emu)
}
