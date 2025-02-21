package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func substract(a, b int) int {
	return a - b
}

func addAndSubstract(a, b int) (int, int) {
	return a + b, a - b
}

func sum(nums ...int) int {
	var total int = 0
	for _, num := range nums {
		total += num
	}
	return total
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var addResult int = add(3, 5)
	fmt.Println(addResult)

	var substractResult int = substract(4, 6)
	fmt.Println(substractResult)

	addRes, substractRes := addAndSubstract(1, 3)
	fmt.Println(addRes, substractRes)

	nums := []int{1, 2, 3, 4, 5}
	var sumOfNums int = sum(nums...)
	fmt.Println(sumOfNums)

	fmt.Println(fibonacci(7))
}
