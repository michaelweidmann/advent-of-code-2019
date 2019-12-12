package day02

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func initializeMemory(noun int, verb int, intcodes []int) {
	intcodes[1] = noun
	intcodes[2] = verb
}

func runProgram(intcodes []int) int {
ProgramLoop:
	for i := 0; i < len(intcodes); i += 4 {
		instruction := intcodes[i]
		firstAddress := intcodes[i+1]
		secondAddress := intcodes[i+2]
		resultAddress := intcodes[i+3]

		switch instruction {
		case 1:
			intcodes[resultAddress] = add(firstAddress, secondAddress, intcodes)
			break
		case 2:
			intcodes[resultAddress] = mutliply(firstAddress, secondAddress, intcodes)
			break
		case 99:
			// fmt.Println("Program finished.")
			break ProgramLoop
		default:
			panic("Invalid program.")
		}
	}

	return intcodes[0]
}

func add(firstAddress int, secondAddress int, intcodes []int) int {
	return intcodes[firstAddress] + intcodes[secondAddress]
}

func mutliply(firstAddress int, secondAddress int, intcodes []int) int {
	return intcodes[firstAddress] * intcodes[secondAddress]
}

func parseFile() []int {
	file, _ := os.Open("day02/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	stringIntcodes := strings.Split(scanner.Text(), ",")
	var intcodes = []int{}

	for _, stringIntcode := range stringIntcodes {
		intcode, _ := strconv.Atoi(stringIntcode)
		intcodes = append(intcodes, intcode)
	}

	return intcodes
}
