package main

import (
	"fmt"
	"os"
	"tetris-optimizer/pkg"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <examples/examples.txt>")
		return
	}

	inputFile := os.Args[1]
	tetrominos, err := pkg.Reader(inputFile)
	if err != nil {
		fmt.Println("ERROR reading file:", err)
		return
	}else if !pkg.Validator(tetrominos) {
		fmt.Println("ERROR: Invalid tetrominos")
		return
	}else{
		pkg.Main(tetrominos)
	}
}
