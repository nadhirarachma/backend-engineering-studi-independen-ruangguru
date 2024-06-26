package main

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var Posts = []Post{
	{ID: 1, Title: "Judul Postingan Pertama", Content: "Ini adalah postingan pertama di blog ini.", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 2, Title: "Judul Postingan Kedua", Content: "Ini adalah postingan kedua di blog ini.", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/posts", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"posts": Posts,
		})
	})

	r.GET("/posts/:id", func(c *gin.Context) {
		id := c.Param("id")
		postId, err := strconv.Atoi(id)
		exist := false

		if err != nil {
			c.JSON(400, gin.H{
				"error": "ID harus berupa angka",
			})
		} else {
			for _, post := range Posts {
				if post.ID == postId {
					exist = true
					c.JSON(200, gin.H{"post": post})
				}
			} 

			if !exist {
				c.JSON(404, gin.H{
					"error": "Postingan tidak ditemukan",
				})
			}
		}
	})

	r.POST("/posts", func(c *gin.Context) {
		var post Post

        err := c.ShouldBindJSON(&post)
        if err != nil {
            c.JSON(400, gin.H{"error": "Invalid request body"})
        } else {
			post.ID = len(Posts) + 1
			Posts = append(Posts, post)

			c.JSON(201, gin.H{
				"message" : "Postingan berhasil ditambahkan",
				"post" : post,
			})
		}
	})

	return r
}

func main() {
	r := SetupRouter()

	r.Run(":8080")
}
