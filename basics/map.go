package main

import "fmt"

func main() {
	dict := make(map[string]int)
	dict["id"] = 1
	dict["age"] = 17

	fmt.Println(dict)

	var age int = dict["age"]
	fmt.Println(age)

	fmt.Println(len(dict)) // 2

	delete(dict, "age")
	fmt.Println(len(dict)) // 1

	value, present := dict["age"]
	fmt.Println(present, value)

	value1, present1 := dict["id"]
	fmt.Println(present1, value1)
}
