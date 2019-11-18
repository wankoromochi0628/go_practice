package main

import "fmt"

func main() {
	ch := make(chan string, 3)
	ch <- "a"
	fmt.Println(len(ch))
	ch <- "b"
	fmt.Println(len(ch))
	close(ch)

	for c := range ch {
		fmt.Println(c)
	}
}
