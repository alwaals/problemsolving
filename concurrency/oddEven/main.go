package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	quit := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)
	go odd(ch1, ch2, quit, &wg)
	go even(ch1, ch2, quit, &wg)
	wg.Wait()
}
func odd(ch1 chan int, ch2 chan int, quit chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	ch2 <- 1
	for {
		select {
		case i := <-ch1:
			fmt.Println("Even:", i)
			if i == 10 {
				quit <- true
				close(ch1)
				close(ch2)
				return
			}
			ch2 <- i + 1
		case <-quit:
			return
		}
	}
}
func even(ch1 chan int, ch2 chan int, quit chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case i := <-ch2:
			fmt.Println("Odd:", i)
			ch1 <- i + 1
		case <-quit:
			return
		}
	}
}
