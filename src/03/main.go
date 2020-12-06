package main

import (
	"common"
	"os"
	"path"
)

func get_encounters(down_pattern int, right_pattern int, input_lines []string) int {
	var cols = len(input_lines[0])
	var rows = len(input_lines)
	encounters := 0
	for counter := 0; counter*down_pattern < rows; counter++ {
		if input_lines[counter] == "" {
			continue
		}
		//encounters = append(encounters, line[3*counter%cols])
		if input_lines[counter*down_pattern][right_pattern*counter%cols] == '#' {
			encounters++
		}
	}
	return encounters
}

func main() {
	wd, _ := os.Getwd()
	act_folder := path.Join(wd, "src", "03")
	input_lines := common.Load_file_content(act_folder)
	encounter_slice := []int{get_encounters(1, 1, input_lines),
		get_encounters(1, 3, input_lines),
		get_encounters(1, 5, input_lines),
		get_encounters(1, 7, input_lines),
		get_encounters(2, 1, input_lines),
	}
	result := 1
	for i := 0; i < len(encounter_slice); i++ {
		result *= encounter_slice[i]
	}
	print(result)
}
