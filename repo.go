package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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
		if !strings.Contains(walker.Path(), ".git") {
			files = append(files, walker.Path())
		}
	}

	return files[1:]

}

func getFileContents(filename string) []byte {

	data, err := ioutil.ReadFile(filename)
	check(err)

	return data

}
