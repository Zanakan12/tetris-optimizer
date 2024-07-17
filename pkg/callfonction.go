package pkg

import (
	"fmt"
	"math"
)

func Main(tetrominos [][]string) {

	board := newBoard(int(math.Ceil(math.Sqrt(float64(len(tetrominos) * 4)))))
	trimmedTetrominos := Remove(tetrominos)
	resolvedBoard := Solver(trimmedTetrominos, board)

	printBoard(resolvedBoard)
}

// newBoard creates a new board of the given size filled with ".".
func newBoard(size int) [][]string {
	board := make([][]string, size)
	for i := range board {
		board[i] = make([]string, size)
		for j := range board[i] {
			board[i][j] = "."
		}
	}
	return board
}

// printBoard prints the board with symbols representing different tetrominos.
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
		"I": "❎",
		"J": "🆗",
		"K": "🆘",
		"L": "🃏",
		".": "⬛",
	}

	for _, row := range board {
		for _, char := range row {
			if symbol, exists := symbols[char]; exists {
				fmt.Print(symbol)
			} else {
				fmt.Print(char)
			}
		}
		fmt.Println()
	}
}
