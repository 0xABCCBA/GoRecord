package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

func vals() (int, int) {
	return 3, 7
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}


func main() {
	v := plus(1, 3)
	fmt.Println("1 + 3 =", v)
	a, b := vals()
	fmt.Println(a, b)

	sum(1, 2)
	sum(1, 2, 3, 4)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

