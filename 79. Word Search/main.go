package main

var dirs = [4][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	visit := make([][]bool, m)
	for i := range visit {
		visit[i] = make([]bool, n)
	}

	var dfs func(int, int, int) bool
	dfs = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}

		if k == len(word)-1 {
			return true
		}

		visit[i][j] = true
		for _, dir := range dirs {
			nextI, nextJ := i+dir[0], j+dir[1]
			if nextI >= 0 && nextI < m && nextJ >= 0 && nextJ < n && !visit[nextI][nextJ] {
				if dfs(nextI, nextJ, k+1) {
					return true
				}
			}
		}

		visit[i][j] = false

		return false
	}

	for i, row := range board {
		for j := range row {
			if board[i][j] == word[0] && dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
