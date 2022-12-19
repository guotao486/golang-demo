/*
 * @Author: GG
 * @Date: 2022-12-19 10:24:26
 * @LastEditTime: 2022-12-19 11:09:49
 * @LastEditors: GG
 * @Description: 解释器模式
 * @FilePath: \设计模式\解释器模式\main.go
 *
 */
package main

import (
	"fmt"
	"strings"
)

// 语句接口
type Expression interface {
	Interpret(context string) bool
}

// 终端语句类
type TerminalExpression struct {
	Data string
}

// 实例化终端语句类
func NewTerminal(data string) *TerminalExpression {
	return &TerminalExpression{
		Data: data,
	}
}

// 终端语句类的解释器
func (te *TerminalExpression) Interpret(context string) bool {
	//判断字符串 context 中是否包含个子串 te.Data 。包含或者 te.Data 为空则返回true
	if strings.Contains(context, te.Data) {
		return true
	}
	return false
}

// 或语句类
type OrExpression struct {
	Expr1 Expression
	Expr2 Expression
}

// 实例化或语句类
func NewOr(expr1, expr2 Expression) *OrExpression {
	return &OrExpression{
		Expr1: expr1,
		Expr2: expr2,
	}
}

// 或语句 解释器
func (or *OrExpression) Interpret(context string) bool {
	return or.Expr1.Interpret(context) || or.Expr2.Interpret(context)
}

// 并语句
type AndExpression struct {
	Expr1 Expression
	Expr2 Expression
}

// 实例化 并语句
func NewAnd(expr1, expr2 Expression) *AndExpression {
	return &AndExpression{
		Expr1: expr1,
		Expr2: expr2,
	}
}

// 并语句 解释器
func (an *AndExpression) Interpret(context string) bool {
	return an.Expr1.Interpret(context) && an.Expr2.Interpret(context)
}

func main() {

	// or------
	lee := NewTerminal("Lee")
	wang := NewTerminal("wang")
	isMale := NewOr(lee, wang)

	b := isMale.Interpret("Lee")
	fmt.Printf("b: %v\n", b)
	fmt.Println("--------")

	yang := NewTerminal("yang")
	married := NewTerminal("married")
	isMarried := NewAnd(yang, married)
	b2 := isMarried.Interpret("married yang")
	fmt.Printf("b2: %v\n", b2)
}
