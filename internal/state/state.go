package state

import (
	"github.com/k4rldoherty/rss-blog-aggregator/internal/config"
	"github.com/k4rldoherty/rss-blog-aggregator/internal/database"
)

type State struct {
	Cfg *config.Config
	Db  *database.Queries
}

func NewState(cfg *config.Config, db *database.Queries) *State {
	return &State{
		cfg,
		db,
	}
}
