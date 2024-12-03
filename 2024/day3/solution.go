package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/MatDevel0per/Advent-of-code-solutions/utils"
)

func main() {
	input := utils.ReadInput()
	fmt.Println("Part 1: ", multiplyInputPart1(input))
	fmt.Println("Part 2: ", multiplyInputPart2(input))
}

func inputCleaner(input []string, part int) []string {
	var matches []string
	// Man I love regex.
	patternPart1 := regexp.MustCompile(`mul\(\d+,\d+\)`)
	patternPart2 := regexp.MustCompile(`(mul\(\d+,\d+\))|(do\(\))|(don't\(\))`)
	for _, str := range input {
		switch part {
		case 1:
			matches = append(matches, patternPart1.FindAllString(str, -1)...)
		case 2:
			matches = append(matches, patternPart2.FindAllString(str, -1)...)
		}
	}
	return matches
}

func multiplyInputPart1(input []string) int {
	var answer = 0
	input = inputCleaner(input, 1)
	for _, str := range input {
		numbers := strings.Split(str[4:len(str)-1], `,`)
		n1, _ := strconv.Atoi(numbers[0])
		n2, _ := strconv.Atoi(numbers[1])
		answer += (n1 * n2)
	}
	return answer
}
func multiplyInputPart2(input []string) int {
	var answer = 0
	input = inputCleaner(input, 2)
	enabled := true
	for _, str := range input {
		switch str {
		case "do()":
			enabled = true
			continue
		case "don't()":
			enabled = false
			continue
		default:
			if enabled {
				numbers := strings.Split(str[4:len(str)-1], `,`)
				n1, _ := strconv.Atoi(numbers[0])
				n2, _ := strconv.Atoi(numbers[1])
				answer += (n1 * n2)
			}
		}
	}
	return answer
}
