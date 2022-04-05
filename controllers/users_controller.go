package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/inciner8r/go_blog/configs"
	"github.com/inciner8r/go_blog/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const secretKey = "secret"

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
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": result.InsertedID})
}

func Login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var requestUser, mongoUser models.User

	defer cancel()

	if err := c.BindJSON(&requestUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}
	if err := usersCollection.FindOne(ctx, bson.M{"email": requestUser.Email}).Decode(&mongoUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(mongoUser.Password), []byte(requestUser.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    mongoUser.ID.Hex(),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	fmt.Println(claims)
	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}

	c.SetCookie("jwt", token, int((1 * time.Hour).Seconds()), "/", "localhost", false, true)

	c.JSON(http.StatusCreated, gin.H{"data": token})
}

func User(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	cookie, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"data": err.Error()})
		return
	}
	claims := token.Claims.(*jwt.StandardClaims)
	objID, err := primitive.ObjectIDFromHex(claims.Issuer)
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{"data": err.Error()})
		return
	}

	//pass objectid to search using _id
	usersCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	c.JSON(http.StatusAccepted, gin.H{"data": user})
}

func Logout(c *gin.Context) {
	c.SetCookie("jwt", "", int(-(1 * time.Hour).Seconds()), "/", "localhost", false, true)

	c.JSON(http.StatusAccepted, gin.H{"message": "logged out"})
}
