package main

import (
	"common"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func main() {
	wd, _ := os.Getwd()
	act_folder := path.Join(wd, "src", "02")
	input_lines := common.Load_file_content(act_folder)
	//print(input_lines)
	valid_founds := 0
	valid_founds_part_two := 0
	for _, line := range input_lines {
		if line == "" {
			continue
		}
		colon_split := strings.Split(line, ": ")
		value_part := colon_split[1]
		rule_part := colon_split[0]
		space_split := strings.Split(rule_part, " ")
		rule_letter := space_split[1]
		rule_range := strings.Split(space_split[0], "-")
		rule_range_min, _ := strconv.Atoi(rule_range[0])
		rule_range_max, _ := strconv.Atoi(rule_range[1])
		occourence := 0
		for _, val := range value_part {
			if string(val) == rule_letter {
				occourence++
			}
		}
		if occourence >= rule_range_min && occourence <= rule_range_max {
			valid_founds++
		}
		if len(value_part) >= rule_range_max {
			if (string(value_part[rule_range_max-1]) == rule_letter) != (string(value_part[rule_range_min-1]) == rule_letter) {
				valid_founds_part_two++
			}
		}
	}
	fmt.Printf("Found %d valid passwords\n", valid_founds)
	fmt.Printf("Found %d valid passwords in part two.\n", valid_founds_part_two)
}
