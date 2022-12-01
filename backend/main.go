package main

import (
	"golang-simple-bank/database"
	"golang-simple-bank/middleware"
	"golang-simple-bank/routes"
	"os"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
)

var entriesCollection *mongo.Collection = database.OpenCollection(database.Client, "entries")
var transfersCollection *mongo.Collection = database.OpenCollection(database.Client, "transfers")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middleware.Authentication())

	routes.AccountRouter(router)
	routes.EntriesRouter(router)
	routes.TransfersRouter(router)

	router.Run(":", port)
}
