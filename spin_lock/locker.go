package spinlock

import (
	"sync/atomic"
)

type Spinlock struct {
	value int32
}

func (sl *Spinlock) Lock() {
	for {
		if atomic.CompareAndSwapInt32(&sl.value, 0, 1) {
			return
		}
	}
}

func (sl *Spinlock) Unlock() {
	sl.value = 0
}
