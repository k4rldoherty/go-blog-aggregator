package middleware

import (
	"context"

	"github.com/k4rldoherty/rss-blog-aggregator/internal/command"
	"github.com/k4rldoherty/rss-blog-aggregator/internal/database"
	"github.com/k4rldoherty/rss-blog-aggregator/internal/state"
)

func MiddlewareLoggedIn(handler func(s *state.State, cmd command.Command, user database.User) error) func(*state.State, command.Command) error {
	return func(s *state.State, c command.Command) error {
		user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, c, user)
	}
}
