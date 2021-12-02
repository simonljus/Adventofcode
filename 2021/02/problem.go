package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseString(s string) string {
	return s
}
func problem1(instructions []string) int {
	var direction string
	var velocity int
	horizontal := 0
	depth := 0
	for _, instruction := range instructions {
		fmt.Sscanf(instruction, "%s %d", &direction, &velocity)
		if direction == "forward" {
			horizontal += velocity
		} else if direction == "up" {
			depth -= velocity
		} else if direction == "down" {
			depth += velocity
		}
	}
	return horizontal * depth
}
func problem2(instructions []string) int {
	var direction string
	var velocity int
	horizontal := 0
	depth := 0
	aim := 0
	for _, instruction := range instructions {
		fmt.Sscanf(instruction, "%s %d", &direction, &velocity)
		if direction == "forward" {
			horizontal += velocity
			depth += aim * velocity
		} else if direction == "up" {
			aim -= velocity
		} else if direction == "down" {
			aim += velocity
		}
	}
	return horizontal * depth
}
func readFile(filename string) []string {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	instructions := []string{}
	for scanner.Scan() {
		instruction := parseString(scanner.Text())
		instructions = append(instructions, instruction)
	}
	return instructions
}
func solve(filename string) (int, int) {
	instructions := readFile(filename)
	p1 := problem1(instructions)
	p2 := problem2(instructions)
	return p1, p2
}
func main() {
	tp1, tp2 := solve("test_input.txt")
	p1, p2 := solve("input.txt")
	fmt.Println(fmt.Sprintf("Test Problem 1: %d , expected %d", tp1, 150))
	fmt.Println(fmt.Sprintf("Test Problem 2: %d , expected %d", tp2, 900))
	fmt.Println(fmt.Sprintf("Problem 1: %d , expected %d", p1, 1728414))
	fmt.Println(fmt.Sprintf("Problem 2: %d , expected %d", p2, 1765720035))
}
