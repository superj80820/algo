package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Action int

const (
	UP Action = iota + 1
	DOWN
	LEFT
	RIGHT
)

var gameProcessor = map[Action]struct {
	Move        func(input [][]int)
	Merge       func(input [][]int)
	AddRandCell func(input [][]int)
}{
	UP:    {Move: upMove, Merge: upMerge, AddRandCell: addRandCell},
	DOWN:  {Move: downMove, Merge: downMerge, AddRandCell: addRandCell},
	LEFT:  {Move: leftMove, Merge: leftMerge, AddRandCell: addRandCell},
	RIGHT: {Move: rightMove, Merge: rightMerge, AddRandCell: addRandCell},
}

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	ch := make(chan Action)

	go func() {
		for {
			Action := bufio.NewScanner(os.Stdin)
			Action.Scan()
			switch strings.ToLower(Action.Text()) {
			case "w":
				ch <- UP
			case "s":
				ch <- DOWN
			case "a":
				ch <- LEFT
			case "d":
				ch <- RIGHT
			}
		}
	}()

	// input := [][]int{
	// 	{0, 4, 0, 0},
	// 	{2, 0, 0, 0},
	// 	{2, 0, 0, 0},
	// 	{8, 8, 2, 2},
	// }

	input := make([][]int, 4)
	for row := range input {
		input[row] = make([]int, 4)
	}

	input = randInput(input)

	printBoard(input)

	for {
		action := <-ch
		fmt.Println("-------")
		gameProcess(input, action)
		printBoard(input)
	}

}

func randInput(input [][]int) [][]int {
	rowRand1, colRand1 := rand.Intn(len(input)-1), rand.Intn(len(input[0])-1)
	var rowRand2, colRand2 int
	for rowRand1 == rowRand2 && colRand1 == colRand2 {
		rowRand2, colRand2 = rand.Intn(len(input)-1), rand.Intn(len(input[0])-1)
	}
	input[rowRand1][colRand1] = getRandomNum()
	input[rowRand2][colRand2] = getRandomNum()
	return input
}

func printBoard(input [][]int) {
	for _, line := range input {
		fmt.Println(line)
	}
}

func gameProcess(input [][]int, action Action) {
	gameProcessor[action].Move(input)
	gameProcessor[action].Merge(input)
	gameProcessor[action].Move(input)
	gameProcessor[action].AddRandCell(input)
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
		var randIdx int
		if len(randomCells) > 1 {
			randIdx = rand.Intn(len(randomCells) - 1)
		}
		randomCell := randomCells[randIdx]
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
