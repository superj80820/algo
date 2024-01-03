package main

import (
	"fmt"
	"os"

	"github.com/superj80820/algo/script/repository"
	"github.com/superj80820/algo/script/usecase"
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

	examRepo := repository.CreateExamRepo("./")
	fileRepo := repository.CreateFileRepo("../neetcode", topicOrderData)
	readMeRepo := repository.CreateReadMeMDRepo("../README.md")

	examUseCase := usecase.CreateExamUseCase(readMeRepo, examRepo, fileRepo)

	switch action {
	case "update-readme":
		examUseCase.UpdateReadMe()
	case "create-exam":
		if err := examUseCase.CreateExam(2, 2, 1); err != nil {
			panic(fmt.Sprintf("%+v", err))
		}
	default:
		panic("no use action argument")
	}
}
