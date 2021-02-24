package pool

import "sync"

// 新建一个协程池
func NewGoroutinePool(num int) *GoroutinePool {
	g := &GoroutinePool{
		mutex:      sync.RWMutex{},
		number:     num,
		busyNumber: 0,
	}
	return g
}

// 协程池
type GoroutinePool struct {
	mutex      sync.RWMutex // 锁
	number     int          // 数量
	busyNumber int          // 工作数量
}

// 是否全部空闲
func (this *GoroutinePool) Free() bool {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return this.busyNumber == 0
}

// 是否全部繁忙
func (this *GoroutinePool) Busy() bool {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return this.busyNumber == this.number
}

// 是否有空闲协程
func (this *GoroutinePool) AnyFree() bool {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return this.busyNumber < this.number
}

// 是否有空闲协程
func (this *GoroutinePool) hasFree() bool {
	return this.busyNumber < this.number
}

// 让一个协程进行某个动作
func (this *GoroutinePool) Do(task func()) bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if !this.hasFree() {
		return false
	}
	go func() {
		task()
		this.mutex.Lock()
		defer this.mutex.Unlock()
		this.busyNumber--
	}()
	this.busyNumber++
	return true
}

// 当前工作线程数量
func (this *GoroutinePool) GetBusyNumber() int {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return this.busyNumber
}
