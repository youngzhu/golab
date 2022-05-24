package main

// https://golang.google.cn/doc/tutorial/web-service-gin

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 加了这个所谓的标签，可以指定json串中的字段名
// 如果没有，就是结构里的名字，大写不符合json的风格
type album struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func addAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		log.Fatal(err)
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	idNum, _ := strconv.Atoi(id)
	for _, a := range albums {
		if a.ID == idNum {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", addAlbums)
	router.GET("/albums/:id", getAlbumByID)

	router.Run("localhost:8080")
}
