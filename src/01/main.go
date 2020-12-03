package main

import (
	"common"
	"fmt"
	"os"
	"path"
	"strconv"
)

func Find(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func main() {
	act_folder, _ := os.Getwd()
	input_lines := common.Load_file_content(path.Join(act_folder, "src", "01"))
	//fmt.Println(input_lines)
	var integers = []int{}
	for _, e := range input_lines {
		//make integers
		e_int, _ := strconv.Atoi(e)
		integers = append(integers, e_int)
	}
	//print(integers)
	for _, e := range integers {
		remainder := 2020 - e
		//print(remainder)
		_, is_found := Find(integers, remainder)
		if is_found {
			fmt.Printf("Found match with %d and %d. Product is %d", remainder, e, remainder*e)
			break
		}
	}
	for _, e := range integers{
		first_remainder := 2020 -e
		
	}
}
