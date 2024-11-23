package semaphore

import (
	"sync/atomic"
)

type Semaphore struct {
	n int32 //you need to write here
}

func NewSemaphore(v int) *Semaphore {
	return &Semaphore{
		int32(v), //you need to write here
	}

}

// sem_wait(consumer) : -1
func (sem *Semaphore) P() {
	atomic.AddInt32(&sem.n, 1)

}

// sem_post(producter) : +1
func (sem *Semaphore) V() {
	atomic.AddInt32(&sem.n, -1)

}
