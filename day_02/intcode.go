package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	Add = iota + 1
	Multiply
	Stop = 99
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var numbers []int
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		stringNumbers := strings.Split(string(line), ",")
		for _, value := range stringNumbers {
			number, err := strconv.Atoi(strings.TrimSpace(value)); if err != nil {
				fmt.Errorf("line expected to be a number, but got: %s", line)
			}
			numbers = append(numbers, number)
		}
	}

	decodedInput := intcode(numbers)
	fmt.Printf("value at position 0 after decoding: %d\n", decodedInput[0])
}


func intcode(input []int) []int {
	for i := 0; i < len(input); i += 4 {
		var opcode = input[i]

		switch opcode {
		case Add:
			input[input[i+3]] = input[input[i+1]] + input[input[i+2]]
		case Multiply:
			input[input[i+3]] = input[input[i+1]] * input[input[i+2]]
		case Stop:
			return input
		default:
			return input
		}
	}
	return input
}
