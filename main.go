package main

import (
	"fmt"
	"math"
	"os"
	"tetris-optimizer/pkg"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		return
	}

	inputFile := os.Args[1]
	tetrominos, err := pkg.ReadInputFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	if !isValidTetrominos(tetrominos) {
		fmt.Println("ERROR")
		return
	}

	boardSize := calculateBoardSize(len(tetrominos))
	board := createEmptyBoard(boardSize)

	compressedTetrominos := pkg.CutUnusedLines(tetrominos)
	resolvedBoard := pkg.Resolve(compressedTetrominos, board)

	printBoard(resolvedBoard)
}

func isValidTetrominos(tetrominos [][]string) bool {
	return pkg.IsTetrominoValid(tetrominos) && len(tetrominos) > 0
}

func calculateBoardSize(numTetrominos int) int {
	return int(math.Ceil(math.Sqrt(float64(numTetrominos * 4))))
}

func createEmptyBoard(size int) [][]string {
	board := make([][]string, size)
	for i := range board {
		board[i] = make([]string, size)
		for j := range board[i] {
			board[i][j] = "."
		}
	}
	return board
}

func printBoard(board [][]string) {
	symbols := map[string]string{
		"A": "🟦",
		"B": "🟥",
		"C": "🟩",
		"D": "🟧",
		"E": "🟪",
		"F": "🟨",
		"G": "🟫",
		"H": "⬜",
		".": "⬛",
	}

	for _, row := range board {
		for _, char := range row {
			fmt.Print(symbols[char])
		}
		fmt.Println()
	}
}
