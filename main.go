package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"code.google.com/p/go-uuid/uuid"
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

	uuidRepo := uuid.New()

	// XXX don't forget to cleanup after we're finished
	err := os.Mkdir("/tmp/"+uuidRepo, 0644)
	check(err)

	// XXX this should move up into the initial if/else
	if repoPtr != "" {
		cloneCmd := exec.Command("git", "clone", repoPtr, "/tmp/"+uuidRepo)
		cloneOut, cloneErr := cloneCmd.Output()
		check(cloneErr)
		fmt.Println(string(cloneOut))

		// XXX temp code to show that the clone worked
		lsCmd := exec.Command("ls", "-a", "-l", "/tmp/"+uuidRepo)
		lsOut, lsCmdErr := lsCmd.Output()
		check(lsCmdErr)
		fmt.Println(string(lsOut))

		// XXX example calls
		t := getFiles("/tmp/"+uuidRepo)
		fmt.Println(t)

		// XXX example calls
		u := getFileContents(t[0])
		fmt.Println(string(u))
	}

}
