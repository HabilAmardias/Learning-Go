package main

import "fmt"

func sum(a int) {
	a++
}

func sumPtr(aptr *int) {
	*aptr++
}

func main() {
	var a int = 5
	fmt.Println(a)

	sum(a)
	fmt.Println(a) //5

	sumPtr(&a)
	fmt.Println(a) //6

	fmt.Println(&a) //memory address
}
