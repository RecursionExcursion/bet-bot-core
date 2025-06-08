package bbcore

import (
	"time"

	"github.com/RecursionExcursion/bet-bot-core/internal"
)

func CollectData() (internal.FirstShotData, error) {

	teams, err := internal.CollectTeamsAndRosters()
	if err != nil {
		return internal.FirstShotData{}, err
	}
	games, err := internal.CollectGames()
	if err != nil {
		return internal.FirstShotData{}, err
	}

	data := internal.FirstShotData{
		Created: time.Now().UnixMilli(),
		Teams:   teams,
		Games:   games,
	}

	return data, nil
}
