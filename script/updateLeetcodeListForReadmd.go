package main

import (
	"fmt"
	"strings"

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
	fileHandler := file.CreateFileHandler(topicOrderData, "../neetcode", "../README.md")

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
			md.WriteString(fileInfo.Difficulty)
			md.WriteString("|")
			md.WriteString(strings.Join(fileInfo.OtherTags, ", "))
			md.WriteString("|\n")
		}
	}

	fileHandler.WriteReadMe(md.String())
}
