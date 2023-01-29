/*
 * @Author: GG
 * @Date: 2023-01-29 09:27:11
 * @LastEditTime: 2023-01-29 10:14:33
 * @LastEditors: GG
 * @Description: redis 连接
 * @FilePath: \redis\main.go
 *
 */
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func ExampleClient() {
	rdb := redis.NewClient(
		&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "root",
			DB:       0,
		},
	)

	err := rdb.Set(ctx, "name", "redis-demo", 60*time.Second).Err()
	if err != nil {
		panic(err)
	}
	// err = rdb.SetEX(ctx, "name2", "redis-demo2", 60*time.Second).Err()
	// if err != nil {
	// 	panic(err)
	// }

	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("val:", val)

	// val, err = rdb.GetEx(ctx, "name2", 60*time.Second).Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("val: %v\n", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exits")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2:", val2)
	}
}

func main() {
	ExampleClient()
}
