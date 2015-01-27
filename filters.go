package main

import (
	"fmt"
)

func parse(files []string, templates map[string]*Template, flagBools []bool) {

	for _, file := range files {
		d := "going to parse..." + file
		if !flagBools[0] && flagBools[2] {
			fmt.Println(colorize(d))
		}
		filterDistribution(templates, file, flagBools)
	}
	if !flagBools[0] && flagBools[2] {
		fmt.Println("all done")
	}

}
