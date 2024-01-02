// tags: backtracking, star3, medium

// time complexity: O(2^m)
// space complexity: O(m)
// `n` is len(candidates)
// `m` is target, and tree's depth
// `minCand` is min(candidates)
// ref: https://leetcode.com/problems/combination-sum/solutions/937255/python-3-dfs-backtracking-two-dp-methods-explanations/
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	dfs(0, target, candidates, []int{}, &res)
	return res
}

func dfs(idx, target int, candidates, result []int, res *[][]int) {
	if target == 0 {
		newResult := make([]int, len(result))
		copy(newResult, result)
		*res = append(*res, newResult)
		return
	} else if target < 0 || idx >= len(candidates) {
		return
	}

	result = append(result, candidates[idx])
	dfs(idx, target-candidates[idx], candidates, result, res)
	result = result[:len(result)-1]
	dfs(idx+1, target, candidates, result, res)
}

// time complexity: O(n^(m/minCand+1))
// space complexity: O(m/minCand)
// `n` is len(candidates)
// `m` is target
// `minCand` is min(candidates)
// `m/minCard+1` is tree's depth
// ref: https://leetcode.com/problems/combination-sum/solutions/937255/python-3-dfs-backtracking-two-dp-methods-explanations/
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates)
	dfs(0, target, candidates, []int{}, &res)
	return res
}

func dfs(position, target int, candidates, result []int, res *[][]int) {
	if target == 0 {
		newResult := make([]int, len(result))
		copy(newResult, result)
		*res = append(*res, newResult)
		return
	} else if target < 0 || position >= len(candidates) {
		return
	}

	for i := position; i < len(candidates); i++ {
		result = append(result, candidates[i])
		dfs(i, target-candidates[i], candidates, result, res)
		result = result[:len(result)-1]
	}
}