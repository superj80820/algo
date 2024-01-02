package file

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	mdLeetcodeListStart = "<!-- leetcode list start -->"
	mdLeetCodeListEnd   = "<!-- leetcode list end -->"
)

var difficultyMap = map[string]bool{
	"easy":   true,
	"medium": true,
	"hard":   true,
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
	Difficulty       string
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
	neetcodeFolderPath string
	readMeFilePath     string
	topicOrder         *topicOrder
	createFileInfo     func(fileName string, tagsInfo string) (*FileInfo, error)
}

type fileHandlerInterface interface {
	ReadFileInfos() []*FileInfo
	WriteReadMe(content string)
}

func CreateFileHandler(topicOrderData []string, neetcodeFolderPath, readMeFilePath string) fileHandlerInterface {
	visited := make(map[string]bool)
	for _, val := range topicOrderData {
		visited[val] = true
	}
	topicOrder := &topicOrder{
		Data:    topicOrderData,
		Visited: visited,
	}

	return &fileHandler{
		neetcodeFolderPath: neetcodeFolderPath,
		readMeFilePath:     readMeFilePath,
		topicOrder:         topicOrder,
		createFileInfo:     createFileInfo(topicOrder),
	}
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
			difficulty string
			otherTags  []string
		)
		for _, tag := range originOtherTags {
			if len(tag) >= len("star*") && tag[:len("star*")-1] == "star" {
				starString := string(tag[len("star*")-1])
				star, err = strconv.Atoi(starString)
				if err != nil {
					return nil, errors.Wrap(err, "star tag format error")
				}
			} else if difficultyMap[tag] {
				difficulty = tag
			} else {
				otherTags = append(otherTags, tag)
			}
		}

		fmt.Println(name, otherTags, len(otherTags), originOtherTags, tagsInfo)
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
