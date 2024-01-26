package utils

import (
	"context"
	"cricketCrawler/model"
	"errors"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() *mongo.Client {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("Missing MONGO_URI environment variable")
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Error connecting to MongoDB", err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB server", err)
	}
	return client
}

func GetDatabase(client *mongo.Client) *mongo.Database {
	databaseName := os.Getenv("DATABASE_NAME")
	if databaseName == "" {
		log.Fatal("Missing DATABASE_NAME environment variable")
	}
	return client.Database(databaseName)
}

func GetCollection(db *mongo.Database) *mongo.Collection {
	collectionName := os.Getenv("COLLECTION_NAME")
	if collectionName == "" {
		log.Fatal("Missing COLLECTION_NAME environment variable")
	}
	return db.Collection(collectionName)
}

func StoreVideo(db *mongo.Client, video model.Video) error {
	collection := GetCollection(GetDatabase(db))

	filter := bson.D{{Key: "id", Value: video.ID}}

	// Check for existing video before upserting
	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("video already exists in the database")
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "title", Value: video.Title},
			{Key: "description", Value: video.Description},
			{Key: "published_at", Value: video.PublishedAt},
			{Key: "thumbnails", Value: video.Thumbnails},
		}},
	}

	_, err = collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
	return err
}

func GetPaginatedVideos(db *mongo.Client, page, pageSize int) ([]model.Video, int64, error) {

	// Calculate the number of videos to skip based on page and pageSize
	skip := (page - 1) * pageSize

	findOptions := options.Find().
		SetSort(bson.D{{Key: "published_at", Value: -1}}).
		SetSkip(int64(skip)).
		SetLimit(int64(pageSize))

	// empty filter
	filter := bson.D{{}}

	collection := GetCollection(GetDatabase(db))

	cursor, err := collection.Find(context.Background(), filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(context.Background())

	var videos []model.Video
	if err = cursor.All(context.Background(), &videos); err != nil {
		return nil, 0, err
	}

	totalVideos, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {	
		return nil, 0, err
	}

	return videos, totalVideos, nil
}

func GetVideosByTitle(db *mongo.Client, query string, page, pageSize int) ([]model.Video, int64, error) {
    // Calculate the number of videos to skip based on page and pageSize
	skip := (page - 1) * pageSize

	findOptions := options.Find().
		SetSort(bson.D{{Key: "published_at", Value: -1}}).
		SetSkip(int64(skip)).
		SetLimit(int64(pageSize))

	// apply filter
    filter := bson.M{"title": bson.M{"$regex": query, "$options": "i"}} // Case-insensitive search

	collection := GetCollection(GetDatabase(db))

	cursor, err := collection.Find(context.Background(), filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(context.Background())

	var videos []model.Video
	if err = cursor.All(context.Background(), &videos); err != nil {
		return nil, 0, err
	}

    totalVideos, err := collection.CountDocuments(context.Background(), filter)
    if err != nil {
        return nil, 0, err
    }

	return videos, totalVideos, nil
}