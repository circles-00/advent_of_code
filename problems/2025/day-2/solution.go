package main

import (
	"strconv"
	"strings"

	"aoc_2024/utils"
)

type Input struct {
	intervals []string
}

func isInvalidID(id string) bool {
	if len(id)%2 != 0 {
		return false
	}

	return id[:len(id)/2] == id[len(id)/2:]
}

func getInvalidIdsFromInterval(interval string) []int {
	intervalSplit := strings.Split(interval, "-")
	start, _ := strconv.Atoi(intervalSplit[0])
	end, _ := strconv.Atoi(intervalSplit[1])
	invalidIds := make([]int, 0)

	currentID := intervalSplit[0]

	for start <= end {
		if isInvalidID(currentID) {
			invalidIds = append(invalidIds, start)
		}

		start++
		currentID = strconv.Itoa(start)
	}

	return invalidIds
}

func partOne(input interface{}) string {
	parsedInput := input.(Input)
	sum := 0

	for _, interval := range parsedInput.intervals {
		invalidIds := getInvalidIdsFromInterval(interval)

		for _, id := range invalidIds {
			sum += id
		}
	}

	return strconv.Itoa(sum)
}

func partTwo(input interface{}) string {
	return ""
}

func processInput(input string) Input {
	intervals := strings.Split(strings.TrimSpace(input), ",")

	return Input{
		intervals: intervals,
	}
}

func main() {
	input := utils.ReadInputRaw()

	utils.Run(processInput(input), []utils.Problem{{Problem: partOne}, {Problem: partTwo}})
}
