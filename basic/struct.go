package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Bob", age: 12})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name, s.age)
	sp := &s
	fmt.Println(sp.name)

	sp.age = 51
	fmt.Println(s)

}
