package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

const postgres_url = "postgres://postgres:123456@localhost:5432/albums"

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	// os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), postgres_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var albums []Album
	rows, err := conn.Query(context.Background(), "select * from albums")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.Id, &album.Title, &album.Artist, &album.Price); err != nil {
			return
		}
		albums = append(albums, album)
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {

	conn, err := pgx.Connect(context.Background(), postgres_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	id := c.Param("id")

	var album Album
	err = conn.QueryRow(context.Background(), "select * from albums where id=$1", id).Scan(&album.Id, &album.Artist, &album.Title, &album.Price)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func postAlbums(c *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	conn, err := pgx.Connect(context.Background(), postgres_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	err = conn.QueryRow(context.Background(), "INSERT INTO albums (title, artist, price) VALUES ($1, $2, $3) RETURNING Id", newAlbum.Title, newAlbum.Artist, newAlbum.Price).Scan(&newAlbum.Id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error saving album"})
		return
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

var albums_static = []Album{
	{Id: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{Id: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{Id: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbumById_static(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for an album whose ID value matches the parameter.
	for _, a := range albums_static {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbum_static(c *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums_static = append(albums_static, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
