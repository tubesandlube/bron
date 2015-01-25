package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"

//	"github.com/gophergala/bron/filters"
)

var (
	blessedPtr   string
	dashboardPtr string
	repoPtr      string
	repoPathPtr  string
	verbosePtr   int
	vizPtr       bool
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	flag.StringVar(&blessedPtr, "blessedPath", "/go/src/github.com/yaronn/blessed-contrib", "Path where blessed-contrib is installed")
	flag.StringVar(&dashboardPtr, "dashboard", "example", "Name of dashboard to use for visualization")
	flag.StringVar(&repoPtr, "repo", "", "Git repository to scan")
	flag.StringVar(&repoPathPtr, "path", "", "Git repository file path to scan")
	flag.IntVar(&verbosePtr, "v", 1, "verbosity level")
	flag.BoolVar(&vizPtr, "viz", false, "Visualize the results, requires blessed")

	flag.Parse()

	if repoPtr == "" && repoPathPtr == "" {
		fmt.Println("please specify either a repo or a path to a git repo to scan")
	} else if repoPtr != "" && repoPathPtr != "" {
		fmt.Println("please specify only either a repo or a path to a git repo to scan, not both")
	} else {
		if verbosePtr > 0 {
			fmt.Println("going to scan repository", repoPtr, "...")
		}
	}

	if repoPtr != "" {
		uuidRepo := cloneRepo(repoPtr)

		// XXX example calls through all commits
		x, _ := getCommits(uuidRepo)
		for _, commit := range x {
			checkoutCommit(uuidRepo, commit)
			// XXX simple channel starts, for now
			files := getFiles(uuidRepo)
			parse(files)
		}
		checkoutCommit(uuidRepo, x[0])

		// XXX test template parsing
		templates := templateParse("templates")
		fmt.Println(templates)

		if vizPtr {
			// get data for dashboard
			languages := "["
			languageLines := "["
			languageMap := countLinesPerLanguage(uuidRepo)
			for key := range languageMap {
				languages += "'"+key+"', "
				languageLines += "'"+strconv.Itoa(languageMap[key])+"', "
			}
			languages = languages[0:len(languages)-2]+"]"
			languageLines = languageLines[0:len(languageLines)-2]+"]"

			authorMap := countAuthorCommits(uuidRepo)
			authors := "["
			for key := range authorMap {
				authors += "['"+key+"', '"+strconv.Itoa(authorMap[key])+"'], "
			}
			authors = authors[0:len(authors)-2]+"]"

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
	}

}
