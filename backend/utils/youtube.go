package utils

import (
	"context"
	"cricketCrawler/model"
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var expiredAPIKeys []string
var currKey string

func InitYouTubeClient(apiKey string) *youtube.Service {
	currKey = apiKey
	client, err := youtube.NewService(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating YouTube service: %v", err)
	}
	return client
}

func FetchAPIKeys() []string {
	apiKeyStr := os.Getenv("API_KEYS")
	if apiKeyStr == "" {
		log.Fatalf("Missing API_KEYS environment variable")
	}
	apiKeys := strings.Split(apiKeyStr, ",")
	// Filter out expired keys
    filteredAPIKeys := make([]string, 0)
    for _, key := range apiKeys {
        if !contains(expiredAPIKeys, key) {
            filteredAPIKeys = append(filteredAPIKeys, key)
        }
    }

    return filteredAPIKeys
}

func FetchVideos(youtubeClient *youtube.Service, db *mongo.Client, query string) error {

	if err := ExecuteFetch(youtubeClient, db, query); err != nil {
		log.Printf("Failed to download videos: %v", err)
		return err
	} else {
		log.Printf("Successfully downloaded videos")
	}

	return nil
}

func ExecuteFetch(youtubeClient *youtube.Service, db *mongo.Client, query string) error {
	call := youtubeClient.Search.List([]string{"snippet"}).
		Q(query).
		MaxResults(20)

	response, err := call.Do()
	if err != nil {
		// Check for quota exhaustion and switch API key
		if apiErr, ok := err.(*googleapi.Error); ok && apiErr.Code == 403 {
			log.Printf("API key %s has been exhausted", currKey)

			log.Printf("Switching API key...")

			expiredAPIKeys = append(expiredAPIKeys, currKey)

			newAPIKey, err := UpdateAPIKey()
			if err != nil {
				return err // All keys are exhausted
			}

			currKey = newAPIKey

			youtubeClient = InitYouTubeClient(newAPIKey)

			return ExecuteFetch(youtubeClient, db, query)
		}
		return err
	}

	var videos []model.Video
	for _, item := range response.Items {
		publishedAt, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
		if err != nil {
			log.Printf("Error parsing published date: %v", err)
			continue
		}

		video := model.Video{
			VideoID:      item.Id.VideoId,
			Title:        item.Snippet.Title,
			Description:  item.Snippet.Description,
			ChannelTitle: item.Snippet.ChannelTitle,
			PublishedAt:  publishedAt,
			ThumbnailURL: item.Snippet.Thumbnails.High.Url,
		}
		videos = append(videos, video)
	}

	if err := StoreVideos(db, videos, query); err != nil {
		return err
	}

	return nil
}

func StoreVideos(db *mongo.Client, videos []model.Video, query string) error {
	duplicateCount := 0
	for _, video := range videos {
		if err := StoreVideo(db, video, query); err != nil {
			log.Printf("Error storing video: %v", err)
			if err.Error() == "video already exists in the database" {
				duplicateCount++
			}
		}
	}
	log.Printf("Duplicate videos found in this API call: %d", duplicateCount)
	return nil
}

func UpdateAPIKey() (string, error) {
    availableAPIKeys := make([]string, 0)

    allAPIKeys := FetchAPIKeys()

    // Filter
    for _, key := range allAPIKeys {
        if !contains(expiredAPIKeys, key) {
            availableAPIKeys = append(availableAPIKeys, key)
        }
    }

    if len(availableAPIKeys) == 0 {
        return "", errors.New("all API keys are quota exhausted")
    }

    newAPIKey := availableAPIKeys[0]

    return newAPIKey, nil
}

func contains(list []string, key string) bool {
    for _, item := range list {
        if item == key {
            return true
        }
    }
    return false
}
