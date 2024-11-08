package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"Artist"`
	Price float64 `json:"price"`
}

// slice de albums para almacenar datos iniciales de álbumes.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context){
	//convirte la lista a json
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context){
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id{
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	//convirte la lista a json
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album no encontrado mij@"})
}

func postAlbums(c *gin.Context){
	var newAlbum album
	//Recibe una estrucutra y retorna un error
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	fmt.Println("Tamos activos")

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome zunguita a la API pa",
		})
	})

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run(":8080")
}