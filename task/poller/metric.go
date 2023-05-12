package poller

import "sync/atomic"

type metric struct {
	busyWorkers uint64
}

func newMetric() *metric {
	return &metric{}
}

// 忙碌的进程 +1
func (m *metric) IncBusyWorkers() uint64 {
	return atomic.AddUint64(&m.busyWorkers, 1)
}

// 忙碌的进程 -1
func (m *metric) DecBusyWorkers() uint64 {
	return atomic.AddUint64(&m.busyWorkers, ^uint64(0))
}

// 读取忙碌的进程数量
func (m *metric) BusyWorkers() uint64 {
	return atomic.LoadUint64(&m.busyWorkers)
}
