package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseString(s string) string {
	return s
}
func binaryToInt(binary string) int {
	integer, _ := strconv.ParseInt(binary, 2, 64)
	return int(integer)
}
func problem1(instructions []string) int {
	counts := make([]int, len(instructions[0]))

	for _, bits := range instructions {
		for i, bit := range bits {
			if bit == '1' {
				counts[i] += 1
			}
		}
	}
	limit := len(instructions) / 2
	mostCommon := ""
	leastCommon := ""
	for _, count := range counts {
		if count > limit {
			leastCommon += "0"
			mostCommon += "1"
		} else {
			leastCommon += "1"
			mostCommon += "0"
		}
	}
	gamma := binaryToInt(mostCommon)
	epsilon := binaryToInt(leastCommon)
	return gamma * epsilon
}
func problem2(instructions []string) int {
	return mostCommon(instructions) * leastCommon(instructions)
}
func mostCommon(instructions []string) int {
	iterator := instructions
	for i := 0; i < len(instructions[0]); i++ {
		if len(iterator) <= 1 {
			break
		}
		ones, zeroes := split(iterator, i)
		if len(ones) >= len(zeroes) {
			iterator = ones
		} else {
			iterator = zeroes
		}
	}
	return binaryToInt(iterator[0])
}
func leastCommon(instructions []string) int {
	iterator := instructions
	for i := 0; i < len(instructions[0]); i++ {
		if len(iterator) <= 1 {
			break
		}
		ones, zeroes := split(iterator, i)
		if len(ones) < len(zeroes) {
			iterator = ones
		} else {
			iterator = zeroes
		}
	}
	return binaryToInt(iterator[0])
}
func split(instructions []string, index int) ([]string, []string) {
	var ones []string
	var zeroes []string
	for _, instruction := range instructions {
		if instruction[index] == '1' {
			ones = append(ones, instruction)
		} else {
			zeroes = append(zeroes, instruction)
		}
	}
	return ones, zeroes
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
	fmt.Println(fmt.Sprintf("Test Problem 1: %d , expected %d", tp1, 198))
	fmt.Println(fmt.Sprintf("Test Problem 2: %d , expected %d", tp2, 230))
	fmt.Println(fmt.Sprintf("Problem 1: %d , expected %d", p1, 845186))
	fmt.Println(fmt.Sprintf("Problem 2: %d , expected %d", p2, 4636702))
}
