package main

import (
	"cricketCrawler/api"
	"cricketCrawler/utils"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize MongoDB and YouTube client
	db := utils.InitDB()
	apiKeys := utils.GetAPIKeys()
	youtube := utils.InitYouTubeClient(apiKeys[0])

	// Initiate background fetching
	go func() {
		for {
			if err := utils.FetchAndStoreVideos(youtube, db, "cricket"); err != nil {
				log.Printf("Error fetching and storing videos: %v", err)
			}
			time.Sleep(20 * time.Second) 
		}
	}()

	router := gin.Default()

	// Enable CORS for frontend
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	// Routes
	router.GET("/videos", func(c *gin.Context) {
		api.GetPaginatedVideosHandler(db, c)
	})

	router.GET("/search", func(c *gin.Context) {
		api.GetVideosHandler(db, c)
	})

	router.GET("/ping", func(c *gin.Context) {
		api.PingHandler(c)
	})

	router.Run(":8080")
}
