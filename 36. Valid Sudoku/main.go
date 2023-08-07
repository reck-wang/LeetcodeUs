package main

func isValidSudoku(board [][]byte) bool {
	// 创建这三个数组用于计数
	rows := [9][9]bool{}
	cols := [9][9]bool{}
	blocks := [3][3][9]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			digit := board[i][j] - '0' - 1
			if rows[i][digit] || cols[j][digit] || blocks[i/3][j/3][digit] {
				return false
			}

			rows[i][digit] = true
			cols[j][digit] = true
			blocks[i/3][j/3][digit] = true
		}
	}

	return true
}
