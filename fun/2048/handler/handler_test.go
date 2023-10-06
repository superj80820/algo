package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/superj80820/algo/fun/2048/enum"
)

type GameHandlerWithMock struct {
	GameHandler
}

func (game *GameHandlerWithMock) Process(action enum.Action) {
	game.Move(action)
	game.Merge(action)
	game.Move(action)
}

func TestGameAction(t *testing.T) {
	GameHandlerWithMock := GameHandlerWithMock{}

	GameHandlerWithMock.NewDefaultGame()

	qs := []struct {
		Action   enum.Action
		Expected [][]int
	}{
		{
			Action: enum.UP,
			Expected: [][]int{
				{2, 4, 1024, 2},
				{4, 12, 8, 16},
				{8, 48, 32, 24},
				{12, 0, 2, 0},
			},
		},
		{
			Action: enum.DOWN,
			Expected: [][]int{
				{2, 0, 1024, 0},
				{4, 4, 8, 2},
				{8, 12, 32, 16},
				{12, 48, 2, 24},
			},
		},
		{
			Action: enum.LEFT,
			Expected: [][]int{
				{2, 1024, 0, 0},
				{8, 8, 2, 0},
				{8, 12, 32, 16},
				{12, 48, 2, 24},
			},
		},
		{
			Action: enum.RIGHT,
			Expected: [][]int{
				{0, 0, 2, 1024},
				{0, 0, 16, 2},
				{8, 12, 32, 16},
				{12, 48, 2, 24},
			},
		},
	}

	for _, q := range qs {
		GameHandlerWithMock.Process(q.Action)
		for row := range GameHandlerWithMock.Data {
			for col := range GameHandlerWithMock.Data[0] {
				assert.Equal(t, q.Expected[row][col], GameHandlerWithMock.Data[row][col])
			}
		}
		assert.True(t, GameHandlerWithMock.CheckAvailable())
		assert.False(t, GameHandlerWithMock.CheckWin())
	}
}
