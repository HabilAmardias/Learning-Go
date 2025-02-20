package main

import "fmt"

func main() {
	var a [5]float32
	fmt.Println(a) //[0,0,0,0,0]

	a[1] = 50
	fmt.Println(a) //[0,50,0,0,0]

	fmt.Println(len(a)) //5

	b := [3]int{1, 2, 3}
	fmt.Println(b)

	twoDim := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(twoDim)
}
