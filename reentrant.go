package greetrant

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// RecursiveMutex wraps a Mutex for reentrancy.
type RecursiveMutex struct {
	sync.Mutex
	owner     int64 // goroutine id
	recursion int32 // The number of times this goroutine has re-entered Lock.
}

func (m *RecursiveMutex) Lock() {
	gid := GoID()
	// If the goroutine currently holding the lock is the goroutine called this time,
	// it means reentry
	if atomic.LoadInt64(&m.owner) == gid {
		m.recursion++
		return
	}
	m.Mutex.Lock()
	// The goroutine that acquires the lock calls for the first time,
	// records its goroutine id, and adds 1 to the number of calls
	atomic.StoreInt64(&m.owner, gid)
	m.recursion = 1
}

func (m *RecursiveMutex) Unlock() {
	gid := GoID()
	// The goroutine that does not hold the lock tries to release the lock, wrong use.
	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.owner, gid))
	}
	// The number of calls is reduced by 1.
	m.recursion--
	// If the goroutine has not been completely released, return directly.
	if m.recursion != 0 {
		return
	}
	// The last time this goroutine was called, the lock needs to be released.
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}
