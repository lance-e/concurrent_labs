package spinlock

import (
	"sync/atomic"
)

type Spinlock struct {
	value int32
}

func (sl *Spinlock) Lock() {
	for !atomic.CompareAndSwapInt32(&sl.value, 0, 1) {

	}
}

func (sl *Spinlock) Unlock() {
	atomic.StoreInt32(&sl.value, 0)
}

//使用了原子操作
//查了一下因为好像是多核同时goroutine时，如果同时检查到lock值为false
//就会同时设置为true，导致自旋锁不可用。
//找kimi问用的原子操作
//原子操作确保了在多处理器或多核心系统中，某个操作在执行过程中不会被其他处理器或核心上的线程或goroutine中断。
//不知道还有没有别的修改办法
//求解答！！！
