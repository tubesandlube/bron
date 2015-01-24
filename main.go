package main

import (
	"fmt"
	"flag"

	//"github.com/gophergala/bron/filters"
)

func main() {

	repoPtr    := flag.String("repo", "https://github.com/gophergala/bron", "Git repository to scan")
	verbosePtr := flag.Int("v", 1, "verbosity level")

	flag.Parse()

	if *verbosePtr > 0 {
		fmt.Println("going to scan repository", *repoPtr, "...")
	}

}
