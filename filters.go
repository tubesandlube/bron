package main

import (

	"fmt"

)

// XXX abstract this into two-stage (or more)?
type filterOutputs struct {
	file string
	commitId string
	lines int
	linesCode int
	linesComments int
	linesBlank int
}

func parse(files []string) {

	//c := make(chan result)
	//errc := make(chan error, 1)
	for _, file := range files {
		fmt.Println("going to parse...", file)
		go filterDistribution(file)
		//fmt.Printf("output of dist, %v\n", l)
		//fmt.Println("File:", file, ":", countLines(file))
	}

}
