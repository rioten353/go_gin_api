package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetAlbumbs(context *gin.Context) {
	context.IndentedJSON(200, Albumbs)
}

func PostAlbumb(context *gin.Context) {
	var albumb Albumb
	if err := context.BindJSON(&albumb); err != nil {
		context.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}
	Albumbs = append(Albumbs, albumb)
	context.IndentedJSON(201, albumb)
}

func GetAlbumbById(context *gin.Context) {
	id := context.Param("id")
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		context.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}
	for _, albumb := range Albumbs {
		if albumb.Id == int(ID) {
			context.IndentedJSON(200, albumb)
			return
		}
	}
	context.IndentedJSON(404, gin.H{"error": "albumb not found"})
}
