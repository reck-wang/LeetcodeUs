package main

// method one: use hashtable
//func totalNQueens(n int) int {
//	cols, pies, nais := make(map[int]bool), make(map[int]bool), make(map[int]bool)
//	var ans int
//
//	var dfs func(int)
//	dfs = func(row int) {
//		if row >= n {
//			ans++
//		}
//
//		for col := 0; col < n; col++ {
//			var pie, nai = row + col, row - col
//			if cols[col] || pies[pie] || nais[nai] {
//				continue
//			}
//
//			cols[col] = true
//			pies[pie] = true
//			nais[nai] = true
//
//			dfs(row + 1)
//
//			cols[col] = false
//			pies[pie] = false
//			nais[nai] = false
//		}
//	}
//
//	dfs(0)
//	return ans
//}

// method two: use binary
func totalNQueens(n int) int {
	var ans int
	var dfs func(int, int, int, int)

	dfs = func(row int, cols int, pies int, nais int) {
		if row >= n {
			ans++
			return
		}

		bits := (^(cols | pies | nais)) & ((1 << n) - 1) // 得到当前所有空位取反，然后筛掉前面多余的1

		for bits > 0 {
			p := bits & (-bits) // 取最后一位1
			dfs(row+1, cols|p, (pies|p)<<1, (nais|p)>>1)
			bits &= bits - 1 // 移除最后一位1
		}
	}

	dfs(0, 0, 0, 0)
	return ans
}
