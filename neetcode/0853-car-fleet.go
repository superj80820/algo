// tags: stack, star1, medium

// time complexity: O(n * logn)
// space complexity: O(n)
func carFleet(target int, position []int, speed []int) int {
	targetFloat64 := float64(target)

	positionWithSpeed := make([]*PositionWithSpeed, len(position))
	for idx := range position {
		positionWithSpeed[idx] = &PositionWithSpeed{float64(position[idx]), float64(speed[idx])}
	}

	sort.Slice(positionWithSpeed, func(i, j int) bool {
		return positionWithSpeed[i].Position > positionWithSpeed[j].Position
	})

	stack := make([]*PositionWithSpeed, 0, len(position))
	stack = append(stack, positionWithSpeed[0])
	for _, val := range positionWithSpeed[1:] {
		top := stack[len(stack)-1]
		curTime := (targetFloat64 - val.Position) / val.Speed
		topTime := (targetFloat64 - top.Position) / top.Speed
		if curTime > topTime {
			stack = append(stack, val)
		}
	}

	return len(stack)
}

type PositionWithSpeed struct {
	Position float64
	Speed    float64
}