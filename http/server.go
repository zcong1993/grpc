package main

import (
	"context"
	"github.com/apex/log"
	"github.com/gin-gonic/gin"
	"github.com/zcong1993/grpc/client"
	pb "github.com/zcong1993/grpc/echo"
	"net/http"
	"strconv"
	"io"
)

func main() {
	clt, conn, err := client.CreateEchoClient("")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()
	router := gin.Default()
	router.GET("/echo/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		a, err := strconv.ParseInt(age, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		resp, err := clt.Echo(context.Background(), &pb.EchoRequest{Name: name, Age: a})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"Name": resp.Name,
			"Age":  resp.Age,
		})
	})

	router.GET("/again/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		a, err := strconv.ParseInt(age, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		resp, err := clt.Echo(context.Background(), &pb.EchoRequest{Name: name, Age: a})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"Name": resp.Name,
			"Age":  resp.Age,
		})
	})

	router.GET("/stream", func(c *gin.Context) {
		stream, err := clt.EchoStream(context.Background(), &pb.EchoRequest{"l", 18})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
		var resp []pb.EchoResponse
		for {
			r, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
			}
			resp = append(resp, *r)
		}
		c.JSON(http.StatusOK, gin.H{
			"resps": resp,
		})
	})

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": []string{"xsxsxs", "cdcdcdc"},
		})
	})

	router.Run()
}
