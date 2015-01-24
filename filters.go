package main

import (

	"fmt"

)

func parse(files []string) {

	for _, file := range files {
		fmt.Println("going to parse...", file)
		l := filterDistribution(file)
		fmt.Printf("output of dist, %v\n", l)
		//fmt.Println("File:", file, ":", countLines(file))
	}

}
