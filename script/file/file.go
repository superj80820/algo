package file

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/superj80820/algo/script/domain"
)

const (
	mdLeetcodeListStart = "<!-- leetcode list start -->"
	mdLeetCodeListEnd   = "<!-- leetcode list end -->"
)

var difficultyMap = map[string]domain.DifficultyType{
	domain.DifficultyEasy.String():   domain.DifficultyEasy,
	domain.DifficultyMedium.String(): domain.DifficultyMedium,
	domain.DifficultyHard.String():   domain.DifficultyHard,
}

type topicOrder struct {
	Data    []string
	Visited map[string]bool
}

var singletonTopicOrder *topicOrder

type FileInfo struct {
	ID               int
	Name             string
	MainTag          string
	OtherTags        []string
	HasTags          bool
	Star             int
	Difficulty       domain.DifficultyType
	CurrentScore     int
	IsFreeInLeetcode bool
}

func (f FileInfo) StarToEmoji() string {
	var build strings.Builder
	for i := 0; i < f.Star; i++ {
		build.WriteString("⭐")
	}
	return build.String()
}

type fileHandler struct {
	examFilePath       string
	neetcodeFolderPath string
	readMeFilePath     string
	topicOrder         *topicOrder
	createFileInfo     func(fileName string, tagsInfo string) (*FileInfo, error)
}

type fileHandlerInterface interface {
	ReadFileInfosByScore() (map[domain.DifficultyType][]*FileInfo, error)
	WriteExam(exam *domain.Exam) error
	ReadFileInfos() []*FileInfo
	WriteReadMe(content string)
}

func CreateFileHandler(topicOrderData []string, neetcodeFolderPath, readMeFilePath, examFilePath string) fileHandlerInterface {
	visited := make(map[string]bool)
	for _, val := range topicOrderData {
		visited[val] = true
	}
	topicOrder := &topicOrder{
		Data:    topicOrderData,
		Visited: visited,
	}

	return &fileHandler{
		examFilePath:       examFilePath,
		neetcodeFolderPath: neetcodeFolderPath,
		readMeFilePath:     readMeFilePath,
		topicOrder:         topicOrder,
		createFileInfo:     createFileInfo(topicOrder),
	}
}

type scoreType struct {
	id    int
	score int
}

func (f *fileHandler) WriteExam(exam *domain.Exam) error {
	fileName := "exam" + "-" + time.Now().Format("2006-01-02")
	jsonData, err := json.MarshalIndent(exam, "", "\t")
	if err != nil {
		return errors.Wrap(err, "json marshal failed")
	}
	if err := os.WriteFile(f.examFilePath+"/"+fileName+".json", jsonData, 0644); err != nil {
		return errors.Wrap(err, "write file failed")
	}
	return nil
}

func (f *fileHandler) ReadFileInfosByScore() (map[domain.DifficultyType][]*FileInfo, error) {
	fileInfos := f.ReadFileInfos()

	files, err := os.ReadDir(f.examFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "read files failed")
	}
	historyScore := make(map[int]int)
	for _, file := range files {
		fileName := file.Name()
		if !(len(fileName) >= 5 && fileName[:5] == "exam-") {
			continue
		}
		data, err := os.ReadFile(f.examFilePath + "/" + fileName)
		if err != nil {
			return nil, errors.Wrap(err, "read files failed")
		}
		var exam domain.Exam
		if err := json.Unmarshal(data, &exam); err != nil {
			return nil, errors.Wrap(err, "json unmarshal failed")
		}
		if exam.CreateTime.Sub(time.Now()).Hours() < 72 {
			continue
		}

		for _, val := range exam.Easy {
			if !val.Done {
				continue
			}
			historyScore[val.ID]--
			historyScore[val.ID] += val.Unfamiliar
		}
		for _, val := range exam.Medium {
			if !val.Done {
				continue
			}
			historyScore[val.ID]--
			historyScore[val.ID] += val.Unfamiliar
		}
		for _, val := range exam.Hard {
			if !val.Done {
				continue
			}
			historyScore[val.ID]--
			historyScore[val.ID] += val.Unfamiliar
		}
	}

	fileInfoMap := make(map[int]*FileInfo, len(fileInfos))
	scores := make([]*scoreType, len(fileInfos))
	for idx, fileInfo := range fileInfos {
		score := fileInfo.Star
		score += historyScore[fileInfo.ID]

		scores[idx] = &scoreType{
			id:    fileInfo.ID,
			score: score,
		}

		fileInfoMap[fileInfo.ID] = fileInfo
	}

	for i := range scores { // shuffle scores
		j := rand.Intn(i + 1)
		scores[i], scores[j] = scores[j], scores[i]
	}

	sort.SliceStable(scores, func(i, j int) bool { // TODO: use heap
		return scores[i].score > scores[j].score
	})

	fileInfosByScore := make(map[domain.DifficultyType][]*FileInfo)
	for _, score := range scores {
		fileInfo := fileInfoMap[score.id]
		fileInfo.CurrentScore = score.score
		fileInfosByScore[fileInfo.Difficulty] = append(fileInfosByScore[fileInfo.Difficulty], fileInfo)
	}

	return fileInfosByScore, nil
}

