package gatiserv

import (
	sgc7game "github.com/zhs007/slotsgamecore7/game"
)

// BasicService - basic service
type BasicService struct {
	Game sgc7game.IGame
}

// NewBasicService - new a BasicService
func NewBasicService(game sgc7game.IGame) *BasicService {
	return &BasicService{
		Game: game,
	}
}

// Config - get configuration
func (sv *BasicService) Config() *sgc7game.Config {
	return sv.Game.GetConfig()
}
