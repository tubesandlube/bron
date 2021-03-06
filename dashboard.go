package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func checkData(flagStrings []string) bool {

	data, err := ioutil.ReadFile("db/" + flagStrings[2] + "/" + flagStrings[1] + ".data")
	if err != nil {
		return false
	}
	lines := strings.Split(string(data), "\n")

	chErr := os.Chdir(flagStrings[0])
	if chErr != nil {
		return false
	}
	for _, line := range lines {
		if line != "\n" && line != "" {
			dt := strings.Split(line, "|")
			if len(dt) < 2 {
				return false
			}
			updateData("dashboards/"+flagStrings[1]+"/dashboard.js", dt[0], dt[1])
		}
	}

	return true

}

func saveData(flagStrings []string, vals ...string) {

	err := os.MkdirAll("db/"+flagStrings[2], 0644)
	check(err)

	data := ""
	for _, val := range vals {
		data += val + "\n"
	}
	d1 := []byte(data)

	err = ioutil.WriteFile("db/"+flagStrings[2]+"/"+flagStrings[1]+".data", d1, 0644)
	check(err)

}

func updateData(filename string, varName string, data string) {

	input, err := ioutil.ReadFile(filename)
	check(err)

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "$"+varName) {
			lines[i] = "var " + varName + " = " + data
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(filename, []byte(output), 0644)
	check(err)

}

func tableData(rowVals []int, invRows map[int]string) string {

	table := "["

	for i := len(rowVals) - 1; i >= 0; i-- {
		table += "['" + strings.Replace(invRows[rowVals[i]], "'", "\\'", -1) + "', '" + strconv.Itoa(rowVals[i]) + "'], "
	}
	table = table[0:len(table)-2] + "]"

	return table

}

func barChartData(barVals []int, invBars map[int]string) (string, string) {

	x := "["
	y := "["

	for _, k := range barVals {
		x += "'" + strings.Replace(invBars[k], "'", "\\'", -1) + "', "
		y += "'" + strconv.Itoa(k) + "', "
	}
	x = x[0:len(x)-2] + "]"
	y = y[0:len(y)-2] + "]"

	return x, y

}

func bubbleSort(arr []int) []int {

	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr

}

func sortMap(m map[string]int) ([]int, map[int]string) {

	// inverting map
	invMap := make(map[int]string, len(m))
	for k, v := range m {
		invMap[v] = k
	}

	// sorting
	sortedKeys := make([]int, len(invMap))
	i := 0
	for k := range invMap {
		sortedKeys[i] = k
		i++
	}

	return bubbleSort(sortedKeys), invMap

}

func updateDashboardData(uuidRepo string, flagStrings []string, flagBools []bool) {

	// get data for dashboard
	if (!flagBools[0] && flagBools[2]) || flagBools[1] {
		fmt.Printf("\rprocessing languages ...")
	}
	languages, languageLines := barChartData(sortMap(countLinesPerLanguage(uuidRepo)))
	if (!flagBools[2] && flagBools[2]) || flagBools[1] {
		fmt.Printf("\rprocessing languages ... done.\n")
		fmt.Printf("\rprocessing authors ...")
	}
	authors := tableData(sortMap(countAuthorCommits(uuidRepo)))
	if (!flagBools[0] && flagBools[2]) || flagBools[1] {
		fmt.Printf("\rprocessing authors ... done.\n")
	}

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
	for i := len(x) - 1; i >= 0; i-- {
		if (!flagBools[0] && flagBools[2]) || flagBools[1] {
			var percent float64
			percent = float64(len(x)) / float64(100)
			if percent > 0 {
				fmt.Printf("\rprocessing commits ... %.2g%% complete", float64((len(x)-i))/percent)
			}
		}
		checkoutCommit(uuidRepo, x[i])

		lineCount := 0
		files := getFiles(uuidRepo)

		templates := templateParse("templates", flagBools)
		parse(files, templates, flagBools)

		for _, file := range files {
			lineCount += countLines(file)
		}
		languageCount := 0
		langMap := countLanguages(uuidRepo)
		for key := range langMap {
			languageCount += langMap[key]
		}

		numLanguagesDataX += "'" + y[x[i]]["timestamp"] + "', "
		numLanguagesDataY += "'" + strconv.Itoa(languageCount) + "', "
		numLinesDataX += "'" + y[x[i]]["timestamp"] + "', "
		numLinesDataY += "'" + strconv.Itoa(lineCount) + "', "
		numAuthorsDataX += "'" + y[x[i]]["timestamp"] + "', "
		numAuthorsDataY += "'" + strconv.Itoa(countAuthorsByCommits(uuidRepo, x[i])) + "', "
		numFilesDataX += "'" + y[x[i]]["timestamp"] + "', "
		numFilesDataY += "'" + strconv.Itoa(countFiles(uuidRepo)) + "', "
	}
	if (!flagBools[0] && flagBools[2]) || flagBools[1] {
		fmt.Printf("\rprocessing commits ... 100.00%% complete      ")
		fmt.Println()
	}
	checkoutCommit(uuidRepo, x[0])

	numLanguagesData := "{" + numLanguagesDataX[0:len(numLanguagesDataX)-2] + "], " + numLanguagesDataY[0:len(numLanguagesDataY)-2] + "]" + "}"
	numLinesData := "{" + numLinesDataX[0:len(numLinesDataX)-2] + "], " + numLinesDataY[0:len(numLinesDataY)-2] + "]" + "}"
	numAuthorsData := "{" + numAuthorsDataX[0:len(numAuthorsDataX)-2] + "], " + numAuthorsDataY[0:len(numAuthorsDataY)-2] + "]" + "}"
	numFilesData := "{" + numFilesDataX[0:len(numFilesDataX)-2] + "], " + numFilesDataY[0:len(numFilesDataY)-2] + "]" + "}"

	saveData(flagStrings, "languages|"+languages, "languageLines|"+languageLines, "authors|"+authors, "numLanguagesData|"+numLanguagesData, "numLinesData|"+numLinesData, "numAuthorsData|"+numAuthorsData, "numFilesData|"+numFilesData)

	if flagBools[3] {
		checkData(flagStrings)
		showDashboard(flagStrings)
	}
}

func showDashboard(flagStrings []string) {

	chErr := os.Chdir(flagStrings[0])
	check(chErr)
	binary, lookErr := exec.LookPath("node")
	check(lookErr)
	args := []string{"node", "./dashboards/" + flagStrings[1] + "/dashboard.js"}
	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	check(execErr)

}
