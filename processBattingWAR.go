package main

import (
	"sort"
)

type warItem struct {
	warScore        float64
	numberOfPlayers int64
}

type warItems []*warItem

func (w warItems) Len() int           { return len(w) }
func (w warItems) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }
func (w warItems) Less(i, j int) bool { return w[i].warScore > w[j].warScore }

func ProcessBattingWAR(players PlayersList, result *LeagueResult) {
	result.BattingWARMaps = make(warItems, 0)
	tempBattingWARMaps := make(map[float64]int64)
	for _, p := range players {
		if p.AB > MinimumAtBats {
			numberOfPlayers, ok := tempBattingWARMaps[p.BWAR]
			if ok == false {
				tempBattingWARMaps[p.BWAR] = 1
			} else {
				tempBattingWARMaps[p.BWAR] = numberOfPlayers + 1
			}
		}
	}

	for key, val := range tempBattingWARMaps {
		wItem := new(warItem)
		wItem.numberOfPlayers = val
		wItem.warScore = key

		result.BattingWARMaps = append(result.BattingWARMaps, wItem)
	}

	sort.Sort(result.BattingWARMaps)

}
