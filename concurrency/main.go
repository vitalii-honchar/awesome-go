package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	c := provideValue(&wg)
	wg.Wait()

	fmt.Printf("Value = %v\n", <-c)
}

func provideValue(wg *sync.WaitGroup) <-chan int {
	ch := make(chan int, 1)
	go func() {
		defer fmt.Println("Channel closed")
		defer close(ch)
		defer wg.Done()

		ch <- 10
	}()
	return ch
}
