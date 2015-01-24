package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"code.google.com/p/go-uuid/uuid"
	"github.com/kr/fs"
)

func getFiles(repoPath string) []string {

	files := []string{}

	walker := fs.Walk(repoPath)
	for walker.Step() {
		if err := walker.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		if !strings.Contains(walker.Path(), ".git") && strings.Contains(walker.Path(), ".") {
			files = append(files, walker.Path())
		}
	}

	return files

}

func getFileContents(filename string) []byte {

	data, err := ioutil.ReadFile(filename)
	check(err)

	return data

}

func cloneRepo(repo string) string {

	uuidRepo := uuid.New()

	// XXX don't forget to cleanup after we're finished
	err := os.Mkdir("/tmp/"+uuidRepo, 0644)
	check(err)

	// XXX this should move up into the initial if/else
	if repo != "" {
		cloneCmd := exec.Command("git", "clone", repo, "/tmp/"+uuidRepo)
		cloneOut, cloneErr := cloneCmd.Output()
		check(cloneErr)
		fmt.Println(string(cloneOut))
	}

	return "/tmp/"+uuidRepo
}
