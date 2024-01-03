package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFileInfo(t *testing.T) {
	topicOrderData := []string{
		"arrays&hashing",
		"two-pointers",
		"sliding-window",
		"stack",
		"binary-search",
		"linked-list",
		"trees",
		"tries",
		"heap(priority-queue)",
		"backtracking",
		"graphs",
		"advanced-graphs",
		"1d-dp",
		"2d-dp",
		"greedy",
		"intervals",
		"math&geometry",
		"bit-manipulation",
		"todo",
	}
	visited := make(map[string]bool)
	for _, val := range topicOrderData {
		visited[val] = true
	}
	topicOrder := &topicOrder{
		Data:    topicOrderData,
		Visited: visited,
	}

	fileInfo, err := createFileInfo(topicOrder)("0001-file-name.go", `// tags: 1d-dp, star3, medium
    // code...
    `)
	assert.Nil(t, err)

	assert.Equal(t, "medium", fileInfo.Difficulty)
	assert.Equal(t, 3, fileInfo.Star)
}
