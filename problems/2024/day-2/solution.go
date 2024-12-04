package main

import (
	"strconv"
	"strings"

	"aoc_2024/utils"
)

type Input struct {
	list [][]int
}

func excludeIndex(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return append([]int{}, slice...)
	}

	result := make([]int, 0, len(slice)-1)
	result = append(result, slice[:index]...)
	result = append(result, slice[index+1:]...)
	return result
}

func isListSafe(l []int) bool {
	increasing := -1
	isSafe := true

	for i, v := range l {
		if i == len(l)-1 {
			continue
		}

		diff := v - l[i+1]
		absDiff := utils.AbsInt(diff)

		if i == 0 && diff < 0 {
			increasing = 1
		}

		if i == 0 && diff > 0 {
			increasing = 0
		}

		if diff == 0 || absDiff > 3 {
			isSafe = false
			break
		}

		if (increasing == 0 && diff < 0) || (increasing == 1 && diff > 0) {
			isSafe = false
			break
		}
	}

	return isSafe
}

func partOne(input interface{}) string {
	parsedInput := input.(Input)
	sum := 0

	for _, l := range parsedInput.list {
		isSafe := isListSafe(l)

		if isSafe {
			sum++
		}
	}

	return strconv.Itoa(sum)
}

func partTwo(input interface{}) string {
	parsedInput := input.(Input)
	sum := 0

	for _, l := range parsedInput.list {
		increasing := -1
		unsafeIdx := -1

		for i, v := range l {
			if i == len(l)-1 {
				continue
			}

			diff := v - l[i+1]
			absDiff := utils.AbsInt(diff)

			if i == 0 && diff < 0 {
				increasing = 1
			}

			if i == 0 && diff > 0 {
				increasing = 0
			}

			if unsafeIdx == -1 && (diff == 0 || absDiff > 3) {
				unsafeIdx = i
				continue
			}

			if unsafeIdx == -1 && ((increasing == 0 && diff < 0) || (increasing == 1 && diff > 0)) {
				unsafeIdx = i
				continue
			}
		}

		if unsafeIdx == -1 {
			sum++
			continue
		}

		newArr1 := excludeIndex(l, unsafeIdx)
		newArr2 := excludeIndex(l, unsafeIdx-1)
		newArr3 := excludeIndex(l, unsafeIdx+1)

		isSafe := isListSafe(newArr1)

		if !isSafe {
			isSafe = isListSafe(newArr2)
		}

		if !isSafe {
			isSafe = isListSafe(newArr3)
		}

		if isSafe {
			sum++
		}
	}

	return strconv.Itoa(sum)
}

func processInput(input []string) Input {
	processedInput := make([][]int, 0)

	for _, line := range input {
		split := strings.Fields(line)
		parsedLine := make([]int, 0)

		for _, pL := range split {
			i, _ := strconv.ParseInt(pL, 10, 0)

			parsedLine = append(parsedLine, int(i))
		}

		processedInput = append(processedInput, parsedLine)
	}

	return Input{
		list: processedInput,
	}
}

func main() {
	input := utils.ReadInput()

	utils.Run(processInput(input), []utils.Problem{{Problem: partOne}, {Problem: partTwo}})
}
