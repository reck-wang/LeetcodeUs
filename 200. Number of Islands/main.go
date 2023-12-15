package main

var dirs = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

// Method one: DFS
//func numIslands(grid [][]byte) int {
//	m, n := len(grid), len(grid[0])
//	var ans int
//
//	var dfs func(i, j int)
//	dfs = func(i, j int) {
//		grid[i][j] = '0'
//		for _, dir := range &dirs {
//			nextI, nextJ := i+dir[0], j+dir[1]
//			if nextI < m && nextI >= 0 && nextJ < n && nextJ >= 0 && grid[nextI][nextJ] == '1' {
//				dfs(nextI, nextJ)
//			}
//		}
//	}
//
//	for i := 0; i < m; i++ {
//		for j := 0; j < n; j++ {
//			if grid[i][j] == '1' {
//				ans++
//				dfs(i, j)
//			}
//		}
//	}
//
//	return ans
//}

// Method two: BFS
//func numIslands(grid [][]byte) int {
//	m, n := len(grid), len(grid[0])
//	var queue [][2]int
//	var ans int
//
//	var bfs func(i, j int)
//	bfs = func(i, j int) {
//		queue = append(queue, [2]int{i, j})
//		for len(queue) != 0 {
//			temp := queue[len(queue)-1]
//			queue = queue[:len(queue)-1]
//			grid[temp[0]][temp[1]] = '0'
//			for _, dir := range &dirs {
//				nextI, nextJ := temp[0]+dir[0], temp[1]+dir[1]
//				if nextI >= 0 && nextI < m && nextJ >= 0 && nextJ < n && grid[nextI][nextJ] == '1' {
//					queue = append(queue, [2]int{nextI, nextJ})
//				}
//			}
//		}
//	}
//
//	for i := 0; i < m; i++ {
//		for j := 0; j < n; j++ {
//			if grid[i][j] == '1' {
//				ans++
//				bfs(i, j)
//			}
//		}
//	}
//
//	return ans
//}

type UnionFind struct {
	parents []int
	rands   []int
	count   int
}

func NewUnionFind(grid [][]byte) *UnionFind {
	m, n := len(grid), len(grid[0])
	u := &UnionFind{
		parents: make([]int, m*n),
		rands:   make([]int, m*n),
		count:   0,
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				var temp = i*n + j
				u.parents[temp] = temp
				u.count++
			}
		}
	}
	return u
}

func (u *UnionFind) Find(i int) int {
	if u.parents[i] != i {
		u.parents[i] = u.Find(u.parents[i])
	}
	return u.parents[i]
}

func (u *UnionFind) Union(i, j int) {
	rootI, rootJ := u.Find(i), u.Find(j)
	if rootI != rootJ {
		if u.rands[rootI] > u.rands[rootJ] {
			u.parents[rootJ] = rootI
		} else if u.rands[rootI] < u.rands[rootJ] {
			u.parents[rootI] = rootJ
		} else {
			u.parents[rootJ] = rootI
			u.rands[rootI]++
		}
		u.count--
	}
}

// Method three: Union&find
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	u := NewUnionFind(grid)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				for _, dir := range &dirs {
					nextI, nextJ := i+dir[0], j+dir[1]
					if nextI >= 0 && nextI < m && nextJ >= 0 && nextJ < n && grid[nextI][nextJ] == '1' {
						u.Union(i*n+j, nextI*n+nextJ)
					}
				}
			}
		}
	}

	return u.count
}
