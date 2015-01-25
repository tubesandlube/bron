package main

import (

	"fmt"
	"io/ioutil"

)

//func filterDistribution(contentFile string, wg *sync.WaitGroup) map[string]int {
//func filterDistribution(contentFile string, c chan string, wg *sync.WaitGroup) {
//func filterDistribution(contentFile string, c chan string) {
//func filterDistribution(contentFile string, in <-chan *string, out chan<- *string) {

func filterDistribution(contentFile string) map[string]int {

	counts := map[string]int{}

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
