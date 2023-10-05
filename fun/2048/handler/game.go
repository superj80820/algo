package handler

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/superj80820/algo/fun/2048/enum"
)

const WinNum int = 2048

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

var singletonGameHandler *GameHandler

type GameHandler struct {
	Data  [][]int
	Score int
}

type gameProcessor struct {
	Move        func(input [][]int)
	Merge       func(input [][]int) int
	AddRandCell func(input [][]int)
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
	fmt.Println(size / 2)
	game.Data = randInput(game.Data, size/2)
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
			strLine[col] = fillNum(strconv.Itoa(val), 4)
		}
		strLines[row] = strings.Join(strLine, " ")
	}
	return strings.Join(strLines, "\n")
}

func fillNum(str string, fillLen int) string {
	strLen := len(str)
	if strLen >= fillLen {
		return str
	}
	var res string
	for i := 0; i < fillLen-strLen; i++ {
		res += "\u00A0"
	}
	return res + str
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

func CreateGameHandler() *GameHandler {
	return &GameHandler{}
}

func randInput(input [][]int, maxCount int) [][]int {
	randRows, randCols := rand.Perm(len(input)), rand.Perm(len(input[0]))
	for idx := 0; idx < maxCount; idx++ {
		input[randRows[idx]][randCols[idx]] = getRandomNum()
	}
	return input
}

type BoardTravelIterator struct {
	curRow int
	curCol int
	rowMax int
	colMax int
	isDone bool
	action enum.Action
}

func CreateBoardTravelIterator(rowMax, colMax int, action enum.Action) *BoardTravelIterator {
	var curRow, curCol int
	switch action {
	case enum.RIGHT:
		curRow = 0
		curCol = colMax - 1
	case enum.LEFT:
		curRow = 0
		curCol = 0
	case enum.UP:
		curRow = 0
		curCol = 0
	case enum.DOWN:
		curRow = rowMax - 1
		curCol = 0
	}

	return &BoardTravelIterator{
		curRow: curRow,
		curCol: curCol,
		rowMax: rowMax,
		colMax: colMax,
		action: action,
	}
}

func (b *BoardTravelIterator) Next() ([]int, bool, bool) {
	curRow, curCol := b.curRow, b.curCol
	var isBegin bool
	switch b.action {
	case enum.RIGHT:
		if b.isDone {
			return []int{}, false, true
		}
		if b.curCol == b.colMax-1 {
			isBegin = true
		}
		if b.curCol > 0 {
			b.curCol--
		} else if b.curCol == 0 {
			b.curCol = b.colMax - 1
			if b.curRow < b.rowMax-1 {
				b.curRow++
			} else if b.curRow == b.rowMax-1 {
				b.isDone = true
			}
		}
	case enum.LEFT:
		if b.isDone {
			return []int{}, false, true
		}
		if b.curCol == 0 {
			isBegin = true
		}
		if b.curCol < b.colMax-1 {
			b.curCol++
		} else if b.curCol == b.colMax-1 {
			b.curCol = 0
			if b.curRow < b.rowMax-1 {
				b.curRow++
			} else if b.curRow == b.rowMax-1 {
				b.isDone = true
			}
		}
	case enum.UP:
		if b.isDone {
			return []int{}, false, true
		}
		if b.curRow == 0 {
			isBegin = true
		}
		if b.curRow < b.rowMax-1 {
			b.curRow++
		} else if b.curRow == b.rowMax-1 {
			b.curRow = 0
			if b.curCol < b.colMax-1 {
				b.curCol++
			} else if b.curCol == b.colMax-1 {
				b.isDone = true
			}
		}
	case enum.DOWN:
		if b.isDone {
			return []int{}, false, true
		}
		if b.curRow == b.rowMax-1 {
			isBegin = true
		}
		if b.curRow > 0 {
			b.curRow--
		} else if b.curRow == 0 {
			b.curRow = b.rowMax - 1
			if b.curCol < b.colMax-1 {
				b.curCol++
			} else if b.curCol == b.colMax-1 {
				b.isDone = true
			}
		}
	}
	return []int{curRow, curCol}, isBegin, false
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
		fmt.Println(position, isBegin, isDone)
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
		game.Data[randomCell[0]][randomCell[1]] = getRandomNum()
	}
}

func getRandomNum() int {
	randNum := r.Float64()
	if randNum < 0.75 {
		return 2
	}
	return 4
}
