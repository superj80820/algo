// tags: greedy, star2

// func canCompleteCircuit(gas []int, cost []int) int {
//     diff := make([]int, len(gas))
//     var diffSum int
//     for idx := range gas {
//         diff[idx] = gas[idx]-cost[idx]
//         diffSum += diff[idx]
//     }
//     if diffSum < 0 {
//         return -1
//     }

//     var res, total int
//     for idx, val := range diff {
//         preTotal := total
//         total += val
//         if preTotal == 0 && total > 0 {
//             res = idx
//         } else if total < 0 {
//             total = 0
//         }
//     }
//     return res
// }

func canCompleteCircuit(gas []int, cost []int) int {
	var minCity, minDiff, curDiff int
	for i := range gas {
		if curDiff < minDiff {
			minDiff = curDiff
			minCity = i
		}
		curDiff += gas[i] - cost[i]
	}
	if curDiff < 0 {
		return -1
	}
	return minCity
}