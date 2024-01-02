package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

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

const (
	mdLeetcodeListStart = "<!-- leetcode list start -->"
	mdLeetCodeListEnd   = "<!-- leetcode list end -->"
)

var difficultyMap = map[string]bool{
	"easy":   true,
	"medium": true,
	"hard":   true,
}

func (f FileInfo) StarToEmoji() string {
	var build strings.Builder
	for i := 0; i < f.Star; i++ {
		build.WriteString("⭐")
	}
	return build.String()
}

type topicOrder struct {
	Data    []string
	Visited map[string]bool
}

var singletonTopicOrder *topicOrder

func GetSingletonTopicOrder() *topicOrder {
	if singletonTopicOrder != nil {
		return singletonTopicOrder
	}
	data := []string{
		"arrays&hashing",
		"two-pointers",
		"sliding-window",
		"stack",
		"binary-search",
		"linked-list",
		"trees",
		"tries",
		"heap(priority-queue)",
		"backtracking",
		"graphs",
		"advanced-graphs",
		"1d-dp",
		"2d-dp",
		"greedy",
		"intervals",
		"math&geometry",
		"bit-manipulation",
		"todo",
	}
	visited := make(map[string]bool)
	for _, val := range data {
		visited[val] = true
	}
	singletonTopicOrder = &topicOrder{
		Data:    data,
		Visited: visited,
	}
	return singletonTopicOrder
}

func readFileInfos() []*FileInfo {
	files, err := os.ReadDir("./neetcode")
	if err != nil {
		panic(err)
	}
	fileInfos := make([]*FileInfo, 0, len(files))
	for _, file := range files {
		fileName := file.Name()
		if fileName[len(fileName)-8:] == "_test.go" {
			continue
		}
		data, err := os.ReadFile("./neetcode/" + fileName)
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
		fileInfo, err := CreateFileInfo(fileName, firstLine)
		if err != nil {
			panic(err)
		}
		fileInfos = append(fileInfos, fileInfo)
	}
	return fileInfos
}

func writeReadMe(content string) {
	data, err := os.ReadFile("./README.md")
	if err != nil {
		panic(err)
	}
	mdFile := string(data)
	newMDFile := mdFile[:strings.Index(mdFile, ""+mdLeetcodeListStart+"")-1] +
		content +
		mdFile[strings.Index(mdFile, mdLeetCodeListEnd)+len(mdLeetCodeListEnd):]

	os.WriteFile("./README.md", []byte(newMDFile), 0644)
}

func main() {
	fileInfos := readFileInfos()

	fileInfosByTag := make(map[string][]*FileInfo)
	for _, fileInfo := range fileInfos {
		fileInfosByTag[fileInfo.MainTag] = append(fileInfosByTag[fileInfo.MainTag], fileInfo)
	}

	var md strings.Builder

	md.WriteString(fmt.Sprintf("\n%s\n", mdLeetcodeListStart))
	md.WriteString("## Leetcode\n\n")
	for _, topic := range GetSingletonTopicOrder().Data {
		tag := topic
		md.WriteString("### " + topic + "\n")
		md.WriteString("| Name | Star | Tags |" + "\n")
		md.WriteString("| -------- | -------- | -------- |" + "\n")
		for _, fileInfo := range fileInfosByTag[tag] {
			md.WriteString("|")
			md.WriteString(fmt.Sprintf("[%d. %s](https://leetcode.com/problems/%s/)", fileInfo.ID, fileInfo.Name, fileInfo.Name))
			md.WriteString("|")
			md.WriteString(fileInfo.StarToEmoji())
			md.WriteString("|")
			md.WriteString(strings.Join(fileInfo.OtherTags, ", "))
			md.WriteString("|\n")
		}
	}
	md.WriteString(mdLeetCodeListEnd)

	writeReadMe(md.String())
}

func CreateFileInfo(fileName, tagsInfo string) (*FileInfo, error) {
	id, err := strconv.Atoi(fileName[:4])
	if err != nil {
		return nil, errors.Wrap(err, "parse id failed")
	}
	name := fileName[5:strings.Index(fileName, ".go")]

	var otherTags []string
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
	firstLineEndIdx := strings.Index(tagsInfo, "\n")
	tags := strings.Split(tagsInfo[len(target):firstLineEndIdx], ", ")
	if GetSingletonTopicOrder().Visited[tags[0]] {
		mainTag = tags[0]
		if len(tags) > 1 {
			otherTags = tags[1:]
		}
	} else {
		return nil, errors.New(fmt.Sprint("find no definition tag: ", fileName))
	}

	var (
		star       int
		difficulty string
	)
	for _, tag := range otherTags {
		if tag[:len("star*")-1] == "star" {
			starString := string(tag[len("star*")-1])
			star, err = strconv.Atoi(starString)
			if err != nil {
				return nil, errors.Wrap(err, "star tag format error")
			}
		} else if difficultyMap[tag] {
			difficulty = tag
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
