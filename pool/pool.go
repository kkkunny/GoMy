package pool

import "sync"

// 新建一个协程池
func NewGoroutinePool(num int) *GoroutinePool {
	g := &GoroutinePool{
		mutex:      &sync.RWMutex{},
		complete:   make(chan int, num),
		number:     num,
		busyNumber: 0,
	}
	return g
}

// 协程池
// 线程不安全
type GoroutinePool struct {
	mutex      *sync.RWMutex // 条件锁
	complete   chan int      // 任务完成信号
	number     int           // 数量
	busyNumber int           // 工作数量
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

// 任务
func (this *GoroutinePool) task(task func()) {
	task()
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.busyNumber--
	this.complete <- 1
}

// 做任务
func (this *GoroutinePool) doTask(task func()) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	go this.task(task)
	this.busyNumber++
}

// 用空闲协程进行一项任务
func (this *GoroutinePool) Do(task func()) {
loop:
	for {
		select {
		case <-this.complete:
			break loop
		default:
			if this.AnyFree() {
				break loop
			}
		}
	}
	this.doTask(task)
}

// 当前工作线程数量
func (this *GoroutinePool) GetBusyNumber() int {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return this.busyNumber
}
