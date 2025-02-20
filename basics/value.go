package main

import "fmt"

func main() {
	var a int = 6
	var b int = 5
	fmt.Println(a + b) //will print 11

	var firstName string = "Muhammad "
	var middleName string = "Habil "
	var lastName string = "Amardias"
	fmt.Println(firstName + middleName + lastName)

	var isAdmin bool = true
	fmt.Println(isAdmin)

	var c float32 = 6.0
	fmt.Println(a + c) //will throw error, missmatch type

	const pi float32 = 3.14

	pi = 2.7 //will throw an error
}
