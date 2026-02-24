package main

import (
	"context"
	"fmt"
	"gator-go/internal/database"
	"time"

	"github.com/google/uuid"
)

func addFeedHandler(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("command addfeed requires 2 args as Name and Url")
	}
	name := cmd.Args[0]
	url := cmd.Args[1]

	currentUser, err := s.db.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    currentUser.ID,
	})
	if err != nil {
		return fmt.Errorf("failed to create feed: %w", err)
	}

	s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	})

	fmt.Printf("Feed created with following data\nID: %s\nName: %s\nURL: %s\n", feed.ID, feed.Name, feed.Url)

	return nil
}
