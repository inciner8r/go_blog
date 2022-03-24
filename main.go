package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/inciner8r/go_blog/configs"
	"go.mongodb.org/mongo-driver/mongo"
)

type blog struct {
	Title       string `json:"title"`
	Datetime    string `json:"datetime"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

func main() {
	// := client.Database("blog")
	//blogsCollection := db.Collection("blogs")
	var db *mongo.Client = configs.ConnectDB()
	configs.GetCollection(db, "blogs")

	r := gin.Default()
	r.GET("/new", postBlog)
	r.Run("localhost:4000")
}

// func postBlog(w http.ResponseWriter, r *http.Request) {
// 	b := blog{}
// 	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
// 		fmt.Println(err)
// 		http.Error(w, "Error decoidng response object", http.StatusBadRequest)
// 		os.Exit(1)
// 	}
// }
func postBlog(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}
func CreateBlog(b blog, blogsCollection *mongo.Collection, ctx context.Context) (string, error) {
	result, err := blogsCollection.InsertOne(ctx, b)
	if err != nil {
		return "0", err
	}
	return fmt.Sprintf("%v", result.InsertedID), err
}
