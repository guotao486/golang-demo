package entity

import (
	"fmt"
	"task/task"
	"time"
)

type OrderExpire struct {
}

func NewOrderExpire() OrderExpire {
	return OrderExpire{}
}

func (oe OrderExpire) Run(taskS *task.TaskStore) error {
	fmt.Println("orderExpire run...")
	fmt.Printf("time:%s task id:%d title:%s group:%s ", time.Now(), taskS.ID, taskS.Title, taskS.Group)
	return nil
}

type OrderCreate struct{}

func NewOrderCreate() OrderCreate {
	return OrderCreate{}
}

func (oc OrderCreate) Run(taskS *task.TaskStore) error {
	fmt.Println("orderCreate run...")
	fmt.Printf("time:%s task id:%d title:%s group:%s ", time.Now(), taskS.ID, taskS.Title, taskS.Group)
	return nil
}
