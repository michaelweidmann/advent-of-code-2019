package day03

import (
	"fmt"
	"math"
	"strconv"
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

	secondWirePath := parseFile()
	currentLine := line{from: coordinate{0, 0}, to: coordinate{0, 0}}

	for _, step := range secondWirePath {
		direction := step[0]
		value, _ := strconv.Atoi(step[1:])

		calculateIntersections(direction, value, &currentLine)
	}

	shortestDistance := math.MaxInt64

	for _, element := range intersections {
		if element < shortestDistance && element != 0 {
			shortestDistance = element
		}
	}
	fmt.Printf("Solution: %d\n", shortestDistance)
}

func calculateIntersections(direction byte, value int, currentLine *line) {
	currentLine.from.x = currentLine.to.x
	currentLine.from.y = currentLine.to.y

	switch direction {
	case 'R':
		currentLine.to.x += value
		checkIntersectionForHorizontalLine(currentLine)
		break
	case 'L':
		currentLine.to.x -= value
		checkIntersectionForHorizontalLine(currentLine)
		break
	case 'U':
		currentLine.to.y += value
		checkIntersectionForVerticalLine(currentLine)
		break
	case 'D':
		currentLine.to.y -= value
		checkIntersectionForVerticalLine(currentLine)
		break
	}
}

func partTwo() {
	defer util.TimeTrack(time.Now())
}
