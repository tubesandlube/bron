package main

import (
	"io/ioutil"
	"strings"
)

func updateData(filename string, varName string, data string) {

	input, err := ioutil.ReadFile(filename)
	check(err)

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "$"+varName) {
			lines[i] = "var "+varName+" = "+data
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(filename, []byte(output), 0644)
	check(err)

}
