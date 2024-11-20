package main

import (
	"fmt"
	"os"
)

// This program solves a Sudoku puzzle using backtracking. The grid is passed
// as command-line arguments, and the solution (if found) is printed.

var grid = [9][9]int{}
var solutionCount int // To track if there's more than one solution

// draw prints the current Sudoku grid to the console.
func draw() {
	for _, row := range grid {
		for i, value := range row {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(value)
		}
		fmt.Println()
	}
}

// acceptable checks if placing a value in grid[x][y] is valid according to Sudoku rules.
func acceptable(x int, y int, value int) bool {
	return !existsInVertical(x, value) && !existsInHorizontal(y, value) && !existsInBox(x, y, value)
}

// existsInVertical checks if a value is already in the given column.
func existsInVertical(x int, value int) bool {
	for i := 0; i < 9; i++ {
		if grid[i][x] == value {
			return true
		}
	}
	return false
}

// existsInHorizontal checks if a value is already in the given row.
func existsInHorizontal(y int, value int) bool {
	for i := 0; i < 9; i++ {
		if grid[y][i] == value {
			return true
		}
	}
	return false
}

// existsInBox checks if a value is already in the 3x3 sub-grid.
func existsInBox(x int, y int, value int) bool {
	sx, sy := int(x/3)*3, int(y/3)*3
	for dy := 0; dy < 3; dy++ {
		for dx := 0; dx < 3; dx++ {
			if grid[sy+dy][sx+dx] == value {
				return true
			}
		}
	}
	return false
}

// nextCell determines the next cell to process on the Sudoku grid.
func nextCell(x int, y int) (int, int) {
	if x == 8 {
		return 0, y + 1
	}
	return x + 1, y
}

// solution tries to solve the Sudoku puzzle using backtracking and checks for multiple solutions.
func solution(x int, y int) bool {
	if y == 9 {
		solutionCount++
		if solutionCount > 1 {
			return false // Early exit if more than one solution is found
		}
		return true
	}
	if grid[y][x] != 0 {
		return solution(nextCell(x, y))
	}
	for v := 1; v <= 9; v++ {
		if acceptable(x, y, v) {
			grid[y][x] = v
			if solution(nextCell(x, y)) {
				if solutionCount > 1 {
					return false
				}
				return true
			}
			grid[y][x] = 0
		}
	}
	return false
}

// stringToInt converts a rune to an integer if it's a digit.
func stringToInt(char rune) (int, bool) {
	if char >= '0' && char <= '9' {
		return int(char - '0'), true
	}
	return 0, false
}

// main processes the input and starts solving the Sudoku puzzle.
func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}
	for i, arg := range os.Args[1:] {
		if len(arg) != 9 {
			fmt.Println("Error")
			return
		}
		for j, char := range arg {
			if char == '.' {
				grid[i][j] = 0
			} else if num, valid := stringToInt(char); valid {
				grid[i][j] = num
			} else {
				fmt.Println("Error")
				return
			}
		}
	}

	if solution(0, 0) {
		if solutionCount == 1 { // Ensure there's only one solution
			draw()
		} else {
			fmt.Println("Error")
		}
	} else {
		fmt.Println("Error")
	}
}
