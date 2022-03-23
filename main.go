package main

import (
	"context"
	"encoding/json"
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
	Datetime    string `json:"datetime"`
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
	fmt.Println("mongo done")
	//db := client.Database("blog")
	//blogsCollection := db.Collection("blogs")

	// blog := blog{"a", "b", "c", "d"}
	// CreateBook(blog, blogsCollection, context.TODO())

	r := mux.NewRouter()
	r.Path("/new").Methods(http.MethodPost).HandlerFunc(postBlog)
	// handler := cors.Default().Handler(r)
	http.ListenAndServe(":4000", r)
}

func postBlog(w http.ResponseWriter, r *http.Request) {
	b := blog{}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoidng response object", http.StatusBadRequest)
		return
	}
	fmt.Println(b)
}

func CreateBook(b blog, blogsCollection *mongo.Collection, ctx context.Context) (string, error) {
	result, err := blogsCollection.InsertOne(ctx, b)
	if err != nil {
		return "0", err
	}
	return fmt.Sprintf("%v", result.InsertedID), err
}
