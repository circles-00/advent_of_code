package main

import (
	"slices"
	"strconv"
	"strings"

	"aoc_2024/utils"
)

type Input struct {
	rules     []string
	orderings [][]string
}

func partOne(i interface{}) string {
	input := i.(Input)
	rulesMap := make(map[string][]string)

	for _, rule := range input.rules {
		ruleSplit := strings.Split(rule, "|")
		left := ruleSplit[0]
		right := ruleSplit[1]

		rulesMap[left] = append(rulesMap[left], right)
	}

	sum := 0
	for _, ordering := range input.orderings {
		isSafe := true
		for i, r := range ordering {
			if i != 0 {
				for j := 0; j < i; j++ {
					if !slices.Contains(rulesMap[ordering[j]], r) {
						isSafe = false
						break
					}
				}
			}
		}

		if !isSafe {
			continue
		}

		if len(ordering)%2 != 0 {
			middle, _ := strconv.Atoi(ordering[len(ordering)/2])
			sum += middle
		}
	}

	return strconv.Itoa(sum)
}

func partTwo(i interface{}) string {
	input := i.(Input)
	rulesMap := make(map[string][]string)

	for _, rule := range input.rules {
		ruleSplit := strings.Split(rule, "|")
		left := ruleSplit[0]
		right := ruleSplit[1]

		rulesMap[left] = append(rulesMap[left], right)
	}

	sum := 0
	for _, ordering := range input.orderings {
		isSafe := true
		for i, r := range ordering {
			if i != 0 {
				for j := 0; j < i; j++ {
					if !slices.Contains(rulesMap[ordering[j]], r) {
						isSafe = false
						ordering[i], ordering[j] = ordering[j], ordering[i]
					}
				}
			}
		}

		if isSafe {
			continue
		}

		if len(ordering)%2 != 0 {
			middle, _ := strconv.Atoi(ordering[len(ordering)/2])
			sum += middle
		}
	}

	return strconv.Itoa(sum)
}

func processInput(input []string) Input {
	rules := make([]string, 0)
	orderings := make([][]string, 0)

	isOrdering := false
	for _, line := range input {
		if len(line) == 0 {
			isOrdering = true
			continue
		}

		if isOrdering {
			orderingRules := strings.Split(line, ",")
			orderings = append(orderings, orderingRules)
			continue
		}

		rules = append(rules, line)
	}

	return Input{
		rules:     rules,
		orderings: orderings,
	}
}

func main() {
	input := utils.ReadInput()

	utils.Run(processInput(input), []utils.Problem{{Problem: partOne}, {Problem: partTwo}})
}
