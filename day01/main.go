package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Run(part string) {
	if part == "1" {
		partOne()
	} else if part == "2" {
		partTwo()
	}
}

func partOne() {
	file, _ := os.Open("day01/input.txt")
	defer file.Close()

	var sum int = 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		sum += calcFuel(&mass)
	}

	fmt.Printf("Solution: %d\n", sum)
}

func partTwo() {
	file, _ := os.Open("day01/input.txt")
	defer file.Close()

	var sum int = 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())

		for calcFuel(&mass) > 0 {
			sum += mass
		}
	}

	fmt.Printf("Solution: %d", sum)
}

/*---------------------------------------------------------------------*/
/*                         Util Function(s)	                           */
/*---------------------------------------------------------------------*/

func calcFuel(mass *int) int {
	*mass = (*mass / 3) - 2
	return *mass
}
