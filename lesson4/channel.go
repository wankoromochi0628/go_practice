package main

import "fmt"

func goroutine(s []int, c chan int) {
	t := 0

	for _, v := range s {
		t += v
	}

	c <- t
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)

	go goroutine(s, c)
	x := <- c
	fmt.Println(x)
}
