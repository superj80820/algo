package repository

import (
	"encoding/json"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/superj80820/algo/script/domain"
)

type examRepo struct {
	examFilePath string
}

func CreateExamRepo(examFilePath string) domain.ExamRepo {
	return &examRepo{
		examFilePath: examFilePath,
	}
}

func (e *examRepo) Create(exam *domain.Exam) error {
	fileName := "exam" + "-" + time.Now().Format("2006-01-02")
	jsonData, err := json.MarshalIndent(exam, "", "\t")
	if err != nil {
		return errors.Wrap(err, "json marshal failed")
	}
	if err := os.WriteFile(e.examFilePath+"/"+fileName+".json", jsonData, 0644); err != nil {
		return errors.Wrap(err, "write file failed")
	}
	return nil
}

func (e *examRepo) ReadAll() ([]*domain.Exam, error) {
	files, err := os.ReadDir(e.examFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "read files failed")
	}
	var exams []*domain.Exam
	for _, file := range files {
		fileName := file.Name()
		if !(len(fileName) >= 5 && fileName[:5] == "exam-") {
			continue
		}
		data, err := os.ReadFile(e.examFilePath + "/" + fileName)
		if err != nil {
			return nil, errors.Wrap(err, "read files failed")
		}
		var exam domain.Exam
		if err := json.Unmarshal(data, &exam); err != nil {
			return nil, errors.Wrap(err, "json unmarshal failed")
		}

		exams = append(exams, &exam)
	}
	return exams, nil
}

// fn := func(val *domain.ExamInfo, createTime time.Time) {
// 	examInfos = append(examInfos, val)
// }
// for _, val := range exam.Easy {
// 	fn(val, exam.CreateTime)
// }
// for _, val := range exam.Medium {
// 	fn(val, exam.CreateTime)
// }
// for _, val := range exam.Hard {
// 	fn(val, exam.CreateTime)
// }
