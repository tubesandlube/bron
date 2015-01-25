package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func checkData(repoName string) bool {

	// XXX stub

	return false

}

func saveData(repoName string, vals ...string) {

	// XXX stub

}

func updateData(filename string, varName string, data string) {

	input, err := ioutil.ReadFile(filename)
	check(err)

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "$"+varName) {
			lines[i] = "var "+varName+" = "+data
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(filename, []byte(output), 0644)
	check(err)

}

func tableData(rows map[string]int) string {

	table := "["

	for key := range rows {
		table += "['"+key+"', '"+strconv.Itoa(rows[key])+"'], "
	}
	table = table[0:len(table)-2]+"]"

	return table

}

func barChartData(bars map[string]int) (string, string) {

	x := "["
	y := "["

	for key := range bars {
		x += "'"+key+"', "
		y += "'"+strconv.Itoa(bars[key])+"', "
	}
	x = x[0:len(x)-2]+"]"
	y = y[0:len(y)-2]+"]"

	return x, y

}

func updateDashboardData(uuidRepo string, repoPtr string) {

	// get data for dashboard
	languages, languageLines := barChartData(countLinesPerLanguage(uuidRepo))
	authors := tableData(countAuthorCommits(uuidRepo))

	// XXX cleanup line charts
	numLanguagesDataX := "x:["
	numLanguagesDataY := "y:["
	numLinesDataX := "x:["
	numLinesDataY := "y:["
	numAuthorsDataX := "x:["
	numAuthorsDataY := "y:["
	numFilesDataX := "x:["
	numFilesDataY := "y:["

	x, y := getCommits(uuidRepo)
	// XXX x needs to be reversed, note don't simply sort and reverse, order matters
	for _, commit := range x {
		checkoutCommit(uuidRepo, commit)

		lineCount := 0
		files := getFiles(uuidRepo)
		for _, file := range files {
			lineCount += countLines(file)
		}
		languageCount := 0
		langMap := countLanguages(uuidRepo)
		for key := range langMap {
			languageCount += langMap[key]
		}

		numLanguagesDataX += "'"+y[commit]["timestamp"]+"', "
		numLanguagesDataY += "'"+strconv.Itoa(languageCount)+"', "
		numLinesDataX += "'"+y[commit]["timestamp"]+"', "
		numLinesDataY += "'"+strconv.Itoa(lineCount)+"', "
		numAuthorsDataX += "'"+y[commit]["timestamp"]+"', "
		numAuthorsDataY += "'"+strconv.Itoa(countAuthorsByCommits(uuidRepo, commit))+"', "
		numFilesDataX += "'"+y[commit]["timestamp"]+"', "
		numFilesDataY += "'"+strconv.Itoa(countFiles(uuidRepo))+"', "
	}
	checkoutCommit(uuidRepo, x[0])

	numLanguagesData := "{"+numLanguagesDataX[0:len(numLanguagesDataX)-2]+"], "+numLanguagesDataY[0:len(numLanguagesDataY)-2]+"]"+"}"
	numLinesData := "{"+numLinesDataX[0:len(numLinesDataX)-2]+"], "+numLinesDataY[0:len(numLinesDataY)-2]+"]"+"}"
	numAuthorsData := "{"+numAuthorsDataX[0:len(numAuthorsDataX)-2]+"], "+numAuthorsDataY[0:len(numAuthorsDataY)-2]+"]"+"}"
	numFilesData := "{"+numFilesDataX[0:len(numFilesDataX)-2]+"], "+numFilesDataY[0:len(numFilesDataY)-2]+"]"+"}"

	saveData(repoPtr, languages, languageLines, authors, numLanguagesData, numLinesData, numAuthorsData, numFilesData)

	chErr := os.Chdir(blessedPtr)
	check(chErr)
	updateData("dashboards/"+dashboardPtr+"/dashboard.js", "languages", languages)
	updateData("dashboards/"+dashboardPtr+"/dashboard.js", "languageLines", languageLines)
	updateData("dashboards/"+dashboardPtr+"/dashboard.js", "authors", authors)
	updateData("dashboards/"+dashboardPtr+"/dashboard.js", "numLanguagesData", numLanguagesData)
	updateData("dashboards/"+dashboardPtr+"/dashboard.js", "numLinesData", numLinesData)
	updateData("dashboards/"+dashboardPtr+"/dashboard.js", "numAuthorsData", numAuthorsData)
	updateData("dashboards/"+dashboardPtr+"/dashboard.js", "numFilesData", numFilesData)

	binary, lookErr := exec.LookPath("node")
	check(lookErr)
	args := []string{"node", "./dashboards/"+dashboardPtr+"/dashboard.js"}
	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	check(execErr)

}

func showDashboard() {

	chErr := os.Chdir(blessedPtr)
	check(chErr)
	binary, lookErr := exec.LookPath("node")
	check(lookErr)
	args := []string{"node", "./dashboards/"+dashboardPtr+"/dashboard.js"}
	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	check(execErr)

}
