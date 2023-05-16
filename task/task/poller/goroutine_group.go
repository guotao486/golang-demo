/*
 * @Author: GG
 * @Date: 2023-05-12 09:59:57
 * @LastEditTime: 2023-05-12 10:01:01
 * @LastEditors: GG
 * @Description:
 * @FilePath: \task\poller\goroutine_group.go
 *
 */
package poller

import "sync"

type goroutineGroup struct {
	waitGroup sync.WaitGroup
}

func newRoutineGroup() *goroutineGroup {
	return new(goroutineGroup)
}

// 执行
func (g *goroutineGroup) Run(fn func()) {
	g.waitGroup.Add(1)

	go func() {
		defer g.waitGroup.Done()
		fn()
	}()
}

// 等待完毕
func (g *goroutineGroup) Wait() {
	g.waitGroup.Wait()
}
