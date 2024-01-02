package domain

import "time"

type DifficultyType int

const (
	DifficultyUnknown DifficultyType = iota
	DifficultyEasy
	DifficultyMedium
	DifficultyHard
)

func (d DifficultyType) String() string {
	switch d {
	case DifficultyEasy:
		return "easy"
	case DifficultyMedium:
		return "medium"
	case DifficultyHard:
		return "hard"
	case DifficultyUnknown:
		return "unknown"
	default:
		return "unknown"
	}
}

type ExamInfo struct {
	ID           int
	Name         string
	Done         bool
	CurrentScore int
	Unfamiliar   int
}

type Exam struct {
	CreateTime time.Time   `json:"create_time"`
	Easy       []*ExamInfo `json:"easy"`
	Medium     []*ExamInfo `json:"medium"`
	Hard       []*ExamInfo `json:"hard"`
}
