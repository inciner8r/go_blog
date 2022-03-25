package main

import (
	"github.com/gin-gonic/gin"
	"github.com/inciner8r/go_blog/configs"
	"github.com/inciner8r/go_blog/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	var db *mongo.Client = configs.ConnectDB()
	configs.GetCollection(db, "blogs")

	r := gin.Default()
	routes.Routes(r)
	r.Run("localhost:4000")
}

func postBlog(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}

// func CreateBlog(b models.Blog, blogsCollection *mongo.Collection, ctx context.Context) (string, error) {
// 	result, err := blogsCollection.InsertOne(ctx, b)
// 	if err != nil {
// 		return "0", err
// 	}
// 	return fmt.Sprintf("%v", result.InsertedID), err
// }
