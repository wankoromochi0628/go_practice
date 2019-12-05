package main

import (
	"./mylib"
	"fmt"
)

func main() {
	s := []int {1, 2, 3, 4, 5}

	// UpperCaseのpublic変数にはアクセス可
	fmt.Println(mylib.Public)
	// LowerCaseのprivate変数にはアクセス不可
	//fmt.Println(mylib.private)

	fmt.Println(mylib.Average(s))
}
