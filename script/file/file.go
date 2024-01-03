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
	CreateTime       time.Time
	ID               int
	Name             string
	MainTag          string
	OtherTags        []string
	HasTags          bool
	Star             int
	Difficulty       domain.DifficultyType
	UnfamiliarScore  int
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
	ReadFileInfos() ([]*FileInfo, error)
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
	id              int
	unfamiliarScore int
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
	fileInfos, err := f.ReadFileInfos()
	if err != nil {
		return nil, errors.Wrap(err, "read files information failed")
	}

	var fileInfosRemoveNotEnough72Hours []*FileInfo
	for _, fileInfo := range fileInfos {
		if time.Now().Sub(fileInfo.CreateTime).Hours() < 72 {
			continue
		}
		fileInfosRemoveNotEnough72Hours = append(fileInfosRemoveNotEnough72Hours, fileInfo)
	}

	fileInfoMap := make(map[int]*FileInfo, len(fileInfosRemoveNotEnough72Hours))
	scores := make([]*scoreType, len(fileInfosRemoveNotEnough72Hours))
	for idx, fileInfo := range fileInfosRemoveNotEnough72Hours {
		scores[idx] = &scoreType{
			id:              fileInfo.ID,
			unfamiliarScore: fileInfo.UnfamiliarScore,
		}

		fileInfoMap[fileInfo.ID] = fileInfo
	}

	for i := range scores { // shuffle scores
		j := rand.Intn(i + 1)
		scores[i], scores[j] = scores[j], scores[i]
	}

	sort.SliceStable(scores, func(i, j int) bool { // TODO: use heap
		return scores[i].unfamiliarScore > scores[j].unfamiliarScore
	})

	fileInfosByScore := make(map[domain.DifficultyType][]*FileInfo)
	for _, score := range scores {
		fileInfo := fileInfoMap[score.id]
		fileInfo.UnfamiliarScore = score.unfamiliarScore
		fileInfosByScore[fileInfo.Difficulty] = append(fileInfosByScore[fileInfo.Difficulty], fileInfo)
	}

	return fileInfosByScore, nil
}

func (f *fileHandler) ReadFileInfos() ([]*FileInfo, error) {
	historyScoreMap, err := f.readExamFiles()
	if err != nil {
		return nil, errors.Wrap(err, "read exam files failed")
	}

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
			return nil, errors.Wrap(err, "read file failed")
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
			return nil, errors.Wrap(err, "create file information")
		}
		if historyScore, ok := historyScoreMap[fileInfo.ID]; ok {
			fileInfo.CreateTime = historyScore.CreateTime
			fileInfo.UnfamiliarScore = historyScore.UnfamiliarScore
		} else {
			fileInfo.UnfamiliarScore = -1
		}
		fileInfos = append(fileInfos, fileInfo)
	}
	return fileInfos, nil
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

func (f *fileHandler) readExamFiles() (map[int]*domain.ExamHistory, error) {
	files, err := os.ReadDir(f.examFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "read files failed")
	}
	historyScoreMap := make(map[int]*domain.ExamHistory)
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

		fn := func(val *domain.ExamInfo, createTime time.Time) {
			if !val.Done {
				return
			}

			var unfamiliarScore int

			unfamiliarScore--
			unfamiliarScore += val.Unfamiliar

			var historyScore *domain.ExamHistory
			if historyScoreVal, ok := historyScoreMap[val.ID]; ok {
				historyScore = historyScoreVal
			} else {
				historyScore = &domain.ExamHistory{
					ID:         val.ID,
					Name:       val.Name,
					CreateTime: createTime,
				}
			}

			historyScore.UnfamiliarScore += unfamiliarScore

			historyScoreMap[val.ID] = historyScore
		}
		for _, val := range exam.Easy {
			fn(val, exam.CreateTime)
		}
		for _, val := range exam.Medium {
			fn(val, exam.CreateTime)
		}
		for _, val := range exam.Hard {
			fn(val, exam.CreateTime)
		}
	}
	return historyScoreMap, nil
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
