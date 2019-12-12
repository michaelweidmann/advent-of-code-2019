package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/michaelweidmann/advent-of-code-2019/util"
)

func Run(part string) {
	if part == "1" {
		partOne()
	} else if part == "2" {
		partTwo()
	}
}

func partOne() {
	defer util.TimeTrack(time.Now())
	intcodes := parseFile()

	initializeMemory(12, 2, intcodes)
	fmt.Printf("Solution: %d\n", runProgram(intcodes))
}

func partTwo() {
	defer util.TimeTrack(time.Now())
	intcodes := parseFile()

	tmpIntcodes := make([]int, len(intcodes))

	var noun int = 0
	var verb int = 0

	calculateNoun(intcodes, tmpIntcodes, &noun, &verb)
	calculateVerb(intcodes, tmpIntcodes, &noun, &verb)

	fmt.Printf("Solution: %d\n", 100*noun+verb)
}

func calculateNoun(intcodes []int, tmpIntcodes []int, noun *int, verb *int) {
	for ; *noun <= 100; *noun++ {
		copy(tmpIntcodes, intcodes)
		initializeMemory(*noun, *verb, tmpIntcodes)

		if runProgram(tmpIntcodes) >= 19690720 {
			*noun = *noun - 1
			break
		}
	}
}

func calculateVerb(intcodes []int, tmpIntcodes []int, noun *int, verb *int) {
	for ; *verb <= 100; *verb++ {
		copy(tmpIntcodes, intcodes)
		initializeMemory(*noun, *verb, tmpIntcodes)

		if runProgram(tmpIntcodes) == 19690720 {
			break
		}
	}
}

/*---------------------------------------------------------------------*/
/*                         Util Function(s)	                           */
/*---------------------------------------------------------------------*/

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
