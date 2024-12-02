package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type Problem struct {
	Problem func(interface{}) string
}

const (
	green = "\033[32m"
	reset = "\033[0m"
)

func AbsInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func ReadInput() []string {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Cannot read input")
	}

	defer file.Close()

	input := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func Run(input interface{}, problems []Problem) {
	for idx, p := range problems {
		start := time.Now()

		duration := time.Since(start)

		fmt.Printf("Result of part %d:%s %s %s\n", idx+1, green, p.Problem(input), reset)

		fmt.Printf("Execution time:%s %s %s\n", green, duration, reset)
	}
}
