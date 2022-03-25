package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/inciner8r/go_blog/configs"
	"github.com/inciner8r/go_blog/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var blogsCollection *mongo.Collection = configs.GetCollection(configs.DB, "blogs")

func CreateBlog(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var blog models.Blog
	defer cancel()

	if err := c.BindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err})
		return
	}

	newBlog := models.Blog{
		Title:       blog.Title,
		Datetime:    blog.Datetime,
		Description: blog.Description,
		Content:     blog.Content,
	}
	result, err := blogsCollection.InsertOne(ctx, newBlog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": result})

}
