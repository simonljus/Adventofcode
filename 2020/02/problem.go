package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(s string) (int, int, string, string) {
	args := strings.Fields(s)
	interval, charcolon, password := args[0], args[1], args[2]
	char := strings.Replace(charcolon, ":", "", 1)
	intervalArgs := strings.Split(interval, "-")
	min, _ := strconv.Atoi(intervalArgs[0])
	max, _ := strconv.Atoi(intervalArgs[1])
	return min, max, char, password
}

func isValidCount(s string) bool {
	min, max, char, password := parseInput(s)
	count := strings.Count(password, char)
	return count >= min && count <= max
}

func isMatchXor(s string) bool {
	firstPos, secondPos, char, password := parseInput(s)
	firstMatch := password[firstPos-1] == char[0]
	secondMatch := password[secondPos-1] == char[0]
	return firstMatch != secondMatch
}

func problemSolver(f func(s string) bool) int {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	validPasswords := 0
	for scanner.Scan() {
		if f(scanner.Text()) {
			validPasswords = validPasswords + 1
		}
	}
	return validPasswords
}

func testProblem1() {
	a := isValidCount("1-3 a: abcde")
	b := isValidCount("1-3 b: cdefg")
	c := isValidCount("2-9 c: ccccccccc")
	if !a {
		fmt.Println("a should be valid")

	}
	if b {
		fmt.Println("b should be invalid")
	}
	if !c {
		fmt.Println("c should be valid")
	}

}
func testProblem2() {
	a := isMatchXor("1-3 a: abcde")
	b := isMatchXor("1-3 b: cdefg")
	c := isMatchXor("2-9 c: ccccccccc")
	if !a {
		fmt.Println("a should be valid")

	}
	if b {
		fmt.Println("b should be invalid")
	}
	if c {
		fmt.Println("c should be invalid")
	}

}

//546
func problem1() int {
	return problemSolver(isValidCount)
}

//275
func problem2() int {
	return problemSolver(isMatchXor)
}
func main() {
	fmt.Println("Problem 1 test start")
	testProblem1()
	fmt.Println("Problem 1 test end")
	fmt.Println("Problem 2 test start")
	testProblem2()
	fmt.Println("Problem 2 test end")
	fmt.Println("Problem 1:", problem1())
	fmt.Println("Problem 2:", problem2())
}
