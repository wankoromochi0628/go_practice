package main

import (
	"fmt"
	"sync"
)

func producer(ch chan int, i int) {
	ch <- i * 2
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println("process", i * 1000)
		wg.Done()
	}
	// 以下処理についてはコードが終了する前に実行されないため表示されない。
	fmt.Println("consumer is done.")
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}

	go consumer(ch, &wg)
	wg.Wait()
	close(ch)
	fmt.Println("main function is done.")
}

