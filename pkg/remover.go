package pkg

// Remove removes rows and columns that do not contain any '#' from each tetromino.
func Remove(tetrominos [][]string) [][]string {
	cutTetrominos := [][]string{}

	// Helper function to check if a line contains '#'
	containsHashtag := func(line string) bool {
		for _, char := range line {
			if char == '#' {
				return true
			}
		}
		return false
	}

	// Helper function to remove a character at a specific index from a string
	removeCharAtIndex := func(input string, index int) string {
		return input[:index] + input[index+1:]
	}

	// Remove empty rows from each tetromino
	for _, tetromino := range tetrominos {
		var newTetromino []string
		for _, line := range tetromino {
			if containsHashtag(line) {
				newTetromino = append(newTetromino, line)
			}
		}
		cutTetrominos = append(cutTetrominos, newTetromino)
	}

	// Remove empty columns from each tetromino
	for _, tetromino := range cutTetrominos {
		if len(tetromino) == 0 {
			continue
		}
		columnCount := len(tetromino[0])
		for col := columnCount - 1; col >= 0; col-- {
			columnHasHashtag := false
			for row := range tetromino {
				if tetromino[row][col] == '#' {
					columnHasHashtag = true
					break
				}
			}
			if !columnHasHashtag {
				for row := range tetromino {
					tetromino[row] = removeCharAtIndex(tetromino[row], col)
				}
			}
		}
	}

	return cutTetrominos
}
