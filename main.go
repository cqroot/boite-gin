package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cqroot/boite/controllers"
)

//go:embed web/dist
var embedFS embed.FS

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/ui")
	})

	distFS, _ := fs.Sub(embedFS, "web/dist")
	r.StaticFS("/ui", http.FS(distFS))

	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/bytecalc", controllers.CalcByte)
		apiGroup.GET("/foo", func(c *gin.Context) { c.JSON(200, gin.H{"foo": true}) })
		apiGroup.GET("/bar", func(c *gin.Context) { c.JSON(200, gin.H{"bar": true}) })
	}

	r.NoRoute(func(c *gin.Context) {
		c.FileFromFS("/", http.FS(distFS))
	})

	r.Run(":8088")
}
