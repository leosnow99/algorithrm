package main

func isValidSudokuDemo(board [][]byte) bool {

	xLength := len(board)
	yLength := 0
	if xLength > 0 {
		yLength = len(board[0])
	}

	for x := 0; x < xLength; x++ {
		yMap := map[byte]bool{}
		for y := 0; y < yLength; y++ {
			_, ok := yMap[board[x][y]]
			if ok {
				return false
			}

			if board[x][y] != '.' {
				yMap[board[x][y]] = true
			}

			if x%3 == 0 && y%3 == 0 {
				if !validCube(board, x, y, xLength, yLength) {
					return false
				}
			}
		}
	}

	for y := 0; y < yLength; y++ {
		xMap := map[byte]bool{}
		for x := 0; x < xLength; x++ {
			_, ok := xMap[board[x][y]]
			if ok {
				return false
			}

			if board[x][y] != '.' {
				xMap[board[x][y]] = true
			}
		}
	}

	return true
}

func validCube(board [][]byte, x, y, xLength, yLength int) bool {
	cubeMap := map[byte]bool{}

	if x+3 > xLength || y+3 > yLength {
		return false
	}

	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			_, ok := cubeMap[board[i][j]]
			if ok {
				return false
			}

			if board[i][j] != '.' {
				cubeMap[board[i][j]] = true
			}
		}
	}

	return true
}
