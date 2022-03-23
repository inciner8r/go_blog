package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type blog struct {
	Title       string `json:"title"`
	DateTime    string `json:"date"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

func main() {
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
	}
	db := client.Database("blog")
	blogsCollection := db.Collection("blogs")

	blog := blog{"a", "b", "c", "d"}
	CreateBook(blog, blogsCollection, context.TODO())

	r := mux.NewRouter()
	// handler := cors.Default().Handler(r)
	http.ListenAndServe(":4000", r)
}

func CreateBook(b blog, blogsCollection *mongo.Collection, ctx context.Context) (string, error) {
	result, err := blogsCollection.InsertOne(ctx, b)
	if err != nil {
		return "0", err
	}
	return fmt.Sprintf("%v", result.InsertedID), err
}
