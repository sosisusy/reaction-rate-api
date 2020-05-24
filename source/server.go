package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "test")
	})

	r.Run(":80")
}
