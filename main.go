package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type blog struct {
	Title       string `json:"title"`
	Datetime    string `json:"datetime"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

func connectdb() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongoURI := os.Getenv("MONGOURI")

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("mongo done")

	}
	return client
}

func getCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("blog").Collection(collectionName)
	return collection
}
func main() {
	// := client.Database("blog")
	//blogsCollection := db.Collection("blogs")
	var db *mongo.Client = connectdb()
	getCollection(db, "blogs")

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
