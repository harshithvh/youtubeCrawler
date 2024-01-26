package model

type Video struct {
	ID          string        `bson:"_id"`
	Title       string        `bson:"title"`
	Description string        `bson:"description"`
	PublishedAt string        `bson:"published_at"`
	Thumbnails   Thumbnails `bson:"thumbnails"`
}

type Thumbnails struct {
	Default string `bson:"default"`
	Medium  string `bson:"medium"`
	High    string `bson:"high"`
}

