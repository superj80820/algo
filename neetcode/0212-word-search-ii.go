// tags: star3, tries

type Tries struct {
	Children [26]*Tries
	Word     string
}

func (t *Tries) AddWord(word string) {
	cur := t
	for i := range word {
		idx := word[i] - 'a'
		if cur.Children[idx] == nil {
			cur.Children[idx] = NewTries()
		}
		cur = cur.Children[idx]
	}
	cur.Word = word
}

func NewTries() *Tries {
	return &Tries{}
}

var directs = [4][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}

// time complexity: O(m*n*4^L)
// space complexity: O(m*n)
func findWords(board [][]byte, words []string) []string {
	tries := NewTries()
	for _, word := range words {
		tries.AddWord(word)
	}

	var res []string
	var dfs func(row, col int, tries *Tries)
	dfs = func(row, col int, tries *Tries) {
		if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) {
			return
		}
		if board[row][col] == '*' {
			return
		}
		ChildIdx := board[row][col] - 'a'
		if tries.Children[ChildIdx] == nil {
			return
		}

		originWord := board[row][col]
		board[row][col] = '*'
		if tries.Children[ChildIdx].Word != "" {
			res = append(res, tries.Children[ChildIdx].Word)
			tries.Children[ChildIdx].Word = ""
		}
		for _, direct := range directs {
			nextRow, nextCol := row+direct[0], col+direct[1]
			dfs(nextRow, nextCol, tries.Children[ChildIdx])
		}
		board[row][col] = originWord
	}
	for row := range board {
		for col := range board[row] {
			dfs(row, col, tries)
		}
	}

	return res
}