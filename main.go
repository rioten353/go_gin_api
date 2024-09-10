package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/albumbs", GetAlbumbs)
	router.POST("/albumbs", PostAlbumb)
	router.GET("/albumbs/:id", GetAlbumbById)

	fmt.Println("Server is running on port 8080")
	router.Run("localhost:8080")
}
