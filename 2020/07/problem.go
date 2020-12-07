package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bag struct {
	color    string
	quantity int64
}

const SHINYGOLD = "shiny gold"

func toBag(s string) (string, []Bag) {
	words := strings.Split(s, " ")
	maincolor := strings.Join(words[0:2], " ")
	bagsStr := strings.Join(words[4:], " ")
	bagarr := strings.Split(bagsStr, ",")
	bags := []Bag{}
	for _, bagStr := range bagarr {
		bagWords := strings.Split(strings.TrimSpace(bagStr), " ")
		if len(bagWords) < 4 {
			continue
		}
		if quantity, err := strconv.Atoi(bagWords[0]); err != nil {
			fmt.Println(err.Error())
		} else {
			bags = append(bags, Bag{strings.Join(bagWords[1:3], " "), int64(quantity)})
		}

	}
	return maincolor, bags
}
func containsBag(bags []Bag, linked map[string]bool) bool {
	for _, bag := range bags {
		if _, exists := linked[bag.color]; exists {
			return true
		}
	}
	return false
}
func problem2(color string, bagmap map[string][]Bag) int64 {
	seen := make(map[string]int64)
	return getSize(color, bagmap, &seen) - 1
}
func getSize(color string, bagmap map[string][]Bag, seen *map[string]int64) int64 {
	val := *seen
	if storedSize, exists := val[color]; exists {
		return storedSize
	}
	bags, _ := bagmap[color]
	size := int64(1)

	for _, bag := range bags {
		size += bag.quantity * getSize(bag.color, bagmap, &val)
	}
	val[color] = size
	return size
}

func problem1(color string, bagmap map[string][]Bag) int64 {
	linked := make(map[string]bool)
	seen := make(map[string]bool)
	linked[color] = true
	seen[color] = true
	return recSomething(&linked, &seen, &bagmap)
}
func recSomething(linked *map[string]bool, seen *map[string]bool, bagmap *map[string][]Bag) int64 {
	inner := make(map[string]bool)
	searching := *seen
	for key, bags := range *bagmap {
		if _, exists := searching[key]; !exists && containsBag(bags, *linked) {
			inner[key] = true
			searching[key] = true
		}
	}
	if len(inner) != 0 {
		recSomething(&inner, &searching, bagmap)
	}
	return int64(len(*seen) - 1)

}
func createBagMap(filename string) map[string][]Bag {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	bagmap := make(map[string][]Bag)
	for scanner.Scan() {
		text := scanner.Text()
		key, bags := toBag(text)
		if _, exists := bagmap[key]; exists {
			fmt.Println(key, "already exists")
		}
		bagmap[key] = bags
	}
	return bagmap
}

func solve(filename string) (int64, int64) {
	bagmap := createBagMap(filename)
	return problem1(SHINYGOLD, bagmap), problem2(SHINYGOLD, bagmap)
}
func getBestColors(filename string) {
	bagmap := createBagMap(filename)
	bestp1color := []string{}
	bestp1 := int64(0)
	bestp2color := []string{}
	bestp2 := int64(2)
	for k, _ := range bagmap {
		p1 := problem1(k, bagmap)
		p2 := problem2(k, bagmap)
		if p1 >= bestp1 {
			bestp1color = append(bestp1color, k)
			bestp1 = p1
		}
		if p2 >= bestp2 {
			bestp2color = append(bestp2color, k)
			bestp2 = p2
		}
		//fmt.Println(fmt.Sprintf("Color %s: Problem 1: %d ,problem2 %d", k, p1, p2))
	}
	fmt.Println(fmt.Sprintf("Best Color: Problem 1: %s %d ,problem2  %s %d", bestp1color, bestp1, bestp2color, bestp2))
}
func main() {
	p1, p2 := solve("input.txt")
	// 185 89084
	fmt.Println(fmt.Sprintf("Problem 1: %d , problem2 %d", p1, p2))
	// getBestColors("input.txt")
}
