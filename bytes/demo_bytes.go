/*
 * @Author: GG
 * @Date: 2022-08-12 14:41:56
 * @LastEditTime: 2022-08-12 17:27:56
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\bytes\demo_bytes.go
 *
 */
package main

import (
	"bytes"
	"fmt"
)

func testReader() {
	data := "1234567890"

	// 通过[]byte创建reader
	re := bytes.NewReader([]byte(data))

	//返回未读取部分长度
	fmt.Printf("re.Len(): %v\n", re.Len())
	//返回底层数据总长度
	fmt.Printf("re.Size(): %v\n", re.Size())
	fmt.Println("---------------------")

	// 创建2字节大小的变量
	buf := make([]byte, 2)
	for {
		//读取数据
		n, err := re.Read(buf)
		if err != nil {
			break
		}
		fmt.Println(string(buf[:n]))
	}
	fmt.Println("---------------------")

	// 设置偏移量，上面的操作修改了读取位置
	re.Seek(0, 0)
	for {
		//一个字节的读
		b, err := re.ReadByte()
		if err != nil {
			break
		}
		fmt.Println(string(b))
	}

	fmt.Println("------------------------")

	re.Seek(0, 0)
	off := int64(0)
	buf = make([]byte, 2)
	for {
		// 指定偏移量读取,读2长度，从0开始
		n, err := re.ReadAt(buf, off)
		if err != nil {
			break
		}
		off += int64(n)
		fmt.Println(off, string(buf[:n]))
	}
}

func testBuffer() {
	// 声明一个buffer
	var a bytes.Buffer                  //直接定义一个buffer变量，不用初始化，可以直接使用
	b := new(bytes.Buffer)              // 使用new返回一个buffer变量
	c := bytes.NewBuffer([]byte("123")) // 从一个[]byte切片，构造一个buffer
	d := bytes.NewBufferString("hello") // 从一个string变量，构造一个buffer

	// 往buffer写入数据
	b.Write([]byte("hello")) // 将切片d写入buffer尾部
	b.WriteString("world")   // 将字符串写入buffer尾部
	b.WriteByte(byte(1))// 将字符西融入buffer
	b.WriteRune("你好") // 将一个rune类型数据放入buffer尾部

	c := make([]byte,8)
b.Read(c) //一次读取8个byte到c容器中，每次读取新的8个byte覆盖c中原来的内容
b.ReadByte() //读取第一个byte，b的第1个byte被拿掉，赋值给 a => a, _ := b.ReadByte()
b.ReadRune() //读取第一个rune，b的第1个rune被拿掉，赋值给 r =>  r, _ := b.ReadRune()
b.ReadBytes(delimiter byte) //需要一个 byte作为分隔符 ，读的时候从缓冲器里找第一个出现的分隔符（delim），找到后，把从缓冲器头部开始到分隔符之间的所有byte进行返回，作为byte类型的slice，返回后，缓冲器也会空掉一部分
b.ReadString(delimiter byte) // 需要一个byte作为分隔符，读的时候从缓冲器里找第一个出现的分隔符（delim），找到后，把从缓冲器头部开始到分隔符之间的所有byte进行返回， 作为字符串返回 ，返回后，缓冲器也会空掉一部分b.ReadForm(i io.Reader) // 从一个实现io.Reader接口的r，把r里的内容读到缓冲器里 ，n 返回读的数量

file, _ := os.Open(".text.txt")  
buf := bytes.NewBufferString("Hello world")  
buf.ReadFrom(file) 
//将text.txt内容追加到缓冲器的尾部
fmt.Println(buf.String())
清空数据
b.Reset()
转换为字符串
b.String()

}

func main() {
	// 类型强转
	var i int = 1
	var j byte = 2
	j = byte(i)
	fmt.Printf("j: %v\n", j)

	// contains  是否包含
	var s = []byte("hello")
	sublice1 := []byte("hell")
	sublice2 := []byte("Hell")
	fmt.Println(bytes.Contains(s, sublice1)) //true
	fmt.Println(bytes.Contains(s, sublice2)) //false

	// count 字符出现次数
	s = []byte("helloooooo")
	sep1 := []byte("h")
	sep2 := []byte("l")
	sep3 := []byte("o")
	fmt.Println(bytes.Count(s, sep1)) //1
	fmt.Println(bytes.Count(s, sep2)) //2
	fmt.Println(bytes.Count(s, sep3)) //6

	// repeat 重复输出
	b := []byte("h1")
	fmt.Printf("b: %v\n", b)
	fmt.Printf("b: %v\n", string(b))        //hi
	fmt.Println(string(bytes.Repeat(b, 1))) //hi
	fmt.Println(string(bytes.Repeat(b, 3))) //hihihi
	// replace 替换
	s = []byte("hello,world")
	old := []byte("o")
	news := []byte("ee")
	fmt.Println(string(bytes.Replace(s, old, news, 0)))  // hello,world
	fmt.Println(string(bytes.Replace(s, old, news, 1)))  // hellee,world
	fmt.Println(string(bytes.Replace(s, old, news, 2)))  // hellee,weerld
	fmt.Println(string(bytes.Replace(s, old, news, -1))) // hellee,weerld
	// runes 转换成汉字
	s = []byte("你好世界")
	r := bytes.Runes(s)
	fmt.Println(len(s)) //12
	fmt.Println(len(r)) //4

	// join join格式字符连接
	s2 := [][]byte{[]byte("你好"), []byte("世界")}
	sep4 := []byte(",")
	fmt.Println(string(bytes.Join(s2, sep4)))
	sep5 := []byte("#")
	fmt.Println(string(bytes.Join(s2, sep5)))

	testReader()
}
