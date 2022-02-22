package main

import (
	"fmt"
	"runtime"
	"time"
)

func determine_go_runtime() {
	switch go_runtime := runtime.GOOS; go_runtime {
	case "darwin":
		fmt.Println("Go is running on Darwin")
	case "linux":
		fmt.Println("Go is running on Linux")
	default:
		fmt.Printf("Go is running on %s\n", go_runtime)
	}
}

func determine_when_is_saturday() {
	today := time.Now().Weekday()
	saturday := time.Saturday
	switch saturday {
	case today:
		fmt.Println("Today is Saturday")
	case today + 1:
		fmt.Println("Tomorrow is Saturday")
	case today + 2:
		fmt.Println("Day after tomorrow is Saturday")
	default:
		fmt.Println("There are more than 2 days to Saturday")
	}
}

func main() {
	determine_go_runtime()
	determine_when_is_saturday()
}
