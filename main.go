package main

import (
	v1 "github.com/Seunghoon-Oh/cloud-ml-foo-manager/controller/v1"
	"github.com/Seunghoon-Oh/cloud-ml-foo-manager/data"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "ok")
	})

	r.GET("/foos", v1.GetFoos)
	r.POST("/foo", v1.CreateFoo)

	return r
}

func main() {
	data.SetupRedisClient()
	r := setupRouter()
	r.Run(":8082")
}
