package task

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"task/task/poller"
	"time"
)

type Manager struct {
	poller   *poller.Poller
	provider TaskProvider
}

var Provides = make(map[string]TaskProvider)

func NewManager(provideName string, workerNum int) (*Manager, error) {
	provider, ok := Provides[provideName]
	if !ok {
		return nil, fmt.Errorf("POLL: unknown provide %q (forgotten import?)", provideName)
	}
	poller, _ := poller.NewPoller(workerNum)
	return &Manager{provider: provider, poller: poller}, nil
}

// 开始调度
func (m *Manager) Run(ctx context.Context) error {
	for {
		// step 1
		m.poller.Schedule() // 开始调度

		select {
		case <-m.poller.Ready: // goroutine准备好后，这里会有消息,没有则阻塞
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
				task, err := m.fetch(ctx)
				if err != nil {
					log.Println("fetch task error:", err.Error())
					break
				}
				if task == nil {
					// log.Println("fetch task is nil")
					break
				}
				m.poller.IncBusyWorkers() // 正在运行的goroutine +1

				// step 3
				m.provider.Run(task)
				m.poller.Run(func() { // 执行任务
					if err := m.Execute(ctx, task); err != nil {
						m.provider.Fail(task)
						log.Println("execute task error:", err.Error())
					}
				})
				break LOOP
			}
		}
	}
}

// 获取任务
func (m *Manager) fetch(ctx context.Context) (*TaskStore, error) {
	taskS := m.provider.Get()
	return taskS, nil
}

// 执行任务
func (m *Manager) Execute(ctx context.Context, taskS *TaskStore) error {
	defer func() {
		m.poller.DecBusyWorkers() // 执行完成后，忙碌进程数 -1
		m.poller.Schedule()       // 执行完成后调度下一个goroutine
	}()

	select {
	case <-ctx.Done():
		return nil
	default:
		if taskS.Rules == "" {
			return errors.New("task rules is nil")
		}

		rules := make(map[string]interface{})
		json.Unmarshal([]byte(taskS.Rules), &rules)
		entityName, ok := rules["entity"]
		if !ok {
			return errors.New("task rules entity is nil")
		}
		entity, ok := EntityContainer[entityName.(string)]
		if !ok {
			return fmt.Errorf("task entity %s not register", entityName.(string))
		}

		err := entity.Run(taskS)
		if err != nil {
			return err
		}
		m.provider.Finish(taskS)
	}

	return nil
}

// 添加任务
func (m *Manager) AddTask(taskName, groupName string, rules, data map[string]interface{}, typeIota int, ExecutionTime int64) {
	taskStore := NewTaskStore()
	taskStore.Title = taskName
	taskStore.Group = groupName
	taskStore.Rules = GetMapToString(rules)
	taskStore.Data = GetMapToString(data)
	taskStore.Type = typeIota
	taskStore.State = PROCESSED
	taskStore.CreateTime = time.Now().Unix()
	taskStore.UpdateTime = taskStore.CreateTime
	taskStore.ExecutionTime = ExecutionTime

	e := m.provider.Add(taskStore)
	if e != nil {
		log.Println(e)
	}
}

// rules map[string]interface{}
// rules[entity] TaskEntityName
// rules[duration] time.Duration

// 添加间隔任务，每间隔一段时间执行一次
func (m *Manager) AddIntervalTask(taskName, groupName, entityName string, IntervalTime time.Duration, data map[string]interface{}) error {
	rules := make(map[string]interface{})
	rules["entity"] = entityName
	rules["duration"] = IntervalTime
	ExecutionTime := time.Now().Add(IntervalTime).Unix()
	m.AddTask(taskName, groupName, rules, data, IntervalTask, ExecutionTime)
	return nil
}

// 添加延迟任务，延迟一段时间执行
func (m *Manager) AddDelayTask(taskName, groupName, entityName string, DelayTime time.Duration, data map[string]interface{}) error {
	rules := make(map[string]interface{})
	rules["entity"] = entityName
	rules["duration"] = DelayTime
	ExecutionTime := time.Now().Add(DelayTime).Unix()
	m.AddTask(taskName, groupName, rules, data, DelayTask, ExecutionTime)
	return nil
}

// 添加即时任务
func (m *Manager) AddSimpleTask(taskName, groupName, entityName string, data map[string]interface{}) error {
	rules := make(map[string]interface{})
	rules["entity"] = entityName
	ExecutionTime := time.Now().Unix()
	m.AddTask(taskName, groupName, rules, data, SimpleTask, ExecutionTime)
	return nil
}

// 添加定时任务，定时一个时间执行
func (m *Manager) AddTimerTask(taskName, groupName, entityName string, timer time.Time, data map[string]interface{}) error {
	rules := make(map[string]interface{})
	rules["entity"] = entityName
	rules["duration"] = timer
	ExecutionTime := timer.Unix()
	m.AddTask(taskName, groupName, rules, data, TimerTask, ExecutionTime)
	return nil
}

func GetMapToString(m map[string]interface{}) string {
	s, e := json.Marshal(m)
	if e != nil {
		log.Panicln(e)
	}
	return string(s)
}
