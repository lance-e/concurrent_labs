package semaphore

type Semaphore struct {
	v chan int
}

func NewSemaphore(v int) *Semaphore {
	return &Semaphore{
		v: make(chan int, v),
	}
}

// sem_wait(consumer) : -1
func (sem *Semaphore) P() {
	sem.v <- 1
}

// sem_post(producter) : +1
func (sem *Semaphore) V() {
	<-sem.v
}
