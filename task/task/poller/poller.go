package poller

import (
	"sync"
)

type Poller struct {
	routineGroup *goroutineGroup // 并发控制
	workerNum    int             // 记录同时在运行的最大goroutine数

	sync.Mutex
	Ready  chan struct{} // 某个goroutine已准备好
	metric *metric       // 统计当前运行中的goroutine数量
}

func NewPoller(workerNum int) (*Poller, error) {
	return &Poller{
		routineGroup: newRoutineGroup(),
		workerNum:    workerNum,
		Ready:        make(chan struct{}, 1),
		metric:       newMetric(),
	}, nil
}

// 调度器
func (p *Poller) Schedule() {
	p.Lock()
	defer p.Unlock()
	// 正在运行的进程数量 >= 最大运行数量，直接返回
	if int(p.metric.BusyWorkers()) >= p.workerNum {
		return
	}

	select {
	case p.Ready <- struct{}{}: // 只要满足当前goroutine数量小于最大goroutine数量就通知poll去调度goroutine执行任务
	default:
	}
}

// 执行进程数量+1
func (p *Poller) IncBusyWorkers() {
	p.metric.IncBusyWorkers()
}

// 执行完成 进程数量-1
func (p *Poller) DecBusyWorkers() {
	p.metric.DecBusyWorkers()
}

// 执行任务
func (p *Poller) Run(fn func()) {
	p.routineGroup.Run(fn)
}