func (f *fileHandler) ReadFileInfos() []*FileInfo {
	files, err := os.ReadDir(f.neetcodeFolderPath)
	if err != nil {
		panic(err)
	}
	fileInfos := make([]*FileInfo, 0, len(files))
	for _, file := range files {
		fileName := file.Name()
		if fileName[len(fileName)-8:] == "_test.go" {
			continue
		}
		data, err := os.ReadFile(f.neetcodeFolderPath + "/" + fileName)
		if err != nil {
			panic(err)
		}
		fileString := string(data)
		var firstLineEndIdx int
		for idx, str := range fileString {
			if str == '\n' {
				firstLineEndIdx = idx
				break
			}
		}
		firstLine := fileString[:firstLineEndIdx]
		fileInfo, err := f.createFileInfo(fileName, firstLine)
		if err != nil {
			panic(err)
		}
		fileInfos = append(fileInfos, fileInfo)
	}
	return fileInfos
}

func (f *fileHandler) WriteReadMe(content string) {
	content = fmt.Sprintf("\n%s\n", mdLeetcodeListStart) + content + mdLeetCodeListEnd

	data, err := os.ReadFile(f.readMeFilePath)
	if err != nil {
		panic(err)
	}
	mdFile := string(data)
	newMDFile := mdFile[:strings.Index(mdFile, ""+mdLeetcodeListStart+"")-1] +
		content +
		mdFile[strings.Index(mdFile, mdLeetCodeListEnd)+len(mdLeetCodeListEnd):]

	os.WriteFile(f.readMeFilePath, []byte(newMDFile), 0644)
}

func createFileInfo(topicOrder *topicOrder) func(fileName, tagsInfo string) (*FileInfo, error) {
	return func(fileName, tagsInfo string) (*FileInfo, error) {

		id, err := strconv.Atoi(fileName[:4])
		if err != nil {
			return nil, errors.Wrap(err, "parse id failed")
		}
		name := fileName[5:strings.Index(fileName, ".go")]

		var originOtherTags []string
		mainTag := "todo"
		target := "// tags: "
		if len(tagsInfo) < len(target) || target != tagsInfo[:len(target)] {
			return &FileInfo{
				ID:               id,
				Name:             name,
				MainTag:          mainTag,
				IsFreeInLeetcode: true,
			}, nil
		}
		tags := strings.Split(tagsInfo[len(target):], ", ")
		if topicOrder.Visited[tags[0]] {
			mainTag = tags[0]
			if len(tags) > 1 {
				originOtherTags = tags[1:]
			}
		} else {
			return nil, errors.New(fmt.Sprint("find no definition tag: ", fileName))
		}

		var (
			star       int
			difficulty domain.DifficultyType
			otherTags  []string
		)
		for _, tag := range originOtherTags {
			if len(tag) >= len("star*") && tag[:len("star*")-1] == "star" {
				starString := string(tag[len("star*")-1])
				star, err = strconv.Atoi(starString)
				if err != nil {
					return nil, errors.Wrap(err, "star tag format error")
				}
			} else if difficultyType, ok := difficultyMap[tag]; ok {
				difficulty = difficultyType
			} else {
				otherTags = append(otherTags, tag)
			}
		}

		return &FileInfo{
			ID:               id,
			Name:             name,
			MainTag:          mainTag,
			OtherTags:        otherTags,
			HasTags:          true,
			Star:             star,
			Difficulty:       difficulty,
			IsFreeInLeetcode: true,
		}, nil
	}
}

// func (f *fileHandler) createFileInfo(fileName, tagsInfo string) (*fileInfo, error) {
// 	id, err := strconv.Atoi(fileName[:4])
// 	if err != nil {
// 		return nil, errors.Wrap(err, "parse id failed")
// 	}
// 	name := fileName[5:strings.Index(fileName, ".go")]

// 	var otherTags []string
// 	mainTag := "todo"
// 	target := "// tags: "
// 	if len(tagsInfo) < len(target) || target != tagsInfo[:len(target)] {
// 		return &fileInfo{
// 			ID:               id,
// 			Name:             name,
// 			MainTag:          mainTag,
// 			IsFreeInLeetcode: true,
// 		}, nil
// 	}
// 	firstLineEndIdx := strings.Index(tagsInfo, "\n")
// 	tags := strings.Split(tagsInfo[len(target):firstLineEndIdx], ", ")
// 	if f.topicOrder.Visited[tags[0]] {
// 		mainTag = tags[0]
// 		if len(tags) > 1 {
// 			otherTags = tags[1:]
// 		}
// 	} else {
// 		return nil, errors.New(fmt.Sprint("find no definition tag: ", fileName))
// 	}

// 	var (
// 		star       int
// 		difficulty string
// 	)
// 	for _, tag := range otherTags {
// 		if tag[:len("star*")-1] == "star" {
// 			starString := string(tag[len("star*")-1])
// 			star, err = strconv.Atoi(starString)
// 			if err != nil {
// 				return nil, errors.Wrap(err, "star tag format error")
// 			}
// 		} else if difficultyMap[tag] {
// 			difficulty = tag
// 		}
// 	}

// 	return &fileInfo{
// 		ID:               id,
// 		Name:             name,
// 		MainTag:          mainTag,
// 		OtherTags:        otherTags,
// 		HasTags:          true,
// 		Star:             star,
// 		Difficulty:       difficulty,
// 		IsFreeInLeetcode: true,
// 	}, nil
// }
