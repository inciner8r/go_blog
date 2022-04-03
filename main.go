package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/inciner8r/go_blog/configs"
	"github.com/inciner8r/go_blog/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	var db *mongo.Client = configs.ConnectDB()
	configs.GetCollection(db, "blogs")

	r := gin.Default()
	r.Use(cors.Default())
	routes.Routes(r)
	r.Run("localhost:4000")
}

func postBlog(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}
