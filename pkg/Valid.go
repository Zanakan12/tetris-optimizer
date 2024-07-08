package pkg

// IsTetrominoValid checks if the given tetrominos are valid according to the game rules.
func IsTetrominoValid(tetrominos [][]string) bool {
	for _, tetromino := range tetrominos {
		totalConnections := 0
		totalHashtags := 0
		
		for y, row := range tetromino {
			for x, char := range row {
				if char != '#' && char != '.' {
					return false // Invalid character found
				} else if char == '#' {
					totalHashtags++
					connections := 0
					// Check connections in all four directions
					if y > 0 && tetromino[y-1][x] == '#' {
						connections++
					}
					if y < len(tetromino)-1 && tetromino[y+1][x] == '#' {
						connections++
					}
					if x > 0 && tetromino[y][x-1] == '#' {
						connections++
					}
					if x < len(row)-1 && tetromino[y][x+1] == '#' {
						connections++
					}
					
					if connections == 0 {
						return false // A '#' block has no connections, invalid tetromino
					}
					totalConnections += connections
				}
			}
		}

		// A valid tetromino has exactly 4 '#' blocks and at least 6 connections
		if totalConnections < 6 || totalHashtags != 4 {
			return false
		}
	}
	return true
}
