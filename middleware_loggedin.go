package main

import (
	"context"
	"errors"

	"github.com/simonlewi/gator/internal/database"
)

var ErrNotLoggedIn = errors.New("you must be logged in to perform this action")

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		// If no user is found, return an error indicating that the user needs to be logged in
		if s.cfg.CurrentUsername == "" {
			return ErrNotLoggedIn
		}

		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
		if err != nil {
			return err
		}

		// Execute the wrapped handler with the authenticated user
		return handler(s, cmd, user)
	}
}
