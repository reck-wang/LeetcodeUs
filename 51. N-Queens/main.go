package main

import "strings"

func solveNQueens(n int) [][]string {
	cols, pie, nai := map[int]struct{}{}, map[int]struct{}{}, map[int]struct{}{}
	var dfs func(int, int, []int)
	var results = [][]int{}

	dfs = func(n int, curRow int, temp []int) {
		// 递归终止条件
		if curRow >= n {
			newArr := make([]int, n)
			copy(newArr, temp)
			results = append(results, newArr)
			return
		}

		// 遍历每一列
		for col := 0; col < n; col++ {
			// 剪枝
			_, inCols := cols[col]
			_, inPie := pie[curRow+col]
			_, inNai := nai[curRow-col]

			if inCols || inPie || inNai {
				continue
			}

			// 记录当前位置的列、撇、捺
			cols[col] = struct{}{}
			pie[curRow+col] = struct{}{}
			nai[curRow-col] = struct{}{}

			dfs(n, curRow+1, append(temp, col))

			// 回溯记录的列、撇、捺
			delete(cols, col)
			delete(pie, curRow+col)
			delete(nai, curRow-col)
		}
	}

	dfs(n, 0, []int{})

	var ans = make([][]string, len(results))
	for i, res := range results {
		line := make([]string, n)
		for j, q := range res {
			row := strings.Repeat(".", q)
			row += "Q"
			row += strings.Repeat(".", n-q-1)
			line[j] = row
		}
		ans[i] = line
	}

	return ans
}
