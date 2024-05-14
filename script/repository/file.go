package repository

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/superj80820/algo/script/domain"
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

type fileRepo struct {
	neetcodeFolderPath string
	createFileInfo     func(fileName, tagsInfo string) (*domain.FileInfo, error)
	topicOrder         *topicOrder
}

func CreateFileRepo(neetcodeFolderPath string, topicOrderData []string) domain.FileRepo {
	visited := make(map[string]bool)
	for _, val := range topicOrderData {
		visited[val] = true
	}
	topicOrder := &topicOrder{
		Data:    topicOrderData,
		Visited: visited,
	}

	return &fileRepo{
		neetcodeFolderPath: neetcodeFolderPath,
		createFileInfo:     createFileInfo(topicOrder),
		topicOrder:         topicOrder,
	}
}

func (f *fileRepo) ReadAll() ([]*domain.FileInfo, error) {
	files, err := os.ReadDir(f.neetcodeFolderPath)
	if err != nil {
		panic(err)
	}
	fileInfos := make([]*domain.FileInfo, 0, len(files))
	for _, file := range files {
		fileName := file.Name()
		if fileName[len(fileName)-4:] == ".mod" {
			continue
		}
		if fileName[len(fileName)-8:] == "_test.go" {
			continue
		}
		if fileName[len(fileName)-3:] != ".go" {
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
		fileInfos = append(fileInfos, fileInfo)
	}
	return fileInfos, nil
}

func (f *fileRepo) GetTopicsByOrder() []string {
	return f.topicOrder.Data
}

func createFileInfo(topicOrder *topicOrder) func(fileName, tagsInfo string) (*domain.FileInfo, error) {
	return func(fileName, tagsInfo string) (*domain.FileInfo, error) {

		id, err := strconv.Atoi(fileName[:4])
		if err != nil {
			return nil, errors.Wrap(err, "parse id failed")
		}
		name := fileName[5:strings.Index(fileName, ".go")]

		var originOtherTags []string
		mainTag := "todo"
		target := "// tags: "
		if len(tagsInfo) < len(target) || target != tagsInfo[:len(target)] {
			return &domain.FileInfo{
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

		return &domain.FileInfo{
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
