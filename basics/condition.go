package main

import (
	"fmt"
	"time"
)

func checkType(i interface{}) {
	switch i.(type) {
	case bool:
		fmt.Println("Bool")
	case int:
		fmt.Println("Integer")
	case float32:
		fmt.Println("Float")
	case float64:
		fmt.Println("Double")
	default:
		fmt.Println("Don't know type")
	}
}

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

	checkType(true)
}
