package main

import (
	"fmt"

	"github.com/MatDevel0per/Advent-of-code-solutions/utils"
)

func main() {
	input := utils.ReadInput()
	count := countXMASPart1(input)
	fmt.Printf("The word xmas appears %d times.\n", count)
	part2 := countXMASPart2(input)
	fmt.Printf("The X-MAS pattern appears %d times.\n", part2)
}

// current implementation could be improved by only checking the directions on x characters but if it works I'm not touching it.
func countXMASPart1(grid []string) int {
	word := "XMAS"
	rows := len(grid)
	cols := len(grid[0])
	wordLength := len(word)
	count := 0

	// Directions for moving: (row_offset, col_offset)
	directions := [8][2]int{
		{0, 1},   // Horizontal right
		{0, -1},  // Horizontal left
		{1, 0},   // Vertical down
		{-1, 0},  // Vertical up
		{1, 1},   // Diagonal down-right
		{-1, -1}, // Diagonal up-left
		{1, -1},  // Diagonal down-left
		{-1, 1},  // Diagonal up-right
	}

	// Convert grid into a 2D slice for easier access
	gridMatrix := make([][]rune, rows)
	for i := range grid {
		gridMatrix[i] = []rune(grid[i])
	}

	// Check if the word exists in a specific direction
	isWordAt := func(row, col, dirRow, dirCol int) bool {
		// Check the direction for the length of the xmas word
		for i := 0; i < wordLength; i++ {
			// increment the row and column by the letter of the targer word I.e first loop is 0 so doesn't add anything to the check
			// third loop is 2 so will displace check by 2 for example diag down right two right two down.
			newRow := row + i*dirRow
			newCol := col + i*dirCol
			// Guard statement Check it's in bounds of the matrix
			if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
				return false
			}
			// Does the word equal what is expected during this loop of the word
			if gridMatrix[newRow][newCol] != rune(word[i]) {
				return false
			}
		}
		// If it passed the loop then it's safe
		return true
	}

	// Iterate through each letter in the grid
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// Check in all 8 directions
			for _, dir := range directions {
				if isWordAt(row, col, dir[0], dir[1]) {
					count++
				}
			}
		}
	}

	return count
}

// Function to count X-MAS patterns
func countXMASPart2(grid []string) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	// Check if diagonals around center form valid MAS patterns
	isValidXMAS := func(row, col int) bool {
		// Ensure diagonals form valid "MAS" or "SAM" patterns crossing at center `A`
		topLeft := row-1 >= 0 && col-1 >= 0 && (grid[row-1][col-1] == 'M' && grid[row+1][col+1] == 'S' || grid[row-1][col-1] == 'S' && grid[row+1][col+1] == 'M')
		topRight := row-1 >= 0 && col+1 < cols && (grid[row-1][col+1] == 'M' && grid[row+1][col-1] == 'S' || grid[row-1][col+1] == 'S' && grid[row+1][col-1] == 'M')

		// Bottom-left and bottom-right must match the opposite diagonal arms
		return topLeft && topRight
	}

	// Iterate over every letter to find centers
	for row := 1; row < rows-1; row++ {
		for col := 1; col < cols-1; col++ {
			// Check if current cell is `A` and validate diagonals match the x pattern
			if grid[row][col] == 'A' && isValidXMAS(row, col) {
				count++
			}
		}
	}

	return count
}
