/*
 * @Author: GG
 * @Date: 2022-10-10 14:56:35
 * @LastEditTime: 2022-10-10 15:08:01
 * @LastEditors: GG
 * @Description:
 * @FilePath: \grpc\client\gin-client.go
 *
 */
package main

import (
	"fmt"
	"grpc/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	address     = "127.0.0.1:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	r := gin.Default()
	r.GET("/rest/n/:name", func(ctx *gin.Context) {
		n := ctx.Param("name")

		req := &pb.HelloRequest{Name: n}
		res, err := c.SayHello(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprintf(res.Message),
		})
	})
	if err := r.Run(":8052"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
