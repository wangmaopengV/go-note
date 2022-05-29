package main

import (
	"fmt"
	"time"
)

func recv(data chan string) {

	for e := range data {

		fmt.Println(e)
	}

	fmt.Println("close")
}

// 谁写谁close chan
func main() {

	test := make(chan string, 1)
	go recv(test)

	test <- "test"

	time.Sleep(3 * time.Second)

	close(test)

	for e := range test {

		fmt.Println(e)
	}
	time.Sleep(3 * time.Second)
}
