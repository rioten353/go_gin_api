package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", GetAlbumbs)
	router.GET("/albumbs/:id", GetAlbumbById)
	router.POST("/albumbs", PostAlbumb)
	router.PUT("/albumbs/:id", UpdateAlbumbById)
	router.DELETE("/albumbs/:id", DeleAlbumbById)

	fmt.Println("Server is running on port 8080")
	router.Run("localhost:8080")
}
