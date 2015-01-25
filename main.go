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

	// XXX testing
	if repoPtr != "" {
		uuidRepo := cloneRepo(repoPtr)

		// XXX temp code to show that the clone worked
		lsCmd := exec.Command("ls", "-a", "-l", uuidRepo)
		lsOut, lsCmdErr := lsCmd.Output()
		check(lsCmdErr)
		fmt.Println(string(lsOut))

		// XXX example calls
		t := getFiles(uuidRepo)
		fmt.Println(t)

		// XXX example calls
		u := getFileContents(t[0])
		fmt.Println(string(u))

		// XXX example calls
		fmt.Println(countFiles(uuidRepo))
		fmt.Println(countLanguages(uuidRepo))

		// XXX example calls
		x, v := getCommits(uuidRepo)
		fmt.Println(x)
		fmt.Println(v)
		fmt.Println(x[0])
		fmt.Println(x[1])
		y := getDiff(uuidRepo, x[0], x[1])
		fmt.Print(string(y))
		fmt.Println("number of commits:", countCommits(uuidRepo))

		// XXX example calls through all commits
		for _, commit := range x {
			checkoutCommit(uuidRepo, commit)
			fmt.Println("number of files;", countFiles(uuidRepo))
			fmt.Println("langs by files:", countLanguages(uuidRepo))
			files := getFiles(uuidRepo)
			for _, file := range files {
				fmt.Println("File:", file, ":", countLines(file))
			}
			fmt.Println("number of lines:", countLinesPerLanguage(uuidRepo))

			// XXX simple channel starts, for now
			parse(files)

		}
		checkoutCommit(uuidRepo, x[0])
		z := countAuthorCommits(uuidRepo)
		fmt.Println(z)
		s := countAuthors(uuidRepo)
		fmt.Println(s)

		// XXX test template parsing
		templates := templateParse("templates")
		fmt.Println(templates)

		if vizPtr {
			chErr := os.Chdir(blessedPtr)
			check(chErr)

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

			//for _, commit := range x {
			//	checkoutCommit(uuidRepo, commit)
			//}

			updateData("dashboards/"+dashboardPtr+"/dashboard.js", "languages", languages)
			updateData("dashboards/"+dashboardPtr+"/dashboard.js", "languageLines", languageLines)

			// XXX fill in '[]' with real data
			updateData("dashboards/"+dashboardPtr+"/dashboard.js", "authors", "[['','']]")
			updateData("dashboards/"+dashboardPtr+"/dashboard.js", "numLanguagesData", "{x:[''],y:['']}")
			updateData("dashboards/"+dashboardPtr+"/dashboard.js", "numLinesData", "{x:[''],y:['']}")
			updateData("dashboards/"+dashboardPtr+"/dashboard.js", "numAuthorsData", "{x:[''],y:['']}")
			updateData("dashboards/"+dashboardPtr+"/dashboard.js", "numFilesData", "{x:[''],y:['']}")

			binary, lookErr := exec.LookPath("node")
			check(lookErr)
			args := []string{"node", "./dashboards/"+dashboardPtr+"/dashboard.js"}
			env := os.Environ()
			execErr := syscall.Exec(binary, args, env)
			check(execErr)
		}
	}

}
