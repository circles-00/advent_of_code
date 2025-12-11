package main

import (
	"strconv"

	"aoc_2024/utils"
)

type Input struct {
	list []string
}

func getNumberAfterRotation(currentNumber int, rotation string) int {
	direction := rotation[:1]
	distance, err := strconv.Atoi(rotation[1:])
	if err != nil {
		panic("Error parsing distance")
	}

	var newNumber int
	if direction == "L" {
		newNumber = currentNumber - distance
	} else {
		newNumber = currentNumber + distance
	}

	return newNumber % 100
}

func partOne(input interface{}) string {
	parsedInput := input.(Input)
	currentNumber := 50
	password := 0

	for _, line := range parsedInput.list {
		currentNumber = getNumberAfterRotation(currentNumber, line)
		if currentNumber == 0 {
			password++
		}
	}

	return strconv.Itoa(password)
}

func processInput(input []string) Input {
	return Input{
		list: input,
	}
}

func main() {
	input := utils.ReadInput()

	utils.Run(processInput(input), []utils.Problem{{Problem: partOne}})
}
