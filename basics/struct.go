package main

import "fmt"

type member struct {
	name string
	age  int
}

func newMember(name string, age int) member {
	m := member{name: name, age: age}
	return m
}

//method
func (m member) getAge() int {
	return m.age
}

func main() {
	a := newMember("Kevin", 18)
	fmt.Println(a.age)
	fmt.Println(a.name)
	fmt.Println(a.getAge())
}
