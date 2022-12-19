/*
 * @Author: GG
 * @Date: 2022-12-19 14:28:01
 * @LastEditTime: 2022-12-19 16:10:57
 * @LastEditors: GG
 * @Description: 责任链模式
 * @FilePath: \设计模式\责任链模式\main.go
 *
 */
package main

import "fmt"

const (
	// 标准日志等级
	StandardLogLevel = iota

	// info 日志等级
	InfoLogLevel

	// 错误 日志等级
	ErrorLogLevel
)

// 日志接口
type BaseLogger interface {
	PrintLog(level int, message string)
	Write(message string)
}

// 标准日志类
type StandardLogger struct {
	Level      int
	NextLogger BaseLogger
}

// 实例化 标准日志类
func NewStandardLogger() *StandardLogger {
	return &StandardLogger{
		Level:      StandardLogLevel,
		NextLogger: nil,
	}
}

// 标准日志类 写入日志
func (sl *StandardLogger) Write(message string) {
	fmt.Printf("标准日志输出：%s.\n", message)
}

// 标准日志类，输入日志并流向下一个对象方法
func (sl *StandardLogger) PrintLog(level int, message string) {
	if sl.Level == level {
		sl.Write(message)
	}

	if sl.NextLogger != nil {
		sl.NextLogger.PrintLog(level, message)
	}
}

// 标准日志类，设置下一个对象方法
func (sl *StandardLogger) SetNextLogger(logger BaseLogger) {
	sl.NextLogger = logger
}

// 提示日志类
type InfoLogger struct {
	Level      int
	NextLogger BaseLogger
}

func NewInfoLogger() *InfoLogger {
	return &InfoLogger{
		Level:      InfoLogLevel,
		NextLogger: nil,
	}
}

func (il *InfoLogger) Write(message string) {
	fmt.Printf("信息日志输出：%s.\n", message)
}

func (il *InfoLogger) PrintLog(level int, message string) {
	if il.Level == level {
		il.Write(message)
	}
	if il.NextLogger != nil {
		il.NextLogger.PrintLog(level, message)
	}
}

func (il *InfoLogger) SetNextLogger(logger BaseLogger) {
	il.NextLogger = logger
}

type ErrorLogger struct {
	Level      int
	NextLogger BaseLogger
}

func NewErrorLogger() *ErrorLogger {
	return &ErrorLogger{
		Level:      ErrorLogLevel,
		NextLogger: nil,
	}
}

func (el *ErrorLogger) Write(message string) {
	fmt.Printf("错误日志输出：%s.\n", message)
}
func (el *ErrorLogger) PrintLog(level int, message string) {
	if el.Level == level {
		el.Write(message)
	}

	if el.NextLogger != nil {
		el.NextLogger.PrintLog(level, message)
	}
}

func (el *ErrorLogger) SetNextLogger(logger BaseLogger) {
	el.NextLogger = logger
}

// 获取日志对象链
func GetChainOfLogger() BaseLogger {
	errLog := NewErrorLogger()
	infoLog := NewInfoLogger()
	standardLog := NewStandardLogger()

	errLog.SetNextLogger(infoLog)
	infoLog.SetNextLogger(standardLog)

	return errLog
}
func main() {
	logChain := GetChainOfLogger()
	logChain.PrintLog(InfoLogLevel, "这是一条消息")
	logChain.PrintLog(ErrorLogLevel, "这是一条错误日志")
	logChain.PrintLog(StandardLogLevel, "这是一条标准消息")
}
