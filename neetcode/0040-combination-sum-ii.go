// tags: backtracking, star3

import "sort"

// time complexity: O(n^(m/minCand+1))
// space complexity: O(m/minCand)
// `n` is len(candidates)
// `m` is target
// `minCand` is min(candidates)
// `m/minCard+1` is tree's depth
// ref(although they are different problems, the concepts are similar): https://leetcode.com/problems/combination-sum/solutions/937255/python-3-dfs-backtracking-two-dp-methods-explanations/
func combinationSum2(candidates []int, target int) [][]int {
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

	pre := -1
	for i := position; i < len(candidates); i++ {
		if candidates[i] == pre {
			continue
		}
		if candidates[i] > target {
			break
		}
		result = append(result, candidates[i])
		dfs(i+1, target-candidates[i], candidates, result, res)
		result = result[:len(result)-1]
		pre = candidates[i]
	}
}