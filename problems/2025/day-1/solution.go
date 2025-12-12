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

	return ((newNumber % 100) + 100) % 100
}

func countNumberOfZeroJumps(currentNumber int, rotation string) int {
	direction := rotation[:1]
	distance, _ := strconv.Atoi(rotation[1:])

	var firstHit int
	if direction == "L" {
		if currentNumber == 0 {
			firstHit = 100
		} else {
			firstHit = currentNumber
		}
	} else {
		if currentNumber == 0 {
			firstHit = 100
		} else {
			firstHit = 100 - currentNumber
		}
	}

	if distance < firstHit {
		return 0
	}

	return (distance-firstHit)/100 + 1
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

func partTwo(input interface{}) string {
	parsedInput := input.(Input)
	currentNumber := 50
	password := 0

	for _, line := range parsedInput.list {
		numberOfZeroJumps := countNumberOfZeroJumps(currentNumber, line)
		currentNumber = getNumberAfterRotation(currentNumber, line)

		password += numberOfZeroJumps
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

	utils.Run(processInput(input), []utils.Problem{{Problem: partOne}, {Problem: partTwo}})
}
