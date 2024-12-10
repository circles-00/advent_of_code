package main

import (
	"bytes"
	"strconv"
	"strings"

	"aoc_2024/utils"
)

type Input struct {
	raw string
}

func partOne(input interface{}) string {
	parsedInput := input.(Input)
	currMulOp, leftOp, rightOp := bytes.NewBufferString(""), bytes.NewBufferString(""), bytes.NewBufferString("")

	sum := 0

	for _, rune := range parsedInput.raw {
		if (currMulOp.Len() == 0 && rune == 'm') || (currMulOp.Len() == 1 && rune == 'u') || (currMulOp.Len() == 2 && rune == 'l') {
			currMulOp.WriteRune(rune)
			continue
		}

		if currMulOp.String() == "mul" {
			if rune == '(' {
				currMulOp.WriteRune(rune)

				continue
			}

			currMulOp.Reset()
		}

		if strings.Contains(currMulOp.String(), "mul(") {
			if rune == ',' {
				currMulOp.WriteRune(rune)
				continue
			}

			if rune == ')' && leftOp.Len() > 0 && rightOp.Len() > 0 {
				left, _ := strconv.Atoi(leftOp.String())
				right, _ := strconv.Atoi(rightOp.String())
				sum += left * right

				currMulOp.Reset()
				leftOp.Reset()
				rightOp.Reset()
				continue
			}

			if !strings.Contains(currMulOp.String(), ",") && (rune >= 48 && rune <= 57) {
				leftOp.WriteRune(rune)
				continue
			}

			if strings.Contains(currMulOp.String(), ",") && (rune >= 48 && rune <= 57) {
				rightOp.WriteRune(rune)
				continue
			}

			currMulOp.Reset()
			leftOp.Reset()
			rightOp.Reset()
		}

	}

	return strconv.Itoa(sum)
}

func partTwo(input interface{}) string {
	parsedInput := input.(Input)
	currMulOp, leftOp, rightOp, switchOp := bytes.NewBufferString(""), bytes.NewBufferString(""), bytes.NewBufferString(""), bytes.NewBufferString("")

	enabledOp := true

	sum := 0

	for _, rune := range parsedInput.raw {
		if (switchOp.Len() == 0 && rune == 'd') || (switchOp.Len() == 1 && rune == 'o') || (switchOp.Len() == 2 && rune == '(') || (switchOp.Len() == 3 && rune == ')') {
			switchOp.WriteRune(rune)
			continue
		}

		if switchOp.String() == "do()" {
			enabledOp = true
			switchOp.Reset()
		}

		if (switchOp.Len() == 0 && rune == 'd') || (switchOp.Len() == 1 && rune == 'o') || (switchOp.Len() == 2 && rune == 'n') || (switchOp.Len() == 3 && rune == '\'') || (switchOp.Len() == 4 && rune == 't') || (switchOp.Len() == 5 && rune == '(') || (switchOp.Len() == 6 && rune == ')') {
			switchOp.WriteRune(rune)
			continue
		}

		if switchOp.String() == "don't()" {
			enabledOp = false
			switchOp.Reset()
		}

		if !enabledOp {
			continue
		}

		if (currMulOp.Len() == 0 && rune == 'm') || (currMulOp.Len() == 1 && rune == 'u') || (currMulOp.Len() == 2 && rune == 'l') {
			currMulOp.WriteRune(rune)
			continue
		}

		if currMulOp.String() == "mul" {
			if rune == '(' {
				currMulOp.WriteRune(rune)

				continue
			}

			currMulOp.Reset()
		}

		if strings.Contains(currMulOp.String(), "mul(") {
			if rune == ',' {
				currMulOp.WriteRune(rune)
				continue
			}

			if rune == ')' && leftOp.Len() > 0 && rightOp.Len() > 0 {
				left, _ := strconv.Atoi(leftOp.String())
				right, _ := strconv.Atoi(rightOp.String())
				sum += left * right

				currMulOp.Reset()
				leftOp.Reset()
				rightOp.Reset()
				continue
			}

			if !strings.Contains(currMulOp.String(), ",") && (rune >= 48 && rune <= 57) {
				leftOp.WriteRune(rune)
				continue
			}

			if strings.Contains(currMulOp.String(), ",") && (rune >= 48 && rune <= 57) {
				rightOp.WriteRune(rune)
				continue
			}

			currMulOp.Reset()
			leftOp.Reset()
			rightOp.Reset()
		}

	}

	return strconv.Itoa(sum)
}

func processInput(input string) Input {
	return Input{raw: input}
}

func main() {
	input := utils.ReadInputRaw()

	utils.Run(processInput(input), []utils.Problem{{Problem: partOne}, {Problem: partTwo}})
}
