/*
 * @Author: GG
 * @Date: 2023-05-13 10:09:15
 * @LastEditTime: 2023-05-15 10:37:38
 * @LastEditors: GG
 * @Description:
 * @FilePath: \task\task\task.go
 *
 */
package task

type TaskEntity interface {
	Run(taskS *TaskStore) error
}

var EntityContainer = make(map[string]TaskEntity)

func EntityRegister(name string, entity TaskEntity) {
	if entity == nil {
		panic("task: Register entity is nil")
	}
	if _, dup := Provides[name]; dup {
		panic("task: Register called twice for entity " + name)
	}
	EntityContainer[name] = entity
}

type TaskStore struct {
	ID            int
	Title         string
	Group         string
	Rules         string
	Data          string
	Type          int
	State         int
	CreateTime    int64
	ExecutionTime int64
	UpdateTime    int64
}

func NewTaskStore() *TaskStore {
	return &TaskStore{}
}
