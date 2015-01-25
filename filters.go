package main

import (

	"fmt"

)

func parse(files []string) {

	for _, file := range files {
		d := "going to parse..." + file
		fmt.Println(colorize(d))
		filterDistribution(file)
	}
	fmt.Println("all done")

}
