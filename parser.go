package main

import (

	"bufio"
	"fmt"
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

	files   := getFiles(templatePath)
	r, _    := regexp.Compile("^templates/\\w+\\.template$")
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

	sectionr, _ := regexp.Compile("^[\\w\\s]+\\:$")
	//cleanr, _   := regexp.Compile("[^\\w+]")
	//paramr, _   := regexp.Compile("^\\s+(\\w+)\\: (.*)$")

	file, err := os.Open(templateFile)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if(sectionr.MatchString(scanner.Text())) {
			section := scanner.Text()[:len(scanner.Text())-1]
			fmt.Println("found section", section, "from file", templateFile)
		} else {
			parts := regexp.MustCompile(": ").Split(scanner.Text(), 2)
			if(len(parts) > 1) {
				fmt.Println("adding content match", parts[1], "to param", parts[0])
			} else {
				if(len(parts[0]) > 0) {
					fmt.Println("adding single param", parts[0])
				} else {
					//fmt.Println("skipping blank line", parts[0])
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		check(err)
	}

	//fmt.Printf("template has this struct: %v\n", t)

	return t

}
