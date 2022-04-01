package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/inciner8r/go_blog/configs"
	"github.com/inciner8r/go_blog/models"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var usersCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func Register(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	//validate request
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(password),
	}
	result, err := usersCollection.InsertOne(ctx, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"data": result.InsertedID})
}
