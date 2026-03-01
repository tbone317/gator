package main

import (
	"context"

	"github.com/tbone317/gator/internal/database"
)

// middlewareLoggedIn wraps a handler that requires a logged-in user.
func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}
