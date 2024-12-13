package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	//quit := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)
	go consumer(ch, &wg)
	go producer(ch, &wg)
	wg.Wait()
}
func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case i, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println("Received value:", i)
		}
	}
}
func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	//quit <- true
	close(ch)
}
