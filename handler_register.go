package main

import (
	"context"
	"fmt"
	"time"

	"gator-go/internal/database"

	"github.com/google/uuid"
)

func registerHandler(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Register command require exactly 1 argument")
	}

	createUsersParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}

	user, err := s.db.CreateUser(context.Background(), createUsersParams)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	fmt.Printf("Created user: %v\n", user)

	err = s.Config.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("Logged in as %s\n", user.Name)
	return nil

}
