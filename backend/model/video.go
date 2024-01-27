package model

import "time"

type Video struct {
	VideoID       string        `gorm:"primaryKey" bson:"_id"`
	Title         string        `bson:"title"`
	Description   string        `bson:"description"`
	ChannelTitle  string        `bson:"channel_title"`
	PublishedAt   time.Time     `bson:"published_at"`
	ThumbnailURL  string        `bson:"thumbnail_url"`
}