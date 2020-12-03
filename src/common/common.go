package common

import (
	"io/ioutil"
	"path"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Load_file_content(folder string) []string{
	content, err := ioutil.ReadFile(path.Join(folder, "input.txt"))
	check(err)
	//input is expected to be multiple line of string
	return strings.Split(string(content), "\n")
}
