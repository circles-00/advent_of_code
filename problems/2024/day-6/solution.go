package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc_2024/utils"
)

type Input struct {
	matrix [][]string
}

type Position struct {
	row       int
	col       int
	direction string
}

func (i *Input) findStartPosition() *Position {
	for row := range i.matrix {
		for col := range i.matrix {
			if i.matrix[row][col] == "^" || i.matrix[row][col] == ">" || i.matrix[row][col] == "<" || i.matrix[row][col] == "v" {
				return &Position{row: row, col: col, direction: i.matrix[row][col]}
			}
		}
	}

	return &Position{}
}

func (i *Input) checkIfEnd(p *Position) bool {
	if p.direction == "^" {
		return p.row-1 < 0
	}

	if p.direction == "v" {
		return p.row+1 > len(i.matrix)-1
	}

	if p.direction == ">" {
		return p.col+1 > len(i.matrix[p.row])-1
	}

	if p.direction == "<" {
		return p.col-1 < 0
	}

	return false
}

func (i *Input) checkIfObstacle(p *Position) bool {
	if p.direction == "^" {
		return i.matrix[p.row-1][p.col] == "#"
	}

	if p.direction == "v" {
		return i.matrix[p.row+1][p.col] == "#"
	}

	if p.direction == ">" {
		return i.matrix[p.row][p.col+1] == "#"
	}

	if p.direction == "<" {
		return i.matrix[p.row][p.col-1] == "#"
	}

	return false
}

func (i *Input) changeGuardDirection(p *Position) {
	if p.direction == "^" {
		p.direction = ">"
		return
	}

	if p.direction == ">" {
		p.direction = "v"
		return
	}

	if p.direction == "v" {
		p.direction = "<"
		return
	}

	if p.direction == "<" {
		p.direction = "^"
		return
	}
}

func (i *Input) moveGuard(p *Position) bool {
	isUnique := false

	if i.matrix[p.row][p.col] != "X" {
		isUnique = true
		i.matrix[p.row][p.col] = "X"

	}

	if p.direction == "^" {
		p.row--
	}

	if p.direction == ">" {
		p.col++
	}

	if p.direction == "v" {
		p.row++
	}

	if p.direction == "<" {
		p.col--
	}

	return isUnique
}

func partOne(i interface{}) string {
	input := i.(Input)

	currPosition := input.findStartPosition()
	sum := 1

	for !input.checkIfEnd(currPosition) {
		isObstacle := input.checkIfObstacle(currPosition)

		if isObstacle {
			input.changeGuardDirection(currPosition)
		}

		isUnique := input.moveGuard(currPosition)

		if isUnique {
			sum++
		}
	}

	for row := 0; row < 10; row++ {
		for column := 0; column < 10; column++ {
			fmt.Print(input.matrix[row][column], " ")
		}
		fmt.Print("\n")
	}

	return strconv.Itoa(sum)
}

func partTwo(i interface{}) string {
	return ""
}

func processInput(input []string) Input {
	matrix := make([][]string, 0)

	for _, line := range input {
		matrix = append(matrix, strings.Split(line, ""))
	}

	return Input{
		matrix: matrix,
	}
}

func main() {
	input := utils.ReadInput()

	utils.Run(processInput(input), []utils.Problem{{Problem: partOne}, {Problem: partTwo}})
}
