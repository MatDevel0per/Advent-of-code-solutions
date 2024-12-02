package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/MatDevel0per/Advent-of-code-solutions/utils"
)

func main() {
	input := utils.ReadInput()
	fmt.Println("Part 1: ", distanceBetweenPart1(input))
	fmt.Println("Part 2: ", distanceBetweenPart2(input))
}
func splitInput(input []string) ([]int, []int) {
	var a, b []int
	for _, str := range input {
		fields := strings.Fields(str)
		n1, _ := strconv.Atoi(fields[0])
		n2, _ := strconv.Atoi(fields[1])
		a = append(a, n1)
		b = append(b, n2)
	}
	return a, b
}
func splitInput2(input []string) ([]int, map[int]int) {
	var a []int
	var b = map[int]int{}

	for _, str := range input {
		fields := strings.Fields(str)
		n1, _ := strconv.Atoi(fields[0])
		n2, _ := strconv.Atoi(fields[1])
		a = append(a, n1)
		b[n2]++
	}
	return a, b
}
func distanceBetweenPart1(input []string) int {
	var answer = 0
	list1, list2 := splitInput(input)
	sort.Ints(list1)
	sort.Ints(list2)
	for i := 0; i < len(list1); i++ {
		answer += abs(list1[i] - list2[i])
	}
	return answer
}

func distanceBetweenPart2(input []string) int {
	answer := 0
	list1, list2 := splitInput2(input)
	for _, n := range list1 {
		answer += n * list2[n]
	}
	return answer
}

// Absolute number returner
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
