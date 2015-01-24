package main

import (
	"strings"
)

func countFiles(repoPath string) int {

	files := getFiles(repoPath)

	return len(files)

}

func countLanguages(repoPath string) map[string]int {

	// XXX tie-into languages defined in templates

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

func countLines(file string) int {

	data := getFileContents(file)
	lines := strings.Split(string(data), "\n")

	return len(lines)

}

func countLinesPerLanguage(repoPath string) map[string]int {

	counts := map[string]int{}
	files := getFiles(repoPath)
	for _, file := range files {
		ext := strings.Split(file, ".")
		if _, ok := counts[ext[len(ext)-1]]; ok {
			counts[ext[len(ext)-1]] = counts[ext[len(ext)-1]]+countLines(file)
		} else {
			counts[ext[len(ext)-1]] = countLines(file)
		}
	}

	return counts

}

func countCommits(repoPath string) int {

	// XXX stub
	count := 0

	return count
}

func countAuthors(repoPath string) int {

	// XXX stub
	count := 0

	return count

}

func countAuthorCommits(repoPath string) map[string]int {

	// XXX stub
	counts := map[string]int{}

	return counts

}

func countAuthorLines(repoPath string) map[string]int {

	// XXX stub
	counts := map[string]int{}

	return counts

}
