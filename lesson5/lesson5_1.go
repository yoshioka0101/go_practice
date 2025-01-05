package main

import (
	"fmt"
	"time"
)

func goroutine(s string) {
	for i := 0; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func nomal(s string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(s)
		time.Sleep(100 * time.Millisecond) // 同じ待機時間を追加
	}
}

func main() {
	go goroutine("Hello")
	nomal("World")
}
