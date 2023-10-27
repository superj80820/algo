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
	IsFreeInLeetcode bool
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

func main() {
	files, err := os.ReadDir("./neetcode")
	if err != nil {
		panic(err)
	}
	fileInfos := make([]*FileInfo, len(files))
	for idx, file := range files {
		data, err := os.ReadFile("./neetcode/" + file.Name())
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
		fileInfo, err := CreateFileInfo(file.Name(), firstLine)
		if err != nil {
			panic(err)
		}
		fileInfos[idx] = fileInfo
	}

	fileInfosByTag := make(map[string][]*FileInfo)
	for _, fileInfo := range fileInfos {
		fileInfosByTag[fileInfo.MainTag] = append(fileInfosByTag[fileInfo.MainTag], fileInfo)
	}

	var md strings.Builder
	mdLeetcodeListStart, mdLeetCodeListEnd := "<!-- leetcode list start -->", "<!-- leetcode list end -->"
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

	data, err := os.ReadFile("./README.md")
	if err != nil {
		panic(err)
	}
	mdFile := string(data)
	newMDFile := mdFile[:strings.Index(mdFile, ""+mdLeetcodeListStart+"")-1] +
		md.String() +
		mdFile[strings.Index(mdFile, mdLeetCodeListEnd)+len(mdLeetCodeListEnd):]

	os.WriteFile("./README.md", []byte(newMDFile), 0644)
}

func CreateFileInfo(fileName, tagsInfo string) (*FileInfo, error) {
	id, err := strconv.Atoi(fileName[:4])
	if err != nil {
		return nil, errors.Wrap(err, "parse id failed")
	}
	name := fileName[5:strings.Index(fileName, ".go")]

	mainTag := "todo"

	target := "// tags: "
	if len(tagsInfo) < len(target) || target != tagsInfo[:len(target)] { // TODO: york
		return &FileInfo{
			ID:               id,
			Name:             name,
			MainTag:          mainTag,
			IsFreeInLeetcode: true,
		}, nil
	}
	tags := strings.Split(tagsInfo[len(target):], ", ")
	if GetSingletonTopicOrder().Visited[tags[0]] {
		mainTag = tags[0]
	} else {
		panic(fmt.Sprint("find no definition tag: ", fileName))
	}
	var (
		otherTags []string
		star      int
	)
	if len(tags) > 1 && len(tags[1]) >= len("star*") && tags[1][:len("star*")-1] == "star" {
		var err error
		star, err = strconv.Atoi(string(tags[1][len("star*")-1]))
		if err != nil {
			// TODO: york
			fmt.Println(err)
		}
		otherTags = tags[2:]
	} else {
		otherTags = tags[1:]
	}

	return &FileInfo{
		ID:               id,
		Name:             name,
		MainTag:          mainTag,
		OtherTags:        otherTags,
		HasTags:          true,
		Star:             star,
		IsFreeInLeetcode: true,
	}, nil
}
