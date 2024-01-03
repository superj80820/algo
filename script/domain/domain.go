package domain

import "time"

type FileRepo interface {
	ReadAll() ([]*FileInfo, error)
	GetTopicsByOrder() []string
}

type ExamRepo interface {
	ReadAll() ([]*Exam, error)
	Create(exam *Exam) error
}

type ReadMeMDRepo interface {
	Write(content string) error
}

type ExamUseCase interface {
	UpdateReadMe() error
	CreateExam(easyCount, mediumCount, hardCount int) error
}

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

type FileInfo struct {
	ID               int
	Name             string
	MainTag          string
	OtherTags        []string
	HasTags          bool
	Star             int
	Difficulty       DifficultyType
	FamiliarScore    int
	IsFreeInLeetcode bool
}

type ExamInfo struct {
	ID         int
	Name       string
	Done       bool
	Familiar   int
	CreateTime time.Time
}

type Exam struct {
	Easy   []*ExamInfo `json:"easy"`
	Medium []*ExamInfo `json:"medium"`
	Hard   []*ExamInfo `json:"hard"`
}
