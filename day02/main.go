package day02

import (
	"fmt"
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

/*---------------------------------------------------------------------*/
/*                         		Part 1		                           */
/*---------------------------------------------------------------------*/

func partOne() {
	defer util.TimeTrack(time.Now())
	intcodes := parseFile()

	initializeMemory(12, 2, intcodes)
	fmt.Printf("Solution: %d\n", runProgram(intcodes))
}

/*---------------------------------------------------------------------*/
/*                         		Part 2		                           */
/*---------------------------------------------------------------------*/

// The input allows to calculate the noun first and then the verb
// This is due to the big jumps of the result when the noun is increased
// When increasing the verb the result does small steps
// So first the noun is raised to about a correct solution.
// Then a "small" correction is made by adjusting the value of the verb.
func partTwo() {
	defer util.TimeTrack(time.Now())
	intcodes := parseFile()

	tmpIntcodes := make([]int, len(intcodes))

	noun, verb := 0, 0
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
