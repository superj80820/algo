package usecase

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/superj80820/algo/script/domain"
)

func starToEmoji(starCount int) string {
	var build strings.Builder
	for i := 0; i < starCount; i++ {
		build.WriteString("⭐")
	}
	return build.String()
}

type algoUseCase struct {
	readMeMDRepo domain.ReadMeMDRepo
	fileRepo     domain.FileRepo
}

func CreateAlgoUseCase(readMeMDRepo domain.ReadMeMDRepo, fileRepo domain.FileRepo) domain.ExamUseCase {
	return &algoUseCase{
		readMeMDRepo: readMeMDRepo,
		fileRepo:     fileRepo,
	}
}

func (e *algoUseCase) UpdateReadMe() error {
	fileInfos, err := e.fileRepo.ReadAll()
	if err != nil {
		return errors.Wrap(err, "read all files failed")
	}

	fileInfosByTag := make(map[string][]*domain.FileInfo)
	for _, fileInfo := range fileInfos {
		fileInfosByTag[fileInfo.MainTag] = append(fileInfosByTag[fileInfo.MainTag], fileInfo)
	}

	var md strings.Builder
	md.WriteString("## Leetcode\n\n")
	for _, topic := range e.fileRepo.GetTopicsByOrder() {
		tag := topic
		md.WriteString("### " + topic + "\n")
		md.WriteString("| Name | Star | Difficulty | Practice-Count | Tags |" + "\n")
		md.WriteString("| -------- | -------- | -------- | -------- | -------- |" + "\n")
		for _, fileInfo := range fileInfosByTag[tag] {
			md.WriteString("|")
			md.WriteString(fmt.Sprintf("[%d. %s](https://leetcode.com/problems/%s/)", fileInfo.ID, fileInfo.Name, fileInfo.Name))
			md.WriteString("|")
			md.WriteString(starToEmoji(fileInfo.Star))
			md.WriteString("|")
			md.WriteString(fileInfo.Difficulty.String())
			md.WriteString("|")
			md.WriteString(strconv.Itoa(fileInfo.PracticeCount))
			md.WriteString("|")
			md.WriteString(strings.Join(fileInfo.OtherTags, ", "))
			md.WriteString("|\n")
		}
	}

	if err := e.readMeMDRepo.Write(md.String()); err != nil {
		return errors.Wrap(err, "write readme failed")
	}

	return nil
}
