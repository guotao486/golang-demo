/*
 * @Author: GG
 * @Date: 2022-12-19 11:42:08
 * @LastEditTime: 2022-12-19 11:50:13
 * @LastEditors: GG
 * @Description: 策略模式
 * @FilePath: \设计模式\策略模式\main.go
 *
 */
package main

import "fmt"

/*
策略模式(Strategy Pattern)，一个类的行为或其算法可以在运行时改变。

策略模式定义一系列算法，把他们一个一个封装起来，并且使他可相互替换。

下面演示图书打折实例，把每种打折算法都封装起来，并抽象出接口，客户端只依赖接口。
*/

// 打折接口
type IDiscount interface {
	Discount() float32
}

// 图书类
type Book struct {
	Price float32
}

// 读取图书折扣价格
func (b *Book) GetPrice(discount IDiscount) float32 {
	return b.Price * discount.Discount()
}

// 8.5折类
type Discount85 struct{}

// 8.5折扣
func (d *Discount85) Discount() float32 {
	return 0.85
}

// 6.5折类
type Discount65 struct{}

// 6.5折扣
func (d *Discount65) Discount() float32 {
	return 0.65
}

func main() {
	b := Book{100}
	p := b.GetPrice(&Discount85{})
	fmt.Printf("p: %v\n", p)

	p2 := b.GetPrice(&Discount65{})
	fmt.Printf("p2: %v\n", p2)
}
