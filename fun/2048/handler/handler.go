package handler

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/superj80820/algo/fun/2048/enum"
	"github.com/superj80820/algo/fun/2048/util"
)

const WinNum int = 2048

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

var singletonGameHandler *GameHandler

type GameHandler struct {
	Data  [][]int
	Score int
}

func CreateGameHandler() *GameHandler {
	return &GameHandler{}
}

func GetSingleTonGameHandler() *GameHandler {
	if singletonGameHandler == nil {
		singletonGameHandler = CreateGameHandler()
	}
	return singletonGameHandler
}

func (game *GameHandler) NewGame(size int) bool {
	if size <= 1 {
		return false
	}
	game.Data = make([][]int, size)
	game.Score = 0
	for row := range game.Data {
		game.Data[row] = make([]int, size)
	}
	game.Data = game.randInput()
	return true
}

func (game *GameHandler) NewDefaultGame() {
	game.Data = [][]int{
		{1, 4, 1024, 2},
		{4, 12, 6, 6},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
}

func (game *GameHandler) Process(action enum.Action) {
	game.Move(action)
	game.Merge(action)
	game.Move(action)
	game.AddRandCell()
}

func (game GameHandler) PrintBoard() {
	for _, line := range game.Data {
		fmt.Println(line)
	}
}

func (game GameHandler) ScoreToHTMLString() string {
	return "\u00A0\u00A0\u00A0\u00A0\u00A0SCORE: " + strconv.Itoa(game.Score)
}

func (game GameHandler) BoardToHTMLString() string {
	strLines := make([]string, len(game.Data))
	for row, line := range game.Data {
		strLine := make([]string, len(game.Data[0]))
		for col, val := range line {
			strLine[col] = util.FillNum(strconv.Itoa(val), 4)
		}
		strLines[row] = strings.Join(strLine, " ")
	}
	return strings.Join(strLines, "\n")
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
			if game.Data[row][col] == WinNum {
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

func (game *GameHandler) Merge(action enum.Action) {
	switch action {
	case enum.UP:
		game.merge(CreateBoardTravelIterator(len(game.Data), len(game.Data[0]), enum.UP))
	case enum.DOWN:
		game.merge(CreateBoardTravelIterator(len(game.Data), len(game.Data[0]), enum.DOWN))
	case enum.LEFT:
		game.merge(CreateBoardTravelIterator(len(game.Data), len(game.Data[0]), enum.LEFT))
	case enum.RIGHT:
		game.merge(CreateBoardTravelIterator(len(game.Data), len(game.Data[0]), enum.RIGHT))
	}
}

func (game *GameHandler) Move(action enum.Action) {
	switch action {
	case enum.UP:
		game.move(CreateBoardTravelIterator(len(game.Data), len(game.Data[0]), enum.UP))
	case enum.DOWN:
		game.move(CreateBoardTravelIterator(len(game.Data), len(game.Data[0]), enum.DOWN))
	case enum.LEFT:
		game.move(CreateBoardTravelIterator(len(game.Data), len(game.Data[0]), enum.LEFT))
	case enum.RIGHT:
		game.move(CreateBoardTravelIterator(len(game.Data), len(game.Data[0]), enum.RIGHT))
	}
}

func (game *GameHandler) move(travelIterator *BoardTravelIterator) {
	var swapCellsPosition [][]int
	for {
		position, isBegin, isDone := travelIterator.Next()
		if isDone {
			break
		}
		if isBegin {
			swapCellsPosition = [][]int{}
		}
		row, col := position[0], position[1]

		if game.Data[row][col] == 0 {
			swapCellsPosition = append(swapCellsPosition, []int{row, col})
		} else {
			if len(swapCellsPosition) != 0 {
				var swapPosition []int
				swapPosition, swapCellsPosition = swapCellsPosition[0], swapCellsPosition[1:]
				swapCellsPosition = append(swapCellsPosition, []int{row, col})
				game.Data[swapPosition[0]][swapPosition[1]], game.Data[row][col] = game.Data[row][col], game.Data[swapPosition[0]][swapPosition[1]]
			}
		}
	}
}

func (game *GameHandler) merge(travelIterator *BoardTravelIterator) int {
	pre := -1
	var prePosition []int
	for {
		position, isBegin, isDone := travelIterator.Next()
		if isDone {
			break
		}
		if isBegin {
			pre = -1
			prePosition = []int{}
		}
		row, col := position[0], position[1]

		if pre == game.Data[row][col] {
			game.Data[prePosition[0]][prePosition[1]] = game.Data[row][col] + pre
			game.Score += game.Data[prePosition[0]][prePosition[1]]
			game.Data[row][col] = 0
		}
		pre = game.Data[row][col]
		prePosition = position
	}
	return game.Score
}

func (game *GameHandler) AddRandCell() {
	var randomCells [][2]int
	for col := 0; col < len(game.Data[0]); col++ {
		for row := 0; row < len(game.Data); row++ {
			if game.Data[row][col] == 0 {
				randomCells = append(randomCells, [2]int{row, col})
			}
		}
	}
	if len(randomCells) > 0 {
		randomCell := randomCells[rand.Intn(len(randomCells))]
		game.Data[randomCell[0]][randomCell[1]] = game.getRandomNum()
	}
}

func (game *GameHandler) getRandomNum() int {
	randNum := r.Float64()
	if randNum < 0.75 {
		return 2
	}
	return 4
}

func (game *GameHandler) randInput() [][]int {
	maxCount := len(game.Data) / 2
	randRows, randCols := rand.Perm(len(game.Data)), rand.Perm(len(game.Data[0]))
	for idx := 0; idx < maxCount; idx++ {
		game.Data[randRows[idx]][randCols[idx]] = game.getRandomNum()
	}
	return game.Data
}
