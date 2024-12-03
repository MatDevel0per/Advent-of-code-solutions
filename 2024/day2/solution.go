package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MatDevel0per/Advent-of-code-solutions/utils"
)

func main() {
	input := utils.ReadInput()
	fmt.Println("Part 1: ", safeReactorReadingPart1(input))
	fmt.Println("Part 2: ", safeReactorReadingPart2(input))
}

func toInt(input []string) ([]int, error) {
	// Expecting it to be split into array of string fields.
	// Make an empty array with the same length as the input array.
	intList := make([]int, len(input))
	// Loop the input and convert to integer then insert into blank array.
	for i, x := range input {
		values, err := strconv.Atoi(x)
		if err != nil {
			return nil, err
		}
		intList[i] = values
	}
	return intList, nil
}

func safe(inputArray []int) bool {
	// Check if it is ascending so we can apply logic later
	ascending := inputArray[1]-inputArray[0] > 0
	// Start at 1 so we can check the second element against first element first. Iterate to end to validate entire entry
	for i := 1; i < len(inputArray); i++ {
		// Get the increment
		incr := inputArray[i] - inputArray[i-1]
		// Guard statement, If increment is greater than zero but first 2 are not ascending then sequence is not safe
		if incr > 0 != ascending {
			return false
		}
		// If increment is less than zero flip it to positive so we can use one check
		if incr < 0 {
			incr = -incr
		}
		// If the difference is less than 1 or greater than three then it is not safe.
		if incr < 1 || incr > 3 {
			return false
		}
	}
	// If non of the checks return false then reading is safe
	return true
}
func safeByRemoving(list []int) bool {
	for i := 0; i < len(list); i++ {
		//All elements before current index
		newIntSlice := append([]int{}, list[:i]...)
		//All elements after current index
		newIntSlice = append(newIntSlice, list[i+1:]...)
		//check if safe if it is then return true because error allowance is met
		if safe(newIntSlice) {
			return true
		}
	}
	// If none of the removals are safe return false
	return false
}
func safeReactorReadingPart1(input []string) int {
	var answer = 0
	// Loop through inputs
	for _, str := range input {
		// Turn row in array into integer array.
		list, _ := toInt(strings.Fields(str))
		// If the reading is safe increment the safe count.
		if safe(list) {
			answer++
		}
	}
	return answer
}
func safeReactorReadingPart2(input []string) int {
	var answer = 0
	// Loop through inputs
	for _, str := range input {
		// Turn row in array into integer array.
		list, _ := toInt(strings.Fields(str))
		// If the reading is safe or safe by removing one value increment the safe count.
		if safe(list) || safeByRemoving(list) {
			answer++
		}
	}
	return answer
}
