package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/inciner8r/go_blog/configs"
	"github.com/inciner8r/go_blog/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var blogsCollection *mongo.Collection = configs.GetCollection(configs.DB, "blogs")

func CreateBlog(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var blog models.Blog
	defer cancel()

	//validate request
	if err := c.BindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	newBlog := models.Blog{
		Id:          primitive.NewObjectID(),
		Title:       blog.Title,
		Datetime:    blog.Datetime,
		Description: blog.Description,
		Content:     blog.Content,
	}
	result, err := blogsCollection.InsertOne(ctx, newBlog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": result})

}

func GetABlog(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var blog models.Blog
	blogId := c.Param("blogId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(blogId)

	err := blogsCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": blog})
}

func EditABlog(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var blog models.Blog
	blogId := c.Param("blogId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(blogId)

	//validate request
	if err := c.BindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	update := bson.M{"title": blog.Title, "datetime": blog.Datetime, "description": blog.Description, "content": blog.Content}
	result, err := blogsCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		result
	}

	var UpdatedBlog models.Blog
	if result.MatchedCount == 1 {
		err := blogsCollection.FindOne(ctx, bson.M{"id": blogId}).Decode(&UpdatedBlog)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}

func DeleteABlog(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	blogId := c.Param("blogId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(blogId)

	result, err := blogsCollection.DeleteOne(ctx, bson.M{"id": objId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err})
		return
	}
	if result.DeletedCount < 1 {
		c.JSON(http.StatusNotFound, gin.H{"data": "user not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "User successfully deleted!"})
}
func GetAllBlogs(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var blogs []models.Blog
	defer cancel()

	results, err := blogsCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}

}
