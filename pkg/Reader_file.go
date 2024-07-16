package pkg

import (
	"bufio"
	"fmt"
	"os"
)

// Reader reads the tetromino shapes from a file and returns them as a 2D slice of strings.
func Reader(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var tetrominos [][]string
	var currentTetromino []string

	for scanner.Scan() {
		line := scanner.Text()

		// Validate line length
		if len(line) != 4 && line != "" {
			return nil, fmt.Errorf("invalid line length in file")
		}

		if line == "" {
			// Add the current tetromino to the list if it's not empty
			if len(currentTetromino) > 0 {
				tetrominos = append(tetrominos, currentTetromino)
				currentTetromino = nil
			}
		} else {
			currentTetromino = append(currentTetromino, line)
		}
	}

	// Add the last tetromino if the file didn't end with an empty line
	if len(currentTetromino) > 0 {
		tetrominos = append(tetrominos, currentTetromino)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return tetrominos, nil
}
