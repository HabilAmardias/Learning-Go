package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now().Weekday()

	if today == time.Saturday || today == time.Sunday {
		fmt.Println("Weekend")
	} else {
		fmt.Println("Weekday")
	}

	switch today {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}
}
