package task

// 状态枚举
const (
	PROCESSED = iota // 待处理
	RUNNING          // 执行中
	FINISH           // 完成
	FAIL             // 失败
)

// 类型枚举
const (
	IntervalTask = iota // 间隔任务，每间隔时间执行一次
	DelayTask           // 延迟任务，延迟时间执行
	SimpleTask          // 即时执行
	TimerTask           // 定时任务
)
