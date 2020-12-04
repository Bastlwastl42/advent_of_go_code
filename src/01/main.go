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
		if e_int == 0 {
			continue
		}
		integers = append(integers, e_int)
	}
	//print(integers)
	for _, e := range integers {
		remainder := 2020 - e
		//print(remainder)
		_, is_found := Find(integers, remainder)
		if is_found {
			fmt.Printf("Found match with %d and %d. Product is %d\n", remainder, e, remainder*e)
			break
		}
	}
	break_flag := false
	for _, e := range integers {
		for _, f := range integers {
			if e+f > 2020 {
				continue
			}
			for _, g := range integers {
				if e+f+g == 2020 {
					fmt.Printf("Found match with %d, %d, and %d. Product is %d\n",
						e, f, g, e*f*g)
					break_flag = true
					break
				}

			}
			if break_flag {
				break
			}
		}
		if break_flag {
			break
		}

	}
}
