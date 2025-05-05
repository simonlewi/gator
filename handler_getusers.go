package main

import (
	"context"
	"fmt"
)

func getUsersHandler(s *state, cmd command) error {
	ctx := context.Background()
	users, err := s.db.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("error getting users: %w", err)
	}

	if len(users) == 0 {
		fmt.Println("No users found")
		return nil
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUsername {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}
