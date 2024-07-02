package pkg

import "fmt"

// Grid represents the game grid.
type Grid struct {
	size     int
	occupied [][]bool
	letter   [][]rune
}

// createGrid initializes an empty game grid.
func CreateGrid(size int) *Grid {
	grid := &Grid{
		size:     size,
		occupied: make([][]bool, size),
		letter:   make([][]rune, size),
	}

	for i := range grid.occupied {
		grid.occupied[i] = make([]bool, size)
		grid.letter[i] = make([]rune, size)
		for j := range grid.letter[i] {
			grid.letter[i][j] = '.'
		}
	}
	return grid
}

// canPlace checks if a tetromino can be placed on the grid.
func (g *Grid) CanPlace(x, y int, tetromino *Tetromino) bool {
	for _, pos := range tetromino.positions {
		px, py := x+pos.x, y+pos.y
		if px < 0 || py < 0 || px >= g.size || py >= g.size || g.occupied[px][py] {
			return false
		}
	}
	return true
}

// place puts a tetromino on the grid.
func (g *Grid) Place(x, y int, tetromino *Tetromino, letter rune) {
	for _, pos := range tetromino.positions {
		px, py := x+pos.x, y+pos.y
		g.occupied[px][py] = true
		g.letter[px][py] = letter
	}
}

// remove lifts a tetromino off the grid.
func (g *Grid) Remove(x, y int, tetromino *Tetromino) {
	for _, pos := range tetromino.positions {
		px, py := x+pos.x, y+pos.y
		g.occupied[px][py] = false
		g.letter[px][py] = '.'
	}
}

// print displays the current grid state.
func (g *Grid) Print() {
	for _, row := range g.letter {
		fmt.Println(string(row))
	}
}
