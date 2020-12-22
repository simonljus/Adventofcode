package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type interval struct {
	min int64
	max int64
}
type pair struct {
	attribute string
	index     int
}

func parseString(s string) string {
	return s
}
func problem1(instructions []string) int {
	return 0
}
func problem2(instructions []string) int {
	return 0
}
func getIntervals(s string) []interval {
	a := []interval{}
	for _, r := range strings.Split(s, "or") {
		fromTo := strings.Split(strings.TrimSpace(r), "-")
		from, _ := strconv.ParseInt(fromTo[0], 10, 64)
		to, _ := strconv.ParseInt(fromTo[1], 10, 64)
		a = append(a, interval{min: from, max: to})
	}
	return a
}
func parseIntervals(text string, mapptr *map[string][]interval) error {
	split := strings.Split(text, ":")
	m := *mapptr
	m[split[0]] = getIntervals(split[1])
	return nil
}
func parseTicket(ticket string) []int64 {
	ticketarr := []int64{}
	for _, t := range strings.Split(ticket, ",") {
		v, _ := strconv.ParseInt(t, 10, 64)
		ticketarr = append(ticketarr, v)
	}
	return ticketarr
}
func isValueValid(value int64, index int, intervals map[string][]interval) (bool, []string) {
	valueValid := false
	validAttributes := []string{}

	for attribute, ranges := range intervals {
		for _, minmax := range ranges {
			if value <= minmax.max && value >= minmax.min {
				valueValid = true
				validAttributes = append(validAttributes, attribute)
			}
		}
	}
	return valueValid, validAttributes
}
func isTicketValid(ticket []int64, intervals map[string][]interval) (bool, int64, []pair) {
	errorValues := int64(0)
	isValid := true
	allValidPairs := []pair{}
	for index, t := range ticket {
		valueValid, validAttributes := isValueValid(t, index, intervals)
		if !valueValid {
			isValid = false
			errorValues += t
		} else if isValid {
			for _, attribute := range validAttributes {
				allValidPairs = append(allValidPairs, pair{attribute, index})
			}
		}
	}
	return isValid, errorValues, allValidPairs
}
func countValidTickets(tickets [][]int64, intervals map[string][]interval, myTicket []int64) (int64, int64) {
	sum := int64(0)
	pairmap := make(map[string]pair)
	hasInitialized := false
	for _, ticket := range tickets {
		isValid, errorSum, validPairs := isTicketValid(ticket, intervals)
		sum += errorSum
		if isValid {
			if !hasInitialized {
				for _, p := range validPairs {
					pairmap[fmt.Sprintf("%s_%d", p.attribute, p.index)] = p
				}
				hasInitialized = true
			} else {
				updatedMap := make(map[string]pair)
				for _, p := range validPairs {
					key := fmt.Sprintf("%s_%d", p.attribute, p.index)
					_, exists := pairmap[key]
					if exists {
						updatedMap[key] = p
					}
				}
				pairmap = updatedMap
			}
		}
	}
	p2 := solveProblem2(pairmap, myTicket)
	return sum, p2
}
func scanAttrToID(attrToID map[string]map[int]bool) []pair {
	pairsToDelete := []pair{}
	for attr, v := range attrToID {
		if len(v) == 1 {
			for id := range v {
				pairsToDelete = append(pairsToDelete, pair{attr, id})
			}
		}
	}
	return pairsToDelete
}
func scanIDToAttr(idToAttr map[int]map[string]bool) []pair {
	pairsToDelete := []pair{}
	for id, v := range idToAttr {
		if len(v) == 1 {
			for attr := range v {
				pairsToDelete = append(pairsToDelete, pair{attr, id})
			}
		}
	}
	return pairsToDelete
}
func cleanMaps(pairsToDelete []pair, attrToIDPtr *map[string]map[int]bool, idToAttrPtr *map[int]map[string]bool) {
	attrToID := *attrToIDPtr
	idToAttr := *idToAttrPtr
	for _, p := range pairsToDelete {
		for _, v := range attrToID {
			delete(v, p.index)
		}
		delete(attrToID, p.attribute)
		for _, v := range idToAttr {
			delete(v, p.attribute)
		}
		delete(idToAttr, p.index)
	}
}
func solveProblem2(m map[string]pair, myTicket []int64) int64 {
	attrToID := make(map[string]map[int]bool)
	idToAttr := make(map[int]map[string]bool)
	ansMap := make(map[string]int)
	for _, p := range m {
		_, attrExists := attrToID[p.attribute]
		if !attrExists {
			attrToID[p.attribute] = make(map[int]bool)
		}
		_, idExists := idToAttr[p.index]
		if !idExists {
			idToAttr[p.index] = make(map[string]bool)
		}
		idToAttr[p.index][p.attribute] = true
		attrToID[p.attribute][p.index] = true
	}
	performedClean := true
	for performedClean {
		performedClean = false
		validPairs := []pair{}
		validPairs = scanAttrToID(attrToID)
		if len(validPairs) > 0 {
			performedClean = true
		} else {
			validPairs := scanIDToAttr(idToAttr)
			if len(validPairs) > 0 {
				performedClean = true
			}
		}
		cleanMaps(validPairs, &attrToID, &idToAttr)
		for _, p := range validPairs {
			ansMap[p.attribute] = p.index
		}
	}
	product := int64(1)
	for k, v := range ansMap {
		if strings.Contains(k, "departure") {
			product *= myTicket[v]
		}

	}
	return product
}
func readFile(filename string) (map[string][]interval, []int64, [][]int64) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	intervals := make(map[string][]interval)
	stepper := 0
	tickets := [][]int64{}
	myTicket := []int64{}
	for scanner.Scan() {
		text := scanner.Text()

		if len(text) == 0 {
			stepper++
		}
		if stepper == 0 {
			parseIntervals(text, &intervals)
		}
		if stepper == 1 {
			if !strings.Contains(text, "your") {
				myTicket = parseTicket(text)
			}
		}
		if stepper == 2 {
			if !strings.Contains(text, "nearby") {
				tickets = append(tickets, parseTicket(text))
			}
		}
	}

	return intervals, myTicket, tickets
}
func solve(filename string) (int64, int64) {
	intervals, myTicket, tickets := readFile(filename)
	p1, p2 := countValidTickets(tickets, intervals, myTicket)
	return p1, p2
}
func main() {
	tp1, _ := solve("test_input.txt")
	p1, p2 := solve("input.txt")
	fmt.Println(fmt.Sprintf("Test Problem 1: %d , expected %d", tp1, 71))
	fmt.Println(fmt.Sprintf("Problem 1: %d , expected %d", p1, 24110))
	fmt.Println(fmt.Sprintf("Problem 2: %d , expected %d", p2, 6766503490793))
}
