package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchNeighbors(node int) []int {
	var neighbors []int
	switch node {
	case 1:
		neighbors = []int{2, 3, 4}
	case 2:
		neighbors = []int{1, 5}
	case 3:
		neighbors = []int{1, 5}
	case 4:
		neighbors = []int{1, 6}
	case 5:
		neighbors = []int{2, 3, 7}
	case 6:
		neighbors = []int{4, 7}
	case 7:
		neighbors = []int{5, 6}
	}
	time.Sleep(1 * time.Second) // simulate io use time
	return neighbors
}

func searchGraph(node int) {
	queue := []int{node}
	visited := make(map[int]bool)
	visited[node] = true
	for len(queue) != 0 {
		front := dequeue(&queue)
		fmt.Println(front)
		neighbors := fetchNeighbors(front)
		for _, neighbor := range neighbors {
			if _, ok := visited[neighbor]; ok {
				continue
			}
			queue = append(queue, neighbor)
			visited[neighbor] = true
		}
	}
}

func searchGraphAsync(node int) {
	queue := []int{node}
	visited := make(map[int]bool)
	visited[node] = true
	for len(queue) != 0 {
		curLen := len(queue)
		receiveNeighborsCh := make(chan []int)
		wg := new(sync.WaitGroup)
		wg.Add(curLen)
		for i := 0; i < curLen; i++ {
			front := dequeue(&queue)
			fmt.Println(front)
			go func(front int) {
				defer wg.Done()
				receiveNeighborsCh <- fetchNeighbors(front)
			}(front)
		}
		go func() {
			wg.Wait()
			close(receiveNeighborsCh)
		}()
		for neighbors := range receiveNeighborsCh {
			for _, neighbor := range neighbors {
				if _, ok := visited[neighbor]; ok {
					continue
				}
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
		}
	}

}

func dequeue(queue *[]int) int {
	front := (*queue)[0]
	*queue = (*queue)[1:]
	return front
}

func main() {
	// searchGraph(1)
	searchGraphAsync(1)
}
