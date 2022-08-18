/*
 * @Author: GG
 * @Date: 2022-08-13 10:10:35
 * @LastEditTime: 2022-08-18 18:02:55
 * @LastEditors: GG
 * @Description: time标准库
 * @FilePath: \golang-demo\time\demo_time.go
 *
 */
package main

import (
	"crypto/md5"
	"fmt"
	"strings"
	"time"
)

// 格式化输出
func test1() {
	now := time.Now()
	fmt.Printf("now: %v\n", now)
	//now: 2022-08-13 10:11:30.4536717 +0800 CST m=+0.036002101

	// 获取单元值
	year := now.Year()
	fmt.Printf("year: %v\n", year)
	month := now.Month()
	fmt.Printf("month: %v\n", month)
	day := now.Day()
	fmt.Printf("day: %v\n", day)
	hour := now.Hour()
	fmt.Printf("hour: %v\n", hour)
	minute := now.Minute()
	fmt.Printf("minute: %v\n", minute)
	second := now.Second()
	fmt.Printf("second: %v\n", second)

	// 自定义格式输出
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
}

// 时间戳
func test2() {
	now := time.Now()
	fmt.Println("uninx:", now.Unix())

	//纳秒时间戳
	fmt.Println("uninxNano:", now.UnixNano())
}

// 时间戳转换普通格式
func timeStampToString() {
	// 获取时间戳
	timeStamp := time.Now().Unix()
	fmt.Printf("timeStamp: %v\n", timeStamp)

	// 将时间戳转换成时间对象
	timeObj := time.Unix(timeStamp, 0)
	fmt.Printf("timeObj: %v\n", timeObj)

	year := timeObj.Year()
	fmt.Printf("year: %v\n", year)
	month := timeObj.Month()
	fmt.Printf("month: %v\n", month)
	day := timeObj.Day()
	fmt.Printf("day: %v\n", day)
	hour := timeObj.Hour()
	fmt.Printf("hour: %v\n", hour)
	minute := timeObj.Minute()
	fmt.Printf("minute: %v\n", minute)
	second := timeObj.Second()
	fmt.Printf("second: %v\n", second)

	// 自定义格式输出
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
}

// 时间操作--增加
func add(h, m, s time.Duration) {
	now := time.Now()
	fmt.Printf("now: %v\n", now)
	fmt.Println(now.Add(time.Hour*h + time.Minute*m + time.Second*s))

}

// 时间操作--相差多少，以使用对象为标准
func sub() {
	now := time.Now()
	targetTime := now.Add(time.Minute * 10)
	fmt.Println("targetTime 为标准，targetTime比now多10 minute", targetTime.Sub(now))
	fmt.Println("now 为标准，now比targetTime少10 minute", now.Sub(targetTime))
}

// 比较时间是否相等，会考虑时区的影响，因此不同时区标准的时间也可以正确比较
func equal() {
	now := time.Now()
	targetTime := now.Add(time.Minute * 10)
	time2 := time.Now()
	fmt.Println(now.Equal(targetTime)) // false
	fmt.Println(now.Equal(time2))      // true
}

// 比较时间前后，接收者在参数前面返回true 反之 false
// 接收者时间小于参数 返回true 反之 false
func before() {
	now := time.Now()
	targetTime := now.Add(time.Minute * 10)
	fmt.Println("now 在 targetTime 时间前面:", now.Before(targetTime)) //true
	fmt.Println("targetTime 在 now 时间前面：", targetTime.Before(now)) //false
}

// 比较时间前后，接收者在参数后面返回true 反之 false
// 接收者时间大于参数返回true 反之 false
func after() {
	now := time.Now()
	targetTime := now.Add(time.Minute * 10)
	fmt.Println("now 在 targetTime 时间后面:", now.After(targetTime)) //false
	fmt.Println("targetTime 在 now 时间后面：", targetTime.After(now)) //true
}

// 自带格式化输出, 不是y-m-d h:m:s 而是2006-01-02 15:04:05
func format() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05 pm Mon Jan"))

	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

// 解析字符串格式的时间
func formatString() {
	now := time.Now()
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))

}
func main() {
	// test2()
	// timeStampToString()
	// add(48, 2, 3)
	// sub()
	// equal()
	// before()
	// after()
	// format()
	formatString()
	str := Md5Crypt("123", "111", "222")
	fmt.Printf("str: %v\n", str)
}

func Md5Crypt(str string, salt ...interface{}) (CryptStr string) {
	fmt.Printf("str: %v\n", str)
	fmt.Printf("salt: %v\n", salt)
	if l := len(salt); l > 0 {
		slice := make([]string, l+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	fmt.Printf("str: %v\n", str)
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
