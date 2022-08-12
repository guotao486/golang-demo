/*
 * @Author: GG
 * @Date: 2022-08-12 20:48:26
 * @LastEditTime: 2022-08-12 23:46:35
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\sort\demo_sort.go
 *
 */
package main

import (
	"fmt"
	"sort"
)

// type Interface interface {
//     Len() int           // Len方法返回集合中的元素个数
//     Less(i, j int) bool // i>j，该方法返回索引i的元素是否比索引j的元素小、
//     Swap(i, j int)      // 交换i, j的值
// }

// 利用sort库，创建自定义排序
type NewInts []uint

func (n NewInts) Len() int {
	return len(n)
}

func (n NewInts) Less(i, j int) bool {
	fmt.Println(i, j, n[i] < n[j], n)
	return n[i] < n[j]
}

func (n NewInts) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func testUintSort() {
	n := []uint{1, 3, 2}

	sort.Sort(NewInts(n))
	fmt.Println(n) // 1 2 3
}

// 自定义end

// sort默认的3个类型
func testInt() {
	// []int
	f2 := []int{3, 5, 1, 2, 4}
	sort.Ints(f2)
	fmt.Printf("f2: %v\n", f2)
	// f: [1 2 3 4 5]
}
func testFloat64() {
	// []float64
	f1 := []float64{1.1, 4.4, 5.5, 3.3, 2.2}
	sort.Float64s(f1)
	fmt.Printf("f1: %v\n", f1)
	// f: [1.1 2.2 3.3 4.4 5.5]
}
func testString() {
	// []string
	//字符串排序，现比较高位，相同的再比较低位
	// [] string
	ls := sort.StringSlice{
		"100",
		"42",
		"41",
		"3",
		"2",
	}
	fmt.Println(ls) //[100 42 41 3 2]
	sort.Strings(ls)
	fmt.Println(ls) //[100 2 3 41 42]

	//字符串排序，现比较高位，相同的再比较低位
	ls = sort.StringSlice{
		"d",
		"ac",
		"c",
		"ab",
		"e",
	}
	fmt.Println(ls) //[d ac c ab e]
	sort.Strings(ls)
	fmt.Println(ls) //[ab ac c d e]

	//汉字排序，依次比较byte大小
	ls = sort.StringSlice{
		"啊",
		"博",
		"次",
		"得",
		"饿",
		"周",
	}
	fmt.Println(ls) //[啊 博 次 得 饿 周]
	sort.Strings(ls)
	fmt.Println(ls) //[博 周 啊 得 次 饿]

	for _, v := range ls {
		fmt.Println(v, []byte(v))
	}

	//博 [229 141 154]
	//周 [229 145 168]
	//啊 [229 149 138]
	//得 [229 190 151]
	//次 [230 172 161]
	//饿 [233 165 191]
}

// sort默认的3个类型 end
func main() {

}
