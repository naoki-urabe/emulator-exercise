package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// filenameを引数から読み出す
	args := os.Args
	// filanameが引数に設定されていなかったらエラーログ出力
	if len(args) != 2 {
		log.Fatal("not found filename")
	}
	// filename設定
	filename := args[1]
	bin, err := os.ReadFile(filename)
	// error handling: os.ReadFileでerrがnil出なかった場合、エラーログを表示する
	if err != nil {
		log.Fatal(err)
	}

	// 読み込んだ[]byteをstringでキャストし表示する
	fmt.Println(string(bin))
}
