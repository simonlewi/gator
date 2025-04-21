package main

import (
	"context"
	"fmt"
	"os"
)

func HandlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("user set to: %s <name>", cmd.Name)
	}

	username := cmd.Args[0]

	ctx := context.Background()
	_, err := s.db.GetUser(ctx, username)
	if err != nil {
		fmt.Printf("User with name %s doesn't exist\n", username)
		os.Exit(1)
	}

	err = s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("error setting user: %v", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}
