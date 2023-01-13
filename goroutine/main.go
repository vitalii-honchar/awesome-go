package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Updating salutation =", salutation)
		salutation = "welcome"
	}()
	wg.Wait()
	fmt.Println(salutation)
}
