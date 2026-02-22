package main

import (
	"context"
	"fmt"
)

func feedsHandler(s *state, _ command) error {

	feeds, err := s.db.GetAllFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("Feed: %s, URL: %s, User: %s\n", feed.Name, feed.Url, user.Name)
	}

	return nil
}
