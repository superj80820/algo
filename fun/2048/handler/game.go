package handler

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/superj80820/algo/fun/2048/enum"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

type GameHandler struct {
	Processor map[enum.Action]struct {
		Move        func(input [][]int)
		Merge       func(input [][]int)
		AddRandCell func(input [][]int)
	}
	Data [][]int
}

func (game *GameHandler) NewGame(rowSize, colSize int) bool {
	if rowSize <= 1 && colSize <= 1 {
		fmt.Println("adf")
		return false
	}
	game.Data = make([][]int, rowSize)
	for row := range game.Data {
		game.Data[row] = make([]int, colSize)
	}
	game.Data = randInput(game.Data)
	fmt.Println("hi")
	return true
}

func (game *GameHandler) NewDefaultGame() {
	game.Data = [][]int{
		{1, 4, 1024, 1024},
		{4, 12, 6, 6},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
}

func (game *GameHandler) Process(action enum.Action) {
	game.Processor[action].Move(game.Data)
	game.Processor[action].Merge(game.Data)
	game.Processor[action].Move(game.Data)
	game.Processor[action].AddRandCell(game.Data)
}

func (game GameHandler) PrintBoard() {
	for _, line := range game.Data {
		fmt.Println(line)
	}
}

func (game GameHandler) CheckAvailable() bool {
	for row, line := range game.Data {
		for col := range line {
			if game.Data[row][col] == 0 {
				return true
			}
		}
	}

	for row, line := range game.Data {
		for col := range line {
			if game.checkNeighborsIsSame(row, col) {
				return true
			}
		}
	}

	return false
}

func (game GameHandler) CheckWin() bool {
	for row, line := range game.Data {
		for col := range line {
			if game.Data[row][col] == 2048 {
				return true
			}
		}
	}

	return false
}

func (game GameHandler) checkNeighborsIsSame(row, col int) bool {
	target := game.Data[row][col]
	if col-1 >= 0 && game.Data[row][col-1] == target {
		return true
	} else if row-1 >= 0 && game.Data[row-1][col] == target {
		return true
	} else if col+1 < len(game.Data) && game.Data[row][col+1] == target {
		return true
	} else if row+1 < len(game.Data) && game.Data[row+1][col] == target {
		return true
	}
	return false

}

func CreateGameHandler() *GameHandler {
	return &GameHandler{
		Processor: map[enum.Action]struct {
			Move        func(input [][]int)
			Merge       func(input [][]int)
			AddRandCell func(input [][]int)
		}{
			enum.UP:    {Move: upMove, Merge: upMerge, AddRandCell: addRandCell},
			enum.DOWN:  {Move: downMove, Merge: downMerge, AddRandCell: addRandCell},
			enum.LEFT:  {Move: leftMove, Merge: leftMerge, AddRandCell: addRandCell},
			enum.RIGHT: {Move: rightMove, Merge: rightMerge, AddRandCell: addRandCell},
		},
	}
}

func randInput(input [][]int) [][]int {
	rowRand1, colRand1 := rand.Intn(len(input)), rand.Intn(len(input[0]))
	var rowRand2, colRand2 int
	for rowRand1 == rowRand2 && colRand1 == colRand2 {
		rowRand2, colRand2 = rand.Intn(len(input)), rand.Intn(len(input[0]))
	}
	input[rowRand1][colRand1] = getRandomNum()
	input[rowRand2][colRand2] = getRandomNum()
	return input
}

func upMove(input [][]int) {
	for col := 0; col < len(input[0]); col++ {
		var swapCellsRow []int
		for row := 0; row < len(input); row++ {
			if input[row][col] == 0 {
				swapCellsRow = append(swapCellsRow, row)
			} else {
				if len(swapCellsRow) != 0 {
					var swapRow int
					swapRow, swapCellsRow = swapCellsRow[0], swapCellsRow[1:]
					swapCellsRow = append(swapCellsRow, row)
					input[swapRow][col], input[row][col] = input[row][col], input[swapRow][col]
				}
			}
		}
	}
}

func upMerge(input [][]int) {
	for col := 0; col < len(input[0]); col++ {
		pre := -1
		for row := 0; row < len(input); row++ {
			if pre == input[row][col] {
				input[row-1][col] = input[row][col] + pre
				input[row][col] = 0
			}
			pre = input[row][col]
		}
	}
}

func downMove(input [][]int) {
	for col := 0; col < len(input[0]); col++ {
		var swapCellsRow []int
		for row := len(input) - 1; row >= 0; row-- {
			if input[row][col] == 0 {
				swapCellsRow = append(swapCellsRow, row)
			} else {
				if len(swapCellsRow) != 0 {
					var swapRow int
					swapRow, swapCellsRow = swapCellsRow[0], swapCellsRow[1:]
					swapCellsRow = append(swapCellsRow, row)
					input[swapRow][col], input[row][col] = input[row][col], input[swapRow][col]
				}
			}
		}
	}
}

func downMerge(input [][]int) {
	for col := 0; col < len(input[0]); col++ {
		pre := -1
		for row := len(input) - 1; row >= 0; row-- {
			if pre == input[row][col] {
				input[row+1][col] = input[row][col] + pre
				input[row][col] = 0
			}
			pre = input[row][col]
		}
	}
}

func leftMove(input [][]int) {
	for row := 0; row < len(input); row++ {
		var swapCellsCol []int
		for col := 0; col < len(input[0]); col++ {
			if input[row][col] == 0 {
				swapCellsCol = append(swapCellsCol, col)
			} else {
				if len(swapCellsCol) != 0 {
					var swapCol int
					swapCol, swapCellsCol = swapCellsCol[0], swapCellsCol[1:]
					swapCellsCol = append(swapCellsCol, col)
					input[row][swapCol], input[row][col] = input[row][col], input[row][swapCol]
				}
			}
		}
	}
}

func leftMerge(input [][]int) {
	for row := 0; row < len(input); row++ {
		pre := -1
		for col := 0; col < len(input[0]); col++ {
			if pre == input[row][col] {
				input[row][col-1] = input[row][col] + pre
				input[row][col] = 0
			}
			pre = input[row][col]
		}
	}
}

func rightMove(input [][]int) {
	for row := 0; row < len(input); row++ {
		var swapCellsCol []int
		for col := len(input[0]) - 1; col >= 0; col-- {
			if input[row][col] == 0 {
				swapCellsCol = append(swapCellsCol, col)
			} else {
				if len(swapCellsCol) != 0 {
					var swapCol int
					swapCol, swapCellsCol = swapCellsCol[0], swapCellsCol[1:]
					swapCellsCol = append(swapCellsCol, col)
					input[row][swapCol], input[row][col] = input[row][col], input[row][swapCol]
				}
			}
		}
	}
}

func rightMerge(input [][]int) {
	for row := 0; row < len(input); row++ {
		pre := -1
		for col := len(input[0]) - 1; col >= 0; col-- {
			if pre == input[row][col] {
				input[row][col+1] = input[row][col] + pre
				input[row][col] = 0
			}
			pre = input[row][col]
		}
	}
}

func addRandCell(input [][]int) {
	var randomCells [][2]int
	for col := 0; col < len(input[0]); col++ {
		for row := 0; row < len(input); row++ {
			if input[row][col] == 0 {
				randomCells = append(randomCells, [2]int{row, col})
			}
		}
	}
	if len(randomCells) > 0 {
		randomCell := randomCells[rand.Intn(len(randomCells))]
		input[randomCell[0]][randomCell[1]] = getRandomNum()
	}
}

func getRandomNum() int {
	randNum := r.Float64()
	if randNum < 0.75 {
		return 2
	}
	return 4
}