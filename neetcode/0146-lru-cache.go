// tags: linked-list, star3, medium

package neetcode

import "container/list"

type LRUCache struct {
	listData *list.List
	mapData  map[int]*list.Element
	cap      int
}

type Element struct {
	Key   int
	Value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		listData: list.New(),
		mapData:  make(map[int]*list.Element),
		cap:      capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	// 1. get
	element, ok := this.mapData[key]
	if !ok {
		return -1
	}

	// 2. update to first
	this.listData.MoveToFront(element)

	return element.Value.(*Element).Value
}

func (this *LRUCache) Put(key int, value int) {
	element, ok := this.mapData[key]
	if ok {
		this.listData.MoveToFront(element)
		element.Value.(*Element).Value = value
		return
	}
	if this.listData.Len() == this.cap {
		backElement := this.listData.Back()
		delete(this.mapData, backElement.Value.(*Element).Key)
		this.listData.Remove(backElement)
		element = this.listData.PushFront(&Element{
			Key:   key,
			Value: value,
		})
		this.mapData[key] = element
	} else {
		element = this.listData.PushFront(&Element{
			Key:   key,
			Value: value,
		})
		this.mapData[key] = element
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
