package main

import (
	"fmt"
)

func main() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

	s := l[0]

	for _, v := range l {
		if s > v {
			s = v
		}
	}

	fmt.Println(s)

	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}

	t := 0
	for _, v := range m {
		t += v
	}

	fmt.Println(t)
}
