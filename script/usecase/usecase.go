package usecase

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

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

// type fileHandler struct {
// 	examFilePath       string
// 	neetcodeFolderPath string
// 	readMeFilePath     string
// 	topicOrder         *topicOrder
// 	createFileInfo     func(fileName string, tagsInfo string) (*domain.FileInfo, error)
// }

type examUseCase struct {
	readMeMDRepo domain.ReadMeMDRepo
	examRepo     domain.ExamRepo
	fileRepo     domain.FileRepo
}

func CreateExamUseCase(readMeMDRepo domain.ReadMeMDRepo, examRepo domain.ExamRepo, fileRepo domain.FileRepo) domain.ExamUseCase {
	return &examUseCase{
		readMeMDRepo: readMeMDRepo,
		examRepo:     examRepo,
		fileRepo:     fileRepo,
	}
}

type scoreWithID struct {
	id              int
	unfamiliarScore int
}

func (e *examUseCase) CreateExam(easyCount, mediumCount, hardCount int) error {
	fileInfos, err := e.fileRepo.ReadAll()
	if err != nil {
		return errors.Wrap(err, "read all files failed")
	}
	exams, err := e.examRepo.ReadAll()
	if err != nil {
		return errors.Wrap(err, "read all exam files failed")
	}

	fileInfoMap := make(map[int]*domain.FileInfo, len(fileInfos))
	for idx, fileInfo := range fileInfos {
		fileInfoMap[fileInfo.ID] = fileInfos[idx]
	}

	historyUnfamiliarScore := make(map[int]int)
	notEnough72HoursMap := make(map[int]bool)
	fn := func(examInfo *domain.ExamInfo) {
		if examInfo.Done {
			historyUnfamiliarScore[examInfo.ID] += examInfo.Unfamiliar
		}
		if time.Now().Sub(examInfo.CreateTime).Hours() < 72 {
			notEnough72HoursMap[examInfo.ID] = true
		}
	}
	for _, exam := range exams {
		for _, val := range exam.Easy {
			fn(val)
		}
		for _, val := range exam.Medium {
			fn(val)
		}
		for _, val := range exam.Hard {
			fn(val)
		}
	}

	var fileInfosRemoveNotEnough72Hours []*domain.FileInfo
	for idx, fileInfo := range fileInfos {
		if notEnough72HoursMap[fileInfo.ID] {
			continue
		}
		fileInfosRemoveNotEnough72Hours = append(fileInfosRemoveNotEnough72Hours, fileInfos[idx])
	}

	unfamiliarScores := make([]*scoreWithID, len(fileInfosRemoveNotEnough72Hours))
	for idx, fileInfo := range fileInfosRemoveNotEnough72Hours {
		unfamiliarScore := -1
		if val, ok := historyUnfamiliarScore[fileInfo.ID]; ok {
			unfamiliarScore = val
		}

		unfamiliarScores[idx] = &scoreWithID{
			id:              fileInfo.ID,
			unfamiliarScore: unfamiliarScore,
		}
	}
	for i := range unfamiliarScores { // shuffle unfamiliar score
		j := rand.Intn(i + 1)
		unfamiliarScores[i], unfamiliarScores[j] = unfamiliarScores[j], unfamiliarScores[i]
	}
	sort.SliceStable(unfamiliarScores, func(i, j int) bool { // TODO: use heap
		return unfamiliarScores[i].unfamiliarScore > unfamiliarScores[j].unfamiliarScore
	})

	fileInfoMapByDifficulty := make(map[domain.DifficultyType][]*domain.FileInfo)
	for _, val := range unfamiliarScores {
		fileInfo := fileInfoMap[val.id]
		fileInfo.UnfamiliarScore = val.unfamiliarScore
		fileInfoMapByDifficulty[fileInfo.Difficulty] = append(fileInfoMapByDifficulty[fileInfo.Difficulty], fileInfo)
	}

	var examInstance domain.Exam
	for i := 0; i < easyCount && i < len(fileInfoMapByDifficulty[domain.DifficultyEasy]); i++ {
		examInstance.Easy = append(examInstance.Easy, &domain.ExamInfo{
			ID:   fileInfoMapByDifficulty[domain.DifficultyEasy][i].ID,
			Name: fileInfoMapByDifficulty[domain.DifficultyEasy][i].Name,
		})
	}
	for i := 0; i < mediumCount && i < len(fileInfoMapByDifficulty[domain.DifficultyMedium]); i++ {
		examInstance.Medium = append(examInstance.Medium, &domain.ExamInfo{
			ID:   fileInfoMapByDifficulty[domain.DifficultyMedium][i].ID,
			Name: fileInfoMapByDifficulty[domain.DifficultyMedium][i].Name,
		})
	}
	for i := 0; i < hardCount && i < len(fileInfoMapByDifficulty[domain.DifficultyHard]); i++ {
		examInstance.Hard = append(examInstance.Hard, &domain.ExamInfo{
			ID:   fileInfoMapByDifficulty[domain.DifficultyHard][i].ID,
			Name: fileInfoMapByDifficulty[domain.DifficultyHard][i].Name,
		})
	}

	if err := e.examRepo.Create(&examInstance); err != nil {
		return errors.Wrap(err, "create exam failed")
	}

	return nil
}

func (e *examUseCase) UpdateReadMe() error {
	fileInfos, err := e.fileRepo.ReadAll()
	if err != nil {
		return errors.Wrap(err, "read all files failed")
	}
	exams, err := e.examRepo.ReadAll()
	if err != nil {
		return errors.Wrap(err, "read all exam files failed")
	}

	examsHistoryScoreMap := make(map[int]int)
	for _, exam := range exams {
		for _, val := range exam.Easy {
			if !val.Done {
				continue
			}
			examsHistoryScoreMap[val.ID] += val.Unfamiliar
		}
		for _, val := range exam.Medium {
			if !val.Done {
				continue
			}
			examsHistoryScoreMap[val.ID] += val.Unfamiliar
		}
		for _, val := range exam.Hard {
			if !val.Done {
				continue
			}
			examsHistoryScoreMap[val.ID] += val.Unfamiliar
		}
	}
	for idx, fileInfo := range fileInfos {
		if unfamiliarScore, ok := examsHistoryScoreMap[fileInfo.ID]; ok {
			fileInfos[idx].UnfamiliarScore = fileInfo.Star + unfamiliarScore
		} else {
			fileInfos[idx].UnfamiliarScore = -1
		}
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
		md.WriteString("| Name | Star | Difficulty | Unfamiliar | Tags |" + "\n")
		md.WriteString("| -------- | -------- | -------- | -------- | -------- |" + "\n")
		for _, fileInfo := range fileInfosByTag[tag] {
			md.WriteString("|")
			md.WriteString(fmt.Sprintf("[%d. %s](https://leetcode.com/problems/%s/)", fileInfo.ID, fileInfo.Name, fileInfo.Name))
			md.WriteString("|")
			md.WriteString(starToEmoji(fileInfo.Star))
			md.WriteString("|")
			md.WriteString(fileInfo.Difficulty.String())
			md.WriteString("|")
			if fileInfo.UnfamiliarScore != -1 {
				md.WriteString(strconv.Itoa(fileInfo.UnfamiliarScore))
			}
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
