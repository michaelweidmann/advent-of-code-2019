package util

import (
	"fmt"
	"time"
)

func TimeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("It took %s to calculate!", elapsed)
}
