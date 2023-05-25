![geentrant](img/geentrant.png)

English | [中文](README_zh.md)

Lightweight reentrant lock implemented in Go.


## Installation

To install this package, you need to install Go and set your Go workspace first.

1. You first need [Go](https://golang.org/) installed, then you can use the below Go command to install geentrant.

```sh
go get -u github.com/LgoLgo/geentrant
```

2. Import it in your code:

```go
import "github.com/LgoLgo/geentrant"
```

## Quick start

```sh
# assume the following codes in example folder
$ cat example/main.go
```

```go
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
```

## License

This project is under the Apache License 2.0. See the LICENSE file for the full license text.