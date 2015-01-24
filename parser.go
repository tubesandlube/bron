package main

import (

	"bufio"
	"fmt"
	//"io"
	//"io/ioutil"
	"os"
	"regexp"

)

type Template struct {
	Name string
	Extensions []string
	Structures map[string]string
	Types map[string]string
	Names map[string]string
	Keywords map[string]string
	Expressions map[string]string
	Libraries map[string]string
}

func templateParse(templatePath string) map[string]*Template {

	count := 0

	files := getFiles(templatePath)
	r, _ := regexp.Compile("^templates/\\w+\\.template$")
	skip, _ := regexp.Compile("^templates/example\\.template$")

	for _, file := range files {
		// XXX
		if(r.MatchString(file) && ! skip.MatchString(file)) {
			count++
			fmt.Println("found a specific template:", file)
		}
	}
	// XXX fix handling count
	templates := map[string]*Template{}

	fmt.Println("found a few templates:", count)

	for _, file := range files {
		// XXX
		if(r.MatchString(file) && ! skip.MatchString(file)) {
			t := templateLoad(file)
			templates[t.Name] = t
		}
	}

	return templates

}

func templateLoad(templateFile string) *Template {

	t := new(Template)
	t.Name = templateFile

	section, _ := regexp.Compile("^\\w+\\:$")

	file, err := os.Open(templateFile)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if(section.MatchString(scanner.Text())) {
			section := scanner.Text()[:len(scanner.Text())-1]
			fmt.Println("found section", section, "from file", templateFile)
		}
	}

	if err := scanner.Err(); err != nil {
		check(err)
	}

	//dat, err := ioutil.ReadFile(templateFile)
	//check(err)
	
	//fmt.Println("whole file:", dat)
	//fmt.Printf("template has this struct: %v\n", t)

	return t

}
