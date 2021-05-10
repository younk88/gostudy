package main

import (
	"fmt"
	"os"

	"github.com/cryptix/wav"
)

func main() {
	fName := os.Args[1]
	info, _ := os.Stat(fName)
	f, _ := os.Open(fName)

	wr, err := wav.NewReader(f, info.Size())
	fmt.Printf("info: %v, err: %v\n", wr, err)
}