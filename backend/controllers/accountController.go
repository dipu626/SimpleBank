package controllers

import (
	"context"
	"golang-simple-bank/database"
	"golang-simple-bank/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var accountCollection *mongo.Collection = database.OpenCollection(database.Client, "account")

func GetAccounts() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Second)
	}
}

func GetAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		account_number = c.Param("account_number")
		var account models.Account

		err := accountCollection.FindOne(ctx, bson.M{"account_number": account_number}).Decode(&account)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H("error" : "error to get account"))
		}
		c.JSON(http.StatusOK, accaccount)
	}
}

func SingUp() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func HashPassword(password string) string {
	return ""
}

func VerifyPassword(userPassword string, providePassword string) (bool, string) {
	return false, ""
}
