package repository

import (
	"fmt"
	"os"
	"strings"

	"github.com/superj80820/algo/script/domain"
)

const (
	mdLeetcodeListStart = "<!-- leetcode list start -->"
	mdLeetCodeListEnd   = "<!-- leetcode list end -->"
)

type readMeMDRepo struct {
	readMeFilePath string
}

func CreateReadMeMDRepo(readMeFilePath string) domain.ReadMeMDRepo {
	return &readMeMDRepo{
		readMeFilePath: readMeFilePath,
	}
}

func (r *readMeMDRepo) Write(content string) error {
	content = fmt.Sprintf("\n%s\n", mdLeetcodeListStart) + content + mdLeetCodeListEnd

	data, err := os.ReadFile(r.readMeFilePath)
	if err != nil {
		panic(err)
	}
	mdFile := string(data)
	newMDFile := mdFile[:strings.Index(mdFile, ""+mdLeetcodeListStart+"")-1] +
		content +
		mdFile[strings.Index(mdFile, mdLeetCodeListEnd)+len(mdLeetCodeListEnd):]

	os.WriteFile(r.readMeFilePath, []byte(newMDFile), 0644)
	return nil
}
