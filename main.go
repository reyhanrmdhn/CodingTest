package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reyhanrmdhn/CodingTest/db"
)

var (
	ListenAddr = "localhost:8080"
	RedisAddr  = "localhost:6379"
)

func main() {
	database, err := db.NewDatabase(RedisAddr)
	if err != nil {
		log.Fatalf("Failed to connect to redis: %s", err.Error())
	}

	router := initRouter(database)
	router.Run(ListenAddr)
}

func initRouter(database *db.Database) *gin.Engine {
	r := gin.Default()

	r.POST("/user/save", func(c *gin.Context) {
		var userJson db.User
		if err := c.ShouldBindJSON(&userJson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := database.SaveUser(&userJson)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"error": 0, "message": ""})
	})

	return r
}
