package main

import (
	"context"
	"fmt"
)

func followingHandler(s *state, cmd command) error {

	currentUser := s.Config.CurrentUserName

	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return err
	}
	feeds, err := s.db.GetFeedsByUserId(context.Background(), user.ID)
	if err != nil {
		return err
	}
	fmt.Printf("User %s is following the following feeds:\n", currentUser)
	for _, feed := range feeds {
		println(feed.FeedName)
	}
	return nil
}
