package main

import (
	"cricketCrawler/api"
	"cricketCrawler/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/api/youtube/v3"
)

var stopFetching bool = true

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Initialize MongoDB and YouTube client
    db := utils.InitDB()
    apiKeys := utils.FetchAPIKeys()
    youtube := utils.InitYouTubeClient(apiKeys[0])

    router := gin.Default()

    // Enable CORS for frontend
    router.Use(cors.New(cors.Config{
        AllowOrigins: []string{"http://localhost:3000"},
        AllowMethods: []string{"GET"},
    }))

    // Routes
    router.GET("/videos", func(c *gin.Context) {
        api.GetPaginatedVideos(db, c)
    })

    router.GET("/search", func(c *gin.Context) {
        api.GetVideos(db, c)
    })

    router.GET("/ping", func(c *gin.Context) {
        api.PingServer(c)
    })

    router.GET("/start-fetching", func(c *gin.Context) {
        if !stopFetching {
            c.String(http.StatusBadRequest, "Fetching is already in progress")
            return
        }
	    category := c.Query("category")
        stopFetching = false
        go fetchVideosPeriodically(category, youtube, db, c)
        c.String(http.StatusOK, "Fetching started")
    })

    router.GET("/stop-fetching", func(c *gin.Context) {
        if stopFetching {
            c.String(http.StatusBadRequest, "Fetching is not in progress")
            return
        }

        stopFetching = true
        c.String(http.StatusOK, "Fetching stopped")
    })

    router.Run(":8080")
}

func fetchVideosPeriodically(category string, youtube *youtube.Service, db *mongo.Client, c *gin.Context) {
    for !stopFetching {
        if err := utils.FetchVideos(youtube, db, category); err != nil {
            log.Printf("Error fetching videos: %v", err)
            c.String(http.StatusInternalServerError, "Error fetching videos")
            stopFetching = true
            return
        }
        time.Sleep(20 * time.Second)
    }
}
