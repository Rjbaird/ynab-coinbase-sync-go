package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"bairrya.com/go/ynab-coinbase/jobs"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
)


func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")

	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().At("09:00").Do(func(){ jobs.SyncYnabCoinbase() })
	s.StartAsync()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		"message": "try out /ping",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		})
	})


	r.Run(fmt.Sprintf(":%s", port))
}