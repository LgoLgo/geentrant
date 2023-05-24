package greetrant

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func GoID() int64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	// Get the second word of the first line, which is the goroutine id in hex.
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.ParseInt(idField, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
