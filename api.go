package main

import (
	"strings"
)

func countFiles(repoPath string) int {

	files := getFiles(repoPath)

	return len(files)

}

func countLanguages(repoPath string) map[string]int {

	languages := map[string]int{}
	files := getFiles(repoPath)

	for _, file := range files {
		ext := strings.Split(file, ".")
		if _, ok := languages[ext[len(ext)-1]]; ok {
			languages[ext[len(ext)-1]] = languages[ext[len(ext)-1]]+1
		} else {
			languages[ext[len(ext)-1]] = 1
		}
	}

	return languages

}
