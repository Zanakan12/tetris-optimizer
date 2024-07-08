package lib

// CutUnusedLines removes rows and columns that do not contain any '#' from each tetromino.
func CutUnusedLines(tetrominos [][]string) [][]string {
	cutTetrominos := [][]string{}

	// Remove empty rows from each tetromino
	for _, tetromino := range tetrominos {
		var newTetromino []string
		for _, line := range tetromino {
			containsHashtag := false
			for _, char := range line {
				if char == '#' {
					containsHashtag = true
					break
				}
			}
			if containsHashtag {
				newTetromino = append(newTetromino, line)
			}
		}
		cutTetrominos = append(cutTetrominos, newTetromino)
	}

	// Remove empty columns from each tetromino
	for i := range cutTetrominos {
		if len(cutTetrominos[i]) == 0 {
			continue
		}
		columnCount := len(cutTetrominos[i][0])
		for col := columnCount - 1; col >= 0; col-- {
			containsHashtag := false
			for row := range cutTetrominos[i] {
				if cutTetrominos[i][row][col] == '#' {
					containsHashtag = true
					break
				}
			}
			if !containsHashtag {
				for row := range cutTetrominos[i] {
					cutTetrominos[i][row] = removeCharAtIndex(cutTetrominos[i][row], col)
				}
			}
		}
	}

	return cutTetrominos
}

// removeCharAtIndex removes the character at the specified index from a string.
func removeCharAtIndex(input string, index int) string {
	if index < 0 || index >= len(input) {
		return input
	}
	runes := []rune(input)
	runes = append(runes[:index], runes[index+1:]...)
	return string(runes)
}
