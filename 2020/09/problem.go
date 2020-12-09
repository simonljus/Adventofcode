package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseString(s string) int64 {
	if i, err := strconv.ParseInt(s, 10, 64); err != nil {
		fmt.Println(fmt.Sprintf("%s is not a number", s))
		return -1
	} else {
		return i
	}
}
func findsum(list []int64, sum int64) (int, int, error) {
	mem := make(map[int64]int)
	for i, v := range list {
		if j, exists := mem[sum-v]; !exists {
			mem[v] = i
		} else {
			return i, j, nil
		}
	}
	return -1, -1, fmt.Errorf("Could not fund sum %d", sum)
}

func problem1(scanned []int64, sum int64) error {
	_, _, err := findsum(scanned, sum)
	return err
}
func smallestlargest(scanned []int64) int64 {
	smallest := scanned[0]
	largest := scanned[0]
	for _, v := range scanned {
		if v < smallest {
			smallest = v
		}
		if v > largest {
			largest = v
		}
	}
	return smallest + largest
}
func problem2(instructions []int64, target int64) int64 {
	start := 0
	end := 1
	sum := instructions[0]
	for end != len(instructions) {
		sum += instructions[end]
		if sum == target {
			return smallestlargest(instructions[start : end+1])
		} else if sum < target {
			end += 1
		} else {
			start += 1
			sum = instructions[start]
			end = start + 1
		}
	}
	return -1
}
func solve(filename string, preamble int) (int64, int64) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	scanned := []int64{}
	// usedIndexes := make(map[int]bool)
	for i := 0; scanner.Scan(); i++ {
		instruction := parseString(scanner.Text())
		if i >= preamble {
			err := problem1(scanned[i-preamble:], instruction)
			if err != nil {
				p2 := problem2(scanned, instruction)
				return instruction, p2
			}
		}
		scanned = append(scanned, instruction)

	}
	return -1, -1
}
func main() {
	tp1, tp2 := solve("test_input.txt", 5)
	p1, p2 := solve("input.txt", 25)
	fmt.Println(fmt.Sprintf("Test Problem 1: %d , expected %d", tp1, 127))
	fmt.Println(fmt.Sprintf("Test Problem 2: %d , expected %d", tp2, 62))
	fmt.Println(fmt.Sprintf("Problem 1: %d , expected %d", p1, 22406676))
	fmt.Println(fmt.Sprintf("Problem 2: %d , expexted %d", p2, 2942387))
}
