package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run(part string) {
	if part == "1" {
		partOne()
	} else if part == "2" {
		partTwo()
	}
}

func partOne() {
	intcodes := parseFile()

	intcodes[1] = 12
	intcodes[2] = 2

	fmt.Printf("Solution: %d\n", runProgram(intcodes))
}

func partTwo() {
	intcodes := parseFile()

	tmp := make([]int, len(intcodes))
	copy(tmp, intcodes)

BruteForceLoop:
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			tmp[1] = i
			tmp[2] = j

			if runProgram(tmp) == 19690720 {
				fmt.Printf("Solution: %d\n", 100*i+j)
				break BruteForceLoop
			}

			copy(tmp, intcodes)
		}
	}
}

/*---------------------------------------------------------------------*/
/*                         Util Function(s)	                           */
/*---------------------------------------------------------------------*/

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
	firstSummand := intcodes[firstAddress]
	secondSummand := intcodes[secondAddress]

	return firstSummand + secondSummand
}

func mutliply(firstAddress int, secondAddress int, intcodes []int) int {
	firstSummand := intcodes[firstAddress]
	secondSummand := intcodes[secondAddress]

	return firstSummand * secondSummand
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
