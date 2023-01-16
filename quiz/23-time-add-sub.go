package quiz

import (
	"fmt"
	"time"
)

func TimeAddSub() {
	start := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	afterTenSeconds := start.Add(time.Second * 10)

	fmt.Printf("start = %v\n", start)
	fmt.Printf("start.Add(time.Second * 10) = %v\n", afterTenSeconds)

	started := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	ended := time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)

	difference := ended.Sub(started)
	fmt.Printf("difference = %v\n", difference)
}
