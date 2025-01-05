package main

import (
	"fmt"
	"time"
	"sync"
)

func goroutine(s string, wg *sync.WaitGroup) {
	for i := 0; i <= 5; i++{
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
	wg.Done()
}

func nomal(s string) {
	for i := 0; i <= 5; i++{
		time.Sleep(2000 * time.Millisecond)
		fmt.Println(s)
	}
}

func main(){

	var wg sync.WaitGroup
	//wgをいくつ使用するかをカウントしておく
	wg.Add(1)
	go goroutine("world", &wg)
	nomal("hello")
	//ここでWait()が終了されるまでブロックされる
	wg.Wait()
	fmt.Println("END")
}