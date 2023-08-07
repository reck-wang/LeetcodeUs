package main

func generateParenthesis(n int) []string {
	var ans []string
	generate("", &ans, n, n)
	return ans
}

func generate(substr string, ans *[]string, left, right int) {
	if left == 0 && right == 0 {
		*ans = append(*ans, substr)
		return
	}

	if left > 0 {
		generate(substr+"(", ans, left-1, right)
	}

	if right > left {
		generate(substr+")", ans, left, right-1)
	}
}
