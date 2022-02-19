package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cqroot/boite/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/ui")
	})
	r.Static("/ui", "./web/dist")
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/bytecalc", controllers.CalcByte)
		apiGroup.GET("/foo", func(c *gin.Context) { c.JSON(200, gin.H{"foo": true}) })
		apiGroup.GET("/bar", func(c *gin.Context) { c.JSON(200, gin.H{"bar": true}) })
	}

	r.NoRoute(func(c *gin.Context) {
		c.File("./web/dist/index.html")
	})

	r.Run(":8088")
}
