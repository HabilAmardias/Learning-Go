package main

import "fmt"

func main() {
	var i int = 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := range 5 {
		fmt.Println(j)
	}

	var k int = 1
	for {
		if k == 6 {
			break
		}
		fmt.Println(k)
		k++
	}
}
