// tags: binary-search, star2

type TimeMap struct {
	dataset map[string][]*data
}

type data struct {
	value     string
	timestamp int
}

func Constructor() TimeMap {
	return TimeMap{
		dataset: make(map[string][]*data),
	}
}

// time complexity: O(1)
// space complexity: O(1)
func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.dataset[key]; ok {
		this.dataset[key] = append(this.dataset[key], &data{value: value, timestamp: timestamp})
	} else {
		this.dataset[key] = []*data{&data{value: value, timestamp: timestamp}}
	}
}

// time complexity: O(logn)
// space complexity: O(1)
func (this *TimeMap) Get(key string, timestamp int) string {
	if data, ok := this.dataset[key]; ok {
		var dataValue string
		for leftIdx, rightIdx := 0, len(data)-1; leftIdx <= rightIdx; {
			midIdx := (leftIdx + rightIdx) / 2

			if data[midIdx].timestamp <= timestamp {
				dataValue = data[midIdx].value
				leftIdx = midIdx + 1
			} else {
				rightIdx = midIdx - 1
			}
		}
		return dataValue
	} else {
		return ""
	}
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */