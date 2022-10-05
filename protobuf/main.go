/*
 * @Author: GG
 * @Date: 2022-10-05 14:55:09
 * @LastEditTime: 2022-10-05 15:10:45
 * @LastEditors: GG
 * @Description:
 * @FilePath: \protobuf\main.go
 *
 */
package main

import (
	"fmt"
	"protobuf/user"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	testJson()
}

// 序列化和反序列号
func test() {
	userModel := &user.User{
		Uid:   1,
		Name:  "tom",
		Email: "12312312",
	}

	// 序列化成二进制数据
	bytes, _ := proto.Marshal(userModel)
	fmt.Printf("bytes: %v\n", bytes)

	// 反序列化，将二进制转化成结构体
	otherUser := &user.User{}
	proto.Unmarshal(bytes, otherUser)
	fmt.Printf("otherUser.GetUid(): %v\n", otherUser.GetUid())
	fmt.Printf("otherUser.Uid: %v\n", otherUser.Uid)
	fmt.Printf("otherUser.GetName(): %v\n", otherUser.GetName())
	fmt.Printf("otherUser.GetEmail(): %v\n", otherUser.GetEmail())

	//bytes: [8 1 18 3 116 111 109 26 8 49 50 51 49 50 51 49 50]
	//otherUser.GetUid(): 1
	//otherUser.Uid: 1
	//otherUser.GetName(): tom
	//otherUser.GetEmail(): 12312312
}

func testJson() {
	userModel := &user.User{
		Uid:   2,
		Name:  "tom",
		Email: "12312312",
	}
	fmt.Printf("userModel: %v\n", userModel)

	fmt.Printf("userModel.ProtoReflect().Interface(): %v\n", userModel.ProtoReflect().Interface())
	// 转化成json
	jsonString := protojson.Format(userModel.ProtoReflect().Interface())
	fmt.Printf("jsonString: %v\n", jsonString)

	// json转化成结构体
	m := userModel.ProtoReflect().Interface()
	protojson.Unmarshal([]byte(jsonString), m)
	fmt.Printf("m: %v\n", m)

	//userModel: uid:2 name:"tom" email:"12312312"
	//userModel.ProtoReflect().Interface(): uid:2 name:"tom" email:"12312312"
	//jsonString: {
	//"uid": 2,
	//"name": "tom",
	//"email": "12312312"
	//}
	//m: uid:2 name:"tom" email:"12312312"
}
