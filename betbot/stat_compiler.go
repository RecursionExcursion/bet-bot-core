package internal

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/RecursionExcursion/bet-bot-core/internal"
)

type PackagedPlayer struct {
	Id                  string `json:"id"`
	Name                string `json:"name"`
	Team                string `json:"team"`
	FirstToScore        uint8  `json:"firstToScore"`
	FirstShotAttempts   uint8  `json:"firstShotAttempts"`
	ScoreOnFirstAttempt uint8  `json:"scoreOnFirstAttempt"`
	HeadShotHref        string `json:"headShotHref"`
	TeamLogoHref        string `json:"teamLogoHref"`
}

type StatCalculator struct {
	fsd internal.FirstShotData
}

func NewStatCalculator(fsd internal.FirstShotData) *StatCalculator {

	filteredGames := []internal.Game{}

	for _, g := range fsd.Games {
		if g.Season.Slug == "preseason" {
			continue
		}
		filteredGames = append(filteredGames, g)
	}

	fsd.Games = filteredGames

	FindGameInFsd(fsd, strconv.Itoa(401705613))

	return &StatCalculator{
		fsd,
	}
}

func (sc *StatCalculator) CalculateAndPackage() ([]PackagedPlayer, error) {

	log.Println("Calc stats for %v games", len(sc.fsd.Games), 7)

	err := sc.calculateFirstScore()
	if err != nil {
		return []PackagedPlayer{}, err
	}

	err = sc.calculateFirstShotAttempt()
	if err != nil {
		return []PackagedPlayer{}, err
	}

	allData := sc.packageData()
	var filteredData = []PackagedPlayer{}

	for _, pp := range allData {
		if pp.ScoreOnFirstAttempt == 0 &&
			pp.FirstShotAttempts == 0 &&
			pp.FirstToScore == 0 {
			continue
		}
		filteredData = append(filteredData, pp)
	}

	sort.Slice(filteredData, func(a, b int) bool {
		return filteredData[a].FirstToScore > filteredData[b].FirstToScore
	})

	return filteredData, nil
}

func (sc *StatCalculator) calculateFirstScore() error {
	for _, gm := range sc.fsd.Games {

		fs := gm.TrackedEvents.FirstScore

		if fs.Id == "" {
			err := fmt.Errorf("first shot data for game:%v not found", gm.Id)
			log.Println("", err, "calculateFirstScore")
			continue
		}

		playerId := fs.Participants[0].Athlete.Id

		player, err := sc.findPlayerById(playerId)
		if err != nil {
			log.Println(fmt.Sprintf("Player %v for play %v not found", playerId, fs.Text), err)
			continue
		}
		player.BetStats.FirstPointsMade++
	}
	return nil
}

func (sc *StatCalculator) calculateFirstShotAttempt() error {
	for _, gm := range sc.fsd.Games {

		fsa := gm.TrackedEvents.FirstShotAttempt

		if fsa.Id == "" {
			err := fmt.Errorf("first shot attempt data for game:%v not found", gm.Id)
			log.Println("", err)
			continue
		}

		playerId := fsa.Participants[0].Athlete.Id
		player, err := sc.findPlayerById(playerId)
		if err != nil {
			log.Println(fmt.Sprintf("Player %v for play %v not found", playerId, fsa.Text), err)
			continue
		}
		player.BetStats.FirstShotAttempts++
		if fsa.ScoringPlay {
			player.BetStats.ScoreOnFirstAttempt++
		}
	}
	return nil
}

func (sc *StatCalculator) packageData() []PackagedPlayer {
	var packagedPlayers = []PackagedPlayer{}

	for _, t := range sc.fsd.Teams {
		for _, p := range t.Roster {

			pp := PackagedPlayer{
				Id:                  p.Id,
				Name:                p.FullName,
				Team:                t.Name,
				FirstToScore:        p.BetStats.FirstPointsMade,
				FirstShotAttempts:   p.BetStats.FirstShotAttempts,
				ScoreOnFirstAttempt: p.BetStats.ScoreOnFirstAttempt,
				HeadShotHref:        p.Headshot.Href,
				TeamLogoHref:        t.Logos[0].Href,
			}

			packagedPlayers = append(packagedPlayers, pp)
		}
	}
	return packagedPlayers
}

func (sc *StatCalculator) findPlayerById(id string) (*internal.Player, error) {
	var player internal.Player

	for _, t := range sc.fsd.Teams {
		for i := range t.Roster {
			if t.Roster[i].Id == id {
				return &t.Roster[i], nil
			}
		}
	}
	return &player, fmt.Errorf("player %v not found", id)
}

func FindGameInFsd(fsd internal.FirstShotData, id string) (internal.Game, error) {
	for _, g := range fsd.Games {
		if g.Id == id {
			return g, nil
		}
	}
	var g internal.Game
	return g, errors.New("game not found")
}

func FindGame(games []internal.Game, id string) (internal.Game, error) {
	for _, g := range games {
		if g.Id == id {
			return g, nil
		}
	}
	var g internal.Game
	return g, errors.New("game not found")
}
