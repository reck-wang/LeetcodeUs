package main

func totalNQueens(n int) int {
	cols, pies, nais := make(map[int]bool), make(map[int]bool), make(map[int]bool)
	var ans int

	var dfs func(int)
	dfs = func(row int) {
		if row >= n {
			ans++
		}

		for col := 0; col < n; col++ {
			var pie, nai = row + col, row - col
			if cols[col] || pies[pie] || nais[nai] {
				continue
			}

			cols[col] = true
			pies[pie] = true
			nais[nai] = true

			dfs(row + 1)

			cols[col] = false
			pies[pie] = false
			nais[nai] = false
		}
	}

	dfs(0)
	return ans
}
