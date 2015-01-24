package main

import (

	//"bufio"
	"fmt"
	//"io"
	"io/ioutil"
	//"os"
	//"regexp"

)

type Template struct {
	Name string
}

func templateParse(templatePath string) map[string]*Template {

	count := 0

	files := getFiles(templatePath)
	for _, file := range files {
		// XXX
		if(file != "templates/example.template" && file != "templates/README.md") {
			count++
			fmt.Println("found a specific template:", file)
		}
	}
	templates := map[string]*Template{}

	fmt.Println("found a few templates:", count)

	for _, file := range files {
		// XXX
		if(file != "templates/example.template" && file != "templates/README.md") {
			t := templateLoad(file)
			templates[t.Name] = t
		}
	}

	return templates

}

func templateLoad(templateFile string) *Template {

	t := new(Template)
	t.Name = templateFile

	dat, err := ioutil.ReadFile(templateFile)
	check(err)
	fmt.Print(string(dat))

	return t

}
