package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/michaelweidmann/advent-of-code-2019/day01"
	"github.com/michaelweidmann/advent-of-code-2019/day02"
	"github.com/michaelweidmann/advent-of-code-2019/day03"
)

func main() {
	arguments := os.Args

	if len(arguments) != 3 {
		printUsage(arguments[0])
		os.Exit(1)
	}

	dayNumber, _ := strconv.Atoi(arguments[1])

	switch dayNumber {
	case 1:
		day01.Run(arguments[2])
		break
	case 2:
		day02.Run(arguments[2])
		break
	case 3:
		day03.Run(arguments[2])
		break
	default:
		printUsage(arguments[0])
		os.Exit(1)
	}
}

func printUsage(programName string) {
	fmt.Printf("Usage: %s [<day-number>] [<part-number>]\n\n", programName)
	fmt.Println("Allowed <day-number> values are: 1 to 24")
	fmt.Println("Allowed <part-number> values are: 1 or 2")
}
