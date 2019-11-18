package main

import (
	"fmt"
	"sync"
)

func world (w string) {
	for i := 0; i < 5; i++ {
		//time.Sleep(100 * time.Millisecond)
		fmt.Println(w)
	}
}

/*
 * wg.Done() : 並列処理の終了を通知する
 */
func hello (h string, wg *sync.WaitGroup) {
	// defer : 処理の最後に実行することを宣言する
	defer wg.Done()

	for i := 0; i < 5; i++ {
		//time.Sleep(100 * time.Millisecond)
		fmt.Println(h)
	}
}

/*
 * sync.WaitGroup : 複数の並列処理(goroutine)がある場合、全ての並列処理の終了を待つための値
 * wg.Add(n) : WaitGroupに並列処理の個数がn個存在することを教える
 * wg.Wait() : wg.Add(n) で与えた個数分の wg.Done()が通知されるまで処理を終了させない
 */
func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	world("world!")
	go hello("hello,", &wg)

	wg.Wait()
}