package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var heighRegex = regexp.MustCompile(`^(\d){2,3}((cm)|in)$`)
var colorRegex = regexp.MustCompile(`^#[a-f0-9]{6}$`)
var idRegex = regexp.MustCompile(`^[0-9]{9}$`)
var eyeColors = map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}

type Passport struct {
	byr,
	iyr,
	eyr,
	hgt,
	hcl,
	ecl,
	pid,
	cid string
}

func problem1(filename string) int {
	return solve(filename, false)
}
func problem2(filename string) int {
	return solve(filename, true)
}
func solve(filename string, strict bool) int {
	passports := readFile(filename, strict)
	count := 0
	for _, passport := range passports {
		if err := allFieldsExist(passport, strict); err == nil {
			count = count + 1
		} else {
			// fmt.Println("Could not set value", err.Error())
		}
	}
	return count
}

func readFile(filename string, strict bool) []Passport {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	var passports []Passport
	pparr := []string{}
	for scanner.Scan() {
		if text := scanner.Text(); len(text) > 0 {
			pparr = append(pparr, strings.Fields(text)...)
		} else {
			passports = append(passports, createPassport(pparr, strict))
			pparr = []string{}
		}
	}
	if len(pparr) > 0 {
		passports = append(passports, createPassport(pparr, strict))
	}
	return passports
}
func allFieldsExist(passport Passport, strict bool) error {
	if err := validateString("byr", passport.byr, isByrValid, strict); err != nil {
		return err
	}
	if err := validateString("iyr", passport.iyr, isIyrValid, strict); err != nil {
		return err
	}
	if err := validateString("eyr", passport.eyr, isEyrValid, strict); err != nil {
		return err
	}
	if err := validateString("hgt", passport.hgt, isHeightValid, strict); err != nil {
		return err
	}
	if err := validateString("hcl", passport.hcl, isHairColorValid, strict); err != nil {
		return err
	}
	if err := validateString("ecl", passport.ecl, isEyeColorValid, strict); err != nil {
		return err
	}
	if err := validateString("pid", passport.pid, isPidValid, strict); err != nil {
		return err
	}
	return nil
}

func checkEmpty(attribute string, value string) error {
	if len(value) > 0 {
		return nil
	}
	return fmt.Errorf("atribute [%s] is required", attribute)
}

func isNumberValid(value string, min, max int) error {
	if i, err := strconv.Atoi(value); err != nil {
		return fmt.Errorf("could not parse [%s] to number", value)
	} else if i < min || i > max {
		return fmt.Errorf("value [%s] is not between %d and %d", value, min, max)
	}
	return nil

}
func isHeightValid(value string) error {
	if !heighRegex.MatchString(value) {
		return fmt.Errorf("value [%s] does not match height regex", value)
	}
	if strings.Contains(value, "in") {
		return isNumberValid(value[:len(value)-len("in")], 59, 76)
	} else {
		return isNumberValid(value[:len(value)-len("cm")], 150, 193)
	}

}
func isHairColorValid(value string) error {
	if !colorRegex.MatchString(value) {
		return fmt.Errorf("value [%s] does not match color regex", value)
	}
	return nil
}
func isEyeColorValid(value string) error {
	if _, exists := eyeColors[value]; !exists {
		return fmt.Errorf("value [%s] is not a valid eye color option", value)
	}
	return nil
}
func isPidValid(value string) error {
	if !idRegex.MatchString(value) {
		return fmt.Errorf("value [%s] is not a valid id", value)
	}
	return nil
}
func isByrValid(value string) error {
	return isNumberValid(value, 1920, 2002)
}
func isIyrValid(value string) error {
	return isNumberValid(value, 2010, 2020)
}
func isEyrValid(value string) error {
	return isNumberValid(value, 2020, 2030)
}
func createPassport(pparr []string, strict bool) Passport {
	passport := Passport{}
	for _, kv := range pparr {

		kvArr := strings.Split(kv, ":")
		value := kvArr[1]
		attribute := kvArr[0]
		err := setValue(attribute, value, &passport, strict)
		if err != nil {
			// fmt.Println("Could not set value", err.Error())
		}

	}
	return passport
}
func isCidValid(value string) error {
	return nil
}
func validateString(attribute string, value string, validator func(s string) error, strict bool) error {
	if len(value) == 0 {

		return fmt.Errorf("Attribute %s is required", attribute)
	}
	if strict {
		return validator(value)
	}
	return nil
}
func setValue(attribute string, value string, passport *Passport, strict bool) error {
	var f func(string) error
	var err error
	var setter *string
	switch attribute {
	case "byr":
		f = isByrValid
		setter = &passport.byr
	case "ecl":
		f = isEyeColorValid
		setter = &passport.ecl
	case "eyr":
		f = isEyrValid
		setter = &passport.eyr
	case "hcl":
		f = isHairColorValid
		setter = &passport.hcl
	case "hgt":
		f = isHeightValid
		setter = &passport.hgt
	case "iyr":
		f = isIyrValid
		setter = &passport.iyr
	case "pid":
		f = isPidValid
		setter = &passport.pid
	case "cid":
		f = isCidValid
		setter = &passport.cid
	default:
		fmt.Println("unknown field", attribute)
	}
	if f != nil && strict {
		err = f(value)
	}
	if err != nil {
		return err
	}
	if setter != nil {
		*setter = value
	}
	return nil
}

func testValid(value string, expected bool, validator func(s string) error) {
	if err := validator(value); ((err == nil) && !expected) || (err != nil && expected) {
		if err != nil {
			fmt.Println(fmt.Sprintf("Test failed: expected %t, was [%s] for value %s", expected, err.Error(), value))
		} else {
			fmt.Println(fmt.Sprintf("Test failed: expected %t, was [%t] for value %s", expected, true, value))
		}

	}
}
func testValidators() {
	testValid("2002", true, isByrValid)
	testValid("2003", false, isByrValid)

	testValid("60in", true, isHeightValid)
	testValid("76in", true, isHeightValid)
	testValid("190cm", true, isHeightValid)
	testValid("190in", false, isHeightValid)
	testValid("190", false, isHeightValid)

	testValid("#123abc", true, isHairColorValid)
	testValid("#123abz", false, isHairColorValid)
	testValid("123abc", false, isHairColorValid)

	testValid("brn", true, isEyeColorValid)
	testValid("wat", false, isEyeColorValid)

	testValid("000000001", true, isPidValid)
	testValid("0123456789", false, isPidValid)
}

func test(filename string, problemNr int, expected int, problemFunc func(s string) int) {
	if actual := problemFunc(filename); actual != expected {
		fmt.Println(fmt.Sprintf("Test Problem %d failed: expected %d, was %d", problemNr, expected, actual))
	}
}

func main() {
	test("test_input.txt", 1, 2, problem1)
	test("test_input_2_invalid.txt", 2, 0, problem2)
	test("test_input_2_valid.txt", 2, 4, problem2)
	testValidators()
	fmt.Println("Problem1: ", problem1("input.txt"))
	fmt.Println("Problem2: ", problem2("input.txt"))
}
