package worker

import "sync"

const (
	StatusRun  = "RUN"  // 运行中
	StatusQuit = "QUIT" // 退出
)

// 新建一个工作者
func NewWorker(num int, fun func(), loop ...bool) *Worker {
	worker := &Worker{
		threadNum: num,
		task:      fun,
		wait:      sync.WaitGroup{},
	}
	if len(loop) != 0 {
		worker.loop = loop[0]
	}
	return worker
}

// 工作者
type Worker struct {
	threadNum int            // 协程数
	task      func()         // 任务
	status    string         // 状态
	loop      bool           // 是否循环
	wait      sync.WaitGroup // 等待锁
}

// 获取协程数
func (this *Worker) GetThreadNum() int {
	return this.threadNum
}

// 获取状态
func (this *Worker) GetStatus() string {
	return this.status
}

// 设置协程数
func (this *Worker) SetThreadNum(num int) {
	this.threadNum = num
}

// 运行函数
func (this *Worker) runFunc() func() {
	var fun func()
	if this.loop {
		fun = func() {
			for {
				if this.status == StatusRun {
					this.task()
				}
			}
		}
	} else {
		fun = func() {
			if this.status == StatusRun {
				this.task()
			}
		}
	}
	return func() {
		defer this.wait.Done()
		fun()
	}
}

// 运行
func (this *Worker) Start() {
	this.status = StatusRun
	if this.task == nil {
		return
	}
	for i := 0; i < this.threadNum; i++ {
		this.wait.Add(1)
		go this.runFunc()()
	}
}

// 停止
func (this *Worker) Exit() {
	this.status = StatusQuit
}

// 等待
func (this *Worker) Wait() {
	this.wait.Wait()
}
