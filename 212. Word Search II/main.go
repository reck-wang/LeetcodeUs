package main

type Node struct {
	child map[byte]*Node
}

var dirs = [4][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}

func findWords(board [][]byte, words []string) []string {
	root := &Node{child: map[byte]*Node{}}

	for _, word := range words {
		node := root
		for i := 0; i < len(word); i++ {
			if _, has := node.child[word[i]]; !has {
				node.child[word[i]] = &Node{child: map[byte]*Node{}}
			}
			node = node.child[word[i]]
		}
		node.child['#'] = nil
	}

	m, n := len(board), len(board[0])
	visit := make([][]bool, m)
	for i := range visit {
		visit[i] = make([]bool, n)
	}

	var ans = map[string]struct{}{}

	var dfs func(int, int, []byte, *Node)
	dfs = func(i, j int, cur []byte, curNode *Node) {
		cur_word := append(cur, board[i][j])
		cur_node := curNode.child[board[i][j]]

		if _, has := cur_node.child['#']; has {
			ans[string(cur_word)] = struct{}{}
		}

		visit[i][j] = true
		for _, dir := range dirs {
			nextI, nextJ := i+dir[0], j+dir[1]
			if nextI >= 0 && nextI < m && nextJ >= 0 && nextJ < n && !visit[nextI][nextJ] {
				if _, has := cur_node.child[board[nextI][nextJ]]; has {
					dfs(nextI, nextJ, cur_word, cur_node)
				}
			}
		}
		visit[i][j] = false
	}

	for i, row := range board {
		for j, ch := range row {
			if _, has := root.child[ch]; has {
				dfs(i, j, []byte{}, root)
			}
		}
	}

	var res []string
	for s := range ans {
		res = append(res, s)
	}

	return res
}
