package api

import (
	"cricketCrawler/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// server check
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// retrieves a paginated list of videos from the database.
func GetPaginatedVideosHandler(db *mongo.Client, c *gin.Context) {

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pageSize parameter"})
		return
	}

	// Retrieve videos from the database
	videos, totalVideos, err := utils.GetPaginatedVideos(db, page, pageSize)
	if err != nil {
		log.Printf("Error retrieving videos: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve videos"})
		return
	}

	totalPages := (totalVideos + int64(pageSize) - 1) / int64(pageSize)

	if page > int(totalPages) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "OOPs! No more videos to show"})
		return
	}

	// Build the response JSON
	response := gin.H{
		"pagination": gin.H{
			"currentPage": page,
			"pageSize":    pageSize,
			"totalPages":  totalPages,
			"totalVideos": totalVideos,
		},
		"videos": videos,
	}

	c.JSON(http.StatusOK, response)
}

func GetVideosHandler(db *mongo.Client, c *gin.Context) {

	searchQuery := c.Query("query")

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pageSize parameter"})
		return
	}

    videos, totalVideos, err := utils.GetVideosByTitle(db, searchQuery, page, pageSize)
	if err != nil {
		log.Printf("Error retrieving videos: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve videos"})
		return
	}

	totalPages := (totalVideos + int64(pageSize) - 1) / int64(pageSize)

	if page > int(totalPages) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "OOPs! No more videos to show"})
		return
	}

	// Build the response JSON
	response := gin.H{
		"pagination": gin.H{
			"currentPage": page,
			"pageSize":    pageSize,
			"totalPages":  totalPages,
			"totalVideos": totalVideos,
		},
		"videos": videos,
	}

	c.JSON(http.StatusOK, response)
}
