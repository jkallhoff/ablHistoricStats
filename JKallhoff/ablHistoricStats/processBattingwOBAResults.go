package main

import (
	"fmt"
	"sort"
)

type wobaItem struct {
	score   int64
	players float64
	woba    float64
}

type wobaResultContainer []*wobaItem

func (w wobaResultContainer) Len() int           { return len(w) }
func (w wobaResultContainer) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }
func (w wobaResultContainer) Less(i, j int) bool { return w[i].score < w[j].score }

func ProcessBattingWOBAMaps(players PlayersList, result *LeagueResult) {

	resultContainer := make(wobaResultContainer, 0)

	for _, p := range players {
		if p.AB > MinimumAtBats {
			if _, ok := result.BattingWOBAMaps[p.Contact]; ok {
				result.BattingWOBAMaps[p.Contact] = append(result.BattingWOBAMaps[p.Contact], p.WOBA)
			} else {
				result.BattingWOBAMaps[p.Contact] = make([]float64, 0)
				result.BattingWOBAMaps[p.Contact] = append(result.BattingWOBAMaps[p.Contact], p.WOBA)
			}
		}
	}

	for s, w := range result.BattingWOBAMaps {
		wResult := new(wobaItem)
		resultContainer = append(resultContainer, wResult)

		wResult.score = s

		counter := 0.0
		total := 0.0

		for _, v := range w {
			counter = counter + 1
			total = total + v
		}

		wResult.players = counter
		wResult.woba = total / counter
	}

	sort.Sort(resultContainer)

	for _, v := range resultContainer {
		fmt.Printf("Score: %d Players: %0.0f wOBA: %10.3f\n", v.score, v.players, v.woba)
	}
}
