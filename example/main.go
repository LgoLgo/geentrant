package main

import (
	"fmt"
	"sync"

	"github.com/LgoLgo/geentrant"
)

func main() {
	var wg sync.WaitGroup
	rm := &greetrant.RecursiveMutex{}

	// This first goroutine locks and unlocks the recursive mutex recursively 5 times.
	wg.Add(1)
	go func() {
		for i := 0; i < 5; i++ {
			rm.Lock()
			fmt.Println("goroutine 1 locked")
		}
		for i := 0; i < 5; i++ {
			rm.Unlock()
			fmt.Println("goroutine 1 unlocked")
		}
		wg.Done()
	}()

	// This second goroutine tries to unlock the mutex without locking it, which should fail with a panic.
	wg.Add(1)
	go func() {
		defer func() {
			if r := recover(); r == nil {
				fmt.Println("Unexpected result: should have panicked")
			}
			wg.Done()
		}()
		rm.Unlock()
	}()
	wg.Wait()
}
