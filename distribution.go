package main

import (

	"fmt"
	"io/ioutil"

)

func filterDistribution(contentFile string) map[string]int {

	counts := map[string]int{}

// determine language
// strip comments & count lines
// strip white space and count
// count remainder

//	coder, _ := regexp.Compile("^[^\\s]

	file, err := ioutil.ReadFile(contentFile)
	check(err)
	fmt.Println("found", len(file), "characters in file", contentFile)
	//string(file)

	counts["loc"] = len(file)

	
	//return counts

}
