package semaphore

import (
	"sync"
)

//给cond关联
var Mutex = sync.Mutex{}

type Semaphore struct {
	value uint32
	cond  *sync.Cond
}

func NewSemaphore(v uint32) *Semaphore {
	return &Semaphore{
		value: v,
		cond:  sync.NewCond(&Mutex),
	}
}

// sem_wait(consumer) : -1
func (sem *Semaphore) P() {
	Mutex.Lock()

	for sem.value == 0 {
		sem.cond.Wait()
	}

	sem.value--
	Mutex.Unlock()
}

// sem_post(producter) : +1
func (sem *Semaphore) V() {
	Mutex.Lock()

	sem.value++
	sem.cond.Signal()
	Mutex.Unlock()
}
