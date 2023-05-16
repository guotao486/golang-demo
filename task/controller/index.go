package controller

import (
	"net/http"
	"task/task"
	"time"
)

type IndexController struct {
	task *task.Manager
}

func NewIndexController(taskManager *task.Manager) *IndexController {
	return &IndexController{
		task: taskManager,
	}
}

func (index *IndexController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	// 定时任务
	index.task.AddIntervalTask("orderCreate", "订单成功，发送积分", "orderCreate", 10*time.Second, data)
	// 延时任务
	index.task.AddDelayTask("orderExpire", "订单失效", "orderExpire", 60*time.Second, data)
	// 立即执行
	index.task.AddSimpleTask("orderCreate", "订单成功通知", "orderCreate", data)
	// 定时执行
	t, _ := time.Parse("2006-01-02 15:04:05", "2023-05-16 16:15:00")
	index.task.AddTimerTask("orderCreate", "订单售后", "orderCreate", t, data)

}

func (index *IndexController) OrderExpire(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	index.task.AddDelayTask("orderExpire", "订单失效", "orderExpire", 60*time.Second, data)
}
