package main

import (
	"flag"
	"fmt"

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
			updateDashboardData(uuidRepo)
		}
	}

}
