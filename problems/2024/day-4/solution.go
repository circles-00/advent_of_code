package main

import (
	"strconv"
	"strings"

	"aoc_2024/utils"
)

type Input struct {
	matrix [][]string
}

func (i Input) Up(row, col int) int {
	if row < 3 {
		return 0
	}

	if i.matrix[row-1][col] == "M" && i.matrix[row-2][col] == "A" && i.matrix[row-3][col] == "S" {
		return 1
	}

	return 0
}

func (i Input) Down(row, col int) int {
	if row > len(i.matrix)-4 {
		return 0
	}

	if i.matrix[row+1][col] == "M" && i.matrix[row+2][col] == "A" && i.matrix[row+3][col] == "S" {
		return 1
	}

	return 0
}

func (i Input) Left(row, col int) int {
	if col < 3 {
		return 0
	}

	if i.matrix[row][col-1] == "M" && i.matrix[row][col-2] == "A" && i.matrix[row][col-3] == "S" {
		return 1
	}

	return 0
}

func (i Input) Right(row, col int) int {
	if col > len(i.matrix[row])-4 {
		return 0
	}

	if i.matrix[row][col+1] == "M" && i.matrix[row][col+2] == "A" && i.matrix[row][col+3] == "S" {
		return 1
	}

	return 0
}

func (i Input) UpLeft(row, col int) int {
	if (row < 3) || (col < 3) {
		return 0
	}

	if i.matrix[row-1][col-1] == "M" && i.matrix[row-2][col-2] == "A" && i.matrix[row-3][col-3] == "S" {
		return 1
	}

	return 0
}

func (i Input) UpRight(row, col int) int {
	if (row < 3) || (col > len(i.matrix[row])-4) {
		return 0
	}

	if i.matrix[row-1][col+1] == "M" && i.matrix[row-2][col+2] == "A" && i.matrix[row-3][col+3] == "S" {
		return 1
	}

	return 0
}

func (i Input) DownLeft(row, col int) int {
	if (col < 3) || (row > len(i.matrix)-4) {
		return 0
	}

	if i.matrix[row+1][col-1] == "M" && i.matrix[row+2][col-2] == "A" && i.matrix[row+3][col-3] == "S" {
		return 1
	}

	return 0
}

func (i Input) DownRight(row, col int) int {
	if (col > len(i.matrix[row])-4) || (row > len(i.matrix)-4) {
		return 0
	}

	if i.matrix[row+1][col+1] == "M" && i.matrix[row+2][col+2] == "A" && i.matrix[row+3][col+3] == "S" {
		return 1
	}

	return 0
}

func (i Input) isXmas(row, col int) bool {
	mCount := 0
	sCount := 0

	if row > 0 && row < len(i.matrix)-1 && col > 0 && col < len(i.matrix[row])-1 {
		if (i.matrix[row-1][col-1] == "M" && i.matrix[row+1][col+1] == "M") || (i.matrix[row-1][col-1] == "S" && i.matrix[row+1][col+1] == "S") || (i.matrix[row+1][col-1] == "M" && i.matrix[row-1][col+1] == "M") || (i.matrix[row+1][col-1] == "S" && i.matrix[row-1][col+1] == "S") {
			return false
		}

		// UpLeft
		if i.matrix[row-1][col-1] == "M" {
			mCount++
		}

		if i.matrix[row-1][col-1] == "S" {
			sCount++
		}

		// UpRight
		if i.matrix[row-1][col+1] == "M" {
			mCount++
		}

		if i.matrix[row-1][col+1] == "S" {
			sCount++
		}

		// DownRight
		if i.matrix[row+1][col+1] == "M" {
			mCount++
		}

		if i.matrix[row+1][col+1] == "S" {
			sCount++
		}

		// DownLeft
		if i.matrix[row+1][col-1] == "M" {
			mCount++
		}

		if i.matrix[row+1][col-1] == "S" {
			sCount++
		}

		return mCount == 2 && sCount == 2
	}

	return false
}

func partOne(input interface{}) string {
	parsedInput := input.(Input)

	sum := 0

	for rowId, row := range parsedInput.matrix {
		for colId, char := range row {
			if char == "X" {
				sum += parsedInput.Up(rowId, colId)
				sum += parsedInput.Down(rowId, colId)
				sum += parsedInput.Left(rowId, colId)
				sum += parsedInput.Right(rowId, colId)
				sum += parsedInput.UpLeft(rowId, colId)
				sum += parsedInput.UpRight(rowId, colId)
				sum += parsedInput.DownLeft(rowId, colId)
				sum += parsedInput.DownRight(rowId, colId)
			}
		}
	}

	return strconv.Itoa(sum)
}

func partTwo(input interface{}) string {
	parsedInput := input.(Input)
	sum := 0

	for rowId, row := range parsedInput.matrix {
		for colId, char := range row {
			if char == "A" {
				if parsedInput.isXmas(rowId, colId) {
					sum++
				}
			}
		}
	}

	return strconv.Itoa(sum)
}

func processInput(input []string) Input {
	matrix := make([][]string, 0)

	for _, line := range input {
		s := strings.Split(line, "")

		matrix = append(matrix, s)
	}

	return Input{
		matrix: matrix,
	}
}

func main() {
	input := utils.ReadInput()

	utils.Run(processInput(input), []utils.Problem{{Problem: partOne}, {Problem: partTwo}})
}
