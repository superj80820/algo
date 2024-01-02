// tags: math&geometry, medium

type DetectSquares struct {
	pointCount map[[2]int]int
	points     [][]int
}

func Constructor() DetectSquares {
	return DetectSquares{
		pointCount: make(map[[2]int]int),
	}
}

// time complexity: O(1)
// space complexity: O(1)
func (this *DetectSquares) Add(point []int) {
	this.pointCount[[2]int{point[0], point[1]}]++
	this.points = append(this.points, point)
}

// time complexity: O(n)
// space complexity: O(1)
func (this *DetectSquares) Count(point []int) int {
	var res int
	for _, curPoint := range this.points {
		if abs(point[0]-curPoint[0]) == abs(point[1]-curPoint[1]) && point[0] != curPoint[0] && point[1] != curPoint[1] {
			res += this.pointCount[[2]int{point[0], curPoint[1]}] * this.pointCount[[2]int{curPoint[0], point[1]}]
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */