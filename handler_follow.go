package main

import (
	"context"
	"fmt"
	"gator-go/internal/database"
	"time"

	"github.com/google/uuid"
)

func followHandler(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Follow command require exactly 1 argument")
	}

	currentUser, err := s.db.GetUser(context.Background(), s.Config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Failed to get current user: %w", err)
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Failed to get feed by URL: %w", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("Failed to create feed follow: %w", err)
	}

	fmt.Printf("User %s is now following feed with URL %s\n", currentUser.Name, feed.Url)

	return nil

}
