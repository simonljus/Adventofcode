package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countUniques(responses []string) map[string]int {
	counts := make(map[string]int)
	for _, response := range responses {
		for _, ans := range strings.Split(response, "") {
			if v, exists := counts[ans]; !exists {
				counts[ans] = 1
			} else {
				counts[ans] = v + 1
			}
		}
	}
	return counts
}
func countGroup(responses []string) (int, int) {
	return countMap(len(responses), countUniques(responses))
}
func countMap(nResponses int, answers map[string]int) (int, int) {
	p2 := 0
	for _, v := range answers {
		if v == nResponses {
			p2 += 1
		}
	}
	return len(answers), p2
}

func solve(filename string) (int, int) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	p1Sum := 0
	p2Sum := 0
	answers := []string{}

	for scanner.Scan() {
		if text := scanner.Text(); len(text) > 0 {
			answers = append(answers, text)
		} else {
			p1, p2 := countGroup(answers)
			p1Sum += p1
			p2Sum += p2
			answers = []string{}
		}
	}
	if len(answers) > 0 {
		p1, p2 := countGroup(answers)
		p1Sum += p1
		p2Sum += p2
	}
	return p1Sum, p2Sum
}
func main() {
	p1Actual, p2Actual := solve("test_input.txt")
	fmt.Println(fmt.Sprintf("Test cases: Problem 1: expexted %d, actual %d , Problem 2: expected,%d actual  %d ", 11, p1Actual, 6, p2Actual))
	p1, p2 := solve("input.txt")
	// 6504,3351
	fmt.Println(fmt.Sprintf("Problem 1: %d, Problem 2: %d", p1, p2))

}
