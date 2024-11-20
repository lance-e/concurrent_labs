package spinlock

type Spinlock struct {
	lock chan bool
}

func NewSpinlock() *Spinlock {
	return &Spinlock{//这里一开始没用指针，后面找KIMI要意见的时候KIMI说用指针会更好，这样就可以创建多个锁，我想确实是，就用了
		lock: make(chan bool, 1)
	}
}

func (sl *Spinlock) Lock() {
	sl.lock <- true 
}

func (sl *Spinlock) Unlock() {
	<-sl.lock
}
