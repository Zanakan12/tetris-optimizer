package main

import (
	"fmt"
	"os"
	"tetris-optimizer/pkg"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <path_file>")
		os.Exit(1)
	}

	lines, err := pkg.ReadInput(os.Args[1])
	if err != nil {
		fmt.Println("Failed to read input:", err)
		os.Exit(1)
	}

	tetrominos, err := pkg.ParseTetrominos(lines)
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(0)
	}

	gridSize := 10 // Define grid size based on input or requirements
	grid := pkg.CreateGrid(gridSize)

	if grid.CanPlace(0, 0, &tetrominos[0]) {
		grid.Place(0, 0, &tetrominos[0], 'A')
		grid.Print()
		grid.Remove(0, 0, &tetrominos[0])
	}
}
