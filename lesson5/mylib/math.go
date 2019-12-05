package mylib

import "golang.org/x/tools/go/ssa/interp/testdata/src/fmt"

// UpperCase : public
var Public string = "Public"
// LowerCase : private
var private string = "private"

func Average(s []int)int {
	total := 0

	for _, i := range s {
		total += i
	}

	fmt.Println(Public)
	fmt.Println(private)

	return int(total/len(s))
}