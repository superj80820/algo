// tags: star3, advanced-graphs, PR

import "sort"

// time complexity: O(e^d)
// space complexity: O(e)
// `e` is length of tickets
// `d` is max number of outgoing flights for airport
func findItinerary(tickets [][]string) []string {
	adj := make(map[string][]string)
	for _, ticket := range tickets {
		adj[ticket[0]] = append(adj[ticket[0]], ticket[1])
	}
	for _, val := range adj {
		sort.Slice(val, func(i, j int) bool {
			return val[i] < val[j]
		})
	}

	var dfs func(cur string) bool
	res := []string{"JFK"}
	dfs = func(cur string) bool {
		if len(res) == len(tickets)+1 {
			return true
		}
		if _, ok := adj[cur]; !ok {
			return false
		}

		temp := make([]string, len(adj[cur]))
		copy(temp, adj[cur])
		for _, val := range temp {
			adj[cur] = adj[cur][1:]
			res = append(res, val)
			if dfs(val) {
				return true
			}
			res = res[:len(res)-1]
			adj[cur] = append(adj[cur], val)
		}

		return false
	}
	dfs("JFK")

	return res
}

import "sort"

// time complexity: O(eloge)
// space complexity: O(e)
// `e` is length of tickets
func findItinerary(tickets [][]string) []string {
	adj := make(map[string][]string)
	for _, ticket := range tickets {
		adj[ticket[0]] = append(adj[ticket[0]], ticket[1])
	}
	for _, val := range adj {
		sort.Slice(val, func(i, j int) bool {
			return val[i] > val[j]
		})
	}

	var (
		res []string
		dfs func(cur string) bool
	)
	dfs = func(cur string) bool {
		for len(adj[cur]) != 0 {
			var pop string
			pop, adj[cur] = adj[cur][len(adj[cur])-1], adj[cur][:len(adj[cur])-1]
			dfs(pop)

		}
		res = append(res, cur)
		return false
	}
	dfs("JFK")

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}
