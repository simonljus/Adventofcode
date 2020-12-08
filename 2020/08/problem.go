package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	opcode string
	count  int
}

func existsInSet(i int, set map[int]bool) bool {
	_, exists := set[i]
	return exists
}

func getOpcode(inst Instruction, pos int, swapIndex int) string {
	if pos == swapIndex {
		if inst.opcode == "jmp" {
			return "nop"
		} else if inst.opcode == "cop" {
			return "jmp"
		}
	}
	return inst.opcode
}
func runInstructions(instructions []Instruction, swapIndex int) (int, bool) {
	acc := 0
	pos := 0
	seen := make(map[int]bool)
	for !existsInSet(pos, seen) {
		seen[pos] = true
		inst := instructions[pos]
		opcode := getOpcode(inst, pos, swapIndex)
		switch opcode {
		case "acc":
			acc += inst.count
			pos += 1
		case "jmp":
			pos += inst.count
		case "nop":
			pos += 1
		default:
			fmt.Println(fmt.Sprintf("UNKNOWN OPCODE %s AT POSITION %d", opcode, pos))
			pos += 1
		}

		if pos == len(instructions) {
			return acc, true
		}
	}
	return acc, false
}

func createInstructions(filename string) []Instruction {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	instructions := []Instruction{}
	for scanner.Scan() {
		instruction := toInstruction(scanner.Text())
		instructions = append(instructions, instruction)
	}
	return instructions
}

func toInstruction(s string) Instruction {
	words := strings.Split(s, " ")
	opcode := words[0]
	count := 0
	if c, err := strconv.Atoi(words[1]); err != nil {
		fmt.Println(err.Error(), c)

	} else {
		count = c
	}
	return Instruction{opcode, count}
}
func problem1(instructions []Instruction) int {
	acc, _ := runInstructions(instructions, -1)
	return acc
}
func problem2(instructions []Instruction) (int, bool) {
	for i, instruction := range instructions {
		if opcode := getOpcode(instruction, i, i); opcode != instruction.opcode {
			if acc, completed := runInstructions(instructions, i); completed {
				return acc, completed
			}
		}
	}
	return -1, false
}

func solve(filename string) (int, int) {
	instructions := createInstructions(filename)
	p1 := problem1(instructions)
	p2, _ := problem2(instructions)
	return p1, p2
}

func main() {
	tp1, tp2 := solve("test_input.txt")
	p1, p2 := solve("input.txt")
	fmt.Println(fmt.Sprintf("Test Problem 1: %d , expected %d", tp1, 5))
	fmt.Println(fmt.Sprintf("Test Problem 2: %d , expected %d", tp2, 8))
	fmt.Println(fmt.Sprintf("Problem 1: %d , expected %d", p1, 1600))
	fmt.Println(fmt.Sprintf("Problem 2: %d , expexted %d", p2, 1543))
}
