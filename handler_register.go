package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/simonlewi/gator/internal/database"
)

func registerHandler(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	ctx := context.Background()
	username := cmd.Args[0]

	_, err := s.db.GetUser(ctx, username)
	if err == nil {
		fmt.Printf("User with name %s already exists\n", username)
		os.Exit(1)
	} else if err != sql.ErrNoRows {
		// This is some other error, not just "user not found"
		return fmt.Errorf("error checking if user exists: %v", err)
	}

	createUser, err := s.db.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
	})
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	fmt.Printf("User %s created successfully with ID %s\n", createUser.Name, createUser.ID)

	err = s.cfg.SetUser(createUser.Name)
	if err != nil {
		return fmt.Errorf("error setting user: %v", err)
	}
	fmt.Println("User switched successfully!")

	return nil
}
