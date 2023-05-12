package poller

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

type Poller struct {
	routineGroup *goroutineGroup // 并发控制
	workerNum    int             // 记录同时在运行的最大goroutine数

	sync.Mutex
	ready  chan struct{} // 某个goroutine已准备好
	metric *metric       // 统计当前运行中的goroutine数量
}

func NewPoller(workerNum int) *Poller {
	return &Poller{
		routineGroup: newRoutineGroup(),
		workerNum:    workerNum,
		ready:        make(chan struct{}, 1),
		metric:       newMetric(),
	}
}

// 调度器
func (p *Poller) schedule() {
	p.Lock()
	defer p.Unlock()
	// 正在运行的进程数量 >= 最大运行数量，直接返回
	if int(p.metric.BusyWorkers()) >= p.workerNum {
		return
	}

	select {
	case p.ready <- struct{}{}: // 只要满足当前goroutine数量小于最大goroutine数量就通知poll去调度goroutine执行任务
	default:
	}
}

func (p *Poller) Poll(ctx context.Context) error {
	for {
		// step 1
		p.schedule() // 开始调度

		select {
		case <-p.ready: // goroutine准备好后，这里会有消息,没有则阻塞
		case <-ctx.Done():
			return nil
		}

	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			default:
				// step 2
				task, err := p.fetch(ctx)
				if err != nil {
					log.Println("fetch task error:", err.Error())
					break
				}
				fmt.Printf("task: %v\n", task)
				p.metric.IncBusyWorkers() // 正在运行的goroutine +1

				// step 3
				p.routineGroup.Run(func() { // 执行任务
					if err := p.execute(ctx, task); err != nil {
						log.Println("execute task error:", err.Error())
					}
				})

				break LOOP
			}
		}
	}
}

// 获取任务
func (p *Poller) fetch(ctx context.Context) (string, error) {
	time.Sleep(1 * time.Second)
	return "task", nil
}

// 执行任务
func (p *Poller) execute(ctx context.Context, task string) error {
	defer func() {
		p.metric.DecBusyWorkers() // 执行完成后，忙碌进程数 -1
		p.schedule()              // 执行完成后调度下一个goroutine
	}()
	return nil
}
