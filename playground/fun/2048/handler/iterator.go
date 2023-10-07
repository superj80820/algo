package handler

import "github.com/superj80820/algo/playground/fun/2048/enum"

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
