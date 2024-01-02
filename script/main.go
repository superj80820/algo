package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/superj80820/algo/script/domain"
	"github.com/superj80820/algo/script/file"
)

var topicOrderData = []string{
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

func main() {
	action := os.Getenv("ACTION")

	switch action {
	case "update-readme":
		updateLeetcodeListForReadMe()
	case "create-exam":
		if err := createExam(2, 2, 1); err != nil {
			panic(fmt.Sprintf("%+v", err))
		}
	default:
		panic("no use action argument")
	}
}

func createExam(easyCount, mediumCount, hardCount int) error {
	fileHandler := file.CreateFileHandler(topicOrderData, "../neetcode", "../README.md", "./")
	fileInfosByScore, err := fileHandler.ReadFileInfosByScore()
	if err != nil {
		return errors.Wrap(err, "read file information by score failed")
	}

	examInstance := domain.Exam{
		CreateTime: time.Now(),
	}
	for i := 0; i < easyCount && i < len(fileInfosByScore[domain.DifficultyEasy]); i++ {
		examInstance.Easy = append(examInstance.Easy, &domain.ExamInfo{
			ID:           fileInfosByScore[domain.DifficultyEasy][i].ID,
			Name:         fileInfosByScore[domain.DifficultyEasy][i].Name,
			CurrentScore: fileInfosByScore[domain.DifficultyEasy][i].CurrentScore,
		})
	}
	for i := 0; i < mediumCount && i < len(fileInfosByScore[domain.DifficultyMedium]); i++ {
		examInstance.Medium = append(examInstance.Medium, &domain.ExamInfo{
			ID:           fileInfosByScore[domain.DifficultyMedium][i].ID,
			Name:         fileInfosByScore[domain.DifficultyMedium][i].Name,
			CurrentScore: fileInfosByScore[domain.DifficultyMedium][i].CurrentScore,
		})
	}
	for i := 0; i < hardCount && i < len(fileInfosByScore[domain.DifficultyHard]); i++ {
		examInstance.Hard = append(examInstance.Hard, &domain.ExamInfo{
			ID:           fileInfosByScore[domain.DifficultyHard][i].ID,
			Name:         fileInfosByScore[domain.DifficultyHard][i].Name,
			CurrentScore: fileInfosByScore[domain.DifficultyHard][i].CurrentScore,
		})
	}

	if err := fileHandler.WriteExam(&examInstance); err != nil {
		return errors.Wrap(err, "write exam failed")
	}

	return nil
}

func updateLeetcodeListForReadMe() {
	fileHandler := file.CreateFileHandler(topicOrderData, "../neetcode", "../README.md", "./")

	fileInfos := fileHandler.ReadFileInfos()

	fileInfosByTag := make(map[string][]*file.FileInfo)
	for _, fileInfo := range fileInfos {
		fileInfosByTag[fileInfo.MainTag] = append(fileInfosByTag[fileInfo.MainTag], fileInfo)
	}

	var md strings.Builder
	md.WriteString("## Leetcode\n\n")
	for _, topic := range topicOrderData {
		tag := topic
		md.WriteString("### " + topic + "\n")
		md.WriteString("| Name | Star | Difficulty | Tags |" + "\n")
		md.WriteString("| -------- | -------- | -------- | -------- |" + "\n")
		for _, fileInfo := range fileInfosByTag[tag] {
			md.WriteString("|")
			md.WriteString(fmt.Sprintf("[%d. %s](https://leetcode.com/problems/%s/)", fileInfo.ID, fileInfo.Name, fileInfo.Name))
			md.WriteString("|")
			md.WriteString(fileInfo.StarToEmoji())
			md.WriteString("|")
			md.WriteString(fileInfo.Difficulty.String())
			md.WriteString("|")
			md.WriteString(strings.Join(fileInfo.OtherTags, ", "))
			md.WriteString("|\n")
		}
	}

	fileHandler.WriteReadMe(md.String())
}
