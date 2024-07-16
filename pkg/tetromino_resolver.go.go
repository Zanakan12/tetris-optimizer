package pkg

type Position struct {
	X int
	Y int
}

type BoardSave struct {
	board    [][]string
	positions [][]Position
}

// Creates a deep copy of the board.
func copyBoard(board [][]string) [][]string {
	newBoard := make([][]string, len(board))
	for i, row := range board {
		newRow := make([]string, len(row))
		copy(newRow, row)
		newBoard[i] = newRow
	}
	return newBoard
}

// Checks if the tetromino can be placed on the board at the given position.
func canPlaceTetromino(tetromino []string, board [][]string, y int, x int) bool {
	for dy, row := range tetromino {
		for dx, char := range row {
			if char == '#' {
				if y+dy >= len(board) || x+dx >= len(board[0]) || board[y+dy][x+dx] != "." {
					return false
				}
			}
		}
	}
	return true
}

// Places the tetromino on the board at the given position.
func placeTetromino(tetromino []string, board [][]string, y int, x int, tetrominoIndex int) [][]string {
	for dy, row := range tetromino {
		for dx, char := range row {
			if char == '#' {
				startOfAlphabet := 65 + tetrominoIndex
				if startOfAlphabet > 90 {
					board[y+dy][x+dx] = "X"
				} else {
					board[y+dy][x+dx] = string(rune(startOfAlphabet))
				}
			}
		}
	}
	return board
}

// Main function to resolve the placement of tetrominos on the board.
func Resolve(tetrominos [][]string, board [][]string) [][]string {
	boardSaves := []BoardSave{} // To save the state of the board at each step

	for i := 0; i < len(tetrominos); i++ {
		hasBeenPlaced := false
		isAtLastPosition := true
		hasDuplicatePosition := false
		canPlaceTetrominoFlag := false
		tetrominoPositions := [][]Position{}

		for x := 0; x < len(board[0]) && !hasBeenPlaced; x++ {
			for y := 0; y < len(board) && !hasBeenPlaced; y++ {
				if canPlaceTetromino(tetrominos[i], board, y, x) {
					canPlaceTetrominoFlag = true
					hasDuplicate := false

					// Check for duplicate positions
					if len(boardSaves) > 0 && len(boardSaves) == i+1 {
						for _, p := range boardSaves[len(boardSaves)-1].positions[len(boardSaves[len(boardSaves)-1].positions)-1] {
							if p.X == x && p.Y == y {
								hasDuplicate = true
								hasDuplicatePosition = true
								break
							}
						}
					}

					if len(boardSaves) > 0 {
						tetrominoPositions = boardSaves[len(boardSaves)-1].positions
					}

					if hasDuplicate {
						continue
					}

					if !hasDuplicatePosition {
						tetrominoPositions = append(tetrominoPositions, []Position{})
					}

					hasBeenPlaced = true

					localBoard := copyBoard(board)
					tetrominoPositions[len(tetrominoPositions)-1] = append(tetrominoPositions[len(tetrominoPositions)-1], Position{x, y})

					if !hasDuplicatePosition {
						isAtLastPosition = false
						boardSaves = append(boardSaves, BoardSave{localBoard, tetrominoPositions})
					}

					board = placeTetromino(tetrominos[i], board, y, x, i)
				}
			}
		}

		if !canPlaceTetrominoFlag {
			isAtLastPosition = false
		}

		if !hasBeenPlaced {
			i-- // Go back to the previous tetromino

			if len(boardSaves) <= 1 {
				// If no space is left, create a larger board
				i = -1
				boardSize := len(board) + 1
				board = make([][]string, boardSize)
				for i := range board {
					board[i] = make([]string, boardSize)
					for j := range board[i] {
						board[i][j] = "."
					}
				}
				boardSaves = []BoardSave{}
			}

			if len(boardSaves) > 1 {
				i-- // Go back to the previous tetromino

				if isAtLastPosition {
					boardSaves = boardSaves[:len(boardSaves)-1]
					board = copyBoard(boardSaves[len(boardSaves)-1].board)
				} else {
					board = copyBoard(boardSaves[len(boardSaves)-1].board)
				}
			}
		}
	}

	return board
}
