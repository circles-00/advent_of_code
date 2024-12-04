package main

import (
	"sort"
	"strconv"
	"strings"

	"aoc_2024/utils"
)

type Input struct {
	list1 []int
	list2 []int
}

func partOne(input interface{}) string {
	parsedInput := input.(Input)

	sort.Ints(parsedInput.list1)
	sort.Ints(parsedInput.list2)
	sum := 0

	for i := 0; i < len(parsedInput.list1); i++ {
		diff := utils.AbsIntDiff(parsedInput.list1[i], parsedInput.list2[i])
		sum += diff
	}

	return strconv.Itoa(sum)
}

func partTwo(input interface{}) string {
	parsedInput := input.(Input)

	frequency := make(map[int]int, 0)

	for _, value := range parsedInput.list2 {
		frequency[value]++
	}

	similiarity := 0

	for _, value := range parsedInput.list1 {
		similiarity += (value * frequency[value])
	}

	return strconv.Itoa(similiarity)
}

func processInput(input []string) Input {
	parsedInput := make([][]int, 2)

	for _, line := range input {
		numbers := strings.Fields(line)
		num1, _ := strconv.ParseInt(numbers[0], 10, 0)
		num2, _ := strconv.ParseInt(numbers[1], 10, 32)

		parsedInput[0] = append(parsedInput[0], int(num1))
		parsedInput[1] = append(parsedInput[1], int(num2))
	}

	return Input{
		list1: parsedInput[0],
		list2: parsedInput[1],
	}
}

func main() {
	input := utils.ReadInput()

	utils.Run(processInput(input), []utils.Problem{{Problem: partOne}, {Problem: partTwo}})
}
