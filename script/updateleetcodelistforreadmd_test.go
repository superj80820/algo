package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFileInfo(t *testing.T) {
	fileInfo, err := CreateFileInfo("0001-file-name.go", `// tags: 1d-dp, star3, medium
    // code...
    `)
	assert.Nil(t, err)

	assert.Equal(t, "medium", fileInfo.Difficulty)
	assert.Equal(t, 3, fileInfo.Star)
}
