package day03

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

// Used for the value to a key in a map because it does not allocate any memory
type voidValue struct{}

type coordinate struct {
	x, y int
}

type line struct {
	from, to coordinate
}

// Two maps (because there are no sets) to save all lines of the first wire
// As value the voidValue is used because it does allocate any memory
var firstWireHorizontalLines map[line]voidValue = make(map[line]voidValue)
var firstWireVerticalLines map[line]voidValue = make(map[line]voidValue)

var void voidValue

var intersections []int

func checkIntersectionForVerticalLine(secondWireLine *line) {
	for firstWireLine, _ := range firstWireHorizontalLines {
		if isBetween(secondWireLine.from.x, firstWireLine.from.x, firstWireLine.to.x) &&
			isBetween(firstWireLine.from.y, secondWireLine.from.y, secondWireLine.to.y) {
			intersections = append(intersections, calculateManhattanDistance(secondWireLine.from.x, firstWireLine.from.y))
		}
	}
}

func checkIntersectionForHorizontalLine(secondWireLine *line) {
	for firstWireLine, _ := range firstWireVerticalLines {
		if isBetween(secondWireLine.from.y, firstWireLine.from.y, firstWireLine.to.y) &&
			isBetween(firstWireLine.from.x, secondWireLine.from.x, secondWireLine.to.x) {
			intersections = append(intersections, calculateManhattanDistance(firstWireLine.from.x, secondWireLine.from.y))
		}
	}
}

// Checks if a number is in between a range given by two points
func isBetween(number int, firstEndpoint int, secondEndpoint int) bool {
	max := math.Max(float64(firstEndpoint), float64(secondEndpoint))
	min := firstEndpoint

	if max == float64(firstEndpoint) {
		min = secondEndpoint
	}

	return min <= number && float64(number) <= max
}

func calculateManhattanDistance(x int, y int) int {
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func parseFile() []string {
	// Open file
	file, _ := os.Open("day03/input.txt")
	defer file.Close()

	// Scan file
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	path := strings.Split(scanner.Text(), ",")

	parseFirstWire(path)

	scanner.Scan()
	return strings.Split(scanner.Text(), ",")
}

func parseFirstWire(path []string) {
	currentPosition := coordinate{0, 0}

	for _, step := range path {
		direction := step[0]
		value, _ := strconv.Atoi(step[1:])

		newPosition := coordinate{currentPosition.x, currentPosition.y}

		switch direction {
		case 'R':
			newPosition.x += value
			firstWireHorizontalLines[line{currentPosition, newPosition}] = void
			break
		case 'L':
			newPosition.x -= value
			firstWireHorizontalLines[line{currentPosition, newPosition}] = void
			break
		case 'U':
			newPosition.y += value
			firstWireVerticalLines[line{currentPosition, newPosition}] = void
			break
		case 'D':
			newPosition.y -= value
			firstWireVerticalLines[line{currentPosition, newPosition}] = void
			break
		}

		currentPosition = newPosition
	}
}
