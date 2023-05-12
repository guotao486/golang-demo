/*
 * @Author: GG
 * @Date: 2023-05-12 10:59:57
 * @LastEditTime: 2023-05-12 11:18:35
 * @LastEditors: GG
 * @Description:
 * @FilePath: \task\test.go
 *
 */
package main

import (
	"context"
	"fmt"
	"task/poller"

	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	fmt.Println("start")
	goleak.VerifyTestMain(m)
}

func TestPoller(t *testing.T) {
	producer := poller.NewPoller(5)
	producer.Poll(context.Background())
}
