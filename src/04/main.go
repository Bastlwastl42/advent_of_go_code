package main

import (
	"common"
	"os"
	"path"
	"strconv"
	"strings"
)

type validatedPassportEntry interface {
	convertToField(passport_entry)
	validateField()
}

type passport_entry struct {
	raw_entry          string
	entry_valid        bool
	has_been_validated bool
}

func (act_passport_entry passport_entry) convertRawString(line string) {
	act_passport_entry.raw_entry = strings.Split(line, ":")[1]
}

func (act_passport_entry passport_entry) getRawString() string {
	return act_passport_entry.raw_entry
}

func (act_passport_entry passport_entry) validateEntry(is_valid bool) {
	act_passport_entry.entry_valid = is_valid
	act_passport_entry.has_been_validated = true
}

type hcl struct {
	passport_entry
	hair_color string ""
	first_val  byte
}

func (act_hcl hcl) convertToField(base_passport_entry passport_entry) {
	act_hcl.passport_entry = base_passport_entry
	act_hcl.first_val = base_passport_entry.getRawString()[0]
	act_hcl.hair_color = base_passport_entry.getRawString()[1:]
}

func (act_hcl hcl) validateField() {
	_, err := strconv.ParseInt(act_hcl.hair_color, 16, 0)
	if err == nil {
		act_hcl.passport_entry.validateEntry(true)
	} else {
		act_hcl.passport_entry.validateEntry(false)
	}
}

type ecl struct {
	passport_entry
	eye_color string
}

func (act_ecl ecl) convertToField(base_passport_entry passport_entry) {
	act_ecl.passport_entry = base_passport_entry
	act_ecl.eye_color = base_passport_entry.getRawString()
}

func (act_ecl ecl) validateField() {}

type iyr struct {
	passport_entry
	issue_year int
}

func (act_iyr iyr) convertToField(base_passport_entry passport_entry) {
	act_iyr.passport_entry = base_passport_entry
	act_iyr.issue_year, _ = strconv.Atoi(base_passport_entry.getRawString())
}

func (act_iyr iyr) validateField() {}

type pid struct {
	passport_entry
	pid int
}

func (act_pid pid) convertToField(base_passport_entry passport_entry) {
	act_pid.passport_entry = base_passport_entry
	act_pid.pid, _ = strconv.Atoi(base_passport_entry.getRawString())
}

func (act_pid pid) validateField() {}

type hgt struct {
	passport_entry
	height int
	unit   string
}

func (act_hgt hgt) convertToField(base_passport_entry passport_entry) {
	act_hgt.passport_entry = base_passport_entry
	heigt_string := base_passport_entry.getRawString()
	if string(heigt_string[len(heigt_string)-2]) == "cm" {
		// height given in cm
		act_hgt.unit = "cm"
		act_hgt.height, _ = strconv.Atoi(heigt_string[:len(heigt_string)-2])
	} else if string(heigt_string[len(heigt_string)-2]) == "in" {
		// height given in cm
		act_hgt.unit = "in"
		act_hgt.height, _ = strconv.Atoi(heigt_string[:len(heigt_string)-2])
	} else {
		act_hgt.passport_entry.has_been_validated = true
		act_hgt.passport_entry.entry_valid = false
	}

}

func (act_hgt hgt) validateField() {}

type eyr struct {
	passport_entry
	expiration_year int
}

func (act_eyr eyr) convertToField(base_passport_entry passport_entry) {
	act_eyr.passport_entry = base_passport_entry
	act_eyr.expiration_year, _ = strconv.Atoi(base_passport_entry.getRawString())
}

func (act_eyr eyr) validateField() {}

type byr struct {
	passport_entry
	year_of_birth int
}

func (act_byr byr) convertToField(base_passport_entry passport_entry) {
	act_byr.passport_entry = base_passport_entry
	act_byr.year_of_birth, _ = strconv.Atoi(base_passport_entry.getRawString())
}

func (act_byr byr) validateField() {}

type cid struct {
	passport_entry
	citizien_id int
}

func (act_cid cid) convertToField(base_passport_entry passport_entry) {
	act_cid.passport_entry = base_passport_entry
	act_cid.citizien_id, _ = strconv.Atoi(base_passport_entry.getRawString())
}

func (act_cid cid) validateField() {}

type passport struct {
	raw_entries   map[int]string
	final_entries map[string]validatedPassportEntry
}

func (act_pass passport) determine_entry_type(field_name string, genericPassportEntry passport_entry) {
	var new_entry validatedPassportEntry
	switch field_name {
	case "hcl":
		new_entry = hcl{}
	case "ecl":
		new_entry = ecl{}
	case "iyr":
		new_entry = iyr{}
	case "pid":
		new_entry = pid{}
	case "hgt":
		new_entry = hgt{}
	case "eyr":
		new_entry = eyr{}
	case "byr":
		new_entry = byr{}
	case "cid":
		new_entry = cid{}
	}
	new_entry.convertToField(genericPassportEntry)
	new_entry.validateField()
	act_pass.final_entries[field_name] = new_entry

}

func main() {
	wd, _ := os.Getwd()
	act_folder := path.Join(wd, "src", "04")
	input_lines := common.Load_file_content(act_folder)
	tobesplit := strings.Join(input_lines[:], "\n")
	final_input_lines := strings.Split(tobesplit, "\n\n")
	for _, entry := range final_input_lines {
		var act_pass passport
		act_pass.raw_entries = make(map[int]string)
		split_entry := strings.Split(strings.ReplaceAll(entry, "\n", " "), " ")
		for counter, e := range split_entry {
			act_pass.raw_entries[counter] = e
			var generic_passport_entry passport_entry
			field_name := strings.Split(e, ":")[0]
			generic_passport_entry.raw_entry = strings.Split(e, ":")[1]
			generic_passport_entry.has_been_validated = false
			generic_passport_entry.entry_valid = false
			act_pass.determine_entry_type(field_name, generic_passport_entry)
		}
	}
}
