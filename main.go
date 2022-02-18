package main

import (
	"github.com/gin-gonic/gin"

	"github.com/cqroot/boite/controllers"
)

func main() {
	r := gin.Default()

	r.Static("/ui", "./web/dist")
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/bytecalc", controllers.CalcByte)
		apiGroup.GET("/foo", func(c *gin.Context) { c.JSON(200, gin.H{"foo": true}) })
		apiGroup.GET("/bar", func(c *gin.Context) { c.JSON(200, gin.H{"bar": true}) })
	}

	r.Run(":8088")
}
