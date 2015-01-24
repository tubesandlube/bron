package main

import (
	"flag"
	"fmt"
	"os/exec"

	//"github.com/gophergala/bron/filters"
)

var (
	repoPtr     string
	repoPathPtr string
	verbosePtr  int
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	flag.StringVar(&repoPtr, "repo", "", "Git repository to scan")
	flag.StringVar(&repoPathPtr, "path", "", "Git repository file path to scan")
	flag.IntVar(&verbosePtr, "v", 1, "verbosity level")

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

		// XXX example calls through all commits
		for _, commit := range x {
			checkoutCommit(uuidRepo, commit)
			fmt.Println(countFiles(uuidRepo))
			fmt.Println(countLanguages(uuidRepo))
		}

		// XXX test template parsing
		templates := templateParse("templates")
		fmt.Println(templates)
	}

}
