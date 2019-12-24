package main

func combine(n int, k int) [][]int {
	remain := n+1-k
	combination := make([]int, k)
	var res [][]int

	var dfs func(int, int)
	dfs = func(idx, begin int) {
		if idx == k {
			temp := make([]int, k)
			copy(temp, combination)
			res = append(res, temp)
			return
		}

		for i := begin; i <= remain+idx; i++ {
			combination[idx] = i
			dfs(idx+1, i+1)
		}
	}

	dfs(0, 1)
	return res
}
