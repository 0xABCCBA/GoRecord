package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 1
	for i <= 5 {
		fmt.Println(i)
		i += 1
	}

//	for {
//		fmt.Println("infinite loop")
//	}

	if true {
		fmt.Println("This is true")
	}

	if 7 % 2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	//还有else if



	v := 2
	fmt.Print("write ", v, " as")
	switch v {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("it's the weekend")
		default:
			fmt.Println("it's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12 :
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

}
