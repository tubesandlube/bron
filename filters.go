package main

import (

	"fmt"

)

func parse(files []string, templates map[string]*Template) {

	for _, file := range files {
		d := "going to parse..." + file
		fmt.Println(colorize(d))
		filterDistribution(templates, file)
	}
	fmt.Println("all done")

}
