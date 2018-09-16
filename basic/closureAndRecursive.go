package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func fact(n int) int {
	if n <= 0 {
		return 1
	}

	return n * fact(n - 1)
}

func main() {
	v := intSeq()
	fmt.Println(v)
	fmt.Println(v())
	fmt.Println(v())
	fmt.Println(v())

	v1 := intSeq()
	fmt.Println(v1)
	fmt.Println(v1())
	fmt.Println(v1())
	fmt.Println(v1())

	fmt.Println(fact(10))
}

