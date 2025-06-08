package bbcore

import (
	"time"

	"github.com/RecursionExcursion/bet-bot-core/internal"
)

type FirstShotData struct {
	Created int64           `json:"created"`
	Teams   []internal.Team `json:"teams"`
	Games   []internal.Game `json:"games"`
}

func CollectData() (FirstShotData, error) {

	teams, err := internal.CollectTeamsAndRosters()
	if err != nil {
		return FirstShotData{}, err
	}
	games, err := internal.CollectGames()
	if err != nil {
		return FirstShotData{}, err
	}

	data := FirstShotData{
		Created: time.Now().UnixMilli(),
		Teams:   teams,
		Games:   games,
	}

	return data, nil
}
