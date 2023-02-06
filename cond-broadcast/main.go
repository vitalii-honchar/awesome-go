package main

import (
	"fmt"
	"sync"
	"time"
)

type Button struct {
	Clicked *sync.Cond
}

func main() {
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()

			for {
				c.L.Lock()
				c.Wait()
				fn()
				c.L.Unlock()
			}

		}()
		goroutineRunning.Wait()
	}

	//var clickRegistered sync.WaitGroup
	//clickRegistered.Add(3)

	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window")
		//clickRegistered.Done()
	})

	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		//clickRegistered.Done()
	})

	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		//clickRegistered.Done()
	})

	button.Clicked.Broadcast()
	time.Sleep(2 * time.Second)
	button.Clicked.Broadcast()

	time.Sleep(2 * time.Second)
	button.Clicked.Broadcast()

	//clickRegistered.Wait()
	time.Sleep(5 * time.Second)
}
