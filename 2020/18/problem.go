package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getOperandIndex(s string) int {
	multIndex := strings.IndexByte(s, '*')
	addIndex := strings.IndexByte(s, '+')
	operandIndex := addIndex
	if multIndex >= 0 && (multIndex < operandIndex || operandIndex < 0) {
		operandIndex = multIndex
	}
	return operandIndex
}
func getPrevOperandIndex(s string) int {
	multIndex := strings.LastIndex(s, "*")
	addIndex := strings.LastIndex(s, "+")
	return int(math.Max(float64(multIndex), float64(addIndex)))
}
func getOperands(s string) (int, int) {
	multIndex := strings.IndexByte(s, '*')
	addIndex := strings.IndexByte(s, '+')
	return addIndex, multIndex
}
func getStartEndParenthesisIndex(bytes []byte) (int, int) {
	openCount := 0
	closeCount := 0
	endIndex := -1
	parenthesisIndex := strings.IndexByte(string(bytes), '(')
	if parenthesisIndex < 0 {
		return -1, -1
	}
	for i, b := range bytes[parenthesisIndex:len(bytes)] {
		if b == '(' {
			openCount++
		} else if b == ')' {
			closeCount++
		}
		if openCount == closeCount {
			endIndex = i + parenthesisIndex
			break
		}
	}
	return parenthesisIndex, endIndex
}
func simplify(s string) (string, error) {

	operandIndex := getOperandIndex(s)
	bytes := []byte(s)
	parenthesisStart, parenthesisEnd := getStartEndParenthesisIndex(bytes)
	if parenthesisStart >= 0 {
		val, err := simplify(string(bytes[parenthesisStart+1 : parenthesisEnd]))
		if err != nil {
			return val, err
		}
		str := string(bytes[0:parenthesisStart]) + val + string(bytes[parenthesisEnd+1:len(bytes)])
		return simplify(str)
	} else if operandIndex >= 0 {
		str, err := simplifyOperand(bytes, operandIndex)
		if err != nil {
			return str, err
		}
		return simplify(str)
	} else {
		return s, nil
	}
}
func simplifyOperand(bytes []byte, operandIndex int) (string, error) {
	start := bytes[0:operandIndex]
	prevOperand := getPrevOperandIndex(string(start))
	head := []byte{}
	if prevOperand != -1 {
		head = start[0 : prevOperand+1]
		start = start[prevOperand+1 : len(start)]
	}
	futureSearch := string(bytes[operandIndex+1 : len(bytes)])
	nextOperand := getOperandIndex(futureSearch)
	endIndex := len(bytes)
	if nextOperand >= 0 {
		endIndex = nextOperand + operandIndex + 1
	}
	end := bytes[operandIndex+1 : endIndex]
	ans, err := evaluateOperand(start, bytes[operandIndex], end)
	if err != nil {
		return "", err
	}
	str := string(head) + strconv.Itoa(ans) + string(bytes[endIndex:len(bytes)])
	return str, nil
}
func simplify2(s string) (string, error) {
	bytes := []byte(s)
	addIndex, multIndex := getOperands(s)
	operandIndex := multIndex
	parenthesisStart, parenthesisEnd := getStartEndParenthesisIndex(bytes)
	if addIndex >= 0 {
		operandIndex = addIndex
	}
	if parenthesisStart >= 0 {
		val, err := simplify2(string(bytes[parenthesisStart+1 : parenthesisEnd]))
		if err != nil {
			return val, err
		}
		str := string(bytes[0:parenthesisStart]) + val + string(bytes[parenthesisEnd+1:len(bytes)])
		return simplify2(str)
	} else if operandIndex >= 0 {
		str, err := simplifyOperand(bytes, operandIndex)
		if err != nil {
			return str, err
		}
		return simplify2(str)
	}
	return s, nil
}
func evaluateOperand(term1 []byte, operand byte, term2 []byte) (int, error) {
	i, iErr := strconv.ParseInt(string(term1), 10, 64)
	if iErr != nil {
		return -1, iErr
	}
	j, jErr := strconv.ParseInt(string(term2), 10, 64)
	if jErr != nil {
		return -1, jErr
	}
	if operand == '*' {
		return int(i * j), nil
	} else if operand == '+' {
		return int(i + j), nil
	} else {

		return 0, fmt.Errorf("Invalid operand %b", operand)
	}
}
func parseString(s string) string {
	parsed := strings.ReplaceAll(s, " ", "")
	return parsed

}
func solveExpression(s string, simplifier func(string) (string, error)) (int, error) {
	simplified, err := simplifier(parseString(s))
	if err != nil {
		return -1, err
	}
	i, iErr := strconv.ParseInt(simplified, 10, 64)
	if iErr != nil {
		return -1, iErr
	}
	return int(i), nil
}
func expectP1(s string, expected int) {
	actual, err := solveExpression(s, simplify)
	if err != nil {
		fmt.Println(err.Error())
	} else if actual != expected {
		fmt.Println("Not correct", s, "should be", expected, "actual", actual)
	}
}
func testCases() {
	expectP1("2 * 3 + (4 * 5)", 26)
	expectP1("5 + (8 * 3 + 9 + 3 * 4 * 3)", 437)
	expectP1("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240)
	expectP1("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632)
}
func readFile(filename string) (int, int) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	sum1 := 0
	sum2 := 0
	for scanner.Scan() {
		instruction := parseString(scanner.Text())
		ans1, err1 := solveExpression(instruction, simplify)
		if err1 != nil {
			fmt.Println(err1.Error())
			return sum1, sum2
		}
		ans2, err2 := solveExpression(instruction, simplify2)
		if err2 != nil {
			fmt.Println(err2.Error())
			return sum1, sum2
		}
		sum1 += ans1
		sum2 += ans2
	}
	return sum1, sum2
}
func solve(filename string) (int, int) {
	p1, p2 := readFile(filename)
	return p1, p2
}
func main() {
	testCases()
	p1, p2 := solve("input.txt")
	fmt.Println(fmt.Sprintf("Problem 1: %d , expected %d", p1, 4940631886147))
	fmt.Println(fmt.Sprintf("Problem 2: %d , expected %d", p2, 283582817678281))
}
