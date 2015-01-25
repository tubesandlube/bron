package main

import (

	"fmt"
	"io/ioutil"

)

func filterDistribution(templates map[string]*Template, contentFile string) map[string]int {

	counts := map[string]int{}

	language := determineLanguage(templates, contentFile)

	if language != "unknown" {
		fmt.Println("language determined as", language)
	} else {
		fmt.Println("skipping loc filtering for", contentFile, " due to not being able to determine language.")
	}

// determine language
// strip comments & count lines
// strip white space and count
// count remainder

	//coder, _ := regexp.Compile("^[^\\s]

	file, err := ioutil.ReadFile(contentFile)
	check(err)
	fmt.Println("found", len(file), "characters in file", contentFile)

	counts["loc"] = len(file)

	return counts

}
