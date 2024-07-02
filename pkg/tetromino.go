package pkg

import "fmt"

// Tetromino represents the shape and position of a tetromino.
type Tetromino struct {
	positions []Position
}

// Position holds coordinates for tetromino positions.
type Position struct {
	x, y int
}

// parseTetrominos reads tetromino shapes from input data.
func ParseTetrominos(lines []string) ([]Tetromino, error) {
	var tetrominos []Tetromino
	for i := 0; i < len(lines); i += 5 {
		var positions []Position
		for j := 0; j < 4; j++ {
			line := lines[i+j]
			for k, char := range line {
				if char == '#' {
					positions = append(positions, Position{x: j, y: k})
				}
			}
		}
		if len(positions) != 4 {
			return nil, fmt.Errorf("invalid tetromino at index %d", i)
		}
		tetrominos = append(tetrominos, Tetromino{positions: positions})
	}
	return tetrominos, nil
}
