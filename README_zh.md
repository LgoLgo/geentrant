![geentrant](img/geentrant.png)

[English](README.md) | 中文

Go 语言实现的轻量级重入锁。

## 安装

要安装此包，您需要先安装 Go 并设置您的 Go 工作区。

1. 你首先需要安装 [Go](https:golang.org)，然后你可以使用下面的命令来安装 geentrant。

```sh
go get -u github.com/LgoLgo/geentrant
```

2. 将其导入您的代码：

```go
import "github.com/LgoLgo/geentrant"
```

## 快速开始

```sh
# 假设 example 文件夹中有以下代码
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

## 许可证

项目在 Apache 许可证 2.0 版下开源。