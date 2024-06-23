package main

import (
	"time"

	"github.com/Eche5/rssagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

type Feeds struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	UserID    uuid.UUID `json:"user_id"`
	URL       string    `json:"url"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

func databaseFeedsToFeeds(dbFeed database.Feed) Feeds {
	return Feeds{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		UserID:    dbFeed.UserID,
		URL:       dbFeed.Url,
	}
}

func databaseFeedsToFeedsArray(dbFeeds []database.Feed) []Feeds {
	feeds := []Feeds{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedsToFeeds(dbFeed))
	}
	return feeds
}
