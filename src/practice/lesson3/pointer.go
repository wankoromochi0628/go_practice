package main

import (
	"fmt"
)

func questionOne() {
	var i int = 100
	var p *int
	p = &i
	var j int
	j = *p
	fmt.Println(j)
}

func questionTwo() {
	var i int = 100
	var j int = 200
	var p1 *int
	var p2 *int
	p1 = &i
	p2 = &j
	i = *p1 + *p2
	p2 = p1
	j = *p2 + i
	fmt.Println(j)
}

func main() {
	var n int = 100
	// int型
	fmt.Println(n)

	// メモリのアドレス
	fmt.Println(&n)

	var p *int = &n

	//ポインタ
	fmt.Println(p)
	fmt.Println(*p)

	*p = 200

	fmt.Println(p)
	fmt.Println(&p)
	fmt.Println(*p)

	questionOne()
	questionTwo()
}
