package main

type UnionFind struct {
	Count  int
	Parent []int
	Rand   []int
}

func NewUF(n int) *UnionFind {
	u := &UnionFind{
		Count:  n,
		Parent: make([]int, n),
		Rand:   make([]int, n),
	}

	for i := range u.Parent {
		u.Parent[i] = i
	}

	return u
}

func (u *UnionFind) Find(i int) int {
	if i != u.Parent[i] {
		u.Parent[i] = u.Find(u.Parent[i])
	}
	return u.Parent[i]
}

func (u *UnionFind) Union(i, j int) {
	root1, root2 := u.Find(i), u.Find(j)
	if root1 != root2 {
		if u.Rand[root1] > u.Rand[root2] {
			u.Parent[root2] = root1
		} else if u.Rand[root1] < u.Rand[root2] {
			u.Parent[root1] = root2
		} else {
			u.Parent[root2] = root1
			u.Rand[root1]++
		}
		u.Count--
	}
}

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	uf := NewUF(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnected[i][j] == 1 {
				uf.Union(i, j)
			}
		}
	}

	return uf.Count
}
