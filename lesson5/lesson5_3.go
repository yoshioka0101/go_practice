package main

import(
  "fmt"
  "sync"
)

func producer(ch chan int , i int) {
	ch <- i * 2 
}

func cunsumer(ch chan int, wg *sync.WaitGroup) {
	for i := range ch{
		fmt.Println("product", i * 1000)
		wg.Done()
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}

	go cunsumer(ch, &wg)
	wg.Wait()
	close(ch)

}