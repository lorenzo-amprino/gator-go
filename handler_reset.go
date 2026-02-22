package main

import (
	"context"
	"fmt"
)

func resetHandler(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("All users have been reset")
	return nil
}
