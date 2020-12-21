package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`[^0-9]`)

func parseString(s string) string {
	return s
}
func problem1(instructions []string) int {
	return 0
}
func problem2(instructions []string) int {
	return 0
}
func parseWriter(text string) (int64, int64, error) {
	split := strings.Split(text, "=")
	if !strings.Contains(split[0], "mem") {
		return -1, -1, fmt.Errorf("NOT MEMORY")
	}
	memAddress := parseInt(split[0])
	value := parseInt(split[1])
	return memAddress, value, nil
}
func toBinary(value int64) string {
	binary := strconv.FormatInt(int64(value), 2)
	binary = strings.Repeat("0", 36-len(binary)) + binary
	return binary
}
func parseInt(s string) int64 {
	valueStr := string(re.ReplaceAll([]byte(s), []byte("")))
	value, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		return -1
	}
	return value
}
func writeValue(memptr *map[int64][]byte, memAddress int64, value []byte) {
	mem := *memptr
	mem[memAddress] = value
}
func createMask(mask []byte, value int64, override byte) []byte {
	valuearr := []byte(toBinary(value))
	createdMask := []byte{}
	for i, b := range mask {
		bit := b
		if b == override {
			bit = valuearr[i]
		}
		createdMask = append(createdMask, bit)
	}
	return createdMask
}
func parseMask(s string) []byte {
	split := strings.Split(s, "=")
	return []byte(strings.TrimSpace(split[1]))
}
func sumMem(mem map[int64][]byte) int64 {
	sum := int64(0)
	for _, v := range mem {
		value, err := strconv.ParseInt(string(v), 2, 64)
		if err == nil {
			sum += value
		}
	}
	return sum
}
func writeValues(mapptr *map[int64][]byte, memMask []byte, value []byte) {
	memStr := string(memMask)
	index := strings.IndexByte(memStr, 'X')
	if index == -1 {
		memAddress, _ := strconv.ParseInt(memStr, 2, 64)
		writeValue(mapptr, memAddress, value)
	} else {
		with1 := []byte(string(memMask))
		with1[index] = '1'
		writeValues(mapptr, with1, value)
		with0 := []byte(string(memMask))
		with0[index] = '0'
		writeValues(mapptr, with0, value)

	}
}
func readFile(filename string) (int64, int64) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	mem := make(map[int64][]byte)
	mem2 := make(map[int64][]byte)
	mask := []byte{}
	for scanner.Scan() {
		text := scanner.Text()
		memAddress, value, err := parseWriter(text)

		if err == nil {
			valuemask := createMask(mask, value, 'X')
			writeValue(&mem, memAddress, valuemask)
			memMask := createMask(mask, memAddress, '0')
			writeValues(&mem2, memMask, []byte(toBinary(value)))
		} else {

			mask = parseMask(text)
		}

	}
	return sumMem(mem), sumMem(mem2)
}
func solve(filename string) (int64, int64) {
	p1, p2 := readFile(filename)
	return p1, p2
}
func main() {
	/*tp1, _ := solve("test_input.txt")
	fmt.Println(fmt.Sprintf("Test Problem 1: %d , expected %d", tp1, 165))*/
	_, tp2 := solve("test_input2.txt")
	fmt.Println(fmt.Sprintf("Test Problem 2: %d , expected %d", tp2, 208))
	p1, p2 := solve("input.txt")
	fmt.Println(fmt.Sprintf("Problem 1: %d , expected %d", p1, 7611244640053))
	fmt.Println(fmt.Sprintf("Problem 2: %d , expected %d", p2, 3705162613854))
}
