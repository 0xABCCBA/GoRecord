package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")
	time.Sleep(2)
//	var input string
//	fmt.Scanln(&input)
//	fmt.Println("done")

	messages := make(chan string)
	go func() { messages <- "ping"}()
	msg := <-messages
	fmt.Println(msg)

	m2 := make(chan string, 2)
	m2 <- "buffered"
	m2 <- "channel"

	fmt.Println(<-m2)
	fmt.Println(<-m2)

}
