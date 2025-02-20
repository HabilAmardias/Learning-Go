// before start, i want to say, whoever think name "Slice" for datatype that has "slice" operation is good, I hope you have a good day

package main

import (
	"fmt"
	"slices"
)

func main() {
	var a []string
	fmt.Println(a)

	s := make([]int, 3)
	for i := 0; i < 3; i++ {
		s[i] = i + 1
	}
	fmt.Println(s, len(s))

	t := []int{4, 5, 6}
	fmt.Println(t)

	fmt.Println(t[0:2]) //[4,5]

	fmt.Println(slices.Equal(s, t)) //false
}
