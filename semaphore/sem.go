package semaphore

import (
	"sync"
)

type Semaphore struct {
	count     int           // 当前可用的信号量数量
	mu        sync.Mutex    // 互斥锁，用于保护count变量
	waitQueue chan struct{} // 用于阻塞等待信号量的goroutine
}

func NewSemaphore(v int) *Semaphore {
	return &Semaphore{
		count:     v,
		waitQueue: make(chan struct{}, v), // 初始化通道，大小为信号量初始值，用于存储等待的goroutine
	}
}

// P操作：尝试获取一个信号量，如果不足则等待
func (sem *Semaphore) P() {
	sem.mu.Lock()
	if sem.count > 0 {
		sem.count-- // 如果信号量足够，减少计数
		sem.mu.Unlock()
		return
	}
	// 如果信号量不足，当前goroutine进入等待队列
	sem.mu.Unlock()
	sem.waitQueue <- struct{}{}
}

// V操作：释放一个信号量，如果有goroutine在等待，则唤醒一个
func (sem *Semaphore) V() {
	sem.mu.Lock()
	if len(sem.waitQueue) > 0 {
		// 如果有goroutine在等待，唤醒一个
		<-sem.waitQueue
	} else {
		sem.count++ // 如果没有goroutine等待，增加信号量计数
	}
	sem.mu.Unlock()
}

//这个不会写/(ㄒoㄒ)/~~
//用ai跑了一下，只能看懂，但是写不出来
