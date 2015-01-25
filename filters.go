package main

import (

	"fmt"

)

func parse(files []string) {

	for _, file := range files {
		fmt.Println("going to parse...", file)
		filterDistribution(file)
	}
	fmt.Println("all done")

}
